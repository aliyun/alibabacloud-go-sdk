// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowRelationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowRelationsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowRelationsResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowRelationsResponseBody struct {
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
	// 15D9E71C-405B-57D7-BE6E-707C2C7A8E0B
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

func (s UpdateTaskFlowRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowRelationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowRelationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowRelationsResponseBody) SetErrorCode(v string) *UpdateTaskFlowRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowRelationsResponseBody) SetErrorMessage(v string) *UpdateTaskFlowRelationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowRelationsResponseBody) SetRequestId(v string) *UpdateTaskFlowRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowRelationsResponseBody) SetSuccess(v bool) *UpdateTaskFlowRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}
