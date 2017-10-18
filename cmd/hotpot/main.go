package main

import (
	"log"
	"net/http"

	"github.com/mbesancon/hotpot"
)

func main() {
	http.HandleFunc("/process", hotpot.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
