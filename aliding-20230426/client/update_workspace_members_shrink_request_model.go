// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *UpdateWorkspaceMembersShrinkRequest
	GetMembersShrink() *string
	SetTenantContextShrink(v string) *UpdateWorkspaceMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *UpdateWorkspaceMembersShrinkRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceMembersShrinkRequest struct {
	// This parameter is required.
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *UpdateWorkspaceMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateWorkspaceMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
