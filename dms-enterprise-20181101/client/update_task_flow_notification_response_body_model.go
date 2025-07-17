// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowNotificationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowNotificationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowNotificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowNotificationResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowNotificationResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 3BDC762F-2525-5E47-8748-D6C58BDB3B38
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

func (s UpdateTaskFlowNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNotificationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowNotificationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowNotificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowNotificationResponseBody) SetErrorCode(v string) *UpdateTaskFlowNotificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowNotificationResponseBody) SetErrorMessage(v string) *UpdateTaskFlowNotificationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowNotificationResponseBody) SetRequestId(v string) *UpdateTaskFlowNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowNotificationResponseBody) SetSuccess(v bool) *UpdateTaskFlowNotificationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
