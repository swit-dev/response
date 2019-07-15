package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Gin Gin
type Gin struct {
	C *gin.Context
}

//JSONResponse JSONResponse
func (g *Gin) JSONResponse(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  GetMsg(errCode),
		"data": data,
	})
}
func (g *Gin) JSONMarshalResponse(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  GetMsg(errCode),
		"data": data,
	})
}

//ErrorResponse ErrorResponse
func (g *Gin) ErrorResponse(errCode int) {
	g.C.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": errCode,
		"msg":  GetMsg(errCode),
	})
}

//ErrorResponse ErrorResponse
func (g *Gin) ErrorResponseJson(errCode int, data interface{}) {
	g.C.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": errCode,
		"msg":  GetMsg(errCode),
		"data": data,
	})
}

//LoginRedirectlResponse LoginRedirectlResponse
func (g *Gin) LoginRedirectlResponse(returnURL string) {
	g.C.Redirect(http.StatusMovedPermanently, "/login")
	return
}

//RedirectlResponse RedirectlResponse
func (g *Gin) RedirectlResponse(redirectURL string) {
	g.C.Redirect(http.StatusMovedPermanently, redirectURL)
	return
}
