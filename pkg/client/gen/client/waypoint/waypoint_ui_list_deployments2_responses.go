// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// WaypointUIListDeployments2Reader is a Reader for the WaypointUIListDeployments2 structure.
type WaypointUIListDeployments2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUIListDeployments2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUIListDeployments2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUIListDeployments2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUIListDeployments2OK creates a WaypointUIListDeployments2OK with default headers values
func NewWaypointUIListDeployments2OK() *WaypointUIListDeployments2OK {
	return &WaypointUIListDeployments2OK{}
}

/*
WaypointUIListDeployments2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointUIListDeployments2OK struct {
	Payload *models.HashicorpWaypointUIListDeploymentsResponse
}

// IsSuccess returns true when this waypoint Ui list deployments2 o k response has a 2xx status code
func (o *WaypointUIListDeployments2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint Ui list deployments2 o k response has a 3xx status code
func (o *WaypointUIListDeployments2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint Ui list deployments2 o k response has a 4xx status code
func (o *WaypointUIListDeployments2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint Ui list deployments2 o k response has a 5xx status code
func (o *WaypointUIListDeployments2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint Ui list deployments2 o k response a status code equal to that given
func (o *WaypointUIListDeployments2OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointUIListDeployments2OK) Error() string {
	return fmt.Sprintf("[GET /ui/deployments/application/{application.application}][%d] waypointUiListDeployments2OK  %+v", 200, o.Payload)
}

func (o *WaypointUIListDeployments2OK) String() string {
	return fmt.Sprintf("[GET /ui/deployments/application/{application.application}][%d] waypointUiListDeployments2OK  %+v", 200, o.Payload)
}

func (o *WaypointUIListDeployments2OK) GetPayload() *models.HashicorpWaypointUIListDeploymentsResponse {
	return o.Payload
}

func (o *WaypointUIListDeployments2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUIListDeploymentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUIListDeployments2Default creates a WaypointUIListDeployments2Default with default headers values
func NewWaypointUIListDeployments2Default(code int) *WaypointUIListDeployments2Default {
	return &WaypointUIListDeployments2Default{
		_statusCode: code,
	}
}

/*
WaypointUIListDeployments2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointUIListDeployments2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint UI list deployments2 default response
func (o *WaypointUIListDeployments2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint UI list deployments2 default response has a 2xx status code
func (o *WaypointUIListDeployments2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint UI list deployments2 default response has a 3xx status code
func (o *WaypointUIListDeployments2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint UI list deployments2 default response has a 4xx status code
func (o *WaypointUIListDeployments2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint UI list deployments2 default response has a 5xx status code
func (o *WaypointUIListDeployments2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint UI list deployments2 default response a status code equal to that given
func (o *WaypointUIListDeployments2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointUIListDeployments2Default) Error() string {
	return fmt.Sprintf("[GET /ui/deployments/application/{application.application}][%d] Waypoint_UI_ListDeployments2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUIListDeployments2Default) String() string {
	return fmt.Sprintf("[GET /ui/deployments/application/{application.application}][%d] Waypoint_UI_ListDeployments2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUIListDeployments2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUIListDeployments2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}