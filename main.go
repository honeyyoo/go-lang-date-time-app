package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    loc, err := time.LoadLocation("Asia/Kolkata")
    var currentTimeIST string

    if err == nil {
        currentTimeIST = time.Now().In(loc).Format("2006-01-02 15:04:05")
    } else {
        ist := time.FixedZone("IST", 5*3600+1800)
        currentTimeIST = time.Now().In(ist).Format("2006-01-02 15:04:05")
    }

    currentTimeUTC := time.Now().UTC().Format("2006-01-02 15:04:05")

    html := `
    <html>
    <head>
        <title>Date & Time App</title>
        <style>
            body {
                background: linear-gradient(135deg, #74ebd5 0%, #ACB6E5 100%);
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
                color: #333;
                display: flex;
                justify-content: center;
                align-items: center;
                height: 100vh;
                margin: 0;
            }
            .container {
                background: white;
                padding: 30px 50px;
                border-radius: 12px;
                box-shadow: 0 8px 16px rgba(0,0,0,0.2);
                text-align: center;
                font-size: 28px;
            }
            strong {
                color: #0077cc;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <p><strong>Current Date & Time (IST):</strong><br>` + currentTimeIST + `</p>
            <p><strong>Current Time (UTC):</strong><br>` + currentTimeUTC + `</p>
        </div>
    </body>
    </html>`

    fmt.Fprint(w, html)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}