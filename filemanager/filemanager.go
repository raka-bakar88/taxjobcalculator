package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("could not openfile")
		fmt.Println(err)
		return nil, err
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading the file content failed.")
		fmt.Println(err)
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	err = file.Close()
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{inputPath, outputPath}
}
