package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	port := ":" + os.Args[1]

	dir := "."
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	fmt.Printf("Serving %s on port %s\n", dir, port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: serve <port> [<directory>]\n")
	os.Exit(2)
}
