// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *UpdateContextDatabaseMemberRequest
	GetMemberId() *string
	SetRole(v string) *UpdateContextDatabaseMemberRequest
	GetRole() *string
	SetStatus(v string) *UpdateContextDatabaseMemberRequest
	GetStatus() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseMemberRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseMemberRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateContextDatabaseMemberRequest) GetRole() *string {
	return s.Role
}

func (s *UpdateContextDatabaseMemberRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseMemberRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseMemberRequest) SetMemberId(v string) *UpdateContextDatabaseMemberRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateContextDatabaseMemberRequest) SetRole(v string) *UpdateContextDatabaseMemberRequest {
	s.Role = &v
	return s
}

func (s *UpdateContextDatabaseMemberRequest) SetStatus(v string) *UpdateContextDatabaseMemberRequest {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseMemberRequest) SetWorkspaceId(v string) *UpdateContextDatabaseMemberRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseMemberRequest) Validate() error {
	return dara.Validate(s)
}
