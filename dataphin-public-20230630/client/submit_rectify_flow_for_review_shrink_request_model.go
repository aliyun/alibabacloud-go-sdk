// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRectifyFlowForReviewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitRectifyFlowForReviewShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitRectifyFlowForReviewShrinkRequest
	GetOpUserId() *string
	SetSubmitCommandShrink(v string) *SubmitRectifyFlowForReviewShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitRectifyFlowForReviewShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The business request parameters.
	//
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitRectifyFlowForReviewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitRectifyFlowForReviewShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) SetOpTenantId(v int64) *SubmitRectifyFlowForReviewShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) SetOpUserId(v string) *SubmitRectifyFlowForReviewShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) SetSubmitCommandShrink(v string) *SubmitRectifyFlowForReviewShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitRectifyFlowForReviewShrinkRequest) Validate() error {
	return dara.Validate(s)
}
