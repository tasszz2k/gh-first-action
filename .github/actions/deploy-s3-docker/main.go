package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve required input: bucket
	bucket := os.Getenv("INPUT_BUCKET")
	if bucket == "" {
		fmt.Fprintln(os.Stderr, "Error: INPUT_BUCKET is required")
		os.Exit(1)
	}

	// Retrieve optional inputs with defaults.
	bucketRegion := os.Getenv("INPUT_BUCKET_REGION")
	if bucketRegion == "" {
		bucketRegion = "us-east-1"
	}

	distFolder := os.Getenv("INPUT_DIST_FOLDER")
	if distFolder == "" {
		distFolder = "dist"
	}

	// Log the simulated deployment process.
	fmt.Printf("Deploying folder '%s' to bucket '%s' in region '%s'...\n", distFolder, bucket, bucketRegion)

	// Mock deployment: create a fake website URL.
	websiteURL := fmt.Sprintf("http://%s.s3-website-%s.amazonaws.com", bucket, bucketRegion)
	fmt.Printf("Mock deployment successful! Website URL: %s\n", websiteURL)

	// Set the output for GitHub Actions.
	// Note: Although GitHub now recommends using output files,
	// this example uses the traditional "::set-output" syntax for simplicity.
	fmt.Printf("::set-output name=website-url::%s\n", websiteURL)
}