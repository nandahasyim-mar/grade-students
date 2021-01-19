package main

import "fmt"

func gradeStudents(nilai int) (pesan string) {
	if nilai <= 34 {
		pesan = "E"
	} else if nilai <= 49 {
		pesan = "D"
	} else if nilai <= 64 {
		pesan = "C"
	} else if nilai <= 79 {
		pesan = "B"
	} else if nilai <= 100 {
		pesan = "A"
	}

	return
}

func main() {
	fmt.Println(gradeStudents(88))
}
