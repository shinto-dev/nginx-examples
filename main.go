package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-example/proto"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedGreetServiceServer
}

func (s Server) Greet(context.Context, *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("Got new request")
	return &pb.GreetResponse{
		Result: "hello world",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	fmt.Println("Starting server")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
