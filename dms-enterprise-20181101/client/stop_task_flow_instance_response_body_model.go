// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskFlowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopTaskFlowInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopTaskFlowInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopTaskFlowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopTaskFlowInstanceResponseBody
	GetSuccess() *bool
}

type StopTaskFlowInstanceResponseBody struct {
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
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 028BF827-3801-5869-8548-F4A039256308
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

func (s StopTaskFlowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTaskFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskFlowInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopTaskFlowInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopTaskFlowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTaskFlowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopTaskFlowInstanceResponseBody) SetErrorCode(v string) *StopTaskFlowInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopTaskFlowInstanceResponseBody) SetErrorMessage(v string) *StopTaskFlowInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopTaskFlowInstanceResponseBody) SetRequestId(v string) *StopTaskFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskFlowInstanceResponseBody) SetSuccess(v bool) *StopTaskFlowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopTaskFlowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
