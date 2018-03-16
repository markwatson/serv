package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

func applyHeadersHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

func registerHandlers(mux *http.ServeMux, paths map[string]string) {
	var handler http.Handler
	for pattern, root := range paths {
		fileInfo, err := os.Stat(root)

		if err != nil {
			log.Println("Error: path doesn't exist: " + root)
		} else {
			if fileInfo.IsDir() {
				log.Printf("Registering handler with pattern: %s, root path: %s",
					pattern, root)
				handler = applyHeadersHandler(http.FileServer(http.Dir(root)))
				mux.Handle(pattern, http.StripPrefix(pattern, handler))
			} else {
				log.Printf("Registering handler with pattern: %s, file: %s",
					pattern, root)
				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.ServeFile(w, r, root)
				})
				mux.Handle(pattern, applyHeadersHandler(handler))
			}
		}
	}
}

func parsePaths(paths string) (pathMap map[string]string, err error) {
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

	paths, err := parsePaths(*pathsRaw)

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	registerHandlers(mux, paths)
    log.Println("serv 1.0.0 started :)")
	log.Println("Listening on: ", *listen)

	log.Fatal(http.ListenAndServe(*listen, mux))
}
