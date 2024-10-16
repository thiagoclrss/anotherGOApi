package model

type Stock struct {
	ID     uint64 `gorm:"primary_key;auto_increment"`
	Ticker string
	Price  float64
}
