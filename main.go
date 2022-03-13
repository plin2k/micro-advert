package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Advert struct {
	CreatedAt time.Time
	Subject   string
	Message   string
	From      string
}

type Adverts []Advert

var advertList Adverts

var advertChannel chan Advert

var maxItems int

func main() {

	if max, exists := os.LookupEnv("MAX_ITEMS"); exists {
		maxItemsEnv, err := strconv.Atoi(max)
		if err != nil {
			log.Fatal(err)
		}

		maxItems = maxItemsEnv
	} else {
		maxItems = 100
	}

	advertList = make(Adverts, maxItems)
	advertChannel = make(chan Advert, 10)
	//CSRF := csrf.Protect(
	//	[]byte("a-32-byte-long-key-goes-here"),
	//	csrf.RequestHeader("Authenticity-Token"),
	//)

	go pushAdvert(advertChannel)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.StaticFile("/style.css", "resources/style.css")
	router.StaticFile("/script.js", "resources/script.js")

	router.GET("/", index)
	router.POST("/publish", publish)

	//http.ListenAndServe(":8080", CSRF(router))
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func index(c *gin.Context) {
	//c.Header("X-CSRF-Token", csrf.Token(c.Request))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"advList": advertList,
		//csrf.TemplateTag: csrf.TemplateField(c.Request),
	})
}

func publish(c *gin.Context) {
	//c.Header("X-CSRF-Token", csrf.Token(c.Request))
	from := c.Query("from")

	subject := c.PostForm("subject")
	message := c.PostForm("message")

	if subject == "" || message == "" {
		c.Redirect(http.StatusFound, "/")

		return
	}

	advertChannel <- Advert{
		CreatedAt: time.Now(),
		Subject:   subject,
		Message:   message,
		From:      from,
	}

	c.Redirect(http.StatusFound, "/")
}

func pushAdvert(ch <-chan Advert) {

	for val := range ch {
		advertList = append(Adverts{val}, advertList[0:maxItems-1]...)

		log.Println(cap(advertList), len(advertList))
	}
}
