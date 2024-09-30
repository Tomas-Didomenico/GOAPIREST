package main

import (
	server "GOAPI/Server"
)

func main() {
	srv := server.New(":8080")
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
