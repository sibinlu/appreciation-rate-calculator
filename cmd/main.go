package main

import (
	"fmt"
	"os"
	"time"

	"sibinlu/arc/pkg/zillow"
)

func main() {
	fmt.Printf(
		"Hello, Kat! What a Sunny day today! @ %v\n",
		time.Now().Format("Jan 02, 2006 3:04 PM"))

	if len(os.Args) < 2 {
		fmt.Println("And we need the zpid to run!")
		return
	}
	zillow.GetAppreciationRate(os.Args[1])
}
