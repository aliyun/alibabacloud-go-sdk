// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOrderApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitOrderApprovalResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SubmitOrderApprovalResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SubmitOrderApprovalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitOrderApprovalResponseBody
	GetSuccess() *bool
}

type SubmitOrderApprovalResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitOrderApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitOrderApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitOrderApprovalResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitOrderApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitOrderApprovalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitOrderApprovalResponseBody) SetErrorCode(v string) *SubmitOrderApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetErrorMessage(v string) *SubmitOrderApprovalResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetRequestId(v string) *SubmitOrderApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetSuccess(v bool) *SubmitOrderApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) Validate() error {
	return dara.Validate(s)
}
