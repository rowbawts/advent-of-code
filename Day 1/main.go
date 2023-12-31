package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Set the URL and cookie information
	inputURL := "https://adventofcode.com/2023/day/1/input"
	cookieName := "session"
	cookieValue := "53616c7465645f5fd20b7a22d11174d58d0cafe5509daeffcea2ffa9981d375aa701714086583f514df432b152fe7122810a85dc6f7a30ed82348fe4b7acfecd"

	// Create a cookie jar
	jar, _ := cookiejar.New(nil)

	// Add the session cookie to the cookie jar
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	}
	u, _ := url.Parse(inputURL)
	jar.SetCookies(u, []*http.Cookie{cookie})

	// Create an HTTP client with the cookie jar
	client := &http.Client{
		Jar: jar,
	}

	// Make the GET request
	response, err := client.Get(inputURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", response.Status)
		return
	}

	// Read the response body dynamically without relying on content length
	var body []byte
	buffer := make([]byte, 1024) // Adjust the buffer size as needed

	for {
		n, err := response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading response body:", err)
			return
		}
		if n == 0 {
			break
		}
		body = append(body, buffer[:n]...)
	}

	// Split the response body into lines
	lines := strings.Split(string(body), "\n")

	var total = 0

	// Process each line
	for _, line := range lines {
		if len(line) > 0 {
			// Find the last digit in the line
			var lastDigit byte
			for i := len(line) - 1; i >= 0; i-- {
				if unicode.IsDigit(rune(line[i])) {
					lastDigit = line[i]
					break
				}
			}

			// Extract the first digit of the line
			var firstDigit byte
			for _, char := range line {
				if unicode.IsDigit(char) {
					firstDigit = byte(char)
					break
				}
			}

			// Combine the first and last digits into a single string
			combinedDigits := string([]byte{firstDigit, lastDigit})

			// Convert the combined digits to an int
			combinedInt, err := strconv.Atoi(combinedDigits)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}

			total += combinedInt
		}
	}

	fmt.Println(total)
}
