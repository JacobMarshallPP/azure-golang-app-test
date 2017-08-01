package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	})
	log.Println("Ello!")
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil))
}
