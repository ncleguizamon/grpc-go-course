package main

import (
	
	"log"
	"net"
	"fmt"
    "github.com/ncleguizamon/grpc-go-course/greet/greetpb"
	grpc "google.golang.org/grpc"
	
)
type server struct{}

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen! %v" , err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("failed to serve %v", err)
	}


}