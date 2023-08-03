package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	
	fmt.Printf("Real UID: %d\n", os.Getuid())
	fmt.Printf("Real GID: %d\n", os.Getgid())

	userInfo, err := user.LookupId(fmt.Sprintf("%d", os.Getuid()))
	fmt.Printf("Real user home: %s\n", userInfo.HomeDir)
	fmt.Printf("Error: %v\n", err)

	fmt.Printf("Effec UID: %d\n", os.Geteuid())
        fmt.Printf("Effec GID: %d\n", os.Getegid())

	userInfo, err = user.LookupId(fmt.Sprintf("%d", os.Geteuid()))
	fmt.Printf("Effec user home: %s\n", userInfo.HomeDir)
	fmt.Printf("Error: %v\n", err)

	fmt.Println("Hello World!")
}
