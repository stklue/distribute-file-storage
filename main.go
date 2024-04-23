package main

import (
	"fmt"
	"log"

	"github.com/stklue/diststore/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	transport := p2p.NewTCPTransport(tcpOpts)

	go func() {
		msg := <-transport.Consume()
		fmt.Println(msg)
	}()

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
}
