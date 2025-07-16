// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *UpdatePermissionShrinkRequest
	GetDentryUuid() *string
	SetMembersShrink(v string) *UpdatePermissionShrinkRequest
	GetMembersShrink() *string
	SetOptionShrink(v string) *UpdatePermissionShrinkRequest
	GetOptionShrink() *string
	SetRoleId(v string) *UpdatePermissionShrinkRequest
	GetRoleId() *string
	SetTenantContextShrink(v string) *UpdatePermissionShrinkRequest
	GetTenantContextShrink() *string
}

type UpdatePermissionShrinkRequest struct {
	// example:
	//
	// kDnRL6jAJMLgNkw7tBnw5aY4VyMoPYe1
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	OptionShrink  *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// READER
	RoleId              *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdatePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *UpdatePermissionShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *UpdatePermissionShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *UpdatePermissionShrinkRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *UpdatePermissionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdatePermissionShrinkRequest) SetDentryUuid(v string) *UpdatePermissionShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *UpdatePermissionShrinkRequest) SetMembersShrink(v string) *UpdatePermissionShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdatePermissionShrinkRequest) SetOptionShrink(v string) *UpdatePermissionShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *UpdatePermissionShrinkRequest) SetRoleId(v string) *UpdatePermissionShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *UpdatePermissionShrinkRequest) SetTenantContextShrink(v string) *UpdatePermissionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdatePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
