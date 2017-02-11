package main

/* TODO
	abstract port in http.ListenAndServe
  add /healthz context
	add /version context
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const app = "HelloStrata"
const version = "1.0.0"
const port = "9000"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Hello Strata from %s", hostname)

}

func main() {
	log.Printf("Starting %s v.%s on port %s...", app, version, port)

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9000", nil)
}
