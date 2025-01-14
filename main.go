package main /* обязательный пакет который сообщает
компилятору что он должен компилироваться. */

import (
	"fmt" // пакеты для работы кода.
	"net/http"
)

type User struct {
	name  string // структура указанных данных которые можно изменять
	age   uint16
	money int16
}

func (u User) Userinfo() string { // u User параметр со строкой, userinfo() это метод
	return fmt.Sprintf("Пользовательское имя: %s. Возраст: %d. "+
		"Монеты: %d", u.name, u.age, u.money) // возвращаем формат строки и подстравиваем в проценты параметр
}

func home_page(z http.ResponseWriter, r *http.Request) {
	citizen := User{"user", 25, 50}    // Имя - возраст - монеты
	fmt.Fprintf(z, citizen.Userinfo()) // запуск метода Userinfo(), который
}

func contacts_page(z http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(z, "Контакт") // То что выводит на страницу, пока что неактивна.
}

func handleRec() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page) // Обработка главной страницы и адрес ссылки.
	http.ListenAndServe(":80", nil)              // Запуск локального сервера порт 80 /localhost
}

func main() {

	handleRec() // запуск функции handleRec которая отвечает за хост и обработку.

}

// запуск через go run main.go cmd или vs code (терминал), зайти на локальный хост.
