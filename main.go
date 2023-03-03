package main

import (
	"fmt"

	"github.com/jalayrupera/coffee-app/data"
)

func main() {
	fmt.Println("Welcome to Coffee App")
	data.ConnectToMongo()

	fmt.Println()
}
