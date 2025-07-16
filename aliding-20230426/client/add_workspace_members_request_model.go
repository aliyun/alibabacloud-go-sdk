// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*AddWorkspaceMembersRequestMembers) *AddWorkspaceMembersRequest
	GetMembers() []*AddWorkspaceMembersRequestMembers
	SetTenantContext(v *AddWorkspaceMembersRequestTenantContext) *AddWorkspaceMembersRequest
	GetTenantContext() *AddWorkspaceMembersRequestTenantContext
	SetWorkspaceId(v string) *AddWorkspaceMembersRequest
	GetWorkspaceId() *string
}

type AddWorkspaceMembersRequest struct {
	Members       []*AddWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *AddWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequest) GetMembers() []*AddWorkspaceMembersRequestMembers {
	return s.Members
}

func (s *AddWorkspaceMembersRequest) GetTenantContext() *AddWorkspaceMembersRequestTenantContext {
	return s.TenantContext
}

func (s *AddWorkspaceMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceMembersRequest) SetMembers(v []*AddWorkspaceMembersRequestMembers) *AddWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceMembersRequest) SetTenantContext(v *AddWorkspaceMembersRequestTenantContext) *AddWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *AddWorkspaceMembersRequest) SetWorkspaceId(v string) *AddWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceMembersRequest) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceMembersRequestMembers struct {
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
	// EDITOR
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AddWorkspaceMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *AddWorkspaceMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *AddWorkspaceMembersRequestMembers) GetRoleType() *string {
	return s.RoleType
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberId(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberType(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetRoleType(v string) *AddWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddWorkspaceMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddWorkspaceMembersRequestTenantContext) SetTenantId(v string) *AddWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddWorkspaceMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
