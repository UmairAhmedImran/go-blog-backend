package main

import (
	"fmt"
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found or there is some issue")
	}

	router := gin.Default()
	
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
