// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowOwnerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowOwnerResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowOwnerResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowOwnerResponseBody struct {
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
	// 482C61C1-2537-5BFB-8E58-34D9F17AD3C3
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

func (s UpdateTaskFlowOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowOwnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowOwnerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowOwnerResponseBody) SetErrorCode(v string) *UpdateTaskFlowOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowOwnerResponseBody) SetErrorMessage(v string) *UpdateTaskFlowOwnerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowOwnerResponseBody) SetRequestId(v string) *UpdateTaskFlowOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowOwnerResponseBody) SetSuccess(v bool) *UpdateTaskFlowOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
