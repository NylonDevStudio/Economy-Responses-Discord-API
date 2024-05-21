package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	ID   int `json:"id"`
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {	file, err := os.Open("csv/beg.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var beg []beg
	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		posne, _ := strconv.Atoi(record[1])
		response := Response{
			ID:   id,
			Posne: posne,
			Response:  record[2]
		}
		begR = append(begR, response)
	}

	fmt.Println(begR)

    router := gin.Default()
    router.GET("/beg", getBeg)
    router.GET("/searchlocal", getSearchlocal)

    router.Run("localhost:8080")
}

// GET - Beg Responses
func getBeg(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, beg)
}

// IGNORE BELOW

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}