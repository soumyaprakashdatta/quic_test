package main

import (
	"log"
	"net/http"
)

const (
	httpPort string = ":8080"
)

func main() {
	log.Print("starting quic_test app")
	startHttpServer()
	select {}
}

func startHttpServer() {
	http.HandleFunc("/test", handleTest)
	log.Printf("starting http server at %s", httpPort)
	http.ListenAndServe(httpPort, nil)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}