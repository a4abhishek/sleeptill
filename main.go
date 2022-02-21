package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("required exactly 1 argument, example: \n%s %q\n", os.Args[0], time.Now().Format(time.RFC3339))
		return
	}

	t, err := time.Parse(time.RFC3339, os.Args[1])
	if err != nil {
		fmt.Printf("failed parsing time string, requires RFC3339 timestamp format, example: \n%s %q\n", os.Args[0], time.Now().Format(time.RFC3339))
		return
	}

	<-time.After(t.Sub(time.Now()))
}
