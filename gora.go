package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func strfNow() string {
	now := time.Now()
	const layout = "2006-01-02 15:04:05"
	return now.Format(layout)
}

func suffix(iter int, length int) string {
	if iter == length {
		return "\n"
	} else {
		return ",\n"
	}
}

func displayHeaders(headers http.Header) {
	fmt.Println("\"headers\": {")
	iter, length := 0, len(headers)
	for key, value := range headers {
		iter++
		fmt.Printf("    \"" + key + "\":\"" + value[0] + "\"" + suffix(iter, length))
	}
	fmt.Println("}")
	fmt.Println()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(strfNow() + " " + r.RequestURI)

	displayHeaders(r.Header)

	fmt.Fprint(w, "ok")
}

func main() {

	port := flag.Int("port", 7999, "listen port number")
	flag.Parse()

	fmt.Println()
	fmt.Println("'gora' has started.")
	fmt.Printf("listening: ':%d/gora/*'\n\n", *port)

	http.HandleFunc("/gora/", handler)

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
