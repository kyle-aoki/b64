package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

const ErrMsg = "b64 encode <string>\nb64 decode <string>"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	args := os.Args[1:]
	if len(args) < 2 {
		panic(ErrMsg)
	}
	if args[0] == "encode" {
		s := args[1]
		se := base64.StdEncoding.EncodeToString([]byte(s))
		fmt.Print(se)
	} else if args[0] == "decode" {
		sDec, err := base64.StdEncoding.DecodeString(args[1])
		if err != nil {
			panic(err)
		}
		fmt.Print(string(sDec))
	} else {
		panic(ErrMsg)
	}
}
