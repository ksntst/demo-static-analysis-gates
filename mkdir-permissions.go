package main

import (
	"os"
)

func main() {
	err := os.Mkdir("/path/to/new/directory", 0777)
	if err != nil {
		return
	}
}
