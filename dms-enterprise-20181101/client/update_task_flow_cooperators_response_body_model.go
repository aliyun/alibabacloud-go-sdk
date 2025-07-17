// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowCooperatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowCooperatorsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowCooperatorsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowCooperatorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowCooperatorsResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowCooperatorsResponseBody struct {
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
	// D05B3EE1-B6D3-5B17-8CA6-A8054828E5B2
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

func (s UpdateTaskFlowCooperatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowCooperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowCooperatorsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowCooperatorsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowCooperatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowCooperatorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowCooperatorsResponseBody) SetErrorCode(v string) *UpdateTaskFlowCooperatorsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponseBody) SetErrorMessage(v string) *UpdateTaskFlowCooperatorsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponseBody) SetRequestId(v string) *UpdateTaskFlowCooperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponseBody) SetSuccess(v bool) *UpdateTaskFlowCooperatorsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsResponseBody) Validate() error {
	return dara.Validate(s)
}
