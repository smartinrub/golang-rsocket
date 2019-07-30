package main

import (
	"context"
	"log"

	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
)

func main() {
	cli, err := rsocket.Connect().
		Resume().
		Fragment(1024).
		SetupPayload(payload.NewString("Hello", "World")).
		Transport("ws://127.0.0.1:8080").
		Start(context.Background())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	result, err := cli.RequestResponse(payload.NewString("Hola", "Mundo")).Block(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("response:", result)

}
