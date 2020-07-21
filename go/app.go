package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port string

func init() {
	port = "8080"
}

type Joke struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	Value string `json:"value"`
}

func main() {

	http.HandleFunc("/schedule-trigger", func(w http.ResponseWriter, r *http.Request) {

		log.Println("trigger invoked")

		var j Joke

		response, err := http.Get("https://api.chucknorris.io/jokes/random?category=sport")
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &j)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(j.Value)
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}
