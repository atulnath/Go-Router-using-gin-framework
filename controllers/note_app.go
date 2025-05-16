package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// app_controller defines the controller for handling note-related API requests.
type app_controller struct{}

// NewAppController initializes the note-related routes and attaches them to the Gin router.
func NewAppController(router *gin.Engine) {
	controllers := &app_controller{}
	// Create a route group for note-related endpoints under "/note_app"
	notes := router.Group("/note_app")
	notes.GET("/", controllers.GetAllNotes())    // GET /note_app to retrieve all notes
	notes.POST("/", controllers.CreateNewNote()) // POST /note_app to create a new note
}

// GetAllNotes returns a handler function for retrieving all notes.
// It responds with a JSON message indicating the GET request was received.
func (n *app_controller) GetAllNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"notes": "Note request (Get) received",
		})
	}
}

// CreateNewNote returns a handler function for creating a new note.
// It responds with a JSON message indicating the POST request was successful.
func (n *app_controller) CreateNewNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"notes": "Note request (Post) Created",
		})
	}
}
