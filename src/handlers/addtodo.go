package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Todo struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"created_at"`
}

func init() {
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: &region},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)
}

func AddTodo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(fmt.Sprintf("AWS_REGION: %v \n", os.Getenv("AWS_REGION")))
	fmt.Println(fmt.Sprintf("TODO_TABLE_NAME: %v \n", os.Getenv("TODOS_TABLE_NAME")))

	// TODO: use aws sdk to connect to dynamoDB

	// TODO: parse request body and write to dynamoDB

	fmt.Println("[addtodo] Received body: ", request.Body)
	return events.APIGatewayProxyResponse{
		Body: request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(AddTodo)
}