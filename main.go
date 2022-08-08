package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {

	var startNode []node

	for {

		printData(startNode)

		fmt.Println("1 - Добавить запись в книжку")
		fmt.Println("2 - Удалить запись из книжки")
		fmt.Println("3 - Сохранить в файл")
		fmt.Println("4 - Загрузить из файла")
		fmt.Println("5 - Выход")

		var readKey int
		fmt.Fscanln(os.Stdin, &readKey)

		switch readKey {

		case 1:

			startNode = append(startNode, *new(node))
			startNode[len(startNode)-1].Number = len(startNode) - 1

			add(&startNode[len(startNode)-1])

		case 2:

			startNode = deleteNode(startNode)

		case 3:

			save(startNode)
			readKey = 10 //костыль, без него надпись введите имя файла появляется снова и снова

		case 4:
			startNode = load(startNode)
			readKey = 10

		case 5:

			os.Exit(0)

		default:

			fmt.Print("Вы ввели неверное значение\n\n")

		}
	}
}

type node struct {
	Number      int
	Name        string
	Lastname    string
	PhoneNumber string
} //Создание структуры 1 ноды

func printData(startNode []node) {

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("|Номер|                 Имя|             Фамилия|Номер телефона|")
	fmt.Println("---------------------------------------------------------------")

	for _, i := range startNode {

		fmt.Printf("|%5d|%20s|%20s|%14s|\n", i.Number, i.Name, i.Lastname, i.PhoneNumber)
		fmt.Println("---------------------------------------------------------------")

	}

}

func add(Node *node) {

	var name string
	var lastname string
	var phone string

	fmt.Println("Введите имя")
	fmt.Fscanln(os.Stdin, &name)

	fmt.Println("Введите фамилию")
	fmt.Fscanln(os.Stdin, &lastname)

	fmt.Println("Введите номер телефона")
	fmt.Fscanln(os.Stdin, &phone)

	Node.Name = name
	Node.Lastname = lastname
	Node.PhoneNumber = phone

}

func deleteNode(startNode []node) []node {

	var deleteNum int

	fmt.Println("Введите номер записи для удаления")
	fmt.Fscanln(os.Stdin, &deleteNum)

	for num, _ := range startNode {

		if num > deleteNum {

			tmp := &startNode[num]
			tmp.Number = tmp.Number - 1

		}
	}

	return append(startNode[:deleteNum], startNode[deleteNum+1:]...)

}

func save(startNode []node) {

	var fileName string

	fmt.Println("Введите имя файла")
	fmt.Fscan(os.Stdin, &fileName)

	fileName += ".json"
	file, _ := os.Create(fileName)

	jsonData, _ := json.Marshal(startNode)
	file.Write(jsonData)

	file.Close()

}

func load(startNode []node) []node {

	var fileName string

	fmt.Println("Введите имя файла из которого считать: ")
	fmt.Fscan(os.Stdin, &fileName)

	file, err := os.Open(fileName)

	if err != nil {

		fmt.Println("Ошибка при открытии файла ", fileName)
		os.Exit(1)

	}

	var buf = make([]byte, 64)
	var tmp []byte

	for {

		num, err := file.Read(buf)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		tmp = append(tmp, buf[:num]...)
	}

	err = json.Unmarshal(tmp, &startNode)

	if err != nil {
		panic(err)
	}

	return startNode
}
