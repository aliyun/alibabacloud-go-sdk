// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIOrderApprovalCommentSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetAIOrderApprovalCommentSSERequest
	GetOrderId() *int64
	SetSessionId(v string) *GetAIOrderApprovalCommentSSERequest
	GetSessionId() *string
}

type GetAIOrderApprovalCommentSSERequest struct {
	// example:
	//
	// 420****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 48363552-225c-4c93-aeab-ea9b9d064b96
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetAIOrderApprovalCommentSSERequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIOrderApprovalCommentSSERequest) GoString() string {
	return s.String()
}

func (s *GetAIOrderApprovalCommentSSERequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetAIOrderApprovalCommentSSERequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAIOrderApprovalCommentSSERequest) SetOrderId(v int64) *GetAIOrderApprovalCommentSSERequest {
	s.OrderId = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSERequest) SetSessionId(v string) *GetAIOrderApprovalCommentSSERequest {
	s.SessionId = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSERequest) Validate() error {
	return dara.Validate(s)
}
