package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var serverID int

func init() {
	serverID = rand.Intn(8999) + 1000
}

type Response struct {
	ID int `json:"id"`
}

func getRandomId(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", serverID)
}

func main() {
	http.HandleFunc("/", getRandomId)
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
