// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *DeletePermissionShrinkRequest
	GetDentryUuid() *string
	SetMembersShrink(v string) *DeletePermissionShrinkRequest
	GetMembersShrink() *string
	SetRoleId(v string) *DeletePermissionShrinkRequest
	GetRoleId() *string
	SetTenantContextShrink(v string) *DeletePermissionShrinkRequest
	GetTenantContextShrink() *string
}

type DeletePermissionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a9E05BDRVQRkezKGCE3nlwPDJ63zgkYA
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MANAGER
	RoleId              *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeletePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DeletePermissionShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *DeletePermissionShrinkRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeletePermissionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeletePermissionShrinkRequest) SetDentryUuid(v string) *DeletePermissionShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *DeletePermissionShrinkRequest) SetMembersShrink(v string) *DeletePermissionShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *DeletePermissionShrinkRequest) SetRoleId(v string) *DeletePermissionShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *DeletePermissionShrinkRequest) SetTenantContextShrink(v string) *DeletePermissionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeletePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
