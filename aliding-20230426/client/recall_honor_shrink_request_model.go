// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *RecallHonorShrinkRequest
	GetTenantContextShrink() *string
	SetHonorId(v string) *RecallHonorShrinkRequest
	GetHonorId() *string
	SetOrgId(v int64) *RecallHonorShrinkRequest
	GetOrgId() *int64
	SetUserId(v string) *RecallHonorShrinkRequest
	GetUserId() *string
}

type RecallHonorShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21660610
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 345391052
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RecallHonorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecallHonorShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *RecallHonorShrinkRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *RecallHonorShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *RecallHonorShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *RecallHonorShrinkRequest) SetTenantContextShrink(v string) *RecallHonorShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *RecallHonorShrinkRequest) SetHonorId(v string) *RecallHonorShrinkRequest {
	s.HonorId = &v
	return s
}

func (s *RecallHonorShrinkRequest) SetOrgId(v int64) *RecallHonorShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *RecallHonorShrinkRequest) SetUserId(v string) *RecallHonorShrinkRequest {
	s.UserId = &v
	return s
}

func (s *RecallHonorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
