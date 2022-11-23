package api

const (
	search_url    = "https://www.jiosaavn.com/api.php?p=1&_format=json&_marker=0&api_version=4&ctx=android&n=10&__call=search.getResults&q="
	song_url      = "https://www.jiosaavn.com/api.php?__call=song.getDetails&cc=in&ctx=android&_format=json&_marker=0&pids="
	albumlist_url = "https://www.jiosaavn.com/api.php?__call=autocomplete.get&_format=json&_marker=0&cc=in&includeMetaTags=1&query="
	album_url     = "https://www.jiosaavn.com/api.php?__call=content.getAlbumDetails&ctx=android&_format=json&_marker=0&albumid="
	playists_url  = "https://www.jiosaavn.com/api.php?__call=autocomplete.get&_format=json&_marker=0&cc=in&includeMetaTags=1&query="
	playlist_url  = "https://www.jiosaavn.com/api.php?__call=playlist.getDetails&_format=json&cc=in&_marker=0_&listid="
)

var (
	// Key is the key used to decrypt the song
	key = []byte("38346591")
)
