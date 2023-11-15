package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func main() {
	json.Marshal("s")
	fmt.Println(runtime.Version())
}
