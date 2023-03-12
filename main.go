package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/matheusrosmano/check-pendrive-capacity/helpers"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	path := os.Getenv("DIRECTORY_PATH")

	filesInDirectory, err := helpers.TotalFilesInDir(path)
	if err != nil {
		panic(err)
	}
	i := filesInDirectory + 1

	for {
		fileName := "file" + strconv.Itoa(i)
		err := helpers.CreateFile(1, path, fileName)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Println(fmt.Sprintf("File \"%s\" created ...", fileName))
		i++
	}
}
