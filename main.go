package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home handler was awakened.")
	})
	http.ListenAndServe(":8000", nil)
}
