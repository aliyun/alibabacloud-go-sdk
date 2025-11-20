// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*DeleteWorkspaceMembersRequestMembers) *DeleteWorkspaceMembersRequest
	GetMembers() []*DeleteWorkspaceMembersRequestMembers
	SetTenantContext(v *DeleteWorkspaceMembersRequestTenantContext) *DeleteWorkspaceMembersRequest
	GetTenantContext() *DeleteWorkspaceMembersRequestTenantContext
	SetWorkspaceId(v string) *DeleteWorkspaceMembersRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceMembersRequest struct {
	// This parameter is required.
	Members       []*DeleteWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *DeleteWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequest) GetMembers() []*DeleteWorkspaceMembersRequestMembers {
	return s.Members
}

func (s *DeleteWorkspaceMembersRequest) GetTenantContext() *DeleteWorkspaceMembersRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteWorkspaceMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceMembersRequest) SetMembers(v []*DeleteWorkspaceMembersRequestMembers) *DeleteWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetTenantContext(v *DeleteWorkspaceMembersRequestTenantContext) *DeleteWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetWorkspaceId(v string) *DeleteWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequest) Validate() error {
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

type DeleteWorkspaceMembersRequestMembers struct {
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
}

func (s DeleteWorkspaceMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *DeleteWorkspaceMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *DeleteWorkspaceMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type DeleteWorkspaceMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteWorkspaceMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkspaceMembersRequestTenantContext) SetTenantId(v string) *DeleteWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
