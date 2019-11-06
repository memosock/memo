package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func CheckTcpError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	CheckTcpError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	CheckTcpError(err)
	_ = outFile.Close()
}