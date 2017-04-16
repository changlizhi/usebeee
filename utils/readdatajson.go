package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAll(file_path string) string {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	fd, err := ioutil.ReadAll(f)
	return string(fd)
}

func main() {
	fmt.Println(ReadAll("../static/data.json"))
}
