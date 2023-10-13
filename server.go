package main

import (
	"fmt"
	"grpc_hello_world/hello"
	"grpc_hello_world/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &service.HelloService{})

	fmt.Println("Server started...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
