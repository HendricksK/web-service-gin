package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
// would usually live in a DB, would probably use POSTGRES via HEROKU
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/name/:name", getAlbumByTitle)
	router.GET("/albums/artist/:artist", getAlbumsByArtist)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds a new album to the array
func postAlbums(c *gin.Context) {
	var newAlbum album

	// binding the json context to a new object, which is of type album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// get album by Name
func getAlbumByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, a := range albums {
		if a.Title == title {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func getAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")
	var albumList []album

	for _, a := range albums {
		if a.Artist == artist {
			albumList = append(albumList, a)
		}
	}

	if len(albumList) > 0 {
		c.IndentedJSON(http.StatusOK, albumList)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("No albums found for artist %v", artist)})
}
