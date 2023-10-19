package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed ui/public
var content embed.FS

func main() {
	port := flag.Int("port", 8000, "The port the http server should listen on.")
	swaggerFile := flag.String("swagger-file", "?", "The path to the swagger-file")
	flag.Parse()

	if *swaggerFile == "?" {
		println("You must supply a parameter to -swagger-file option for the swagger json")
		os.Exit(1)
	}

	if _, err := os.Stat(*swaggerFile); os.IsNotExist(err) {
		println(fmt.Sprintf("Unable to find swagger json: %s", *swaggerFile))
		os.Exit(1)
	}

	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/Users/jcorbett/dev/partner-api/src/tsoa/swagger.json")
	})

	var root, _ = fs.Sub(content, "ui/public")
	http.Handle("/", http.FileServer(http.FS(root)))

	println(fmt.Sprintf("Docs available at http://localhost:%d", *port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
