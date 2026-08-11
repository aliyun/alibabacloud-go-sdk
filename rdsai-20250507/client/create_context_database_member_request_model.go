// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerateInitialKey(v bool) *CreateContextDatabaseMemberRequest
	GetGenerateInitialKey() *bool
	SetInitialKeyName(v string) *CreateContextDatabaseMemberRequest
	GetInitialKeyName() *string
	SetMemberName(v string) *CreateContextDatabaseMemberRequest
	GetMemberName() *string
	SetRole(v string) *CreateContextDatabaseMemberRequest
	GetRole() *string
	SetWorkspaceId(v string) *CreateContextDatabaseMemberRequest
	GetWorkspaceId() *string
}

type CreateContextDatabaseMemberRequest struct {
	// example:
	//
	// true
	GenerateInitialKey *bool `json:"GenerateInitialKey,omitempty" xml:"GenerateInitialKey,omitempty"`
	// example:
	//
	// my-key
	InitialKeyName *string `json:"InitialKeyName,omitempty" xml:"InitialKeyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Alice
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// member
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateContextDatabaseMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseMemberRequest) GetGenerateInitialKey() *bool {
	return s.GenerateInitialKey
}

func (s *CreateContextDatabaseMemberRequest) GetInitialKeyName() *string {
	return s.InitialKeyName
}

func (s *CreateContextDatabaseMemberRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseMemberRequest) GetRole() *string {
	return s.Role
}

func (s *CreateContextDatabaseMemberRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseMemberRequest) SetGenerateInitialKey(v bool) *CreateContextDatabaseMemberRequest {
	s.GenerateInitialKey = &v
	return s
}

func (s *CreateContextDatabaseMemberRequest) SetInitialKeyName(v string) *CreateContextDatabaseMemberRequest {
	s.InitialKeyName = &v
	return s
}

func (s *CreateContextDatabaseMemberRequest) SetMemberName(v string) *CreateContextDatabaseMemberRequest {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseMemberRequest) SetRole(v string) *CreateContextDatabaseMemberRequest {
	s.Role = &v
	return s
}

func (s *CreateContextDatabaseMemberRequest) SetWorkspaceId(v string) *CreateContextDatabaseMemberRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseMemberRequest) Validate() error {
	return dara.Validate(s)
}
