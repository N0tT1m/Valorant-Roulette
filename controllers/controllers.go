package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "Valorant Roulette",
	})
}

func Agent(c *gin.Context) {
	c.HTML(http.StatusOK, "agentpicker.gohtml", gin.H{
		"title": "Agent Picker",
	})
}

func Team(c *gin.Context) {
	c.HTML(http.StatusOK, "teampicker.gohtml", gin.H{
		"title": "Team Picker",
	})
}

func Map(c *gin.Context) {
	c.HTML(http.StatusOK, "mappicker.gohtml", gin.H{
		"title": "Map Picker",
	})
}
