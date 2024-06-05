package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func (h *Handler) FizzBuzzHandler(ctx *gin.Context) {
	queryFrom := ctx.Query("from")
	from, err := strconv.Atoi(queryFrom)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "query 'from' must be number",
			"cause":   "invalid parameter",
		})
		return
	}

	queryTo := ctx.Query("to")
	to, err := strconv.Atoi(queryTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "query 'to' must be number",
			"cause":   "invalid parameter",
		})
		return
	}

	if to-from+1 > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "range must equal or less than 100",
			"cause":   "invalid parameter",
		})
		return
	}

	if to-from < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "parameter 'to' must be equal or greater than 'from'",
			"cause":   "invalid parameter",
		})
		return
	}

	result := h.Service.DoFizzBuzz(ctx.Request.Context(), from, to)
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
