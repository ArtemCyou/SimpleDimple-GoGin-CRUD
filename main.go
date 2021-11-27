package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type errorBad struct {
	ErrorB string `json:"error"`
}

// создаем albums slice и записываем данные.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//getAlbums возвращает список всех альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums добавляет альбом из JSON, полученный в теле запроса.
func postAlbum(c *gin.Context) {
	var newAlbum album

//Вызов BindJSON, чтобы привязать полученный JSON к	newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorBad{"bad_request"})
		return
	}
//Добавляем newAlbum в срез.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID находит альбом, id которого совпадает с id
//отправленный клиентом, затем возвращает этот альбом в качестве ответа.
func getAlbumById(c *gin.Context)  {
	id:= c.Param("id")

	for _, a:= range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"сообщение":"альбом не найден"})
}

func deleteAlbumByID(c *gin.Context)  {
	id:= c.Param("id")
	for i, a:= range albums{
		if a.ID == id{
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"сообщение":"альбом не найден"})
}

func updateAlbumByID(c *gin.Context)  {
	id:= c.Param("id")
	for i, a:= range albums{
		if a.ID == id{
			c.BindJSON(&a)
			albums[i] = a
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"сообщение":"альбом не найден"} )
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById) //В Gin двоеточие перед элементом в пути означает, что
	// этот элемент является параметром пути.
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.Run("localhost: 8080")
}
