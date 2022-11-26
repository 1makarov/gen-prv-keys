package file

import (
	"os"
	"sync"
)

type File struct {
	sync *sync.Mutex
	file *os.File
}

func New(name string, flag int) (*File, error) {
	fi, err := os.OpenFile(name, flag, 0644)
	if err != nil {
		return nil, err
	}

	file := &File{
		sync: new(sync.Mutex),
		file: fi,
	}

	return file, nil
}

func (f *File) Close() error {
	return f.file.Close()
}

func (f *File) Write(text string) error {
	defer f.sync.Unlock()
	f.sync.Lock()

	_, err := f.file.WriteString(text + "\n")

	return err
}
