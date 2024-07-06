# spp
spp - signing post-processor for the Usenet Newsreader flnews.

This is a post-processor for the Unix Newsreader [flnews](https://micha.freeshell.org/flnews/), which
allow you to sign your Usenet articles with an Ed25519 key.

The post-processor adds the follwoing Headers to your articles:

X-Ed25519-Sig: b1395ca908d84cc0f0f2e65a9086fa0b6eebfadcdb90b2da6a98e3129e313a6f  
 14ecd905b67e87b663dc791cc9620e6c42fb81c300a7c9f7a5b1e8365c981b09  
X-Ed25519-Pub: 43e6681c5cf3cdc2f9ccf975f8a01b18c5e84bf0ba00605faba9cc0f8757a117  

key pair generation and article verification can be done with my [sve](https://github.com/stefanclaas/sve) programm.

Please note: While the author of flnews said the post-processor for flnews is not (yet) designed
to be used for signing, it works flawlessly and I have tested it thoroughly.


