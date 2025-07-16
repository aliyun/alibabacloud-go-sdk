// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *AddPermissionShrinkRequest
	GetDentryUuid() *string
	SetMembersShrink(v string) *AddPermissionShrinkRequest
	GetMembersShrink() *string
	SetOptionShrink(v string) *AddPermissionShrinkRequest
	GetOptionShrink() *string
	SetRoleId(v string) *AddPermissionShrinkRequest
	GetRoleId() *string
	SetTenantContextShrink(v string) *AddPermissionShrinkRequest
	GetTenantContextShrink() *string
}

type AddPermissionShrinkRequest struct {
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	OptionShrink  *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// This parameter is required.
	RoleId              *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s AddPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *AddPermissionShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *AddPermissionShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *AddPermissionShrinkRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AddPermissionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddPermissionShrinkRequest) SetDentryUuid(v string) *AddPermissionShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *AddPermissionShrinkRequest) SetMembersShrink(v string) *AddPermissionShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddPermissionShrinkRequest) SetOptionShrink(v string) *AddPermissionShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *AddPermissionShrinkRequest) SetRoleId(v string) *AddPermissionShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *AddPermissionShrinkRequest) SetTenantContextShrink(v string) *AddPermissionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
