package main

import (
	"flag"
	"log"
	"math"
	"net/http"
	"strconv"
)

const ProgName string = "Serve"
const MaxPort uint = math.MaxUint16
const MinPort uint = 0

var body *string
var path *string
var port *uint

func bodyHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(*body))
	if err != nil {
		log.Println(ProgName, ": Issue handling endpoint at: ", *path, ". Port: ", *port, ". Body: ", *body, ". Error: ", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main(){
	//Getting arguments
	body = flag.String("body", "OK", "The content to return on request")
	path = flag.String("path", "/", "The path to serve content")
	port = flag.Uint("port", 80, "The port to run the HTTP server on")

	flag.Parse()

	//Validating provided ports: 0 — 2*16-1
	if *port > MaxPort || *port < MinPort {
		log.Fatal(ProgName, ": Invalid port provided. Must be between ", MinPort, " and ", MaxPort, ". Provided: ", *port)
	}

	//Putting up server
	http.HandleFunc(*path, bodyHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(int(*port)), nil)
	if err != nil {
		log.Fatal(ProgName, ": Error starting HTTP. Port: ", *port, ". Path: ", *path, ". Body: ", *body, ". Error: ", err)
	}
}
