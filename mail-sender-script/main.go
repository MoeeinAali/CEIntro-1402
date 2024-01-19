package main

import (
	"bufio"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("emails.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		SendEmail(scanner.Text())
	}
}

func SendEmail(to string) {
	content, _ := ioutil.ReadFile("index.html")
	str := string(content)
	m := gomail.NewMessage()
	m.SetHeader("From", "farzamprogrammer@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Salam!")
	m.SetBody("text/html", str)
	d := gomail.NewPlainDialer("smtp.gmail.com", 465, "farzamprogrammer", "clwipsklodmkxkkk")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Email was not sent: " + to)
		return
	}
	fmt.Println("Email was sent: " + to)
}
