package main

import (
	"github.com/apardee/sam-template/shared"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := shared.MethodMessage(req.HTTPMethod)
	return events.APIGatewayProxyResponse{Body: string(message), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
