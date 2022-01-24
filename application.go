package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
	"net/http"
	"os"
)

func sendMailHandler(c *gin.Context) {
	mailjetClient := mailjet.NewMailjetClient("d4f64bc764c4576540cf8033fd69a1e5", "624fbd6d20aad779ca3db090baa0294a")

	from := mailjet.RecipientV31{
		Email: "yaronp@gmail.com",
		Name:  "yaron",
	}
	to := mailjet.RecipientV31{
		Email: "yaronp@gmail.com",
		Name:  "yaron",
	}

	message := mailjet.InfoMessagesV31{
		From:     &from,
		To:       &mailjet.RecipientsV31{to},
		Subject:  "Hello lusha!",
		TextPart: "This mail is generated by lusha test.\n https://jan20.lushatest.com,\n Cheers, Yaron ",
		HTMLPart: "<p><b>Hi Arik,</b></p><br>This mail is generated by lusha test. https://jan20.lushatest.com<br><br><p><i>Cheers, Yaron.</i></p>",
		CustomID: "AppLushTest",
	}

	messagesInfo := []mailjet.InfoMessagesV31{message}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, res)
		log.Printf("sendMailHandler message send data: %+v\n", res)
	}
}

const DefaultPort = "5000"
const PortEnvParam = "PORT"

func main() {
	port := os.Getenv(PortEnvParam)
	if port == "" {
		port = DefaultPort
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
	r.GET("/", func(c *gin.Context) { c.File("public/index.html") })
	r.POST("/send", sendMailHandler)

	r.Run(":" + port)

}
