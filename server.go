package main

import (
	"golang-rpc/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	log.Println("Creating TCP Server..")
	lis,err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatalf("Error Creating A server: %v", err)
	}

	log.Println("Creating RPC Server..")

	grpcServer := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	log.Println("Listening..")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error While setting up GRPC: %v", err)
	}

}
