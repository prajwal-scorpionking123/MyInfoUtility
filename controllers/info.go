package controllers

import (
	"MyUtilityBackend/helpers"
	"MyUtilityBackend/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var connection = helpers.ConnectToMongoDB()
var collection = connection.Collection("Info")

func AddInfo(c *gin.Context) {
	var info models.Info
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := collection.InsertOne(context.TODO(), info)
	if err != nil {
		helpers.GetError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": res})
}

func GetInfos(c *gin.Context) {
	var infos []models.Info

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helpers.GetError(err, c)
		return
	}

	if err != nil {
		helpers.GetError(err, c)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var info models.Info
		err := cur.Decode(&info)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		infos = append(infos, info)
	}
	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"infos": infos})

}
