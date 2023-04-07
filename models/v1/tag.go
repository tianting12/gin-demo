package v1

import "gin-demo/models"

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []models.Tag) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	models.Db.Model(&models.Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag models.Tag
	models.Db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	models.Db.Create(&models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagByID(id int, name string, modifiedBy string) bool {
	var tag models.Tag
	models.Db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	models.Db.Where("id=?", id).Delete(&models.Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	models.Db.Model(&models.Tag{}).Where("id=?", id).Update(data)
	return true
}
