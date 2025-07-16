// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *DeleteWorkspaceMembersShrinkRequest
	GetMembersShrink() *string
	SetTenantContextShrink(v string) *DeleteWorkspaceMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *DeleteWorkspaceMembersShrinkRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceMembersShrinkRequest struct {
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

func (s DeleteWorkspaceMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *DeleteWorkspaceMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteWorkspaceMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
