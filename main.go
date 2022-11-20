package main

import (
	"github.com/anandyeole/JioSaavnApi_go/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/search/songs/:query", func(c *gin.Context) {
		query := c.Param("query")
		songlist := api.GetSongList(query)
		c.JSON(200, songlist)
	})
	router.GET("/search/album/:query", func(c *gin.Context) {
		query := c.Param("query")
		albumlist := api.GetAlbumList(query)
		c.JSON(200, albumlist)
	})
	router.GET("/search/playlist/:query", func(c *gin.Context) {
		query := c.Param("query")
		playlist := api.GetPlaylists(query)
		c.JSON(200, playlist)
	})
	router.GET("/song/:songID", func(c *gin.Context) {
		songID := c.Param("songID")
		song := api.GetSongDetails(songID)
		c.JSON(200, song)
	})
	router.GET("/album/:albumID", func(c *gin.Context) {
		albumID := c.Param("albumID")
		album := api.GetAlbumDetails(albumID)
		c.JSON(200, album)
	})
	router.GET("/playlist/:playlistID", func(c *gin.Context) {
		playlistID := c.Param("playlistID")
		playlist := api.GetPlaylistDetails(playlistID)
		c.JSON(200, playlist)
	})
	router.Run(":8080")
}
