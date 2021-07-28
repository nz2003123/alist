package models

import (
	"fmt"
	"github.com/Xhofe/alist/conf"
	"time"
)

type File struct {
	Dir           string     `json:"dir" gorm:"index"`
	FileExtension string     `json:"file_extension"`
	FileId        string     `json:"file_id"`
	Name          string     `json:"name" gorm:"index"`
	Type          string     `json:"type"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Category      string     `json:"category"`
	ContentType   string     `json:"content_type"`
	Size          int64      `json:"size"`
	Password      string     `json:"password"`
	Url           string     `json:"url" gorm:"-"`
	ContentHash   string     `json:"content_hash"`
}

func (file *File) Create() error {
	return conf.DB.Create(file).Error
}

func Clear(drive *conf.Drive) error {
	if err := conf.DB.Where("dir = '' AND name = ?", drive.Name).Delete(&File{}).Error; err != nil {
		return err
	}
	return conf.DB.Where("dir like ?", fmt.Sprintf("%s%%", drive.Name)).Delete(&File{}).Error
}

func GetFileByDirAndName(dir, name string) (*File, error) {
	var file File
	if err := conf.DB.Where("dir = ? AND name = ?", dir, name).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func GetFilesByDir(dir string) (*[]File, error) {
	var files []File
	if err := conf.DB.Where("dir = ?", dir).Find(&files).Error; err != nil {
		return nil, err
	}
	return &files, nil
}

func SearchByNameGlobal(keyword string) (*[]File, error) {
	var files []File
	if err := conf.DB.Where("name LIKE ? ", fmt.Sprintf("%%%s%%", keyword)).Find(&files).Error; err != nil {
		return nil, err
	}
	return &files, nil
}

func SearchByNameInDir(keyword string, dir string) (*[]File, error) {
	var files []File
	if err := conf.DB.Where("dir LIKE ? AND name LIKE ? ", fmt.Sprintf("%s%%", dir), fmt.Sprintf("%%%s%%", keyword)).Find(&files).Error; err != nil {
		return nil, err
	}
	return &files, nil
}

func DeleteWithDir(dir string) error {
	return conf.DB.Where("dir like ?", fmt.Sprintf("%s%%", dir)).Delete(&File{}).Error
}
