package main

import (
	"github.com/ncleguizamon/grpc-go-course/greet/greetpb"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	
)
type server struct{}

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:8051")
	if err != nil {
		log.Fatalf("failed to listen! %v" , err)
	}

	s := grpc.Newserver()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err:= s.Server(lis); err !=nil {
		log.Fatalf("failed to serve %v", err)
	}


}