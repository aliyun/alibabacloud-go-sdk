// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCancelIgnoreSuspEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperationCancelIgnoreSuspEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OperationCancelIgnoreSuspEventResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OperationCancelIgnoreSuspEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperationCancelIgnoreSuspEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperationCancelIgnoreSuspEventResponseBody
	GetSuccess() *bool
	SetTimeCost(v int64) *OperationCancelIgnoreSuspEventResponseBody
	GetTimeCost() *int64
}

type OperationCancelIgnoreSuspEventResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was is successful. Other status codes indicate that the request fails. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The handling result of an exception. Valid values:
	//
	// 	- **true**: successful
	//
	// 	- **false**: failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time consumed for the request. Unit: seconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s OperationCancelIgnoreSuspEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperationCancelIgnoreSuspEventResponseBody) GoString() string {
	return s.String()
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperationCancelIgnoreSuspEventResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetCode(v string) *OperationCancelIgnoreSuspEventResponseBody {
	s.Code = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetHttpStatusCode(v int32) *OperationCancelIgnoreSuspEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetMessage(v string) *OperationCancelIgnoreSuspEventResponseBody {
	s.Message = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetRequestId(v string) *OperationCancelIgnoreSuspEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetSuccess(v bool) *OperationCancelIgnoreSuspEventResponseBody {
	s.Success = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) SetTimeCost(v int64) *OperationCancelIgnoreSuspEventResponseBody {
	s.TimeCost = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponseBody) Validate() error {
	return dara.Validate(s)
}
