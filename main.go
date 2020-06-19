package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var correctedData [][]byte
	cArgs := os.Args[1:]
	buf := bufio.NewReader(os.Stdin)

	for {
		data, dataErr := buf.ReadBytes('\n')
		if dataErr != nil && dataErr != io.EOF {
			fmt.Fprintln(os.Stderr, dataErr)
			os.Exit(1)
		}

		dataSplit := bytes.Split(data, []byte(""))

		for i, j := range dataSplit {
			if bytes.Compare(j,[]byte(cArgs[0])) == 0 {
				dataSplit[i] = []byte(cArgs[1])
			}
		}

		correctedData = append(correctedData, bytes.Join(dataSplit,[]byte("")))

		if dataErr == io.EOF {
			break
		}
	}

	for i := range correctedData {
		fmt.Printf("%s", correctedData[i])
	}
}