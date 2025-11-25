package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	handler := func(res http.ResponseWriter, req *http.Request) {
		path := req.PathValue("file")
		if path == "" {
			path = "index.html"
		}
		file, err := renderPage(path)
		if err != nil {
			res.WriteHeader(400)
			fmt.Fprint(res, err)
			return
		}
		fmt.Fprint(res, file)
	}
	http.HandleFunc("GET /{file...}", handler)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
