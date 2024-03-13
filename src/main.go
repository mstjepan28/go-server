package main

import (
	"encoding/json" 
	"log" 
	"net/http"
) 

type Response struct {
	Message string `json:"message"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, world!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Starting server on port 8080")
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
