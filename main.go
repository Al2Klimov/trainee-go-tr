package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	const chunkSize = 4096
	cArgs := os.Args[1:]
	buf := bufio.NewReader(os.Stdin)
	buffer := make([]byte, chunkSize)

	for {
		_, dataErr := buf.Read(buffer)
		if dataErr != nil && dataErr != io.EOF {
			fmt.Fprintln(os.Stderr, dataErr)
			os.Exit(1)
		}

		if dataErr == io.EOF {
			break
		}

		data := bytes.Split(bytes.Trim(buffer, string(0x0)), []byte("\n"))

		for i := range data {
			for j := range data[i] {
				if string(data[i][j]) == cArgs[0] {
					data[i][j] = []byte(cArgs[1])[0]
				}
			}
		}
		fmt.Printf("%s", bytes.Join(data, []byte("\n")))
	}
}
