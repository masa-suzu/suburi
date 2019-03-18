package io

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReadWriteLine(t *testing.T) {

	tests := []struct {
		name string
		in   string
		act  func(*LineReadWriter)
		want string
	}{
		{
			name: "read-and-write-once",
			in:   "1",
			act: func(lrw *LineReadWriter) {
				l := lrw.ReadLine()
				lrw.WriteLine(l)
			},
			want: "1\n",
		},
		{
			name: "read-and-write-twice",
			in:   "a",
			act: func(lrw *LineReadWriter) {
				l := lrw.ReadLine()
				lrw.WriteLine(l)
				lrw.WriteLine(l)
			},
			want: "a\na\n",
		},
	}

	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := bytes.NewBufferString(tt.in)
			w := bytes.NewBufferString("")

			rw := LineReadWriter{
				In:  bufio.NewScanner(r),
				Out: bufio.NewWriter(w),
			}

			tt.act(&rw)

			_ = rw.Out.Flush()

			got := w.String()

			if tt.want != got {
				t.Errorf("\nw.String()\t%#v\nwant\t%#v", got, tt.want)
			}
		})
	}
}
