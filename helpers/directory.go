package helpers

import (
	"os"
)

func TotalFilesInDir(path string) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	return len(files), nil
}
