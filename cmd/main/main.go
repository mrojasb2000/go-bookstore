package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrojasb2000/go-bookstore/pkg/routes"
)

var (
	port = ":9010"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010\n")
	log.Fatal(http.ListenAndServe(port, r))
}
