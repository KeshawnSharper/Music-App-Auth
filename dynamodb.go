package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var svc = dynamodb.New(sess)

func scanUsers() []UserStruct {

	params := &dynamodb.ScanInput{

		TableName: aws.String("MusicAppUsers"),
	}

	var scannedUsers []UserStruct

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	for _, i := range result.Items {
		item := UserStruct{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		scannedUsers = append(scannedUsers, item)
	}

	return scannedUsers
}

func addUser(user UserRequestStruct) UserStruct {
	userLength := len(scanUsers())
	var convertedUser UserStruct
	convertedUser.Email = user.Email
	convertedUser.Password = user.Password
	convertedUser.ID = int(userLength)
	convertedUser.Username = user.Username
	newUser, err := dynamodbattribute.MarshalMap(convertedUser)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      newUser,
		TableName: aws.String("MusicAppUsers"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	return scanUsers()[len(scanUsers())-1]
}
