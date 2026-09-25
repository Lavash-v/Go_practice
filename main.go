package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	topCount := flag.Int("top", 5, "Количество слов для вывода в топе наиболее часто встречающихся")
	linesOnly := flag.Bool("l", false, "Выводить только общее количество строк в файле")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Ошибка: не указан путь к файлу.")
		fmt.Println("Использование: go run main.go [-top=N] [-l] <путь_к_файлу>")
		os.Exit(1)
	}

	filePath := args[0]

	linesCount := 42 
	fmt.Printf("(Читаем файл: %s)\n\n", filePath)

	if *linesOnly {
		fmt.Println(linesCount)
		return
	}

	fmt.Printf("Общее количество строк: %d\n", linesCount)

	fmt.Printf("Топ-%d популярных слов:\n", *topCount)
}
