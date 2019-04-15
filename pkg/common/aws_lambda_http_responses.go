package common

/**
AWSLambdaHTTPResponse...
Declares a library of methods for creation and removal of structures 
dedicated to providing HTTP responses for use in serverless AWS Lambda 
functions.
Provides helper methods for reliable cast-transformation of those 
structs to different data-types.
**/

// AWSLambdaHTTPResponse defines a struct featuring only the necessary 
attributes (from Amazon Web Services documentation).
type AWSLambdaHTTPResponse struct {
	StatusCode int16
	Header     map[string]string
	Body       string
}

// helloWorld returns a Hello World response for testing and debugging 
purposes, and provides a white-box usage example.
func helloWorld() *AWSLambdaHTTPResponse {
	return newAWSLambdaHTTPResponse()
}

// NewAWSLambdaHTTPResponse returns a pointer reference to a new 
AWSLambdaHTTPResponse.
func newAWSLambdaHTTPResponse() *AWSLambdaHTTPResponse {
	return &AWSLambdaHTTPResponse{
		Header: map[string]string{},
	}
}

// fromJSON returns a pointer reference to a new AWSLambdaHTTPResponse 
from a JSON object of matching key-value pairs.
func fromJSON() *AWSLambdaHTTPResponse {
	return &AWSLambdaHTTPResponse{
		Header: map[string]string{},
	}
}

// toJSON builds a JSON formatted string from an AWSLambdaHTTPResponse.
func toJSON() *AWSLambdaHTTPResponse {
	return &AWSLambdaHTTPResponse{
		Header: map[string]string{},
	}
}
