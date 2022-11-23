package api

import (
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Get the song list for a given query from JioSaavn
func GetSongList(query string) []interface{} {
	query = strings.Replace(query, " ", "%20", -1)
	url := search_url + query
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)

	}
	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	songs := result["results"].([]interface{})
	var songlist []interface{}
	songlist = append(songlist, songs...)
	return songlist
}

// get specific song details from the song id
func GetSongDetails(songID string) map[string]any {
	url := song_url + songID
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	song := result[songID].(map[string]any)

	//decrypt the song url
	song["decrypted_media_url"] = DecryptURL(song["encrypted_media_url"].(string))

	//fix the image url
	song["image"] = FixImageURL(song["image"].(string))

	//fix the title
	song["song"] = FixTitle(song["song"].(string))

	return song
}

// get the albums from the query
func GetAlbumList(query string) []interface{} {
	query = strings.Replace(query, " ", "%20", -1)
	url := albumlist_url + query
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	album := result["albums"].(map[string]interface{})
	albums := album["data"].([]interface{})
	var albumlist []interface{}
	albumlist = append(albumlist, albums...)
	return albumlist
}

// get the album details from the album id
func GetAlbumDetails(albumID string) map[string]any {
	url := album_url + albumID
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var album map[string]any
	json.Unmarshal(body, &album)
	for _, song := range album["songs"].([]interface{}) {
		song.(map[string]any)["decrypted_media_url"] = DecryptURL(song.(map[string]any)["encrypted_media_url"].(string))
		song.(map[string]any)["image"] = FixImageURL(song.(map[string]any)["image"].(string))
		song.(map[string]any)["song"] = FixTitle(song.(map[string]any)["song"].(string))
	}
	return album
}

// get the playlists from the query
func GetPlaylists(query string) []interface{} {
	query = strings.Replace(query, " ", "%20", -1)
	url := playists_url + query
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	plist := result["playlists"].(map[string]interface{})
	plists := plist["data"].([]interface{})
	var playlists []interface{}
	playlists = append(playlists, plists...)
	return playlists
}

// get the playlist details from the playlist id
func GetPlaylistDetails(playlistID string) map[string]any {
	url := playlist_url + playlistID
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var playlist map[string]any
	json.Unmarshal(body, &playlist)
	for _, song := range playlist["songs"].([]interface{}) {
		song.(map[string]any)["decrypted_media_url"] = DecryptURL(song.(map[string]any)["encrypted_media_url"].(string))
		song.(map[string]any)["image"] = FixImageURL(song.(map[string]any)["image"].(string))
		song.(map[string]any)["song"] = FixTitle(song.(map[string]any)["song"].(string))
	}
	return playlist
}

// decrypt and format the song url
func DecryptURL(url string) string {
	//fmt.Println(url)
	crypted, _ := base64.RawStdEncoding.DecodeString(url)
	//fmt.Println(crypted)
	decrypted, _ := decryptECB(crypted)
	//fmt.Println(decrypted)
	decryptedurl := strings.Replace(string(decrypted), "http://aac.saavncdn.com", "http://h.saavncdn.com", -1)
	decryptedurl = strings.Replace(decryptedurl, "_96.mp4", ".mp3", -1)
	return decryptedurl
}

// fix title of the song remove unwanted characters
func FixTitle(title string) string {
	title = strings.Replace(title, "&quot;", "", -1)
	title = strings.Replace(title, "&#039;", "'", -1)
	title = strings.Replace(title, "/", "|", -1)
	return title
}

// fix image url
func FixImageURL(url string) string {
	url = strings.Replace(url, "150x150", "500x500", -1)
	return url
}

// PKCS5UnPadding removes padding from the decrypted data
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// decryptECB decrypts the crypted data using the key and iv
func decryptECB(crypted []byte) ([]byte, error) {
	block, _ := des.NewCipher(key)
	origin := make([]byte, len(crypted))
	dst := origin
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		log.Fatal("crypto/cypher input size not valid")
		return nil, errors.New("crypto/cypher input size not valid")
	}

	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}
	origin = PKCS5UnPadding(origin)
	return origin, nil
}
