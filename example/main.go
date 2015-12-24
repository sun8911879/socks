package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/sun8911879/socks"
)

func main() {
	helperFullPath := "socks-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	log.Println("Using icon at %v", iconFullPath)
	err := socks.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return
	}
	err = socks.On("127.0.0.1", "123")
	if err != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return
	}
	socks.Off()
}
