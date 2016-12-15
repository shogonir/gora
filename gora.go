package main

import (
  "fmt"
  "time"
  "flag"
  "net/http"
  "strconv"
)

func strfNow() string {
  now:= time.Now()
  const layout = "2006-01-02 15:04:05"
  return now.Format(layout)
}

func suffix(iter int, length int) string {
  if (iter == length) {
    return "\n"
  } else {
    return ",\n"
  }
}

func writeHeaders(w http.ResponseWriter, headers http.Header) {
  fmt.Fprint(w, "\"headers\":{\n")
  iter, length := 0, len(headers)
  for key, value := range headers {
    iter++
    fmt.Fprintf(w, "\"" + key + "\":\"" + value[0] + "\"" + suffix(iter, length))
  }
  fmt.Fprint(w, "}")
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "{\"time\" : \"" + strfNow() + "\",\n")

  writeHeaders(w, r.Header)

  fmt.Fprint(w, "}")
}

func main() {

  port := flag.Int("port", 7999, "listen port number")
  flag.Parse()

  http.HandleFunc("/gora/", handler)

  err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
  if err != nil {
    fmt.Println(err)
  }
}