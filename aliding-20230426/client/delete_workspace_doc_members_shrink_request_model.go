// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest
	GetMembersShrink() *string
	SetNodeId(v string) *DeleteWorkspaceDocMembersShrinkRequest
	GetNodeId() *string
	SetTenantContextShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *DeleteWorkspaceDocMembersShrinkRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceDocMembersShrinkRequest struct {
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YRBGv0xxx
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// YRBGvyxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceDocMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
