package v1

import "gin-demo/models"

func CheckAuth(username string, password string) bool {
	var auth models.Auth

	models.Db.Select("id").Where(models.Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
