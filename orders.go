package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// Данные для подключения к БД
const (
	host     = "localhost"
	port     = 5481
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Формируем строку подключения
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открываем соединение с PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Ошибка пинга БД: %v", err)
	}

	fmt.Println("✅ Вы успешно подключены к базе данных!")

	// Ввод данных пользователя
	reader := bufio.NewReader(os.Stdin)

	// Ввод имени клиента
	fmt.Print("Введите имя клиента: ")
	customerName, _ := reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	// Ввод города (вывод списка)
	fmt.Println("Выберите город из списка:")
	cities := []string{
		"Moscow", "Vyksa", "Saint_Petersburg", "Krasnodar", "Ekaterinburg",
		"Krasnoyarsk", "Murom", "Samara", "Tula", "Tver",
	}
	for _, city := range cities {
		fmt.Println(city)
	}
	fmt.Print("Введите город: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	// Ввод номера телефона
	fmt.Print("Введите номер телефона: ")
	phoneNumber, _ := reader.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	// Ввод email
	fmt.Print("Введите email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Ввод артикула товара
	fmt.Print("Введите артикул товара: ")
	articleNumber, _ := reader.ReadString('\n')
	articleNumber = strings.TrimSpace(articleNumber)

	// SQL-запрос на вставку данных
	sqlStatement := `
  INSERT INTO orders (customer_name, city, phone_number, email, article_number) 
  VALUES ($1, $2, $3, $4, $5) RETURNING id;
 `
	var orderID int

	// Выполняем SQL-запрос
	err = db.QueryRow(sqlStatement, customerName, city, phoneNumber, email, articleNumber).Scan(&orderID)
	if err != nil {
		log.Fatalf("❌ Ошибка при создании заказа: %v", err)
	}

	// Выводим результат
	fmt.Printf("✅ Заказ успешно создан! ID заказа: %d\n", orderID)
}
