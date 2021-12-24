package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the service...")

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home handler was awakened.")
	})

	log.Print("Service is ready to listen and serve")
	http.ListenAndServe(":8000", nil)
}
