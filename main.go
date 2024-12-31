package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("Hello World")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found or there is some issue")
	}

	fmt.Println("PORT: ", portString)
}
