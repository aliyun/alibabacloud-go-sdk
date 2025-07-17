// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskNameResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskNameResponseBody
	GetSuccess() *bool
}

type UpdateTaskNameResponseBody struct {
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
	// C4604178-3BE1-5973-ACF0-7D561AEEF3A8
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

func (s UpdateTaskNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskNameResponseBody) SetErrorCode(v string) *UpdateTaskNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskNameResponseBody) SetErrorMessage(v string) *UpdateTaskNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskNameResponseBody) SetRequestId(v string) *UpdateTaskNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskNameResponseBody) SetSuccess(v bool) *UpdateTaskNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskNameResponseBody) Validate() error {
	return dara.Validate(s)
}
