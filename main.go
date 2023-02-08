package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("An error occurred when getting environment variables: %v\n", err)
	}

	if os.Getenv("DISCORD_TOKEN") == "" {
		log.Fatalf("Missing DISCORD_TOKEN environment variable\n")
	}
	client, err := disgo.New(os.Getenv("DISCORD_TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		bot.WithEventListenerFunc(func(e *events.Ready) {
			log.Infof("The bot is now ready!\n")
		}),
	)
	if err != nil {
		log.Fatalf("Error creating Discord client: %v\n", err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		log.Fatalf("An error occurred when opening the client: %v\n", err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
