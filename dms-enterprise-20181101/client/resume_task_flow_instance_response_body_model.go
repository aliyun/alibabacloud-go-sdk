// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskFlowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ResumeTaskFlowInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ResumeTaskFlowInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ResumeTaskFlowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeTaskFlowInstanceResponseBody
	GetSuccess() *bool
}

type ResumeTaskFlowInstanceResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 7FF2E325-763F-5E27-9157-C3CFA02F4CBF
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

func (s ResumeTaskFlowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTaskFlowInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResumeTaskFlowInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResumeTaskFlowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeTaskFlowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeTaskFlowInstanceResponseBody) SetErrorCode(v string) *ResumeTaskFlowInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeTaskFlowInstanceResponseBody) SetErrorMessage(v string) *ResumeTaskFlowInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeTaskFlowInstanceResponseBody) SetRequestId(v string) *ResumeTaskFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeTaskFlowInstanceResponseBody) SetSuccess(v bool) *ResumeTaskFlowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeTaskFlowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
