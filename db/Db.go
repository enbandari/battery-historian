package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type TaskRecord struct {
	gorm.Model
	Remark       string `gorm:"size:1024"`
	DeviceModel  string
	DeviceId string
	SDKVersion int
	AnalyzedData string `gorm:"size:65536"`

}

const host = "root:@/battery_task?charset=utf8&parseTime=True&loc=Local"
const sqlType = "mysql"

func InsertRecord(record TaskRecord) {
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	err = db.Create(&record).Error
	if  err != nil {
		panic(err)
	}
	fmt.Println("insert ok: ")
}

func DeleteRecord(record TaskRecord) {
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	err = db.Delete(&record).Error
	if  err != nil {
		panic(err)
	}
	fmt.Println("delete ok: ")
}

func ListAllRecordsWithoutData() []TaskRecord{
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	var records []TaskRecord
	err = db.Select("id,created_at, remark, device_model, sdk_version").Find(&records).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("list ok: ", len(records))
	return records
}

func ListRecordsWithoutData(page uint16, size uint16) []TaskRecord{
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	var records []TaskRecord
	err = db.Where("limit ? offset ?", size, page * size).Select("id", "created_at", "remark", "device_model").Find(&records).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("list ok: ", page, size)
	return records
}

func QueryDataForId(id string) string{
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	var result = TaskRecord{}
	err = db.Select("analyzed_data").Where("id = ?", id).Find(&result).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("query ok, size: ", len(result.AnalyzedData))
	return result.AnalyzedData
}

func ClearRecords() {
	db, err := gorm.Open(sqlType, host)
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TaskRecord{})

	err = db.Delete(&TaskRecord{}).Error
	if  err != nil {
		panic(err)
	}
	fmt.Println("clear ok.")
}
