package main

import (
	"context"
	"log"
	"net"

	"github.com/nathanverse/mygo/apis/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	hello.RegisterGreeterServer(s, server{})
	lis, _ := net.Listen("tcp", "127.0.0.1:8080")
	log.Fatal(s.Serve(lis))
}
