// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*AddWorkspaceDocMembersRequestMembers) *AddWorkspaceDocMembersRequest
	GetMembers() []*AddWorkspaceDocMembersRequestMembers
	SetNodeId(v string) *AddWorkspaceDocMembersRequest
	GetNodeId() *string
	SetTenantContext(v *AddWorkspaceDocMembersRequestTenantContext) *AddWorkspaceDocMembersRequest
	GetTenantContext() *AddWorkspaceDocMembersRequestTenantContext
	SetWorkspaceId(v string) *AddWorkspaceDocMembersRequest
	GetWorkspaceId() *string
}

type AddWorkspaceDocMembersRequest struct {
	// This parameter is required.
	Members []*AddWorkspaceDocMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	NodeId        *string                                     `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *AddWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceDocMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequest) GetMembers() []*AddWorkspaceDocMembersRequestMembers {
	return s.Members
}

func (s *AddWorkspaceDocMembersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AddWorkspaceDocMembersRequest) GetTenantContext() *AddWorkspaceDocMembersRequestTenantContext {
	return s.TenantContext
}

func (s *AddWorkspaceDocMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceDocMembersRequest) SetMembers(v []*AddWorkspaceDocMembersRequestMembers) *AddWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetNodeId(v string) *AddWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetTenantContext(v *AddWorkspaceDocMembersRequestTenantContext) *AddWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetWorkspaceId(v string) *AddWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequest) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceDocMembersRequestMembers struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USER
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EDITOR
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AddWorkspaceDocMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *AddWorkspaceDocMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *AddWorkspaceDocMembersRequestMembers) GetRoleType() *string {
	return s.RoleType
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberId(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetRoleType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceDocMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddWorkspaceDocMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *AddWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
