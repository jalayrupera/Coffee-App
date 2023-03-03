package data

import (
	"context"
	"log"
	"time"

	"github.com/jalayrupera/coffee-app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))

	if err != nil {
		log.Fatal("Could not create mongo client", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DATABASE - Connected to MongoDB")
	return client
}
