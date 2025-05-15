package main

import "fmt"

const NMAX int = 100

type dataP [NMAX]Pengeluaran

// Struct untuk menyimpan data pengeluaran
type Pengeluaran struct {
	Kategori string
	Jumlah   float64
}

func main() {
	var kotaTujuan string
	var durasiPerjalanan int
	var totalBudget float64
	var daftarPengeluaran dataP
	var jumlahPengeluaran int
	var pilihan int
	var keluar bool

	// Input kota tujuan dan durasi
	fmt.Print("Pilih kota tujuan traveling: ")
	fmt.Scan(&kotaTujuan)
	fmt.Print("Berapa hari Anda akan traveling? ")
	fmt.Scan(&durasiPerjalanan)
	fmt.Print("Masukkan budget total Anda : Rp.")
	fmt.Scan(&totalBudget)

	fmt.Println("-------------------------")
	fmt.Println("	GoBudget		 ")

	// Menu utama
	for !keluar {
		pilihan = tampilkanMenu()

		switch pilihan {
		case 1:
			if jumlahPengeluaran >= NMAX {
				fmt.Println("Kapasitas pengeluaran penuh.")
			} else {
				var kategori string
				var jumlahPerHari float64
				// Input kategori pengeluaran dan jumlah per hari
				fmt.Print("Masukkan kategori pengeluaran: ")
				fmt.Scan(&kategori)
				fmt.Print("Masukkan jumlah pengeluaran per hari : Rp.")
				fmt.Scan(&jumlahPerHari)
				totalPengeluaran := jumlahPerHari * float64(durasiPerjalanan)
				daftarPengeluaran[jumlahPengeluaran] = Pengeluaran{Kategori: kategori, Jumlah: totalPengeluaran}
				jumlahPengeluaran++
				fmt.Println("Pengeluaran berhasil ditambahkan.")
			}
		case 2:
			fmt.Println("\nRiwayat Pengeluaran:")
			tampilkanPengeluaran(daftarPengeluaran, jumlahPengeluaran)
		case 3:
			var totalPengeluaran float64
			for i := 0; i < jumlahPengeluaran; i++ {
				totalPengeluaran += daftarPengeluaran[i].Jumlah
			}
			budgetTersisa := totalBudget - totalPengeluaran
			fmt.Printf("Budget Tersisa: Rp%.2f\n", budgetTersisa)
		case 4:
			urutkanPengeluaran(&daftarPengeluaran, jumlahPengeluaran)
			var kategori string
			fmt.Print("Masukkan kategori pengeluaran yang dicari: ")
			fmt.Scan(&kategori)
			index := cariPengeluaran(daftarPengeluaran, jumlahPengeluaran, kategori)
			if index != -1 {
				fmt.Printf("Pengeluaran ditemukan: %s, Rp%.2f\n", daftarPengeluaran[index].Kategori, daftarPengeluaran[index].Jumlah)
			} else {
				fmt.Println("Pengeluaran tidak ditemukan.")
			}
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menampilkan menu dan mengembalikan pilihan pengguna
func tampilkanMenu() int {
	var pilihan int
	fmt.Println("--------- MENU ----------")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Riwayat Pengeluaran")
	fmt.Println("3. Cek Budget Tersisa")
	fmt.Println("4. Cari Pengeluaran ") // Binary Search dan Selection Sort
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Menampilkan riwayat pengeluaran
func tampilkanPengeluaran(daftarPengeluaran dataP, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("- %s: Rp%.2f\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah)
	}
}

// Selection Sort untuk mengurutkan pengeluaran berdasarkan kategori
func urutkanPengeluaran(daftarPengeluaran *dataP, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		min := i
		for j := i + 1; j < jumlah; j++ {
			if daftarPengeluaran[j].Kategori < daftarPengeluaran[min].Kategori {
				min = j
			}
		}
		daftarPengeluaran[i], daftarPengeluaran[min] = daftarPengeluaran[min], daftarPengeluaran[i]
	}
}

// Binary Search untuk mencari pengeluaran berdasarkan kategori
func cariPengeluaran(daftarPengeluaran dataP, jumlah int, kategori string) int {
	kiri, kanan := 0, jumlah-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarPengeluaran[tengah].Kategori == kategori {
			return tengah
		} else if daftarPengeluaran[tengah].Kategori < kategori {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}
