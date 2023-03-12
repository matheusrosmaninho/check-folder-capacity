package helpers

import (
	"fmt"
	"os"
)

func CreateFile(size int64, directory string, fileName string) error {
	totalSize := size * 1024 * 1024 * 1024
	fd, err := os.Create(fmt.Sprintf("%s/%s", directory, fileName))
	if err != nil {
		return err
	}

	_, err = fd.Seek(totalSize-1, 0)
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte{0})
	if err != nil {
		return err
	}
	err = fd.Close()
	if err != nil {
		return err
	}
	return nil
}
