package main

import (
	"io"
	"net/http"
	"runtime"
)

func HelloServer(c *http.Conn, req *http.Request) {
	io.WriteString(c, "hello, world!\n")
}
func main() {
	np := runtime.NumCPU()
	runtime.GOMAXPROCS(np)
	http.Handle("/", http.HandlerFunc(HelloServer))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("ListenAndServe: ", err.String())
	}
}
