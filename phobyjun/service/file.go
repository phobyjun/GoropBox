package service

import (
	"phobyjun/db"
	"phobyjun/model"
)

func CreateFile(fileDto *model.File, userId uint) (*model.File, error) {
	file := model.File{
		FileName: fileDto.FileName,
		FileDir:  fileDto.FileDir,
		UserID:   userId,
	}
	tx := db.Session.Create(&file)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &file, nil
}
