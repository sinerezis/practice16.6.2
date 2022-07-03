package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Волшебство рандома
func init() {
	rand.Seed(time.Now().Unix())
}

// Интерфейс, который должен реализовать клиент
type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

// Структура, описывающая клиента
type Client struct {
	Mu             sync.RWMutex
	AccountBalance int
}

// Функция-конструктор для нового экземпляра
// структуры Client
func NewClient() *Client {
	return &Client{}
}

// Функция автоматического пополнения
// баланса клиента
func (c *Client) TopUp() {

	// Инициализируем переменные
	var RandInt int
	var RandTime float64

	//Бесконечный цикл
	for {
		// Генерим рандомную сумму платежа(1-10)
		// и рандомное время сна(0.5-1.0 секунд)
		RandInt = rand.Intn(10)
		RandTime = ((rand.Float64() * 500) + 500)

		c.Mu.Lock()
		c.AccountBalance += RandInt
		c.Mu.Unlock()
		time.Sleep(time.Duration(RandTime) * time.Millisecond)

	}

}

// Функция автоматического снятие
// средств со счета клиента
// "Hello, Tinkoff Bank!"
func (c *Client) Withdraw() {
	// Инициализируем переменные
	var RandInt int
	var RandTime float64

	//Бесконечный цикл
	for {
		// Генерим рандомную сумму платежа(1-5)
		// и рандомное время сна(0.5-1.0 секунд)
		RandInt = rand.Intn(5)
		RandTime = ((rand.Float64() * 500) + 500)

		c.Mu.Lock()
		if c.AccountBalance-RandInt >= 0 {
			c.AccountBalance -= RandInt
		} else {
			fmt.Println("Снятие средств со счета невозможно - не подключен овердрафт")
		}
		c.Mu.Unlock()
		time.Sleep(time.Duration(RandTime) * time.Millisecond)
	}
}

// Функция пополняет баланс пользователя на
// amount денег
func (c *Client) Deposit(amount int) {
	c.Mu.Lock()
	c.AccountBalance += amount
	c.Mu.Unlock()

}

// Функция пытается снять с баланса пользователя
// amount денег, если баланс <amount -
// возвращает ошибку
func (c *Client) Withdrawal(amount int) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if c.AccountBalance-amount < 0 {
		return fmt.Errorf("Невозможно снять %d со счета, тк не подключен овердрафт", amount)
	}

	return nil
}

// Функция выводит на печать баланс
// поьзователя
func (c *Client) Balance() int {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	return c.AccountBalance
}

// Функция запускает горутины и читает ввод с консоли.
func (c *Client) Reader() {

	//10 горутин, пополняющих счет
	for i := 0; i < 10; i++ {
		go c.TopUp()
	}
	//5 горутин а-ля тинькоф
	for i := 0; i < 5; i++ {
		go c.Withdraw()
	}

	//Сканируем ввод
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		command := sc.Text()
		switch command {

		case "balance":
			fmt.Println(c.Balance())

		case "deposit":
			var amount int
			fmt.Println("Введите сумму депозита: ")
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println(err)
			}
			c.Deposit(amount)

		case "withdrawal":
			var amount int
			fmt.Println("Введите сумму, которую вы хотите снять: ")
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println(err)
			}
			err = c.Withdrawal(amount)
			if err != nil {
				fmt.Println(err)
			}

		case "exit":
			return

		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")

		}
	}

}

func main() {
	c := NewClient()

	c.Reader()

}
