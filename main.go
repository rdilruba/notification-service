package main

import (
	"log"
	"notification-service/config"
	"notification-service/mail"
	"notification-service/message"
)

func main() {

	config := config.InitConfig()

	emailSender := mail.NewEmailSender(config)

	sqsReader, err := message.NewSQSClient(config)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for {
		messages, err := sqsReader.ReceiveMessages()
		if err != nil {
			log.Printf("error: %v", err)
		}
		for _, message := range messages {

			// TODO: simdilik kendi emailinize gonderebilirsiniz, daha sonra customer service ten okuyacagiz
			err := emailSender.SendEmail("dilrubakose@gmail.com", "You have a new order", *message.Body)

			if err != nil {
				log.Printf("error: %v", err)
				continue
			}

			err = sqsReader.DeleteMessage(message.ReceiptHandle)

			if err != nil {
				log.Printf("error: %v", err)
				continue
			}

		}
	}

}
