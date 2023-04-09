package v1

import "gin-demo/models"

func AddArticle(data map[string]interface{}) bool {
	models.Db.Create(&models.Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func ExistArticleById(id int) bool {
	var article models.Article
	models.Db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticle(id int) (article models.Article) {
	models.Db.Where("id = ?", id).First(&article)
	models.Db.Model(&article).Related(&article.Tag)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []models.Article) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticlesTotal(maps interface{}) (count int) {
	models.Db.Model(&models.Tag{}).Where(maps).Count(&count)
	return
}

func DeleteArticle(id int) bool {
	var article models.Article
	models.Db.Where("id = ?", id).Delete(&article)
	return true
}

func EditArticle(id int, data map[string]interface{}) bool {
	models.Db.Model(&models.Article{}).Where("id = ?", id).Update(data)
	return true
}
