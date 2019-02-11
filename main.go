package main

import (
        "flag"
        "fmt"
        "log"
        "net/http"
        "os"

        "github.com/gorilla/mux"
)

func main() {
        log.Printf("Started..")
        port := flag.Int("port", 3000, "-port=2020")
        flag.Parse()

        err := runWebServer(*port)
        if err != nil {
                log.Printf("Error running server: %s", err)
                os.Exit(1)
        }
}

func runWebServer(port int) error {
        log.Printf("Launching Web Server on Port %d", port)
        router := mux.NewRouter()
        router.HandleFunc("/hello/{name}", helloName)
        router.HandleFunc("/hello", helloWorld)
        return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func helloName(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]
        w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
}