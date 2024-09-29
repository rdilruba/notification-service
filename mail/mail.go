package mail

import (
	"fmt"
	"net/smtp"
	"notification-service/config"
)

type EmailSender struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
}

func NewEmailSender(config *config.Config) *EmailSender {
	return &EmailSender{
		SMTPServer: config.SMTPServer,
		Port:       "587",
		Username:   config.EmailFrom,
		Password:   config.EmailPass,
	}
}

func (es *EmailSender) SendEmail(to, subject, body string) error {
	from := es.Username
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	auth := smtp.PlainAuth("", es.Username, es.Password, es.SMTPServer)
	err := smtp.SendMail(es.SMTPServer+":"+es.Port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
