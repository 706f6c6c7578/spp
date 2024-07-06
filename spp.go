package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const signatureHeader = "X-Ed25519-Sig: "
const edpubHeader = "X-Ed25519-Pub: "

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: spp <path to private key file>")
		return
	}

	keyFile := os.Args[1]

	privateKeyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	privateKey, err := hex.DecodeString(string(privateKeyBytes))
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	// Generate public key from private key
	publicKey := privateKey[32:]
	publicKeyHex := hex.EncodeToString(publicKey)

	messageBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read the message from stdin: %v", err)
	}

	// Add public key header to the message
	pubKeyHeader := edpubHeader + publicKeyHex + "\r\n"
	messageWithPubKey := append([]byte(pubKeyHeader), messageBytes...)

	// Include header values in the message body
	messageToSign := appendHeaderValues(messageWithPubKey)

	signature := signMessage(privateKey, messageToSign)
	signatureHex := hex.EncodeToString(signature)

	// Fold the signature into two lines
	header := signatureHeader + signatureHex[:64] + "\r\n " + signatureHex[64:] + "\r\n"
	
	// Prepend the header to the message
	messageWithHeader := append([]byte(header), messageWithPubKey...)

	fmt.Print(string(messageWithHeader))
}

func signMessage(privateKey []byte, message []byte) []byte {
	return ed25519.Sign(privateKey, message)
}

func appendHeaderValues(message []byte) []byte {
	headers := strings.SplitN(string(message), "\r\n\r\n", 2)[0]
	var headerValues []string
	lines := strings.Split(headers, "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, edpubHeader) {
			headerValues = append(headerValues, strings.TrimPrefix(line, edpubHeader))
		}
	}

	// Append header values to the message body
	splitMessage := strings.SplitN(string(message), "\r\n\r\n", 2)
	messageBody := ""
	if len(splitMessage) > 1 {
		messageBody = splitMessage[1]
	}
	return []byte(strings.Join(headerValues, "\r\n") + "\r\n" + messageBody)
}
