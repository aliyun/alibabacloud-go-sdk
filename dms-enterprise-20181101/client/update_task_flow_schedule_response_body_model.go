// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowScheduleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowScheduleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowScheduleResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowScheduleResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 7BF38A13-C181-5B5E-97F1-8643F8A10093
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskFlowScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowScheduleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowScheduleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowScheduleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowScheduleResponseBody) SetErrorCode(v string) *UpdateTaskFlowScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowScheduleResponseBody) SetErrorMessage(v string) *UpdateTaskFlowScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowScheduleResponseBody) SetRequestId(v string) *UpdateTaskFlowScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowScheduleResponseBody) SetSuccess(v bool) *UpdateTaskFlowScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
