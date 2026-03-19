// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIOrderApprovalCommentSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetData() *string
	SetErrorCode(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAIOrderApprovalCommentSSEResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAIOrderApprovalCommentSSEResponseBody
	GetSuccess() *bool
}

type GetAIOrderApprovalCommentSSEResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAIOrderApprovalCommentSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIOrderApprovalCommentSSEResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetData(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.Data = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetErrorCode(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetErrorMessage(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetRequestId(v string) *GetAIOrderApprovalCommentSSEResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) SetSuccess(v bool) *GetAIOrderApprovalCommentSSEResponseBody {
	s.Success = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponseBody) Validate() error {
	return dara.Validate(s)
}
