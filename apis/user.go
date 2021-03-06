package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Gin_Mysql_api/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works.")
}

func CommonAddUser(user models.User) (msg string) {

	firstName := user.FirstName
	lastName := user.LastName

	log.Println(firstName, lastName)

	p := models.User{FirstName: firstName, LastName: lastName}

	ra, err := p.AddUser()
	if err != nil {
		log.Println(err)
	}

	msg = fmt.Sprintf("insert successful %d", ra)
	return msg
}

func AddUserApi(c *gin.Context) {

	var user models.User

	err := c.ShouldBind(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := CommonAddUser(user)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func AddUsersApi(c *gin.Context) {

	var users models.Users

	err := c.ShouldBindJSON(&users)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users.Users {

		var msg string

		if user.FirstName == "" || user.LastName == "" {
			msg = fmt.Sprintf("invalid user.FirstName and user.LastName")

		} else {
			msg = CommonAddUser(user)
		}

		c.JSON(http.StatusOK, gin.H{"msg": msg})

	}
}

func DelUserIdsApi(c *gin.Context) {
	var users models.Users
	// var persons []models.Person
	err := c.BindJSON(&users)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, user := range users.Users {

		id := user.Id

		log.Println(id)

		p := models.User{Id: id}

		ra, err := p.DelUser()
		if err != nil {
			log.Println(err)
		}

		msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}
}

func DelUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}

	ra, err := p.DelUser()
	if err != nil {
		log.Println(err)
	}

	msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func ModUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}

	err := c.Bind(&p)
	if err != nil {
		log.Println(err)
	}

	ra, err := p.ModUser()
	if err != nil {
		log.Println(err)
	}

	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	p := models.User{Id: id}

	user, err := p.GetUser()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsersApi(c *gin.Context) {
	p := models.User{}

	users, err := p.GetUsers()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
