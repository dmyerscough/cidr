package main

import (
	"fmt"
	"os"
	"github.com/dmyerscough/cidr/validator"
	"path"
)

func main() {
	if len(os.Args[1:]) > 1 {
		if err := validator.CheckOverlapping(os.Args[1:]); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Printf("Error: Two subnets are required at a minimum\n\n")
		fmt.Printf("Usage:\n\t./%s 10.0.0.0/24 10.1.0.0/24\n", path.Base(os.Args[0]))
	}
}
