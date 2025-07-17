// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNameAndDescResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowNameAndDescResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowNameAndDescResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowNameAndDescResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowNameAndDescResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowNameAndDescResponseBody struct {
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
	// F73CCB9D-0CF3-5D3D-97B0-D852A8022663
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

func (s UpdateTaskFlowNameAndDescResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNameAndDescResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNameAndDescResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowNameAndDescResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowNameAndDescResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowNameAndDescResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowNameAndDescResponseBody) SetErrorCode(v string) *UpdateTaskFlowNameAndDescResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponseBody) SetErrorMessage(v string) *UpdateTaskFlowNameAndDescResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponseBody) SetRequestId(v string) *UpdateTaskFlowNameAndDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponseBody) SetSuccess(v bool) *UpdateTaskFlowNameAndDescResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowNameAndDescResponseBody) Validate() error {
	return dara.Validate(s)
}
