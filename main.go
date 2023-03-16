package main

import (
	"fmt"
	"simpleSQL/config"
	"simpleSQL/produk"
)

// ini remark ugi
func main() {
	var koneksi = config.InitSQL()
	var mdl = produk.ProdukModel{}
	mdl.SetSQLConnection(koneksi)

	var menu int
	for menu != 99 {
		fmt.Println("1. Input Barang")
		fmt.Println("2. Tampilkan Daftar Barang")
		fmt.Println("4. Delete Barang")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:

		case 2:
			res, err := mdl.GetAllProduct()
			if err != nil {
				fmt.Println("Terjadi sebuah kesalahan")
				break
			}

			for i := 0; i < len(res); i++ {
				fmt.Println(res[i])
			}
		case 4:
			var barcode int
			fmt.Print("Masukkan barcode produk:")
			fmt.Scanln(&barcode)
			err := mdl.DeleteProduk(barcode)
			if err != nil {
				fmt.Println("Terjadi sebuah kesalahan")
				break
			}

			fmt.Println("sukses menghapus data")
		}
	}
}
