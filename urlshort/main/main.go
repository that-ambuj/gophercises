package main

import (
	"fmt"
	"net/http"
	"os"

	"flag"
	"urlshort"
)

func main() {
	yamlPath := flag.String("yamlFile", "", "Path to the YAML file")
	jsonPath := flag.String("jsonFile", "", "Path to the json file")
	flag.Parse()

	if *yamlPath == "" && *jsonPath == "" {
		fmt.Fprintln(os.Stderr, "Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var fileHandler http.HandlerFunc

	if *yamlPath != "" {
		yamlData, err := os.ReadFile(*yamlPath)
		if err != nil {
			panic(err)
		}

		fileHandler, err = urlshort.YAMLHandler(yamlData, mapHandler)
		if err != nil {
			panic(err)
		}
	} else {
		jsonData, err := os.ReadFile(*jsonPath)
		if err != nil {
			panic(err)
		}

		fileHandler, err = urlshort.JSONHandler(jsonData, mapHandler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", fileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
