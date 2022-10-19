package main

import (
	"fmt"
	"github.com/tonnytg/gcp-billing-alert/pkg/webreq"
	"os"
)

func main() {

	project := os.Getenv("GCP_PROJECT")
	region := os.Getenv("GCP_REGION")

	url := "https://bigqueryreservation.googleapis.com/v1/projects/%s/locations/%s/reservations"
	url = fmt.Sprintf(url, project, region)

	body, err := webreq.Get(url, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
