package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var pokupki []string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nСписок покупок.\nКоманды: Добавить, Изменить, Удалить, Список, Команды, Стоп")

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := strings.TrimSpace(strings.ToLower((scanner.Text())))

		switch command {
		case "добавить":
			pokupki = addItem(pokupki, scanner)

		case "изменить":
			pokupki = changeItem(pokupki, scanner)

		case "удалить":
			pokupki = deleteItem(pokupki)

		case "список":
			showList(pokupki)

		case "стоп":
			fmt.Println("\nЗавершение программы")
			return

		case "команды":
			fmt.Println("\nКоманды: Добавить, Изменить, Удалить, Список, Команды, Стоп")
		default:
			fmt.Println("\nНеизвестная команда. Введите 'Команды' чтобы увидеть список доступных.")
		}
	}

}

func addItem(pokupki []string, scanner *bufio.Scanner) []string {
	fmt.Print("\nЧто добавить?\nЕсли не хотите ничего добавлять напишите 'назад'")
	maxAttempts := 3
	for attempts := maxAttempts; attempts > 0; attempts-- {
		fmt.Print("\n> ")

		scanner.Scan()
		item := strings.TrimSpace(scanner.Text())

		if strings.ToLower(item) == "назад" {
			fmt.Println("Возврат в меню")
			return pokupki
		}
		if item == "" {
			fmt.Printf("\nПустой ввод. Ничего не добавлено, осталось попыток: %d.\nЕсли не хотите ничего добавлять напишите 'назад'", attempts-1)
			continue
		}
		pokupki = append(pokupki, item)
		fmt.Printf("\nДобавлено: %s\n\n", item)
		return pokupki
	}
	fmt.Println("Превышено количестов попыток. Возврат в меню")
	return pokupki
}

func changeItem(pokupki []string, scanner *bufio.Scanner) []string {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст.")
		return pokupki
	}
	showList(pokupki)
	fmt.Println("Какой пункт вы хотите изменить?")

	index, ok := inputIndex(len(pokupki))

	if !ok {
		fmt.Println("Превышено количество попыток. Возврат в меню.")
		return pokupki
	}

	fmt.Println("На что хотите заменить?")

	fmt.Print("> ")
	scanner.Scan()
	newItem := strings.TrimSpace(scanner.Text())

	oldItem := pokupki[index-1]
	pokupki[index-1] = newItem
	fmt.Printf("\nЭлемент изменен.\nВы заменили '%s', на '%s'.\n", oldItem, newItem)

	return pokupki
}

func deleteItem(pokupki []string) []string {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст")
		return pokupki
	}

	showList(pokupki)
	fmt.Println("Какой элемент вы хотите удалить")

	index, ok := inputIndex(len(pokupki))

	if !ok {
		fmt.Println("Количество попыток превышено. Возврат в меню")
		return pokupki
	}

	removedItem := pokupki[index-1]
	pokupki = append(pokupki[:index-1], pokupki[index:]...)

	fmt.Printf("Вы удалили %s\n", removedItem)
	return pokupki

}
func showList(pokupki []string) {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст, сначала добавьте что-нибудь")
		return
	}
	fmt.Println("\nВаш список продуктов:")
	for i, value := range pokupki {
		fmt.Printf("%d. %s\n", i+1, value)
	}
}

func inputIndex(limit int) (int, bool) {
	var index int

	maxAttempts := 3

	for attempts := maxAttempts; attempts > 0; attempts-- {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&index); err != nil || index <= 0 || index > limit {
			fmt.Printf("Неккоректный номер, попробуйте снова, осталось попыток: %d\n", attempts-1)
			continue
		}
		return index, true
	}
	return 0, false
}
