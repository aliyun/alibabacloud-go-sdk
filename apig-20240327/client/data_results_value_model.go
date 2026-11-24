// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataResultsValue interface {
	dara.Model
	String() string
	GoString() string
	SetRouteId(v string) *DataResultsValue
	GetRouteId() *string
	SetStatus(v string) *DataResultsValue
	GetStatus() *string
	SetErrorMessage(v string) *DataResultsValue
	GetErrorMessage() *string
}

type DataResultsValue struct {
	// The route ID corresponding to the result, which is the same as the key in results.
	//
	// example:
	//
	// hr-example001
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
	// For asynchronous publishing, Pending indicates that the request is accepted, and Rejected indicates that the request is rejected before submission. For synchronous publishing, Succeeded indicates that the route is successfully published, and Failed indicates that the publishing failed.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The cause of the route being rejected or failing to be published. This parameter is returned when the status is Rejected or Failed.
	//
	// example:
	//
	// ErrorType: Conflict, ErrorCode: DeployRouteForbidden, Message: the route  example-route is not allowed to deploy when it\\"s status is Deploying., Params: [example-route Deploying]
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

func (s DataResultsValue) String() string {
	return dara.Prettify(s)
}

func (s DataResultsValue) GoString() string {
	return s.String()
}

func (s *DataResultsValue) GetRouteId() *string {
	return s.RouteId
}

func (s *DataResultsValue) GetStatus() *string {
	return s.Status
}

func (s *DataResultsValue) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DataResultsValue) SetRouteId(v string) *DataResultsValue {
	s.RouteId = &v
	return s
}

func (s *DataResultsValue) SetStatus(v string) *DataResultsValue {
	s.Status = &v
	return s
}

func (s *DataResultsValue) SetErrorMessage(v string) *DataResultsValue {
	s.ErrorMessage = &v
	return s
}

func (s *DataResultsValue) Validate() error {
	return dara.Validate(s)
}
