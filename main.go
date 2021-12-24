package main

import (
	"github.com/YeslieSnayder/WebServiceGo/handlers"
	"log"
	"net/http"
	"os"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router()
	log.Print("The service is ready to listen and serve")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
