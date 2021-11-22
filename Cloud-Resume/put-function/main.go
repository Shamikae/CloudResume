package main

import (
	// "errors"
	// "io/ioutil"
	// "net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	


	// "context"
    // "fmt"
    "log"

    // "github.com/aws/aws-sdk-go-v2/aws"
    // "github.com/aws/aws-sdk-go-v2/config"
    // "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// var (
// 	// DefaultHTTPGetAddress Default Address
// 	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

// 	// ErrNoIP No IP found in response
// 	ErrNoIP = errors.New("no IP in HTTP response")

// 	// ErrNon200Response non 200 status code in response
// 	ErrNon200Response = errors.New("non 200 Response found")
// )

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// config.WithSharedConfigProfile("Cloud-Resume"),
	// )

	// input := &dynamodb.UpdateItemInput{
	// 	Key: key,
	// 	ExpressionAttributeValues:
	// }

	svc := dynamodb.New(sess)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("cloud-resume"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("visitors"),
			},
		},
		UpdateExpression: aws.String("ADD visitors :inc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				N: aws.String("1"),
			},
		},
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	// resp, err := http.Get(DefaultHTTPGetAddress)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if resp.StatusCode != 200 {
	// 	return events.APIGatewayProxyResponse{}, ErrNon200Response
	// }

	// ip, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if len(ip) == 0 {
	// 	return events.APIGatewayProxyResponse{}, ErrNoIP
	// }

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:       string("{ \"count\": \"2\" }"),		
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
	// Using the SDK's default configuration, loading additional config
    // and credentials values from the environment variables, shared
    // credentials, and shared configuration files
    // cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
    // if err != nil {
    //     log.Fatalf("unable to load SDK config, %v", err)
    // }

    // // Using the Config value, create the DynamoDB client
    // svc := dynamodb.NewFromConfig(cfg)

    // // Build the request with its input parameters
    // resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
    //     Limit: aws.Int32(5),
    // })
    // if err != nil {
    //     log.Fatalf("failed to list tables, %v", err)
    // }

    // fmt.Println("Tables:")
    // for _, tableName := range resp.TableNames {
    //     fmt.Println(tableName)
    // }
}
