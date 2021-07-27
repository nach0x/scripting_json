package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Posts []struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Users []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  Address
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  Company
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	Catchphrase string `json:"catchphrase"`
	Bs          string `json:"bs"`
}

type Albums []struct {
	Userid int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

type Photos []struct {
	Albumid      int    `json:"albumid"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Thumbnailurl string `json:"thumbnailurl"`
}

type Info_user struct {
	Albums []info_albums
	Posts  []info_posts
	Photos []info_photos
}

type info_albums struct {
	Title string `json:"title"`
}

type info_posts struct {
	Title string `json:"title"`
}

type info_photos struct {
	Url string `json:"url`
}

var info_u Info_user

//var info_p info_posts

func main() {
	Post(User())
	Album(User())
	Photo()
	Convert()
	//fmt.Println(info_u)
}

func User() (int, error) {
	params := url.Values{}
	params.Add("email", "Nathan@yesenia.net")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	users := Users{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return 0, err
	}

	//usr := users[0].ID

	//fmt.Println(users[0].ID)

	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}
	return users[0].ID, nil
	//fmt.Println(users[0].Username)
}

func Post(data2 int, err error) {
	data1 := fmt.Sprintf("%d", data2)
	params := url.Values{}
	params.Add("userId", data1)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	for _, post := range posts {
		info_u.Posts = append(info_u.Posts, info_posts{post.Title})
		//fmt.Println(post.Title)
	}
}

func Album(data2 int, err error) {
	data1 := fmt.Sprintf("%d", data2)
	params := url.Values{}
	params.Add("userId", data1)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/albums?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	al := Albums{}
	err = json.Unmarshal(body, &al)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	for _, album := range al {
		info_u.Albums = append(info_u.Albums, info_albums{album.Title})
		//fmt.Println(album.Title)
	}
	//fmt.Println(users[0].Username)
}

func Photo() {
	params := url.Values{}
	params.Add("albumId", "3")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	ph := Photos{}
	err = json.Unmarshal(body, &ph)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	for _, photo := range ph {
		info_u.Photos = append(info_u.Photos, info_photos{photo.Url})
		//fmt.Println(photo.Url)
	}
	//fmt.Println(users[0].Username)
}

func Convert() {
	data, _ := json.Marshal(info_u)

	fmt.Println(string(data))
}
