package io

import "bufio"

// LineReadWriter reads a line from the buffered scanner
// and writes a line to the buffered writer.
type LineReadWriter struct {
	_ interface{}

	In  *bufio.Scanner
	Out *bufio.Writer
}

// ReadLine reads a line from buffered scanner.
// If no next token can not be scanned, returns ""
func (lrw *LineReadWriter) ReadLine() string {
	if !lrw.In.Scan() {
		return ""
	}
	return lrw.In.Text()
}

// WriteLine writes a string with a new line to buffered writer.
func (lrw *LineReadWriter) WriteLine(l string) error {
	_, err := lrw.Out.WriteString(l)
	if err != nil {
		return err
	}

	_, err = lrw.Out.WriteString("\n")
	if err != nil {
		return err
	}
	return nil
}
