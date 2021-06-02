package handlers

import (
	"net/http"

	"github.com/krastomer/netflix-backend/database"
	"github.com/labstack/echo/v4"
)

func BrowseHandlers(e *echo.Group) {
	e.GET("", getBrowseHandler)
	e.GET("/mylist", getMyListHandler)
	e.GET("/movie", getMovieBrowseHandler)
	e.GET("/tvshows", getTVShowsBrowseHandler)
	e.GET("/history", getHistoryHandler)
}

func getBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func getMyListHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	myList := database.GetMyList(v)
	return c.JSON(http.StatusOK, myList)
}

func getHistoryHandler(c echo.Context) error {
	_, v := getUserFromToken(c)
	hm := database.GetHistoryMovie(v)
	return c.JSON(http.StatusOK, hm)
}

func getMovieBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func getTVShowsBrowseHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

// TODO: getHistory
