package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (_ MyReader) Read(buff []byte) (int, error) {
	buff[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
