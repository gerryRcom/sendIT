package main

//

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// loggeddata represents data reported by server.
type loggeddata struct {
	Client string `json:"client"`
	Date   string `json:"date"`
	Server string `json:"server"`
	Space  string `json:"space"`
}

// loggeddata slice to seed data.
var loggeddatas = []loggeddata{
	{Client: "ZeroApps", Date: "20/02/23", Server: "win-server1", Space: "25"},
}

func main() {
	router := gin.Default()
	router.GET("/loggeddatas", getLoggeddata)
	router.GET("/loggeddatas/:client", getLoggeddataByClient)
	router.POST("/loggeddatas", postLoggeddata)

	router.Run("localhost:8080")
}

// getLoggeddata responds with the list of all data as JSON.
func getLoggeddata(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loggeddatas)
}

// postLoggeddata adds an loggeddata from JSON received in the request body.
func postLoggeddata(c *gin.Context) {
	var newLoggeddata loggeddata

	// Call BindJSON to bind the received JSON to
	// newLoggeddata.
	if err := c.BindJSON(&newLoggeddata); err != nil {
		return
	}

	// Add the new loggeddata to the slice.
	loggeddatas = append(loggeddatas, newLoggeddata)
	c.IndentedJSON(http.StatusCreated, newLoggeddata)
}

// getLoggeddataByClient locates the data for a specific client
// parameter sent in the request
func getLoggeddataByClient(c *gin.Context) {
	client := c.Param("client")

	// Loop over the list of data, looking for specific client data
	for _, a := range loggeddatas {
		if a.Client == client {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "client not found"})
}
