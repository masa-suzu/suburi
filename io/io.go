package io

import "bufio"

type LineReadWriter struct {
	_ interface{}

	In  *bufio.Scanner
	Out *bufio.Writer
}

func (lrw *LineReadWriter) ReadLine() string {
	if !lrw.In.Scan() {
		return ""
	}
	return lrw.In.Text()
}

func (lrw *LineReadWriter) WriteLine(l string) {
	_, err := lrw.Out.WriteString(l)
	if err != nil {
		panic(err)
	}

	_, err = lrw.Out.WriteString("\n")
	if err != nil {
		panic(err)
	}
}
