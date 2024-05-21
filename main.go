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

type begR struct {
	ID   int `json:"id"`
	Posne string `json:"posne"`
	Response int `json:"response"`
}

func main() {	

    router := gin.Default()
    router.GET("/beg", getBeg)

    router.Run("localhost:8080")
}

// GET - Beg Responses
func getBeg(c *gin.Context) {
  file, err := os.Open("csv/beg.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var begRos []begR
	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		response, _ := strconv.Atoi(record[1])
		begr := begR{
			ID:   id,
			Posne: record[1],
			Response:  response,
		}
		begRos = append(begRos, begr)
	}

  	fmt.Println(begRos)
    c.IndentedJSON(http.StatusOK, begRos)
}

// IGNORE BELOW

