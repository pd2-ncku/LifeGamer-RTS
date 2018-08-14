package service

import (
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Dummy")
}

func userHandler(w http.ResponseWriter, r *http.Request) {

    if b, err := ioutil.ReadAll(r.Body); err == nil {
        fmt.Printf("%s", b)
    }

    fmt.Fprint(w, "New user add")
}

func Start() {
    log.Println("Starting HTTP listener")

    mux := http.NewServeMux()

    mux.HandleFunc("/", mainHandler)
    mux.HandleFunc("/user", userHandler)

    log.Println("HTTP listening on port 25252")
    http.ListenAndServe(":25252", mux)
}
