package main

import (
	"context"
	"oddysseus/internal/frames"

	"github.com/go-zeromq/zmq4"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Print("TAP SBC System")

	if err := zqmClient(context.Background()); err != nil {
		panic(err)
	}
}

func zqmClient(ctx context.Context) error {
	socket := zmq4.NewDealer(ctx, zmq4.WithAutomaticReconnect(true))
	defer socket.Close()

	webcamHandler := frames.InitializeFrameHandler()

	liveVisionHandler := frames.NewVisionHandler(webcamHandler, socket)

	log.Info().Msg("Connecting to Zeus at tcp://localhost:7207i")
	if err := socket.Dial("tcp://localhost:7207"); err != nil {
		log.Panic().Msg("Cannot find Zeus")
	}

	liveVisionHandler.HandleDealer()

	return nil
}
