package iocpb

import (
	"io"
	"sync"
	"testing"
)

type writer struct {
	io.Writer
}

func (w *writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type reader struct {
	pos int
	len int
	io.Reader
}

func (r *reader) Read(p []byte) (n int, err error) {
	if r.pos == r.len {
		return 0, io.EOF
	}

	read := r.len - r.pos
	if read > len(p) {
		read = len(p)
	}

	r.pos += read
	return read, nil
}

func BenchmarkIoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := &writer{}
		r := &reader{len: 1024 * 1024}

		io.Copy(w, r)
	}
}

func BenchmarkIoCopyBufferWithPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 32*1024)
		},
	}

	for i := 0; i < b.N; i++ {
		w := &writer{}
		r := &reader{len: 1024 * 1024}
		buffer := pool.Get().([]byte)

		io.CopyBuffer(w, r, buffer)

		pool.Put(buffer)
	}
}

func BenchmarkIoCopyBuffer(b *testing.B) {
	buffer := make([]byte, 32*1024)
	for i := 0; i < b.N; i++ {
		w := &writer{}
		r := &reader{len: 1024 * 1024}

		io.CopyBuffer(w, r, buffer)
	}
}
