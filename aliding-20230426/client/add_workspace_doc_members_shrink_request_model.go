// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *AddWorkspaceDocMembersShrinkRequest
	GetMembersShrink() *string
	SetNodeId(v string) *AddWorkspaceDocMembersShrinkRequest
	GetNodeId() *string
	SetTenantContextShrink(v string) *AddWorkspaceDocMembersShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceId(v string) *AddWorkspaceDocMembersShrinkRequest
	GetWorkspaceId() *string
}

type AddWorkspaceDocMembersShrinkRequest struct {
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceDocMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *AddWorkspaceDocMembersShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AddWorkspaceDocMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddWorkspaceDocMembersShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
