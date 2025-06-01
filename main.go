package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    loc, err := time.LoadLocation("Asia/Kolkata")
    var currentTime string

    if err == nil {
        currentTime = time.Now().In(loc).Format("2006-01-02 15:04:05")
    } else {
        ist := time.FixedZone("IST", 5*3600+1800)
        currentTime = time.Now().In(ist).Format("2006-01-02 15:04:05")
    }

    fmt.Fprintf(w, "Current Date & Time (IST): %s\n", currentTime)
    fmt.Fprintf(w, "Current Time (UTC): %s\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}