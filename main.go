package main

import (
	"fmt"
	"log"
	"os"

	v1 "github.com/michaelwp/golang-gmap-places/server/v1"
)

/*
	Google Map - Service
*/

func main(){
	// set host
	host := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	h := v1.Host(host)

	// call the server
	srv, _, resp := h.Server()

	// print log to the screen
	log.Println(resp)
	log.Fatal(srv.ListenAndServe())
}
