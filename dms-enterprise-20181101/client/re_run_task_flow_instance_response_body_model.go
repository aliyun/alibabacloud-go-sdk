// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReRunTaskFlowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ReRunTaskFlowInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReRunTaskFlowInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReRunTaskFlowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReRunTaskFlowInstanceResponseBody
	GetSuccess() *bool
}

type ReRunTaskFlowInstanceResponseBody struct {
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
	// 8CFF2295-8249-5287-B888-DBD4F0D76CB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReRunTaskFlowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReRunTaskFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReRunTaskFlowInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReRunTaskFlowInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReRunTaskFlowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReRunTaskFlowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReRunTaskFlowInstanceResponseBody) SetErrorCode(v string) *ReRunTaskFlowInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReRunTaskFlowInstanceResponseBody) SetErrorMessage(v string) *ReRunTaskFlowInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReRunTaskFlowInstanceResponseBody) SetRequestId(v string) *ReRunTaskFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReRunTaskFlowInstanceResponseBody) SetSuccess(v bool) *ReRunTaskFlowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReRunTaskFlowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
