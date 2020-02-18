package swingsingersbundle

//SwingSinger struct
type SwingSinger struct {
	Name           string `json:"name"`
	Nationality    string `json:"nationality"`
	NumberOfAlbums int    `json:"number_of_albums"`
}

//NewSwingSinger creates a new singer
func NewSwingSinger(name string, nationality string, numberOfAlbums int) *SwingSinger {
	singer := new(SwingSinger)
	singer.Name = name
	singer.Nationality = nationality
	singer.NumberOfAlbums = numberOfAlbums
	return singer
}
