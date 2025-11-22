package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	serveDir := os.DirFS("./srv")
	handler := func(res http.ResponseWriter, req *http.Request) {
		path := req.PathValue("file")
		if path == "" {
			path = "index.html"
		}
		file, err := fs.ReadFile(serveDir, path)
		if err != nil {
			res.WriteHeader(400)
			fmt.Fprint(res, err)
			return
		}
		fmt.Fprint(res, string(file))
	}
	http.HandleFunc("GET /{file...}", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
