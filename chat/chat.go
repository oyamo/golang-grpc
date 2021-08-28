package chat

import (
	"context"
	"log"
)

type Server struct {

}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Println("Received Message from client:", message.Body)
	return &Message{Body: "Affirmative"}, nil
}
