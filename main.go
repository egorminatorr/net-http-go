package main

import (
	"fmt"
	"net-http-go/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.HandleHello)
	
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Oops...")
	}
}
