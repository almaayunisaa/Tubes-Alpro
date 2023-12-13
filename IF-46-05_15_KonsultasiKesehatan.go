package main

import "fmt"

const NMAX int = 200

type pasien struct {
	info [NMAX]dataPasien
	N    int
}

type dokter struct {
	info [NMAX]dataDokter
	N    int
}

type question struct {
	info [NMAX]dataTanya
	N    int
}

type dataPasien struct {
	usn, nama, sex string
	usia           int
	pass           string
}

type dataDokter struct {
	usn, nama, sp string
	usia          int
	pass          string
}

type dataTanya struct {
	tanya string
	jawab dataJawab
	tag   string
	N     int
}

type dataJawab struct {
	isi [NMAX]string
	N   int
}

func header() {
	fmt.Println("*** ------------------------------- ***")
	fmt.Println("---- Aplikasi Konsultasi Kesehatan ----")
	fmt.Println("-- Tugas Besar Algoritma Pemrograman --")
	fmt.Println("*** ------------------------------- ***")
}

func menu(dok *dokter, pas *pasien, q *question) { 
	var opsi, i int
	fmt.Println()
	fmt.Println("*Menu Utama*")
	fmt.Println("1. Halaman Dokter")
	fmt.Println("2. Halaman Pasien")
	fmt.Println("3. List Dokter")
	fmt.Println("4. Keluar")
	fmt.Print("Pilihan : ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		menudokter(dok, pas, q)
	} else if opsi == 2 {
		menupasien(pas, dok, q)
	} else if opsi==3 {
		for i=0; i<dok.N; i++ {
			fmt.Print(i+1, ". ", dok.info[i].nama, ", spesialis ", dok.info[i].sp, "\n")
		}
		menu(dok, pas, q)
	} else if opsi == 4 {
		keluar()
	} else {
		fmt.Println("Error, gunakan input yang benar")
		menu(dok, pas, q)
	}
}

func keluar() {
	fmt.Println("Anda telah keluar dari aplikasi.")
}

func menupasien(T *pasien, D *dokter, q *question) {
	var opsi int
	var yn string
	fmt.Println()
	fmt.Println("*Halaman Pasien*")
	fmt.Println("1. Pasien Tidak Terdaftar")
	fmt.Println("2. Pasien Terdaftar")
	fmt.Print("Pilihan : ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Apakah ingin mendaftar ? (iya/tidak)")
		fmt.Scan(&yn)
		if yn == "iya" {
			inputPasien(T, D, *q)
		} else if yn == "tidak" {
			forumTDKPasien(q, *D, *T)
		} else {
			fmt.Println("Input tidak valid.")
			menupasien(T, D, q)
		}
	} else if opsi == 2 {
		//cari data pasien (login)
		loginPasien(*T, *D, *q)
	} else {
		fmt.Println("Error, gunakan input yang benar")
		menu(D, T, q)
	}
}

func menudokter(T *dokter, A *pasien, q *question) {
	//input data dokter (login)
	var opsi int
	var yn string
	fmt.Println()
	fmt.Println("*Halaman Dokter*")
	fmt.Println("1. Dokter Tidak Terdaftar")
	fmt.Println("2. Dokter Terdaftar")
	fmt.Print("Pilihan : ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Apakah ingin mendaftar ? (iya/tidak)")
		fmt.Scan(&yn)
		if yn == "iya" {
			inputDokter(T, A, *q)
		} else if yn == "tidak" {
			menu(T, A, q)
		} else {
			fmt.Println("Input tidak valid.")
			menudokter(T, A, q)
		}
	} else if opsi == 2 {
		//cari data pasien (login)
		loginDokter(*T, *A, *q)
	} else {
		fmt.Println("Error, gunakan input yang benar")
		menu(T, A, q)
	}
}

func inputPasien(T *pasien, D *dokter, q question) {
	var username, nama, opsi string
	fmt.Println()
	fmt.Println("*Halaman Sign In*")
	fmt.Println("Gunakan underscore (_) untuk pengganti spasi, Gunakan kapital pada huruf awal nama (Co: Azizah)")
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	for nama[0]>='a' && nama[0]<='z' {
		fmt.Println("Gunakan Kapital, atau tekan 1 untuk mengulang")
		fmt.Print("Nama : ")
		fmt.Scan(&nama)
		if nama=="1" {
			menu(D, T, &q)
		}
	}
	T.info[T.N].nama=nama
	fmt.Print("Username : ")
	fmt.Scan(&username)
	for adausnP(*T, username)!=-1 {
		fmt.Println("Username tersebut telah terpakai, silahkan input kembali atau tekan 1 jika ingin keluar.")
		fmt.Print("Username : ")
		fmt.Scan(&username)
		if username=="1" {
			menu(D, T, &q)
		}
	}
	T.info[T.N].usn=username
	fmt.Println("1. Perempuan")
	fmt.Println("2. Laki Laki")
	fmt.Print("Jenis Kelamin : ")
	fmt.Scan(&opsi)
	if opsi == "1" {
		T.info[T.N].sex="Perempuan"
	} else if opsi == "2" {
		T.info[T.N].sex="Laki Laki"
	}
	fmt.Print("Usia : ")
	fmt.Scan(&T.info[T.N].usia)
	fmt.Print("Password : ")
	fmt.Scan(&T.info[T.N].pass)
	fmt.Println("Selamat kamu telah terdaftar!")
	T.N = T.N + 1
	urutPasien(T)
	loginPasien(*T, *D, q)
}

func adausnP(T pasien, q string) int {
	var found int = -1
	var i int
	for i < T.N && found == -1 {
		if T.info[i].usn == q {
			found = i
		}
		i = i + 1
	}
	return found
}

func adausnD(T dokter, q string) int {
	var found int = -1
	var i int
	for i < T.N && found == -1 {
		if T.info[i].usn == q {
			found = i
		}
		i = i + 1
	}
	return found
}

func inputDokter(T *dokter, A *pasien, q question) {
	var username, nama string
	var kode string
	fmt.Println()
	fmt.Println("*Halaman Sign In*")
	fmt.Println("Gunakan underscore (_) untuk pengganti spasi, Gunakan kapital pada huruf awal nama (Co: Azizah)")
	fmt.Print("Kode Pendaftaran : ")
	fmt.Scan(&kode)
	for kode!="100" && kode!="122" && kode!="500" && kode!="111" && kode!="145" {
		fmt.Println("Kode salah, silahkan input kembali atau tekan 1 jika ingin keluar.")
		fmt.Print("Kode Pendaftaran : ")
		fmt.Scan(&kode)
		if kode=="1" {
			menu(T, A, &q)
		}
	}
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	for nama[0]>='a' && nama[0]<='z' {
		fmt.Println("Gunakan Kapital, atau tekan 1 untuk mengulang")
		fmt.Print("Nama : ")
		fmt.Scan(&nama)
		if nama=="1" {
			menu(T, A, &q)
		}
	}
	T.info[T.N].nama=nama
	fmt.Print("Username : ")
	fmt.Scan(&username)
	for adausnD(*T, username)!=-1 {
		fmt.Println("Username tersebut telah terpakai, silahkan input kembali atau tekan 1 jika ingin keluar.")
		fmt.Print("Username : ")
		fmt.Scan(&username)
		if username=="1" {
			menu(T, A, &q)
		}
	}
	T.info[T.N].usn=username
	fmt.Print("Spesialis : ")
	fmt.Scan(&T.info[T.N].sp)
	fmt.Print("Usia : ")
	fmt.Scan(&T.info[T.N].usia)
	fmt.Print("Password : ")
	fmt.Scan(&T.info[T.N].pass)
	fmt.Println("Selamat kamu telah terdaftar!")
	T.N = T.N + 1
	urutDokter(T)
	loginDokter(*T, *A, q)
}

func loginPasien(T pasien, D dokter, q question) {
	fmt.Println()
	fmt.Println("*Halaman Log In Pasien*")
	var usr, psw string
	var found bool = false
	var pilihan, i int
	fmt.Print("Username : ")
	fmt.Scan(&usr)
	fmt.Print("Password : ")
	fmt.Scan(&psw)
	for i < T.N && !found {
		if usr == T.info[i].usn && psw == T.info[i].pass {
			found = true
		}
		i = i + 1
	}
	if found == true {
		fmt.Println("Login berhasil!")
		forumPasien(&q, D, T)
	} else {
		fmt.Println("Maaf, password atau username yang anda masukkan salah.")
		fmt.Println("1. Menu")
		fmt.Println("2. Login Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			menu(&D, &T, &q)
		} else if pilihan == 2 {
			loginPasien(T, D, q)
		}
	}
}
func loginDokter(T dokter, P pasien, q question) { 
	var usr, psw string
	var found bool = false
	var pilihan, i int
	fmt.Println()
	fmt.Println("*Halaman Log In Dokter*")
	fmt.Print("Username : ")
	fmt.Scan(&usr)
	fmt.Print("Password : ")
	fmt.Scan(&psw)
	for i < T.N && !found {
		if usr == T.info[i].usn && psw == T.info[i].pass {
			found = true
		}
		i = i + 1
	}
	if found == true {
		fmt.Println("Login berhasil!")
		forumDokter(&q, T, P)
	} else {
		fmt.Println("Maaf, password atau username yang anda masukkan salah.")
		fmt.Println("1. Menu")
		fmt.Println("2. Login Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			menu(&T, &P, &q)
		} else if pilihan == 2 {
			loginDokter(T, P, q)
		}
	}
}

func forumDokter(T *question, D dokter, P pasien) {
	var pilihan, i int
	fmt.Println()
	fmt.Println("*Forum Dokter*")
	fmt.Println("Gunakan underscore (_) untuk pengganti spasi")
	fmt.Println("1. Cari Pertanyaan")
	fmt.Println("2. Tampilkan Semua Pertanyaan")
	fmt.Println("3. Tampilkan Tag")
	fmt.Println("4. Tampilkan Nama Pasien")
	fmt.Println("5. Log Out")
	fmt.Println("Pilihan : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		cariDokter(T, D, P)
	} else if pilihan == 2 {
		fmt.Println("*List Pertanyaan*")
		for i = 0; i < T.N; i++ {
			fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
		}
		if T.N==0 {
			fmt.Println("Tidak ada pertanyaan")
		}
		forumDokter(T, D, P)
	} else if pilihan==3 {
		fmt.Println("*Tag dari List Pertanyaan(Mungkin terdapat duplikat)*")
		urutTag(*T)
		forumDokter(T, D, P)
	} else if pilihan == 4 {
		fmt.Println("*List Pasien*")
		for i=0; i<P.N; i++ {
			fmt.Print(i+1, ". ", P.info[i].nama, "\n")
		}
		forumDokter(T, D, P)
	} else if pilihan == 5 {
		menu(&D, &P, T)
	} else {
		fmt.Println("Error, gunakan input yang benar")
		forumDokter(T , D , P)
	}
}

func forumPasien(T *question, D dokter, P pasien) {
	var opsi, i int
	fmt.Println()
	fmt.Println("*Forum Pasien*")
	fmt.Println("Gunakan underscore (_) untuk pengganti spasi")
	fmt.Println("1. Bertanya")
	fmt.Println("2. Mencari Pertanyaan")
	fmt.Println("3. Tampilkan Semua Pertanyaan")
	fmt.Println("4. Tampilkan Tag")
	fmt.Println("5. Log Out")
	fmt.Print("Pilihan : ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		inputQ(T, D, P)
	} else if opsi == 2 {
		cariPasien(T, D, P)
	} else if opsi == 3 {
		fmt.Println("*List Pertanyaan*")
		for i = 0; i < T.N; i++ {
			fmt.Print(i+1,". ",  T.info[i].tanya, "\n")
		}
		if T.N==0 {
			fmt.Println("Tidak ada pertanyaan")
		}
		forumPasien(T, D, P)
	} else if opsi==4 {
		fmt.Println("*Tag dari List Pertanyaan(Mungkin terdapat duplikat)*")
		urutTag(*T)
		forumPasien(T, D, P)
	} else if opsi == 5 {
		menu(&D, &P, T)
	} else {
		forumPasien(T , D , P)
	}
}

func forumTDKPasien(T *question, D dokter, P pasien) {
	cariTDK(*T, D, P)
}

func adaq(T question, q string) int {
	var found int = -1
	var i int
	for i < T.N && found == -1 {
		if T.info[i].tanya == q {
			found = i
		}
		i = i + 1
	}
	return found
}

func inputQ(T *question, D dokter, P pasien) {
	var q, hashtag string
	fmt.Println()
	fmt.Println("*Halaman Input Pertanyaan*")
	fmt.Println("Perhatian! Apabila menginput pertanyaan yang menggunakan spasi, gunakanlah underscore(_)")
	fmt.Println("contoh: Apa_itu_tubercolosis?")
	fmt.Print("Pertanyaan : ")
	fmt.Scan(&q)
	fmt.Print("Hashtag : ")
	fmt.Print("#")
	fmt.Scan(&hashtag)
	if adaq(*T, q) != -1 {
		T.info[adaq(*T, q)].N = T.info[adaq(*T, q)].N + 1
	} else {
		T.info[T.N].tanya = q
		T.info[T.N].tag = hashtag
		T.info[T.N].N = T.info[T.N].N + 1
		T.N = T.N + 1
	}
	sortJmlTanya(T)
	forumPasien(T, D, P)
}

func cariDokter(T *question, D dokter, P pasien) {
	var cari, dasar, jawab, nama string
	var carinmr, N, pilihan int
	var found int = -1
	fmt.Println()
	fmt.Println("*Halaman Pencarian*")
	fmt.Print("Cari berdasarkan? (pertanyaan/tag/nomor/JumlahPertanyaan)", "\n")
	fmt.Print("Input : ")
	fmt.Scan(&dasar)
	if dasar == "pertanyaan" {
		fmt.Print("Input pertanyaan : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tanya {
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
				found = i
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ada.")
		}
	} else if dasar == "tag" {
		fmt.Print("Input tag : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tag {
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
				found = -999
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ada.")
		}
	} else if dasar == "nomor" {
		fmt.Print("Input nomor : ")
		fmt.Scan(&carinmr)
		fmt.Println("*Pertanyaan*")
		if cariBinary(*T, carinmr-1) != -1 {
			fmt.Print(carinmr, ". ", T.info[carinmr-1].tanya, "\n")
			for j := 0; j < T.info[carinmr-1].jawab.N; j++ {
				fmt.Println("--", T.info[carinmr-1].jawab.isi[j])
			}
		} else {
			fmt.Println("Pertanyaan tidak ada.")
		}
		found = carinmr-1
	} else if dasar == "JumlahPertanyaan" {
		fmt.Print("Input Jumlah Pertanyaan : ")
		fmt.Scan(&N)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if N == T.info[i].N {
				found=-999
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ada.")
		}
	} else {
		fmt.Println("Error, gunakan input yang benar")
		cariDokter(T , D , P)
	}
	if found != -999 && found != -1 {
		fmt.Println("1. Jawab")
		fmt.Println("2. Lanjut")
		fmt.Println("3. Kembali ke Forum")
		fmt.Println("4. Hapus Pertanyaan")
		fmt.Println("5. Edit")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Print("Masukan jawaban : ")
			fmt.Scan(&jawab)
			fmt.Print("Oleh : ")
			fmt.Scan(&nama)
			T.info[found].jawab.isi[T.info[found].jawab.N] = jawab + "<Dr." + nama + ">"
			T.info[found].jawab.N = T.info[found].jawab.N + 1
			fmt.Println("1. Forum")
			fmt.Println("2. LogOut")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				forumDokter(T, D, P)
			} else if pilihan == 2 {
				menu(&D, &P, T)
			}
		} else if pilihan == 2 {
			cariDokter(T, D, P)
		} else if pilihan == 3 {
			forumDokter(T, D, P)
		} else if pilihan == 4 {
			delQ(T, found)
			forumDokter(T, D, P)
		} else if pilihan == 5{
			edit(T, found)
			forumDokter(T, D, P)
		} else {
			fmt.Println("Error, gunakan input yang benar")
			cariDokter(T , D , P)
		}
	} else if found == -999 {
		fmt.Println("1. Jawab")
		fmt.Println("2. Lanjut")
		fmt.Println("3. Kembali ke Forum")
		fmt.Println("4. Hapus Pertanyaan")
		fmt.Println("5. Edit")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Pertanyaan nomor berapa yang ingin dijawab?")
			fmt.Scan(&pilihan)

			fmt.Print("Masukan jawaban : ")
			fmt.Scan(&jawab)
			fmt.Print("Oleh : ")
			fmt.Scan(&nama)
			T.info[pilihan-1].jawab.isi[T.info[pilihan-1].jawab.N] = jawab + "<Dr." + nama + ">"
			T.info[pilihan-1].jawab.N = T.info[pilihan-1].jawab.N + 1
			fmt.Println("1. Forum")
			fmt.Println("2. LogOut")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				forumDokter(T, D, P)
			} else if pilihan == 2 {
				menu(&D, &P, T)
			}
		} else if pilihan == 2 {
			cariDokter(T, D, P)
		} else if pilihan == 3 {
			forumDokter(T, D, P)
		} else if pilihan == 4 {
			fmt.Println("Pertanyaan nomor berapa yang ingin dihapus?")
			fmt.Scan(&pilihan)
			delQ(T, pilihan-1)
			forumDokter(T, D, P)
		} else if pilihan == 5 {
			fmt.Println("Pertanyaan nomor berapa yang ingin diedit?")
			fmt.Scan(&pilihan)
			edit(T, pilihan-1)
			forumDokter(T, D, P)
		} else {
			fmt.Println("Error, gunakan input yang benar")
			cariDokter(T , D , P)
		}
	} else if found == -1 {
		forumDokter(T, D, P)
	}
}

func cariPasien(T *question, D dokter, P pasien) {
	var cari, dasar, jawab, nama string
	var carinmr, N, pilihan int
	var found int = -1
	fmt.Println()
	fmt.Println("*Halaman Pencarian*")
	fmt.Print("Cari berdasarkan? (pertanyaan/tag/nomor/JumlahPertanyaan)", "\n")
	fmt.Print("Input : ")
	fmt.Scan(&dasar)
	if dasar == "pertanyaan" {
		fmt.Print("Input pertanyaan : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tanya {
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
				found = i
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ditemukan.")
		}
	} else if dasar == "tag" {
		fmt.Print("Input tag : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tag {
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
				found = -999
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ditemukan.")
		}
	} else if dasar == "nomor" {
		fmt.Print("Input nomor : ")
		fmt.Scan(&carinmr)
		fmt.Println("*Pertanyaan*")
		if cariBinary(*T, carinmr-1) != -1 {
			fmt.Print(carinmr, ". ", T.info[carinmr-1].tanya, "\n")
			for j := 0; j < T.info[carinmr-1].jawab.N; j++ {
				fmt.Println("--", T.info[carinmr-1].jawab.isi[j])
			}
		} else {
			fmt.Println("Pertanyaan tidak ada.")
		}
		found = carinmr-1
	} else if dasar == "JumlahPertanyaan" {
		fmt.Print("Input Jumlah Pertanyaan : ")
		fmt.Scan(&N)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if N == T.info[i].N {
				found=-999
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ada.")
		}
	} else {
		fmt.Println("Error, gunakan input yang benar")
		cariPasien(T , D , P)
	}
	if found != -999 && found != -1 {
		fmt.Println("1. Jawab")
		fmt.Println("2. Lanjut")
		fmt.Println("3. Kembali ke Forum")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Print("Masukan jawaban : ")
			fmt.Scan(&jawab)
			fmt.Print("Oleh : ")
			fmt.Scan(&nama)
			T.info[found].jawab.isi[T.info[found].jawab.N] = jawab + "<" + nama + ">"
			T.info[found].jawab.N = T.info[found].jawab.N + 1
			fmt.Println("1. Forum")
			fmt.Println("2. LogOut")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				forumPasien(T, D, P)
			} else if pilihan == 2 {
				menu(&D, &P, T)
			}
		} else if pilihan == 2 {
			cariPasien(T, D, P)
		} else if pilihan == 3 {
			forumPasien(T, D, P)
		} else {
			fmt.Println("Error, gunakan input yang benar")
			cariPasien(T , D , P)
		}
	} else if found == -999 {
		fmt.Println("1. Jawab")
		fmt.Println("2. Lanjut")
		fmt.Println("3. Kembali ke Forum")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Pertanyaan nomor berapa yang ingin dijawab?")
			fmt.Scan(&pilihan)

			fmt.Print("Masukan jawaban : ")
			fmt.Scan(&jawab)
			fmt.Print("Oleh : ")
			fmt.Scan(&nama)
			T.info[pilihan-1].jawab.isi[T.info[pilihan-1].jawab.N] = jawab + "<" + nama + ">"
			T.info[pilihan-1].jawab.N = T.info[pilihan-1].jawab.N + 1
			fmt.Println("1. Forum")
			fmt.Println("2. LogOut")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				forumPasien(T, D, P)
			} else if pilihan == 2 {
				menu(&D, &P, T)
			}
		} else if pilihan == 2 {
			cariPasien(T, D, P)
		} else if pilihan == 3 {
			forumPasien(T, D, P)
		} else {
			fmt.Println("Error, gunakan input yang benar")
			cariPasien(T , D , P)
		}
	} else if found == -1 {
		forumPasien(T, D, P)
	}
}

func cariTDK(T question, D dokter, P pasien) {
	var cari, dasar string
	var carinmr, N, pilihan int
	var found int = -1
	fmt.Println()
	fmt.Println("*Halaman Pencarian*")
	fmt.Print("Cari berdasarkan? (pertanyaan/tag/nomor/JumlahPertanyaan)", "\n")
	fmt.Print("Input : ")
	fmt.Scan(&dasar)
	if dasar == "pertanyaan" {
		fmt.Print("Input pertanyaan : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tanya {
				found = i
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ditemukan.")
		}
	} else if dasar == "tag" {
		fmt.Print("Input tag : ")
		fmt.Scan(&cari)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if cari == T.info[i].tag {
				found = i
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
			}
		}
		if found == -1 {
			fmt.Println("Pertanyaan tidak ditemukan.")
		}
	} else if dasar == "nomor" {
		fmt.Print("Input nomor : ")
		fmt.Scan(&carinmr)
		fmt.Println("*Pertanyaan*")
		if cariBinary(T, carinmr-1) != -1 {
			fmt.Print(carinmr, ". ", T.info[carinmr-1].tanya, "\n")
			for j := 0; j < T.info[carinmr-1].jawab.N; j++ {
				fmt.Println("--", T.info[carinmr-1].jawab.isi[j])
			}
		} else {
			fmt.Println("Pertanyaan tidak ada.")
		}
		found = carinmr-1
	} else if dasar == "JumlahPertanyaan" {
		fmt.Print("Input Jumlah Pertanyaan : ")
		fmt.Scan(&N)
		fmt.Println("*Pertanyaan*")
		for i := 0; i < T.N; i++ {
			if N == T.info[i].N {
				found=-999
				fmt.Print(i+1, ". ", T.info[i].tanya, "\n")
				for j := 0; j < T.info[i].jawab.N; j++ {
					fmt.Println("--", T.info[i].jawab.isi[j])
				}
			}
		}
		if found==-1 {
			fmt.Println("Pertanyaan tidak ditemukan")
		}
	} else {
		fmt.Println("Error, gunakan input yang benar")
		cariTDK(T , D , P)
	}
	fmt.Println("1. Lanjut")
	fmt.Println("2. LogOut")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		cariTDK(T, D, P)
	} else if pilihan == 2 {
		menu(&D, &P, &T)
	} else {
		fmt.Println("Error, gunakan input yang benar")
		cariTDK(T , D , P)
	}
}

func sortJmlTanya(T *question) {
	// SELECTION SORT ASC
	var i, j, idx_min int
	var temp dataTanya
	i = 1
	for i <= T.N-1 {
		idx_min = i - 1
		j = i
		for j < T.N {
			if T.info[idx_min].N < T.info[j].N {
				idx_min = j
			}
			j++
		}
		temp = T.info[idx_min]
		T.info[idx_min] = T.info[i-1]
		T.info[i-1] = temp
		i++
	}
}

func cariBinary(T question, x int) int {
	var kiri, kanan, med int
	var found int = -1
	kiri = 0
	kanan = T.N - 1
	for kiri <= kanan && found == -1 {
		med = (kiri + kanan) / 2
		if x < med {
			kanan = kanan - 1
		} else if x > med {
			kiri = kiri + 1
		} else {
			found = med
		}
	}
	return found
}

func delQ(T *question, found int) {
	for found <= T.N-2 {
		T.info[found] = T.info[found+1]
		found++
	}
	T.N=T.N-1
}

func urutDokter(T *dokter) { 
	var i, j, idx_min int
	var temp dataDokter
	i = 1
	for i <= T.N-1 {
		idx_min = i - 1
		j = i
		for j < T.N {
			if T.info[idx_min].nama > T.info[j].nama {
				idx_min = j
			}
			j++
		}
		temp = T.info[idx_min]
		T.info[idx_min] = T.info[i-1]
		T.info[i-1] = temp
		i++
	}
}

func urutPasien(T *pasien) { // insertion urut nama pasien asc
	var j, i int
	var temp dataPasien
	i = 1
	for i <= T.N-1 {
		j = i
		temp = T.info[j]
		for j > 0 && temp.nama < T.info[j-1].nama {
			T.info[j] = T.info[j-1]
			j--
		}
		T.info[j] = temp
		i++
	}
}

func urutTag(T question) { // insertion urut tag asc 
	var j, i int
	var temp dataTanya
	i = 1
	for i <= T.N-1 {
		j = i
		temp = T.info[j]
		for j > 0 && temp.tag > T.info[j-1].tag {
			T.info[j] = T.info[j-1]
			j--
		}
		T.info[j] = temp
		i++
	}
	for i=0; i<T.N; i++ {
		fmt.Print(i+1, ". ", T.info[i].tag, "\n")
	}
}

func edit(T *question, found int) {
	var pilihan int
	var input string
	fmt.Println("Apa yang ingin diedit?")
	fmt.Println("1. Pertanyaan")
	fmt.Println("2. Tag")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	if pilihan==1 {
		fmt.Print("Input pertanyaan : ")
		fmt.Scan(&input)
		T.info[found].tanya=input
	} else if pilihan == 2 {
		fmt.Print("Input tag : ")
		fmt.Scan(&input)
		T.info[found].tag=input
	}
}

func dummy(D *dokter, P *pasien, T *question) {
	D.info[0].usn="almaw"
	D.info[0].nama="Azizah"
	D.info[0].sp="Jantung"
	D.info[0].usia=19
	D.info[0].pass="almamaw"
	
	D.info[1].usn="pawah"
	D.info[1].nama="Farah"
	D.info[1].sp="Jantung"
	D.info[1].usia=19
	D.info[1].pass="pawawah"
	D.N=D.N+2
	
	P.info[0].usn="almapas"
	P.info[0].nama="Ayunisa"
	P.info[0].sex="Perempuan"
	P.info[0].usia=50
	P.info[0].pass="almamenjadidokter"
	
	P.info[1].usn="pawahpas"
	P.info[1].nama="Aurellia"
	P.info[1].sex="Perempuan"
	P.info[1].usia=50
	P.info[1].pass="pawahmenjadidokter"

	P.info[2].usn="zizah"
	P.info[2].nama="Zizah"
	P.info[2].sex="Perempuan"
	P.info[2].usia=50
	P.info[2].pass="zizahmenjadidokter"
	P.N=P.N+3
	
	T.info[0].tanya="Apa_itu_maag?"
	T.info[0].tag="sakit"
	T.info[0].N=T.info[0].N+1
	T.info[0].jawab.isi[0]="Maag_adalah_penyakit_perut<Dr.Farah>"
	T.info[0].jawab.N=T.info[0].jawab.N+1
	T.info[0].jawab.isi[1]="Terima_Kasih_dok<Ayunisa>"
	T.info[0].jawab.N=T.info[0].jawab.N+1
	T.N=T.N+1
	
	T.info[1].tanya="Dok_perut_saya_sakit"
	T.info[1].tag="maag"
	T.info[1].N=T.info[1].N+3
	T.info[1].jawab.isi[0]="Bisa_konsultasikan_ke_rumahsakit_terdekat_untuk_pemeriksaan<Dr.Azizah>"
	T.info[1].jawab.N=T.info[1].jawab.N+1
	T.info[1].jawab.isi[1]="Terima_Kasih_dok<Aurellia>"
	T.info[1].jawab.N=T.info[1].jawab.N+1
	T.N=T.N+1
	
	T.info[2].tanya="Saya_sedih_dok:("
	T.info[2].tag="sedih"
	T.info[2].N=T.info[2].N+2
	T.info[2].jawab.isi[0]="Untuk_pertanyaan_tersebut_bisa_anda_konsultasikan_ke_psikologi_terdekat<Dr.Farah>"
	T.info[2].jawab.N=T.info[2].jawab.N+1
	T.info[2].jawab.isi[1]="Dokter_ga_seru<Aurellia>"
	T.info[2].jawab.N=T.info[2].jawab.N+1
	T.N=T.N+1
}

func main() {
	var doc dokter
	var pas pasien
	var T question
	header()
	dummy(&doc, &pas, &T)
	sortJmlTanya(&T)
	urutPasien(&pas)
	urutDokter(&doc)
	menu(&doc, &pas, &T)
}
