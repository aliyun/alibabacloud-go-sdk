// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewRectifyFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ReviewRectifyFlowShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ReviewRectifyFlowShrinkRequest
	GetOpUserId() *string
	SetReviewCommandShrink(v string) *ReviewRectifyFlowShrinkRequest
	GetReviewCommandShrink() *string
}

type ReviewRectifyFlowShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The business request parameters.
	//
	// This parameter is required.
	ReviewCommandShrink *string `json:"ReviewCommand,omitempty" xml:"ReviewCommand,omitempty"`
}

func (s ReviewRectifyFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReviewRectifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReviewRectifyFlowShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ReviewRectifyFlowShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ReviewRectifyFlowShrinkRequest) GetReviewCommandShrink() *string {
	return s.ReviewCommandShrink
}

func (s *ReviewRectifyFlowShrinkRequest) SetOpTenantId(v int64) *ReviewRectifyFlowShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ReviewRectifyFlowShrinkRequest) SetOpUserId(v string) *ReviewRectifyFlowShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ReviewRectifyFlowShrinkRequest) SetReviewCommandShrink(v string) *ReviewRectifyFlowShrinkRequest {
	s.ReviewCommandShrink = &v
	return s
}

func (s *ReviewRectifyFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
