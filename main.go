package main

import (
	"fmt"
	"my_project/model"
	"os"
	"strconv"
)

func main() {

	participants := []model.Participant{
		{ID: "", Nama: "Saeful", Alamat: "Jl. Merdeka No.1", Pekerjaan: "Software Engineer", Alasan: "Ingin mempelajari Golang"},
		{ID: "", Nama: "Adilla", Alamat: "Jl. Sudirman No.2", Pekerjaan: "Data Scientist", Alasan: "Tertarik dengan performa Golang"},
		{ID: "", Nama: "Afiza", Alamat: "Jl. Thamrin No.3", Pekerjaan: "Backend Developer", Alasan: "Ingin memperdalam backend dengan Golang"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Silakan masukkan nama sebagai argumen.")
		return
	}

	name := os.Args[1]

	for i := range participants {
		if participants[i].Nama == name {
			fmt.Printf("ID : "+strconv.Itoa(i)+"%s\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n",
				participants[i].ID, participants[i].Nama, participants[i].Alamat, participants[i].Pekerjaan, participants[i].Alasan)
			return
		}
	}

	fmt.Printf("Data untuk nama %s tidak ditemukan.\n", name)
}
