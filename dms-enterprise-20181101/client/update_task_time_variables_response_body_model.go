// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskTimeVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskTimeVariablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskTimeVariablesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskTimeVariablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskTimeVariablesResponseBody
	GetSuccess() *bool
}

type UpdateTaskTimeVariablesResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 39557312-28D5-528F-9554-80C0700EB489
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

func (s UpdateTaskTimeVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskTimeVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskTimeVariablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskTimeVariablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskTimeVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskTimeVariablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskTimeVariablesResponseBody) SetErrorCode(v string) *UpdateTaskTimeVariablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskTimeVariablesResponseBody) SetErrorMessage(v string) *UpdateTaskTimeVariablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskTimeVariablesResponseBody) SetRequestId(v string) *UpdateTaskTimeVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskTimeVariablesResponseBody) SetSuccess(v bool) *UpdateTaskTimeVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskTimeVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
