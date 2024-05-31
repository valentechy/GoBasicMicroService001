package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("GET /hello/{name}", helloWeb)

  log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
  name := r.PathValue("name")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // Create basic user
  user := &User {
    Id: 1,
    Name: "Valentin",
    LastName: "Palacios",
  }

  json.NewEncoder(w).Encode(user)
  fmt.Fprintf(w, "Hello %s\n", name)
}
