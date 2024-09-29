package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion  string
	SQSUrl     string
	SMTPServer string
	EmailFrom  string
	EmailPass  string
	LogLevel   string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	awsRegion := os.Getenv("AWS_REGION")
	sqsUrl := os.Getenv("SQS_URL")

	smtpServer := os.Getenv("SMTP_SERVER")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPass := os.Getenv("EMAIL_PASS")

	logLevel := os.Getenv("LOG_LEVEL")

	return &Config{
		AWSRegion:  awsRegion,
		SQSUrl:     sqsUrl,
		SMTPServer: smtpServer,
		EmailFrom:  emailFrom,
		EmailPass:  emailPass,
		LogLevel:   logLevel,
	}
}
