package wav

import (
	"bytes"
	"fmt"

	"github.com/cryptix/wav"
)

// WAV audio file
type WAV struct {
	reader *wav.Reader
	c      chan []int16
	e      chan error
}

// Parse wav file
func (w *WAV) Parse(data []byte) (err error) {
	f := bytes.NewReader(data)
	w.reader, err = wav.NewReader(
		f,
		f.Size(),
	)
	return
}

// GetSample get audio sample
func (w *WAV) GetSample() (<-chan []int16, <-chan error) {
	if w.reader == nil {
		w.e <- fmt.Errorf("wav info not exist")
		return w.c, w.e
	}
	// todo
	// del consts
	go func() {
		var (
			s   []int32
			err error
		)

		for {
			s, err = w.reader.ReadSampleEvery(2, 0)

			samples := make([]int16, 0, len(s))
			for _, sample := range s {
				samples = append(samples, int16(sample))
			}

			w.c <- samples
			w.e <- err
		}
	}()
	return w.c, w.e
}

// NewWAV return handler wav file
func NewWAV() *WAV {
	return &WAV{
		c: make(chan []int16),
		e: make(chan error),
	}
}
