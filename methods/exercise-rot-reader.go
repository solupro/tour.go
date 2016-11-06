package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (self *rot13Reader) Read(buff []byte) (int, error) {
	tmp := make([]byte, 8)
	n, err := self.r.Read(tmp)
	for i := 0; i < n; i++ {
		if c := tmp[i]; (c >= 'A' && c < 'N') || (c >= 'a' && c < 'n') {
			buff[i] = c + 13
		} else if (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z') {
			buff[i] = c - 13
		} else {
			buff[i] = c
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
