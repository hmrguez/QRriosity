package config

import (
	"backend/internal/db/repository"
	"backend/internal/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"os"
)

type Config struct {
	Port        string
	UserRepo    repository.UserRepository
	ProblemRepo repository.ProblemRepository
	TopicRepo   repository.TopicRepository
	AuthService services.CognitoAuthService
}

func NewLocalConfig() *Config {

	userRepo, err := repository.NewMongoDBUserRepository("mongodb://localhost:27017", "saas", "users")
	if err != nil {
		panic("Couldn't connnect to mongodb: " + err.Error())
	}

	topicRepo, err := repository.NewMongoDBTopicRepository("mongodb://localhost:27017", "saas", "topics")
	if err != nil {
		panic("Couldn't connnect to mongodb: " + err.Error())
	}

	authService := *services.NewCognitoAuthService(
		os.Getenv("COGNITO_APP_CLIENT_ID"),
		os.Getenv("COGNITO_USER_POOL_ID"),
		os.Getenv("COGNITO_USER_POOL_REGION"),
	)

	localLlmService := services.NewLDefaultLMStudioService()

	return &Config{
		Port:        ":9000",
		UserRepo:    userRepo,
		AuthService: authService,
		ProblemRepo: repository.NewMongoDBProblemRepository(localLlmService),
		TopicRepo:   topicRepo,
	}
}

func NewProdConfig() *Config {
	// Create a new AWS session using credentials from .env
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	userRepo, err := repository.NewMongoDBUserRepository("mongodb://localhost:27017", "saas", "users")
	if err != nil {
		panic("Couldn't connnect to mongodb: " + err.Error())
	}

	authService := *services.NewCognitoAuthService(
		os.Getenv("COGNITO_APP_CLIENT_ID"),
		os.Getenv("COGNITO_USER_POOL_ID"),
		os.Getenv("COGNITO_USER_POOL_REGION"),
	)

	llmService := services.NewLLMService()

	return &Config{
		Port:        ":9000",
		UserRepo:    userRepo,
		ProblemRepo: repository.NewMongoDBProblemRepository(llmService),
		TopicRepo:   repository.NewDynamoDBTopicRepository(sess, "Qriosity-Topics"),
		AuthService: authService,
	}
}
