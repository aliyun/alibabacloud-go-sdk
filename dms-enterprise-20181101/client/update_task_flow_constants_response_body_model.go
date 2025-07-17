// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowConstantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskFlowConstantsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskFlowConstantsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskFlowConstantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskFlowConstantsResponseBody
	GetSuccess() *bool
}

type UpdateTaskFlowConstantsResponseBody struct {
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
	// C4CCC000-C193-5A32-B701-573F497BF729
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

func (s UpdateTaskFlowConstantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowConstantsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowConstantsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskFlowConstantsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskFlowConstantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskFlowConstantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskFlowConstantsResponseBody) SetErrorCode(v string) *UpdateTaskFlowConstantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskFlowConstantsResponseBody) SetErrorMessage(v string) *UpdateTaskFlowConstantsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskFlowConstantsResponseBody) SetRequestId(v string) *UpdateTaskFlowConstantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskFlowConstantsResponseBody) SetSuccess(v bool) *UpdateTaskFlowConstantsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskFlowConstantsResponseBody) Validate() error {
	return dara.Validate(s)
}
