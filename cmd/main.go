package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	listenerPort := os.Getenv("HTTP_PORT")
	if listenerPort == "" {
		listenerPort = "8080"
	}
	listenerAddr := os.Getenv("HTTP_ADDR")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		if len(body) == 0 {
			body = []byte("Empty request body!!")
		}
		fmt.Fprintf(w, "Hello from %s\n Request body: %s\n", r.Host, body)

	})
	listenerContext := fmt.Sprintf("%s:%s", listenerAddr, listenerPort)
	http.ListenAndServe(listenerContext, nil)
}
