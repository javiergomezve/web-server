package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewServer(":3030")
	server.Handle(http.MethodGet, "/", HandleRoot)
	server.Handle(http.MethodGet, "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle(http.MethodPost, "/api", HandlePostApi)
	server.Handle(http.MethodPost, "/api/user", UserPost)

	fmt.Println(server.router.rules)

	server.Listen()
}
