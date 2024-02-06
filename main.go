package main

import "starter-go-fiber/bin"

func main() {
	app := bin.NewApp()

	app.Run()
}
