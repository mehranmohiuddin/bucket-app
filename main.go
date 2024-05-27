package main

import (
	"fmt"

	fileUpdate "github.com/mehranmohiuddin/bucket-app/internal/file_update"
)

func main() {
	fmt.Println("Hello from bucket app")

	workflow := fileUpdate.WorkflowImpl{}
	workflow.Process()
}
