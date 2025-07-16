// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *AddWorkspaceMembersShrinkRequest
	GetMembersShrink() *string
	SetTenantContextShrink(v string) *AddWorkspaceMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *AddWorkspaceMembersShrinkRequest
	GetWorkspaceId() *string
}

type AddWorkspaceMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *AddWorkspaceMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddWorkspaceMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *AddWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *AddWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *AddWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
