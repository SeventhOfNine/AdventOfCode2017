package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(f string) []byte {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}
