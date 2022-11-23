package main

import (
	"fmt"
	"log"
	"os"

	v1 "github.com/suhriar/submission_maps_api/server/v1"
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
