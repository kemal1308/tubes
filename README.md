# GoBudget 🧳💸

**GoBudget** adalah aplikasi berbasis terminal yang dibangun dengan bahasa pemrograman **Go (Golang)**, dirancang khusus untuk membantu pengguna mengelola dan melacak pengeluaran selama perjalanan. Aplikasi ini memastikan agar anggaran tetap terkendali melalui pencatatan yang sistematis dan fitur analisis sederhana namun efektif.

## 📌 Fitur Utama

1. **Tambah Pengeluaran**
   - Mencatat transaksi pengeluaran baru dengan data kategori, jumlah, dan metode pembayaran.
   - Menggunakan **sequential search** untuk menghitung total pengeluaran sebelum menyimpan data baru, sehingga pengguna dapat mengetahui apakah masih berada dalam batas anggaran.

2. **Lihat Riwayat Pengeluaran**
   - Menampilkan seluruh daftar pengeluaran yang telah dicatat.
   - Dapat diurutkan berdasarkan jumlah secara **menaik atau menurun** menggunakan algoritma mirip **insertion sort**.

3. **Hitung Total Pengeluaran**
   - Memberikan ringkasan dari total seluruh pengeluaran yang tercatat.
   - Proses penghitungan dilakukan menggunakan **sequential search**.

4. **Cek Budget Tersisa**
   - Menghitung sisa anggaran dengan cara mengurangi total pengeluaran dari budget awal.
   - Proses dilakukan dengan **sequential search**.

5. **Cari Pengeluaran**
   - **Berdasarkan Kategori**: Daftar pengeluaran diurutkan dengan **selection sort**, kemudian dicari menggunakan **binary search**.
   - **Berdasarkan Batas Harga**: Filter pengeluaran di atas nilai tertentu menggunakan **sequential search**.

6. **Total Pengeluaran Berdasarkan Metode Pembayaran**
   - Menganalisis total pengeluaran berdasarkan metode pembayaran (Tunai/Non-Tunai).
   - Dihitung dengan **sequential search**.

7. **Keluar**
   - Menghentikan sesi aplikasi dengan aman dan memberikan pesan penutup.

## ⚙️ Teknologi

- **Bahasa**: Go (Golang)
- **Interface**: CLI (Command Line Interface)
- **Struktur Data**: Array
- **Algoritma Pendukung**:
  - Sequential Search
  - Binary Search
  - Insertion Sort
  - Selection Sort

## 🏗️ Arsitektur Sistem

Aplikasi ini adalah aplikasi berbasis terminal tanpa backend server atau frontend grafis. Semua data dikelola dalam memori menggunakan array, sehingga cocok untuk simulasi dan pembelajaran algoritma dasar dalam pengelolaan data berbasis CLI.

## 🚀 Cara Menjalankan

1. Clone repositori:
   ```bash
   git clone https://github.com/username/GoBudget.git
   cd GoBudget

2. jalankan program :
    ```bash
   go run main.go

## Tim Pengembang GoBudget
**Kelompok 11**
1. Dimas Tri Mahesa
2. Kemal Farouq At-Tirmidzi
