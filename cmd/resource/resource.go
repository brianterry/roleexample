package resource

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Create a CloudControlApi client with additional configuration
	svc2 := cloudcontrolapi.New(req.Session)
	// First we check to see if we are stabilizing.
	// We do this by checking the callback CallbackContext
	v, ok := req.CallbackContext["Stabilizing"].(bool)
	if ok {
		if v {
			out, err := svc2.GetResourceRequestStatus(&cloudcontrolapi.GetResourceRequestStatusInput{
				RequestToken: aws.String(*currentModel.Token),
			})
			if err != nil {
				return reportError(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
			}

			re := *out.ProgressEvent.OperationStatus
			switch re {
			case "IN_PROGRESS":
				return handler.ProgressEvent{
					OperationStatus:      handler.InProgress,
					CallbackDelaySeconds: 5,
					CallbackContext:      req.CallbackContext,
					Message:              "In progress",
					ResourceModel:        currentModel,
				}, nil
			case "FAILED":
				return reportError(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
			case "SUCCESS":
				return handler.ProgressEvent{
					OperationStatus: handler.Success,
					Message:         "Success",
					ResourceModel:   currentModel,
				}, nil
			}
		}
	}
	out, err := svc2.CreateResource(&cloudcontrolapi.CreateResourceInput{
		DesiredState: aws.String("{\"AssumeRolePolicyDocument\":{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":[\"ec2.amazonaws.com\"]},\"Action\":[\"sts:AssumeRole\"]}]},\"RoleName\":\"Test-role\",\"Path\":\"/\",\"Policies\":[{\"PolicyName\":\"root\",\"PolicyDocument\":{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Action\":\"*\",\"Resource\":\"*\"}]}}]}"),
		TypeName:     aws.String("AWS::IAM::Role"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudcontrolapi.ErrCodeNotStabilizedException:
				fmt.Println(cloudcontrolapi.ErrCodeNotStabilizedException, aerr.Error())
			case cloudcontrolapi.ErrCodeNotUpdatableException:
				fmt.Println(cloudcontrolapi.ErrCodeNotUpdatableException, aerr.Error())
			case cloudcontrolapi.ErrCodeThrottlingException:
				fmt.Println(cloudcontrolapi.ErrCodeThrottlingException, aerr.Error())
			case cloudcontrolapi.ErrCodeServiceInternalErrorException:
				fmt.Println(cloudcontrolapi.ErrCodeServiceInternalErrorException, aerr.Error())
			case cloudcontrolapi.ErrCodeServiceLimitExceededException:
				fmt.Println(cloudcontrolapi.ErrCodeServiceLimitExceededException, aerr.Error())
			case cloudcontrolapi.ErrCodeResourceConflictException:
				fmt.Println(cloudcontrolapi.ErrCodeResourceConflictException, aerr.Error())
			case cloudcontrolapi.ErrCodeResourceNotFoundException:
				fmt.Println(cloudcontrolapi.ErrCodeResourceNotFoundException, aerr.Error())
			default:
				return reportError(aerr.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return reportError(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
		}
	}

	currentModel.Token = out.ProgressEvent.RequestToken

	// Finally, we return an inProgress event.
	c := map[string]interface{}{"Stabilizing": true}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		CallbackDelaySeconds: 5,
		CallbackContext:      c,
		Message:              "In progress",
		ResourceModel:        currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "Read complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Read")
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "Update complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	// Construct a new handler.ProgressEvent and return it
	response := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
		ResourceModel:   currentModel,
	}

	return response, nil

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Delete")
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Add your code here:
	// * Make API calls (use req.Session)
	// * Mutate the model
	// * Check/set any callback context (req.CallbackContext / response.CallbackContext)

	/*
	   // Construct a new handler.ProgressEvent and return it
	   response := handler.ProgressEvent{
	       OperationStatus: handler.Success,
	       Message: "List complete",
	       ResourceModel: currentModel,
	   }

	   return response, nil
	*/

	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func reportError(message string, code string) handler.ProgressEvent {
	log.Println(message)
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		HandlerErrorCode: code,
		Message:          "Failed",
	}
}
