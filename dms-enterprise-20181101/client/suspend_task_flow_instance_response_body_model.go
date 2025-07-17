// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskFlowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SuspendTaskFlowInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SuspendTaskFlowInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SuspendTaskFlowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendTaskFlowInstanceResponseBody
	GetSuccess() *bool
}

type SuspendTaskFlowInstanceResponseBody struct {
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
	// 028BF827-3801-5869-8548-F4A039256305
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

func (s SuspendTaskFlowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendTaskFlowInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SuspendTaskFlowInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SuspendTaskFlowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendTaskFlowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendTaskFlowInstanceResponseBody) SetErrorCode(v string) *SuspendTaskFlowInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SuspendTaskFlowInstanceResponseBody) SetErrorMessage(v string) *SuspendTaskFlowInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SuspendTaskFlowInstanceResponseBody) SetRequestId(v string) *SuspendTaskFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendTaskFlowInstanceResponseBody) SetSuccess(v bool) *SuspendTaskFlowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendTaskFlowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
