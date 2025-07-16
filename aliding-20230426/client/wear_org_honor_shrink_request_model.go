// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *WearOrgHonorShrinkRequest
	GetTenantContextShrink() *string
	SetHonorId(v string) *WearOrgHonorShrinkRequest
	GetHonorId() *string
	SetOrgId(v int64) *WearOrgHonorShrinkRequest
	GetOrgId() *int64
	SetUserId(v string) *WearOrgHonorShrinkRequest
	GetUserId() *string
	SetWear(v bool) *WearOrgHonorShrinkRequest
	GetWear() *bool
}

type WearOrgHonorShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21659595
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Wear *bool `json:"wear,omitempty" xml:"wear,omitempty"`
}

func (s WearOrgHonorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorShrinkRequest) GoString() string {
	return s.String()
}

func (s *WearOrgHonorShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *WearOrgHonorShrinkRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *WearOrgHonorShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *WearOrgHonorShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *WearOrgHonorShrinkRequest) GetWear() *bool {
	return s.Wear
}

func (s *WearOrgHonorShrinkRequest) SetTenantContextShrink(v string) *WearOrgHonorShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *WearOrgHonorShrinkRequest) SetHonorId(v string) *WearOrgHonorShrinkRequest {
	s.HonorId = &v
	return s
}

func (s *WearOrgHonorShrinkRequest) SetOrgId(v int64) *WearOrgHonorShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *WearOrgHonorShrinkRequest) SetUserId(v string) *WearOrgHonorShrinkRequest {
	s.UserId = &v
	return s
}

func (s *WearOrgHonorShrinkRequest) SetWear(v bool) *WearOrgHonorShrinkRequest {
	s.Wear = &v
	return s
}

func (s *WearOrgHonorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
