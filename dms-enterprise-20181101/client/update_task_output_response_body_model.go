// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskOutputResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskOutputResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskOutputResponseBody
	GetSuccess() *bool
}

type UpdateTaskOutputResponseBody struct {
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
	// AB524768-8A5F-523A-91BD-1147187FCD62
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

func (s UpdateTaskOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskOutputResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskOutputResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskOutputResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskOutputResponseBody) SetErrorCode(v string) *UpdateTaskOutputResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskOutputResponseBody) SetErrorMessage(v string) *UpdateTaskOutputResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskOutputResponseBody) SetRequestId(v string) *UpdateTaskOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskOutputResponseBody) SetSuccess(v bool) *UpdateTaskOutputResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
