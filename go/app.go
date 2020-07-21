package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string

func init() {
	port = "8080"
}

func main() {

	// Return a collection of topic subscriptions
	http.HandleFunc("/schedule-trigger", func(w http.ResponseWriter, r *http.Request) {
		log.Println("trigger invoked")
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}
