package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeUnicString(file_name string) ([]string, error) {

	file, err := os.Open(file_name)
	if err != nil {
		return []string{}, fmt.Errorf("Файл нельзя прочитать %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	un_lines := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		un_lines[line] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка сканирования файла: %w", err)
	}

	res := make([]string, 0, len(un_lines))
	for line := range un_lines {
		res = append(res, line)
	}
	for i := range res {
		res[i] = strings.ToUpper(res[i])
	}
	return res, err
}

func main() {
	res, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("ошибка создания файла: %v", err)
		return
	}
	name := os.Args[1]
	uniqLines, err := makeUnicString(name)
	for i := range uniqLines {
		res.WriteString(fmt.Sprintf("%s - %d байт\n", uniqLines[i], len(uniqLines[i])))
	}
}
