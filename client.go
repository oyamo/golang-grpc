package main

import (
	"fmt"
	"golang-rpc/chat"
	"google.golang.org/grpc"
	"log"
)

import (
	"context"
)

func main() {
	var conn *grpc.ClientConn
	log.Println("Connecting...")

	conn, err := grpc.Dial(":3030", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	log.Println("Connected to client")

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	for i := 0; i < 10; i++{
		var input string
		log.Print("Enter > ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			log.Fatalf("Error Scanning STDIN: %v", err)
		}

		msg, err := c.SayHello(context.Background(), &chat.Message{Body: input})
		if err != nil {
			log.Fatalf("Error sending Chat: %v", err)
		}

		log.Println(msg)
	}
}
