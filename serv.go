package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type Handler struct {
	Pattern, Root string
}

func registerHandlers(mux *http.ServeMux, paths map[string]string) {
	for pattern, root := range paths {
		log.Printf("Registering handler with pattern: %s, root path: %s",
			pattern, root)
		mux.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(root))))
	}
}

func parsePaths(paths string) (err error, pathMap map[string]string) {
	err = json.Unmarshal([]byte(paths), &pathMap)
	return
}

func main() {
	var listen = flag.String("listen", ":8080",
		"Interface/port to listen on. eg. :8080 or 127.0.0.1:8080")
	var pathsRaw = flag.String("paths", `{"/": "."}`,
		"Paths to serve. A json object with the keys as the url pattern, and "+
			"the value as the root. Default serves current folder.")
	flag.Parse()

	err, paths := parsePaths(*pathsRaw)

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	registerHandlers(mux, paths)
	log.Println("Listening on: ", *listen)

	log.Fatal(http.ListenAndServe(*listen, mux))
}
