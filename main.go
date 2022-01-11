// main.go
package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("test", version, runtime.GOOS, runtime.GOARCH)
}
