package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log/slog"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	bucketName = ""
	fileName   = ""
)

func main() {
	slog.Info("Hello from bucket app")

	ctx := context.Background()

	// Initialize client with service account credentials
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		slog.Error("Failed to create client: %v", err)
		return
	}
	defer client.Close()

	// Get the file from the bucket
	rc, err := client.Bucket(bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		slog.Error("Failed to get object: %v", err)
		return
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		slog.Error("Failed to read object: %v", err)
		return
	}

	// Append the current date to the file content
	currentDate := time.Now().Format(time.RFC3339)
	updatedData := fmt.Sprintf("%s\nUpdated on: %s", string(data), currentDate)

	// Upload the updated file back to the bucket
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	wc.ContentType = "text/plain"
	if _, err := wc.Write([]byte(updatedData)); err != nil {
		slog.Error("Failed to write object: %v", err)
		return
	}
	if err := wc.Close(); err != nil {
		slog.Error("Failed to close writer: %v", err)
		return
	}

	slog.Info("File updated successfully")
}
