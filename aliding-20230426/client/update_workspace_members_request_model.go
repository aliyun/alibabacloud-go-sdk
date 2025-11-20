// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*UpdateWorkspaceMembersRequestMembers) *UpdateWorkspaceMembersRequest
	GetMembers() []*UpdateWorkspaceMembersRequestMembers
	SetTenantContext(v *UpdateWorkspaceMembersRequestTenantContext) *UpdateWorkspaceMembersRequest
	GetTenantContext() *UpdateWorkspaceMembersRequestTenantContext
	SetWorkspaceId(v string) *UpdateWorkspaceMembersRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceMembersRequest struct {
	// This parameter is required.
	Members       []*UpdateWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *UpdateWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequest) GetMembers() []*UpdateWorkspaceMembersRequestMembers {
	return s.Members
}

func (s *UpdateWorkspaceMembersRequest) GetTenantContext() *UpdateWorkspaceMembersRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateWorkspaceMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceMembersRequest) SetMembers(v []*UpdateWorkspaceMembersRequestMembers) *UpdateWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetTenantContext(v *UpdateWorkspaceMembersRequestTenantContext) *UpdateWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetWorkspaceId(v string) *UpdateWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequest) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceMembersRequestMembers struct {
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
	// ONLY_VIEWER
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateWorkspaceMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateWorkspaceMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateWorkspaceMembersRequestMembers) GetRoleType() *string {
	return s.RoleType
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkspaceMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkspaceMembersRequestTenantContext) SetTenantId(v string) *UpdateWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
