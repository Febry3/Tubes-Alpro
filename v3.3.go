package main

import "fmt"

const NMAX int = 10

type mobil struct {
	namaMobil, namaPabrikan                  string
	jumlahUnit, tahunKeluar, jumlahPenjualan int
}

type arrMobil struct {
	M [NMAX]mobil
	n int
}

type pabrikan struct {
	nama                  string
	arrMobil              arrMobil
	jumlahPenjualanPabrik int
}

type arrPabrikan struct {
	Pabrik [NMAX]pabrikan
	n      int
}

func main() {
	var ans int
	var A arrPabrikan

	uimenu()
	fmt.Scan(&ans)
	for ans != 7 {
		for ans < 1 || ans > 7 {

			uimenu()
			fmt.Scan(&ans)
		}

		if ans == 1 {
			tambahData(&A)
		} else if ans == 2 {
			editData(&A)
		}
		if ans != 7 {
			uimenu()
			fmt.Scan(&ans)
		}

	}
	fmt.Print(A)
}

func uimenu() {
	fmt.Println("==============================================")
	fmt.Println("|         Selamat datang di aplikasi         |")
	fmt.Println("|                Dealer Mobil                |")
	fmt.Println("==============================================")
	fmt.Println()
	fmt.Println("=== Main Menu ===")
	fmt.Println()
	fmt.Println("1. Tambah data ")
	fmt.Println("2. Ubah data ")
	fmt.Println("3. Hapus data ")
	fmt.Println("4. Cari data")
	fmt.Println("5. Daftar data")
	fmt.Println("6. Top 3 data tertinggi")
	fmt.Println("7. Keluar")
	fmt.Println()
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiTambahData() {
	fmt.Println("1. Tambah data pabrikan")
	fmt.Println("2. Tambah data mobil")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiEditData() {
	fmt.Println("1. Edit data pabrikan")
	fmt.Println("2. Edit data mobil")
	fmt.Println("3. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func uiEditDataMobil() {
	fmt.Println("1. Edit data nama mobil")
	fmt.Println("2. Edit data jumlah unit")
	fmt.Println("3. Edit data tahun keluar")
	fmt.Println("4. Edit data jumlah penjualan")
	fmt.Println("5. Keluar")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func pilihan() {
	fmt.Println("1. Setuju")
	fmt.Println("2. Batal")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func pilihanEditPabrik() {
	fmt.Println("1. Edit nama pabrik")
	fmt.Println("2. Kelaur")
	fmt.Print("=== Pilihlah berdasarkan angka: ")
}

func tambahData(A *arrPabrikan) {
	var temp pabrikan
	var ans, ans1, idx int
	var tMob arrMobil
	var t string

	uiTambahData()
	fmt.Scan(&ans)
	for ans != 3 {
		for ans < 1 || ans > 3 {
			uiTambahData()
			fmt.Scan(&ans)
		}
		//opsi pilihan 1
		if ans == 1 {

			if A.n < NMAX {
				fmt.Print("Masukkan nama pabrikan : ")
				fmt.Scan(&temp.nama)

				pilihan()
				fmt.Scan(&ans1)

				for ans1 < 1 && ans1 > 2 {
					pilihan()
					fmt.Scan(&ans1)
				}
				if ans1 == 1 {
					A.Pabrik[A.n].nama = temp.nama
					A.n++
					fmt.Println("Data berhasil ditambahkan")
				}
			} else {
				fmt.Println("Data sudah penuh")
			}

			//opsi 2
		} else if ans == 2 {

			cetakPabrik(*A)
			fmt.Println()
			fmt.Print("Masukkan nama pabrik yang ingin diisi :")
			fmt.Scan(&t)
			fmt.Println()
			idx = cariPabrik(*A, t)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				if A.Pabrik[idx].arrMobil.n < NMAX {
					fmt.Println("Ketikan data secara horizontal dan ditambahi spasi dengan format:")
					fmt.Println("Nama mobil | Nama pabrikan | Jumlah unit | Jumlah penjualan | Tahun keluar ")
					fmt.Scan(&tMob.M[tMob.n].namaMobil, &tMob.M[tMob.n].namaPabrikan, &tMob.M[tMob.n].jumlahUnit, &tMob.M[tMob.n].jumlahPenjualan, &tMob.M[tMob.n].tahunKeluar)
					tMob.n++
				} else {
					fmt.Println("Data sudah penuh")
				}

				pilihan()
				fmt.Scan(&ans1)

				//jika pilihan user lebih dari opsi
				for ans1 < 1 && ans1 > 2 {
					pilihan()
					fmt.Scan(&ans1)
				}
				//Jika memilih opsi simpan
				if ans1 == 1 {
					A.Pabrik[idx].arrMobil = tMob
					A.Pabrik[idx].arrMobil.n = tMob.n
					fmt.Println("Data berhasil ditambahkan")
				}

			}

		}
		uiTambahData()
		fmt.Scan(&ans)
	}
}

func editData(A *arrPabrikan) {
	var ans, idx, ans1, ans2, idxM, tempInt int
	var t, temp, tMobil string
	var tPab pabrikan

	uiEditData()
	fmt.Scan(&ans)
	for ans != 3 {
		for ans < 1 || ans > 3 {
			uiEditData()
			fmt.Scan(&ans)
		}

		if ans == 1 {
			cetakPabrik(*A)
			fmt.Print("Masukkan nama pabrik yang ingin diedit : ")
			fmt.Scan(&t)

			idx = cariPabrik(*A, t)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")
				pilihanEditPabrik()
				fmt.Scan(&ans1)
				for ans < 1 || ans > 2 {
					pilihanEditPabrik()
					fmt.Scan(&ans1)
				}

				if ans1 == 1 {
					fmt.Print("Masukkan nama pabrik yang baru :")
					fmt.Scan(&temp)

					pilihan()
					fmt.Scan(&ans2)

					//jika pilihan user lebih dari opsi
					for ans2 < 1 && ans2 > 2 {
						pilihan()
						fmt.Scan(&ans2)
					}
					//Jika memilih opsi simpan
					if ans2 == 1 {
						A.Pabrik[idx].nama = temp
						fmt.Println("Data berhasil diubah")
					}
				}
			}
		} else if ans == 2 {
			cetakPabrik(*A)
			fmt.Print("Masukkan nama pabrikan mobil yang ingin diedit : ")
			fmt.Scan(&t)

			idx = cariPabrik(*A, t)

			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
			} else {
				fmt.Println("Data ditemukan")

				tPab = A.Pabrik[idx]

				cetakMobil(tPab)

				fmt.Print("Masukkan nama mobil yang ingin diubah datanya : ")
				fmt.Scan(&tMobil)
				idxM = cariMobil(tPab, tMobil)

				if idxM == -1 {
					fmt.Println("Data tidak ditemukan")
				} else {
					uiEditDataMobil()
					fmt.Scan(&ans1)

					for ans1 < 1 || ans1 > 5 {
						uiEditDataMobil()
						fmt.Scan(&ans1)
					}

					if ans1 == 1 {
						fmt.Print("Masukkan nama mobil yang baru :")
						fmt.Scan(&temp)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}
						//Jika memilih opsi simpan
						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].namaMobil = temp
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 2 {
						fmt.Print("Masukkan data jumlah unit yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].jumlahUnit = tempInt
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 3 {
						fmt.Print("Masukkan data tahun keluar yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].tahunKeluar = tempInt
							fmt.Println("Data berhasil diubah")
						}
					} else if ans1 == 4 {
						fmt.Print("Masukkan data jumlah penjualan yang baru : ")
						fmt.Scan(&tempInt)

						pilihan()
						fmt.Scan(&ans2)

						//jika pilihan user lebih dari opsi
						for ans2 < 1 && ans2 > 2 {
							pilihan()
							fmt.Scan(&ans2)
						}

						if ans2 == 1 {
							A.Pabrik[idx].arrMobil.M[idxM].jumlahPenjualan = tempInt
							fmt.Println("Data berhasil diubah")
						}
					}
				}

			}

		}
		uiEditData()
		fmt.Scan(&ans)

	}
}

//sequential
func cariPabrik(A arrPabrikan, target string) int {
	for i := 0; i < A.n; i++ {
		if A.Pabrik[i].nama == target {
			return i
		}
	}
	return -1
}

func cetakPabrik(A arrPabrikan) {
	fmt.Println("Daftar Pabrik")
	for i := 0; i < A.n; i++ {
		fmt.Println(i+1, ". ", A.Pabrik[i].nama)
	}
}

//sequential
func cariMobil(A pabrikan, target string) int {
	for i := 0; i < A.arrMobil.n; i++ {
		if A.arrMobil.M[i].namaMobil == target {
			return i
		}
	}
	return -1
}

func cetakMobil(A pabrikan) {
	fmt.Println("Daftar Mobil Pabrik", A.nama)
	fmt.Println("Nama Mobil|Nama Pabrikan|Jumlah Unit|Jumlah Penjualan|Tahun Keluar")
	for i := 0; i < A.arrMobil.n; i++ {
		fmt.Println(i+1, ".", A.arrMobil.M[i].namaMobil, "\t", A.arrMobil.M[i].namaPabrikan, "\t", A.arrMobil.M[i].jumlahUnit, "\t", A.arrMobil.M[i].jumlahPenjualan, "\t", A.arrMobil.M[i].tahunKeluar)
	}
}
