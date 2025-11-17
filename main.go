package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	ID   int
	Name string
	Done bool
}

var nextID = 1
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var pokupki []Item

	fmt.Println("\nСписок покупок.\nКоманды: Добавить, Изменить, Удалить, Список, Отметить, Команды, Стоп")

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := strings.TrimSpace(strings.ToLower((scanner.Text())))

		switch command {
		case "добавить":
			pokupki = addItem(pokupki)

		case "изменить":
			pokupki = changeItem(pokupki)

		case "удалить":
			pokupki = deleteItem(pokupki)

		case "список":
			showList(pokupki)

		case "отметить":
			pokupki = toggleItem(pokupki)

		case "стоп":
			fmt.Println("\nЗавершение программы")
			return

		case "команды":
			fmt.Println("\nКоманды: Добавить, Изменить, Удалить, Список, Отметить, Команды, Стоп")
		default:
			fmt.Println("\nНеизвестная команда. Введите 'Команды' чтобы увидеть список доступных.")
		}
	}

}

func addItem(pokupki []Item) []Item {
	fmt.Print("\nЧто добавить?\nЕсли не хотите ничего добавлять напишите 'назад'")
	maxAttempts := 3
	for attempts := maxAttempts; attempts > 0; attempts-- {
		fmt.Print("\n> ")

		scanner.Scan()
		name := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if name == "назад" {
			fmt.Println("Возврат в меню")
			return pokupki
		}
		if name == "" {
			fmt.Printf("\nПустой ввод. Ничего не добавлено, осталось попыток: %d.\nЕсли не хотите ничего добавлять напишите 'назад'", attempts-1)
			continue
		}

		newItem := Item{
			ID:   nextID,
			Name: name,
			Done: false,
		}

		nextID++

		pokupki = append(pokupki, newItem)
		fmt.Printf("\nДобавлено: %s (ID: %d)\n\n", newItem.Name, newItem.ID)
		return pokupki
	}
	fmt.Println("Превышено количестов попыток. Возврат в меню")
	return pokupki
}

func changeItem(pokupki []Item) []Item {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст.")
		return pokupki
	}

	showList(pokupki)
	fmt.Println("Какой пункт вы хотите изменить?")
	fmt.Print("> ")

	var id int
	if _, err := fmt.Scanln(&id); err != nil {
		fmt.Println("Ошибка ввода")
		return pokupki
	}

	index := findIndexByID(pokupki, id)
	if index == -1 {
		fmt.Println("Такого элемента не существует")
		return pokupki
	}

	fmt.Println("На что хотите заменить?")
	fmt.Print("> ")

	scanner.Scan()
	newName := strings.TrimSpace(scanner.Text())

	if newName == "" {
		fmt.Println("Пустой ввод. Изменение отменено")
		return pokupki
	}

	oldName := pokupki[index].Name
	pokupki[index].Name = newName

	fmt.Printf("\nЭлемент изменен.\nВы заменили '%s', на '%s'.\n", oldName, newName)

	return pokupki
}

func deleteItem(pokupki []Item) []Item {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст")
		return pokupki
	}

	showList(pokupki)
	fmt.Println("Какой элемент вы хотите удалить")

	var id int
	if _, err := fmt.Scanln(&id); err != nil {
		fmt.Println("Ошибка ввода.")
		return pokupki
	}

	index := findIndexByID(pokupki, id)
	if index == -1 {
		fmt.Println("Такого элемента не существует")
		return pokupki
	}

	removedItem := pokupki[index]

	pokupki = append(pokupki[:index], pokupki[index+1:]...)

	for i := range pokupki {
		pokupki[i].ID = i + 1
	}

	fmt.Printf("Вы удалили %s\n", removedItem.Name)
	return pokupki

}
func showList(pokupki []Item) {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст, сначала добавьте что-нибудь")
		return
	}
	fmt.Println("\nВаш список продуктов:")
	for _, item := range pokupki {
		check := "[ ]"
		if item.Done {
			check = "[x]"
		}
		fmt.Printf("%d. %s %s\n", item.ID, item.Name, check)
	}
}

func findIndexByID(items []Item, id int) int {
	for i := range items {
		if items[i].ID == id {
			return i
		}
	}
	return -1
}

func toggleItem(pokupki []Item) []Item {
	if len(pokupki) == 0 {
		fmt.Println("Список пуст.")
		return pokupki
	}

	showList(pokupki)
	fmt.Println("\nВвведите номер элемента который хотите отметить/снять метку.\nЧтобы вернутся назад напишите'назад'")

	var idtest string
	fmt.Print("> ")

	if _, err := fmt.Scanln(&idtest); err != nil {
		fmt.Println("Ошибка ввода, возврат в меню")
		return pokupki
	}

	if strings.ToLower(strings.TrimSpace(idtest)) == "назад" {
		fmt.Println("Возврат в меню")
		return pokupki
	}

	id, _ := strconv.Atoi(idtest)

	index := findIndexByID(pokupki, id)

	if index == -1 {
		fmt.Println("Такого элемента не существует, возврат в меню")
		return pokupki
	}

	pokupki[index].Done = !pokupki[index].Done

	if pokupki[index].Done {
		fmt.Printf("'%s' отмечено как выполнено \n", pokupki[index].Name)
	} else {
		fmt.Printf("\nОтметка снята с '%s' \n", pokupki[index].Name)
	}
	return pokupki
}
