package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/root", rootHandler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
			w.Header().Set(k, v[0])
		}
	}
}
