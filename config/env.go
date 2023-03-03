package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURI() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Could not load ENV file")
	}

	return os.Getenv("MONGOURI")

}
