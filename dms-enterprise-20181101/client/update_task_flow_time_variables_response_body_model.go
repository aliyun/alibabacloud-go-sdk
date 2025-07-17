// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowTimeVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowTimeVariablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowTimeVariablesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowTimeVariablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowTimeVariablesResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowTimeVariablesResponseBody struct {
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

func (s UpdateTaskFlowTimeVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowTimeVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) SetErrorCode(v string) *UpdateTaskFlowTimeVariablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) SetErrorMessage(v string) *UpdateTaskFlowTimeVariablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) SetRequestId(v string) *UpdateTaskFlowTimeVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) SetSuccess(v bool) *UpdateTaskFlowTimeVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
