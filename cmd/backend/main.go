package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/royberkoweee/webgin/v1/pkg/bq"
	"github.com/royberkoweee/webgin/v1/pkg/logger"
)

var l = logger.NewLogger("ginweb")

func main() {
	router := gin.Default()
	router.POST("/query_bq", getBqResults)
	router.Run(":8080")
}

// Using a struct here cause the fields of request should be known and validated
type GetBqQueryResultsPayload struct {
	QueryString string `json:"query_string" binding:"required"`
	ProjectId   string `json:"project_id" binding:"required"`
}

func getBqResults(c *gin.Context) {
	payload := GetBqQueryResultsPayload{}
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var client = bq.NewBQClient(c, payload.ProjectId)
	l.LogInfo("going to run given query on project -> ", payload.ProjectId)
	var rows, schema = bq.GetQueryResults(c, client, payload.QueryString)
	sc := []gin.H{}
	var err = json.Unmarshal(schema, &sc)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"schema": sc, "rows": rows})
}
