package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// URL
	url := ""// domain removed due to potential copyright
	// cache because data is huge
	cacheFileName := "cached_data.json"

	if _, err := os.Stat(cacheFileName); err == nil { // Check if the file is cached locally

		data, err := os.ReadFile(cacheFileName)
		if err != nil {
			fmt.Println("Error reading from cache:", err)
			return
		}

		processJSON(data)
	} else { // URL FETCH
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data from URL:", err)
			return
		}
		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// create cache file
		err = os.WriteFile(cacheFileName, data, 0644)
		if err != nil {
			fmt.Println("Error writing to cache:", err)
			return
		}

		processJSON(data)
	}
}

func processJSON(data []byte) {
	// Unmarshal JSON data
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("JSON Data:", jsonData)
}
