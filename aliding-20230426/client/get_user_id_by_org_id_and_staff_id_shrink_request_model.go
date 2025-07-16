// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOrgIdAndStaffIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v int64) *GetUserIdByOrgIdAndStaffIdShrinkRequest
	GetOrgId() *int64
	SetTenantContextShrink(v string) *GetUserIdByOrgIdAndStaffIdShrinkRequest
	GetTenantContextShrink() *string
}

type GetUserIdByOrgIdAndStaffIdShrinkRequest struct {
	// example:
	//
	// 123456
	OrgId               *int64  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetUserIdByOrgIdAndStaffIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GetUserIdByOrgIdAndStaffIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserIdByOrgIdAndStaffIdShrinkRequest) SetOrgId(v int64) *GetUserIdByOrgIdAndStaffIdShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdShrinkRequest) SetTenantContextShrink(v string) *GetUserIdByOrgIdAndStaffIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
