package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var photoIconSvg = []byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M0 0h24v24H0z" fill="none"/>
<path d="M21 19V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zM8.5 13.5l2.5 3.01L14.5 12l4.5 6H5l3.5-4.5z"/>
</svg>`)

var albumIconSvg = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M10 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2h-8l-2-2z"/>
<path d="M0 0h24v24H0z" fill="none"/></svg>`)

var labelIconSvg = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
<path d="M0 0h24v24H0z" fill="none"/><path d="M17.63 5.84C17.27 5.33 16.67 5 16 5L5 5.01C3.9 5.01 3 5.9 3 7v10c0 1.1.9 1.99 2 1.99L16 19c.67 0 1.27-.33 1.63-.84L22 12l-4.37-6.16z"/></svg>`)

// GET /api/v1/svg/*
func GetSvg(router *gin.RouterGroup) {
	router.GET("/svg/photo", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/svg+xml", photoIconSvg)
	})

	router.GET("/svg/label", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/svg+xml", labelIconSvg)
	})

	router.GET("/svg/album", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/svg+xml", albumIconSvg)
	})
}
