package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiResponse struct {
	Result     bool    `json:"ok"`
	StatusCode int     `json:"error_code"`
	Movies     []Movie `json:"description"`
}

type Movie struct {
	Name        string `json:"#TITLE"`
	ReleaseYear int    `json:"#YEAR"`
	Actors      string `json:"#ACTORS"`
	ImdbURL     string `json:"#IMDB_URL"`
	ImagePoster string `json:"#IMG_POSTER"`
}

type CallResult struct {
	URL        string
	CallDate   string
	Duration   string
	StatusCode int
	Headers    http.Header
	Movies     []Movie
}

func main() {
	endpoint := "https://imdb.iamidiotareyoutoo.com/search?q=Spiderman"
	callCount := 5
	results := make([]CallResult, 0)

	for i := 0; i < callCount; i++ {
		startTime := time.Now()
		response, err := http.Get(endpoint)
		if err != nil {
			fmt.Printf("Error in call %d: %s\n", i+1, err)
			continue
		}

		duration := time.Since(startTime)

		dataInBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading body in call %d: %s\n", i+1, err)
			continue
		}
		response.Body.Close()

		data := ApiResponse{}
		err = json.Unmarshal(dataInBytes, &data)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON in call %d: %s\n", i+1, err)
			continue
		}

		result := CallResult{
			URL:        endpoint,
			CallDate:   time.Now().Format("2006-01-02 15:04:05"),
			Duration:   duration.String(),
			StatusCode: response.StatusCode,
			Headers:    response.Header,
			Movies:     data.Movies,
		}

		results = append(results, result)
	}

	for index, result := range results {
		fmt.Printf("===[ CALL %d ]===\n", index+1)
		fmt.Printf("URL: %s\n", result.URL)
		fmt.Printf("Date: %s\n", result.CallDate)
		fmt.Printf("Duration: %s\n", result.Duration)
		fmt.Printf("Status Code: %d\n", result.StatusCode)
		fmt.Printf("Headers: %v\n", result.Headers)
		fmt.Printf("Movies:\n")

		for _, movie := range result.Movies {
			fmt.Printf("- %s (%d)\n", movie.Name, movie.ReleaseYear)
		}
		fmt.Println("")
	}
}
