package main

import (
	"app.onca.api/server"
	"google.golang.org/appengine"
)

func main() {
	server.Router()
	appengine.Main()
}
