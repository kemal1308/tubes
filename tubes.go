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

		// Hitung total pengeluaran saat ini
		var totalPengeluaran float64
		for i := 0; i < jumlahPengeluaran; i++ {
			totalPengeluaran += daftarPengeluaran[i].Jumlah
		}

		sisaBudget := totalBudget - totalPengeluaran

		// Cek apakah jumlah melebihi total budget atau sisa budget
		if jumlahPerHari > totalBudget || jumlahPerHari > sisaBudget {
			fmt.Printf("Pengeluaran melebihi budget.")
		} else {
			fmt.Println("Metode pembayaran: ")
			fmt.Println("1. Tunai")
			fmt.Println("2. Non Tunai")
			var opsiMetode int
			fmt.Print("Pilih: ")
			fmt.Scan(&opsiMetode)

			if opsiMetode == 1 {
				metodeBayar = "Tunai"
			} else if opsiMetode == 2 {
				metodeBayar = "Non Tunai"
			} else {
				fmt.Println("Opsi tidak ada")
				metodeBayar = "Tidak Valid"
			}

			if metodeBayar != "Tidak Valid" {
				var pengeluaranBaru Pengeluaran
				pengeluaranBaru.Kategori = kategori
				pengeluaranBaru.Jumlah = jumlahPerHari
				pengeluaranBaru.MetodeBayar = metodeBayar

				daftarPengeluaran[jumlahPengeluaran] = pengeluaranBaru
				jumlahPengeluaran++
				fmt.Println("Pengeluaran berhasil ditambahkan.")
			}
		}
		fmt.Println()
    }

		case 2:
			menuRiwayatPengeluaran(daftarPengeluaran, jumlahPengeluaran)
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
func menuRiwayatPengeluaran (daftarPengeluaran dataP, jumlah int){
	var pilihan int
	fmt.Println("\n1. Urutkan pengeluaran dari harga terbesar : ")
	fmt.Println("2. Urutkan pengeluaran dari harga terkecil : ")
	fmt.Print("Mau pilih yang mana? ")
	fmt.Scan(&pilihan)
	fmt.Print()
	
	switch pilihan {
		case 1:
			urutkanPengeluaranMenurun(daftarPengeluaran, jumlah)
		case 2 :
			urutkanPengeluaranMenurun(daftarPengeluaran, jumlah)
		default : 
			fmt.Println("Pilihan tidak ada")
	}
}
//insertion sort untuk pengeluaran menurun
func urutkanPengeluaranMenurun(daftarPengeluaran dataP, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		for j := i + 1; j < jumlah; j++ {
			if daftarPengeluaran[i].Jumlah < daftarPengeluaran[j].Jumlah {
				daftarPengeluaran[i], daftarPengeluaran[j] = daftarPengeluaran[j], daftarPengeluaran[i]
			}
		}
	}
	for i := 0; i < jumlah; i++ {
		fmt.Printf("- %s: Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
	}
}
// insertion sort untuk pengeluaran menaik
func urutkanPengeluaranMenaik(daftarPengeluaran dataP, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		for j := i + 1; j < jumlah; j++ {
			if daftarPengeluaran[i].Jumlah > daftarPengeluaran[j].Jumlah {
				daftarPengeluaran[i], daftarPengeluaran[j] = daftarPengeluaran[j], daftarPengeluaran[i]
			}
		}
	}
	for i := 0; i < jumlah; i++ {
		fmt.Printf("- %s: Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
	}
}


// total pengeluaran
func hitungTotalPengeluaran(daftarPengeluaran dataP, jumlah int) {
	var totalPengeluaran float64
	var i int
	
	for i = 0; i < jumlah; i++ {
		totalPengeluaran += daftarPengeluaran[i].Jumlah
	}
	fmt.Printf("Total Pengeluaran: Rp%.2f\n", totalPengeluaran)
	fmt.Println()
}

// budget tersisa
func cekBudgetTersisa(totalBudget float64, daftarPengeluaran dataP, jumlah int) {
	var totalPengeluaran float64
	var i int
	var budgetTersisa float64
	
	for i = 0; i < jumlah; i++ {
		totalPengeluaran += daftarPengeluaran[i].Jumlah
	}
	budgetTersisa = totalBudget - totalPengeluaran
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
// mengurutkan pengeluaran berdasarkan kategori (Selection Sort)
func urutkanPengeluaranKategori(daftarPengeluaran dataP, jumlah int) {
	var i, minIdx, j int
    for i = 0; i < jumlah-1; i++ {
        minIdx = i
        for j = i + 1; j < jumlah; j++ {
            if daftarPengeluaran[j].Kategori < daftarPengeluaran[minIdx].Kategori {
                minIdx = j
            }
        }
        daftarPengeluaran[i], daftarPengeluaran[minIdx] = daftarPengeluaran[minIdx], daftarPengeluaran[i]
    }
}

// binary search untuk kategori pengeluaran
func cariPengeluaranKategori(daftarPengeluaran dataP, jumlah int) {
    var kategori string
    fmt.Print("Masukkan kategori pengeluaran yang dicari: ")
    fmt.Scan(&kategori)

    urutkanPengeluaranKategori(daftarPengeluaran, jumlah)
	
	var kiri, kanan, tengah int
	var ditemukan bool
    kiri = 0
	kanan = jumlah-1
    ditemukan = false

    fmt.Printf("\nPengeluaran pada kategori: %s\n", kategori)
    for kiri <= kanan {
        tengah = (kiri + kanan) / 2
        if daftarPengeluaran[tengah].Kategori == kategori {
            ditemukan = true
            // Mencari semua pengeluaran dengan kategori yang sama tanpa out of range
            for i := tengah; i >= 0 && daftarPengeluaran[i].Kategori == kategori; i-- {
                fmt.Printf("- Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
            }
            for i := tengah + 1; i < jumlah && daftarPengeluaran[i].Kategori == kategori; i++ {
                fmt.Printf("- Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
            }
            break
        } else if daftarPengeluaran[tengah].Kategori < kategori {
            kiri = tengah + 1
        } else {
            kanan = tengah - 1
        }
    }

    if !ditemukan {
        fmt.Println("Tidak ada pengeluaran pada kategori tersebut.")
    }
}

// mencari pengeluaran berdasarkan harga
func cariPengeluaranHarga(daftarPengeluaran dataP, jumlah int, pilihan int) {
	var batas float64
	var i int
	if pilihan == 2 {
		batas = 100000
	} else if pilihan == 3 {
		batas = 500000
	} else {
		batas = 1000000
	}

	fmt.Printf("\nPengeluaran dengan harga > Rp%.2f:\n", batas)
	for i = 0; i < jumlah; i++ {
		if daftarPengeluaran[i].Jumlah > batas {
			fmt.Printf("- %s: Rp%.2f (Pembayaran: %s)\n", daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah, daftarPengeluaran[i].MetodeBayar)
		}
	}
	fmt.Println()
}
// pengeluaran tiap metode bayar dengan sequential search
func totalPengeluaranPerMetode(daftarPengeluaran dataP, jumlah int) {
	var pilihan int
	fmt.Println("\nPilih metode pembayaran untuk total pengeluaran:")
	fmt.Println("1. Tunai")
	fmt.Println("2. Non Tunai")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	var total float64
	var i int
	
	for i = 0; i < jumlah; i++ {
		if (pilihan == 1 && daftarPengeluaran[i].MetodeBayar == "Tunai") ||
			(pilihan == 2 && daftarPengeluaran[i].MetodeBayar == "Non Tunai") {
			total += daftarPengeluaran[i].Jumlah
		}
	}

	if pilihan == 1 {
		fmt.Printf("Total Pengeluaran Tunai: Rp%.2f\n\n", total)
	} else if pilihan == 2 {
		fmt.Printf("Total Pengeluaran Non Tunai: Rp%.2f\n\n", total)
	}else if pilihan == 3{
		fmt.Print("Kembali ke menu utama")
	}else{
		fmt.Print("Opsi tidak ada")
	}
}

