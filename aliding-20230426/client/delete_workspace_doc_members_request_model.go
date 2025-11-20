// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*DeleteWorkspaceDocMembersRequestMembers) *DeleteWorkspaceDocMembersRequest
	GetMembers() []*DeleteWorkspaceDocMembersRequestMembers
	SetNodeId(v string) *DeleteWorkspaceDocMembersRequest
	GetNodeId() *string
	SetTenantContext(v *DeleteWorkspaceDocMembersRequestTenantContext) *DeleteWorkspaceDocMembersRequest
	GetTenantContext() *DeleteWorkspaceDocMembersRequestTenantContext
	SetWorkspaceId(v string) *DeleteWorkspaceDocMembersRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceDocMembersRequest struct {
	// This parameter is required.
	Members []*DeleteWorkspaceDocMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// YRBGv0xxx
	NodeId        *string                                        `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *DeleteWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// YRBGvyxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequest) GetMembers() []*DeleteWorkspaceDocMembersRequestMembers {
	return s.Members
}

func (s *DeleteWorkspaceDocMembersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteWorkspaceDocMembersRequest) GetTenantContext() *DeleteWorkspaceDocMembersRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteWorkspaceDocMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceDocMembersRequest) SetMembers(v []*DeleteWorkspaceDocMembersRequestMembers) *DeleteWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetNodeId(v string) *DeleteWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetTenantContext(v *DeleteWorkspaceDocMembersRequestTenantContext) *DeleteWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetWorkspaceId(v string) *DeleteWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) Validate() error {
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

type DeleteWorkspaceDocMembersRequestMembers struct {
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

func (s DeleteWorkspaceDocMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *DeleteWorkspaceDocMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}

type DeleteWorkspaceDocMembersRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *DeleteWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
