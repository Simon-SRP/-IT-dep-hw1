package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UP(uniqueLines map[string]bool) []string {
	result := make([]string, 0, len(uniqueLines))
	for line := range uniqueLines {
		result = append(result, line)
	}
	for i := range result {
		result[i] = strings.ToUpper(result[i])
	}
	return result
}

func findUnicString(file_name string) ([]string, error) {

	file, err := os.Open(file_name)
	if err != nil {
		return []string{}, fmt.Errorf("Ошибка открытия файла %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	uniqueLines := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		uniqueLines[line] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка сканирования файла: %w", err)
	}

	result := UP(uniqueLines)
	return result, err
}

func outputsResult(filename string, result *os.File) error {
	uniqLines, err := findUnicString(filename)
	if err != nil {
		return fmt.Errorf("Ошибка поиска уникальных строк: %w", err)
	}

	for _, line := range uniqLines {
		_, err := result.WriteString(fmt.Sprintf("%s - %d байт\n", line, len(line)))
		if err != nil {
			return fmt.Errorf("Ошибка записи в файл: %w", err)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("вы не указали путь к файлу")
		return
	}
	name := os.Args[1]
	result, err := os.Create("resulted.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("ошибка создания файла: %w", err))
		return
	}
	defer result.Close()
	err = outputsResult(name, result)
	if err != nil {
		fmt.Println(fmt.Errorf("ошибка записи в файл: %w", err))
		return
	}
}
