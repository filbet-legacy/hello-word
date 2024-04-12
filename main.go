package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler function returns "Hello World" to the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// Start an HTTP server listening on port 80
	// 注意: 在80端口上监听可能需要管理员权限
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
