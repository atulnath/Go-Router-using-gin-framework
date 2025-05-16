package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

// @title Gin API
// @version 1.0
// @description This is a sample server for a note-taking application.

// main initializes the Gin router, sets up controllers, and starts the HTTP server.
func main() {
	// Create a new Gin router with default middleware (logger, recovery)
	myRouter := gin.Default()

	// Initialize the note application controller and attach routes
	controllers.NewAppController(myRouter)

	// Start the server on port 8080
	myRouter.Run(":8080") // listen and serve on http://localhost:8080
}
