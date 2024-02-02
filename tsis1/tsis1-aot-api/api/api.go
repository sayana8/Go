package api

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Nickname        string `json:"nickname"`
	TopListenedSong string `json:"top_listened_song"`
}
