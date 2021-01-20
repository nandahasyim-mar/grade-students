package main

import "fmt"

func graduates(nama string, nilai int, absen int) (pesan string) {
	if nilai > 70 && absen < 5 {
		pesan = "lulus"
	} else {
		pesan = "tidak lulus"
	}

	return
}

func skleton() string {
	pesan := ""

	word := "Javascript"
	second := " is"
	third := " awesome"
	fourth := " and"
	fifth := " I"
	sixth := " love"
	seventh := " it!"

	pesan = word + second + third + fourth + fifth + sixth + seventh

	return pesan

}

func main() {
	fmt.Println(graduates("Nanda Hasyim Marfianshar", 84, 3))
	fmt.Println(skleton())
}
