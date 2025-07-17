// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowEdgesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowEdgesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowEdgesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowEdgesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowEdgesResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowEdgesResponseBody struct {
	// The error code.
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
	// The request ID.
	//
	// example:
	//
	// 93FC1AE1-EC54-52B1-B146-650180FB82E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskFlowEdgesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowEdgesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowEdgesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowEdgesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowEdgesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowEdgesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowEdgesResponseBody) SetErrorCode(v string) *UpdateTaskFlowEdgesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowEdgesResponseBody) SetErrorMessage(v string) *UpdateTaskFlowEdgesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowEdgesResponseBody) SetRequestId(v string) *UpdateTaskFlowEdgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowEdgesResponseBody) SetSuccess(v bool) *UpdateTaskFlowEdgesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowEdgesResponseBody) Validate() error {
	return dara.Validate(s)
}
