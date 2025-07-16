// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest
	GetMembersShrink() *string
	SetNodeId(v string) *UpdateWorkspaceDocMembersShrinkRequest
	GetNodeId() *string
	SetTenantContextShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *UpdateWorkspaceDocMembersShrinkRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceDocMembersShrinkRequest struct {
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// node_feb8fea0
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xb8bkxxxxxrXJNaL
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceDocMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
