// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*UpdateWorkspaceDocMembersRequestMembers) *UpdateWorkspaceDocMembersRequest
	GetMembers() []*UpdateWorkspaceDocMembersRequestMembers
	SetNodeId(v string) *UpdateWorkspaceDocMembersRequest
	GetNodeId() *string
	SetTenantContext(v *UpdateWorkspaceDocMembersRequestTenantContext) *UpdateWorkspaceDocMembersRequest
	GetTenantContext() *UpdateWorkspaceDocMembersRequestTenantContext
	SetWorkspaceId(v string) *UpdateWorkspaceDocMembersRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceDocMembersRequest struct {
	Members []*UpdateWorkspaceDocMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// node_feb8fea0
	NodeId        *string                                        `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *UpdateWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// xb8bkxxxxxrXJNaL
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequest) GetMembers() []*UpdateWorkspaceDocMembersRequestMembers {
	return s.Members
}

func (s *UpdateWorkspaceDocMembersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateWorkspaceDocMembersRequest) GetTenantContext() *UpdateWorkspaceDocMembersRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateWorkspaceDocMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceDocMembersRequest) SetMembers(v []*UpdateWorkspaceDocMembersRequestMembers) *UpdateWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetNodeId(v string) *UpdateWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetTenantContext(v *UpdateWorkspaceDocMembersRequestTenantContext) *UpdateWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetWorkspaceId(v string) *UpdateWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceDocMembersRequestMembers struct {
	// example:
	//
	// 012345
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// USER
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// example:
	//
	// ONLY_VIEWER
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateWorkspaceDocMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateWorkspaceDocMembersRequestMembers) GetRoleType() *string {
	return s.RoleType
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceDocMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *UpdateWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
