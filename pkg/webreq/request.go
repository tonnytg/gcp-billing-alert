package webreq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// only works with this methods
const (
	MethodGet = "GET"
)

// Get receive an url, you can send headers and timeout parameters for request.
func Get(url string, timeOut int) ([]byte, error) {

	fmt.Println("url:", url)
	// client run everything
	client := &http.Client{
		CheckRedirect: nil,
	}

	// context help us to control timeout for request
	ctx := context.Background()
	if timeOut == 0 {
		timeOut = 10
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	// request with NewRequest permit add headers
	request, err := http.NewRequestWithContext(ctx, MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("GCP_TOKEN")
	bearer := fmt.Sprintf("Bearer %s", token)

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Accept", "Application/json")

	// execute call and return *http.Response type
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	// clone request after process
	defer response.Body.Close()

	// convert information to slice of byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
