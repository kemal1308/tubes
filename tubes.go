package main
import "fmt"

const NMAX int = 100

type dataP [NMAX]Pengeluaran

type Pengeluaran struct {
	Kategori string
	Jumlah   float64
	MetodeBayar string
}

func main() {
	var kotaTujuan string
	var durasiPerjalanan int
	var totalBudget float64
	var daftarPengeluaran dataP
	var jumlahPengeluaran int
	var pilihan int
	var keluar bool

	// menginput tujuan
	fmt.Print("Pilih kota tujuan traveling: ")
	fmt.Scan(&kotaTujuan)
	fmt.Print("Berapa hari Anda akan traveling? ")
	fmt.Scan(&durasiPerjalanan)
	fmt.Print("Masukkan budget total Anda : Rp.")
	fmt.Scan(&totalBudget)

	fmt.Println("-------------------------")
	fmt.Println("\tGoBudget\t\t ")

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
				var metodeBayar string

				fmt.Print("Masukkan kategori pengeluaran: ")
				fmt.Scan(&kategori)

				fmt.Print("Masukkan jumlah pengeluaran : Rp.")
				fmt.Scan(&jumlahPerHari)

				fmt.Print("Metode pembayaran (tunai / non_tunai): ")
				fmt.Scan(&metodeBayar)

				var pengeluaranBaru Pengeluaran
				pengeluaranBaru.Kategori = kategori
				pengeluaranBaru.Jumlah = jumlahPerHari
				pengeluaranBaru.MetodeBayar = metodeBayar

				daftarPengeluaran[jumlahPengeluaran] = pengeluaranBaru
				jumlahPengeluaran++

				fmt.Println("Pengeluaran berhasil ditambahkan.")
				fmt.Println()
    }

		case 2:
			tampilkanPengeluaran(daftarPengeluaran, jumlahPengeluaran)
		case 3:
			hitungTotalPengeluaran(daftarPengeluaran, jumlahPengeluaran)
		case 4:
			cekBudgetTersisa(totalBudget, daftarPengeluaran, jumlahPengeluaran)
		case 5:
			menuCariPengeluaran(daftarPengeluaran, jumlahPengeluaran)
		case 6:
			totalPengeluaranPerMetode(daftarPengeluaran, jumlahPengeluaran)
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi GoBudget")
			keluar = true
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}
}

// menu
func tampilkanMenu() int {
	var pilihan int
	fmt.Println()
	fmt.Println("--------- MENU ----------")
	fmt.Println()
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Riwayat Pengeluaran")
	fmt.Println("3. Hitung Total Pengeluaran")
	fmt.Println("4. Cek Budget Tersisa")
	fmt.Println("5. Cari Pengeluaran")
	fmt.Println("6. Total Pengeluaran Berdasarkan Metode Pembayaran") 
	fmt.Println("7. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan
}

// riwayat pengeluaran
func tampilkanPengeluaran(daftarPengeluaran dataP, jumlah int) {
	fmt.Println("\nRiwayat Pengeluaran:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("- %s: Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
	}
	fmt.Println()
}

// total pengeluaran
func hitungTotalPengeluaran(daftarPengeluaran dataP, jumlah int) {
	var totalPengeluaran float64
	for i := 0; i < jumlah; i++ {
		totalPengeluaran += daftarPengeluaran[i].Jumlah
	}
	fmt.Printf("Total Pengeluaran: Rp%.2f\n", totalPengeluaran)
	fmt.Println()
}

// budget tersisa
func cekBudgetTersisa(totalBudget float64, daftarPengeluaran dataP, jumlah int) {
	var totalPengeluaran float64
	for i := 0; i < jumlah; i++ {
		totalPengeluaran += daftarPengeluaran[i].Jumlah
	}
	budgetTersisa := totalBudget - totalPengeluaran
	fmt.Printf("Budget Tersisa: Rp%.2f\n", budgetTersisa)
	fmt.Println()
}

func menuCariPengeluaran (daftarPengeluaran dataP, jumlah int){
	var pilihan int
	fmt.Println("\n1. Cari berdasarkan kategori harga")
	fmt.Println("2. Cari Pengeluaran dengan harga > 100.000")
	fmt.Println("3. Cari Pengeluaran dengan harga  > 500.000")
	fmt.Println("4. Cari Pengeluaran dengan harga  > 1.000.000")
	fmt.Print("Mau pilih yang mana? ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
		case 1:
			cariPengeluaranKategori(daftarPengeluaran, jumlah)
		case 2 :
			cariPengeluaranHarga(daftarPengeluaran, jumlah, pilihan)
		case 3 :
			cariPengeluaranHarga(daftarPengeluaran, jumlah, pilihan)
		case 4 :
			cariPengeluaranHarga(daftarPengeluaran, jumlah, pilihan)
		default : 
			fmt.Println("Pilihan tidak ada")
	}
}
// cari pengeluaran berdasarkan kategori
func cariPengeluaranKategori(daftarPengeluaran dataP, jumlah int) {
	var kategori string
	fmt.Print("Masukkan kategori pengeluaran yang dicari: ")
	fmt.Scan(&kategori)
	for i := 0; i < jumlah; i++ {
		if daftarPengeluaran[i].Kategori == kategori {
			fmt.Printf("Pengeluaran ditemukan: %s, Rp%.2f\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah)
			fmt.Println()
			return
		}
	}
	fmt.Println("Pengeluaran tidak ditemukan.")
	fmt.Println()
}

// mencari pengeluaran berdasarkan harga
func cariPengeluaranHarga(daftarPengeluaran dataP, jumlah int, pilihan int) {
	var batas float64
	if pilihan == 2 {
		batas = 100000
	} else if pilihan == 3 {
		batas = 500000
	} else {
		batas = 1000000
	}

	fmt.Printf("\nPengeluaran dengan harga > Rp%.2f:\n", batas)
	for i := 0; i < jumlah; i++ {
		if daftarPengeluaran[i].Jumlah > batas {
			fmt.Printf("- %s: Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
		}
	}
	fmt.Println()
}
// pengeluara tiap metode bayar
func totalPengeluaranPerMetode(daftarPengeluaran dataP, jumlah int) {
	var totalTunai, totalNonTunai float64
	for i := 0; i < jumlah; i++ {
		if daftarPengeluaran[i].MetodeBayar == "tunai" {
			totalTunai += daftarPengeluaran[i].Jumlah
		} else if daftarPengeluaran[i].MetodeBayar == "non_tunai" {
			totalNonTunai += daftarPengeluaran[i].Jumlah
		}
	}
	fmt.Printf("\nTotal Pengeluaran Tunai    : Rp%.2f\n", totalTunai)
	fmt.Printf("Total Pengeluaran Non Tunai: Rp%.2f\n\n", totalNonTunai)
}

