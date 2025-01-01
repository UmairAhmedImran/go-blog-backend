package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found or there is some issue")
	}

	fmt.Println("PORT: ", portString)
}
