package main

import (
	"context"
	"log"

	"github.com/nathanverse/mygo/apis/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	client, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	c := hello.NewGreeterClient(client)
	var p peer.Peer
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{
		Name: "hihi",
	}, grpc.Peer(&p))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("reply:", r.Message)

	if info, ok := p.AuthInfo.(credentials.TLSInfo); ok {
		log.Printf("TLS used, version=%x, cipher=%x", info.State.Version, info.State.CipherSuite)
	} else {
		log.Println("NOT using TLS")
	}
}
