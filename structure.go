package main

import "fmt"

type Animal interface {
	breath()
	walk()
}

type Loin struct {
	age int
}

func (l Loin) breath() {
	fmt.Println("Lion breathes!")
}

func (l Loin) walk() {
	fmt.Print("Lion walk!")
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type Group struct {
	Name        string `json:"name"`
	Admin       User   `json:"admin"`
	Users       []User `json:"users"`
	NumberOfPpl int    `json:"number_of_ppl"`
}

func main() {
	// user := User{
	// 	ID:        "1",
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Email:     "john.doe@example.com",
	// }
	// fmt.Println(user.FirstName)
	// user2 := User{}
	// fmt.Println(user2)
	// fmt.Println(user.ID)
	// fmt.Println(&user)
	// fmt.Println(&user.ID)
	var a Animal
	fmt.Println(a)
	a = Loin{age: 10}
	a.breath()
	a.walk()
}
