package produk

import (
	"database/sql"
	"errors"
	"fmt"
)

type ProdukModel struct {
	conn *sql.DB
}

func (pm *ProdukModel) SetSQLConnection(db *sql.DB) {
	pm.conn = db
}

func (pm *ProdukModel) TambahProduk(newProduk Produk) error {
	res, err := pm.conn.Exec("INSERT INTO produk (barcode, nama, qty, harga, input_oleh) values(?,?,10,?,?)", newProduk.Barcode, newProduk.Nama, newProduk.Harga, newProduk.InputOleh)

	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	pm.conn.Close()

	return nil
}

func (pm ProdukModel) GetAllProduct() ([]Produk, error) {
	listProduk := []Produk{}
	rows, err := pm.conn.Query("SELECT barcode, nama, qty, harga, input_oleh FROM produk")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var newProduk = Produk{}
		if err := rows.Scan(&newProduk.Barcode, &newProduk.Nama, &newProduk.Qty, &newProduk.Harga, &newProduk.InputOleh); err != nil {
			fmt.Println(err)
			return nil, err
		}
		listProduk = append(listProduk, newProduk)
	}
	pm.conn.Close()

	return listProduk, nil
}

func (pm *ProdukModel) EditBarang(barcode int, updatedData Produk) error {
	res, err := pm.conn.Exec("UPDATE produk SET nama = ?, harga = ?, qty = ?", barcode, updatedData.Nama, updatedData.Harga, updatedData.Qty)

	if err != nil {
		fmt.Println(err)
		return err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if aff <= 0 {
		return errors.New("terjadi sebuah masalah pada sistem")
	}

	pm.conn.Close()

	return nil
}
