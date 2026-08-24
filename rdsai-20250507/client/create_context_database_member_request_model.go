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
	// Specifies whether to issue the first API key when the member is created. Default value: false.
	//
	// example:
	//
	// true
	GenerateInitialKey *bool `json:"GenerateInitialKey,omitempty" xml:"GenerateInitialKey,omitempty"`
	// The name of the first API key. This parameter takes effect only when GenerateInitialKey is set to true.
	//
	// example:
	//
	// my-key
	InitialKeyName *string `json:"InitialKeyName,omitempty" xml:"InitialKeyName,omitempty"`
	// The member name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The member role. Valid values:
	//
	// - owner
	//
	// - admin
	//
	// - member
	//
	// This parameter is required.
	//
	// example:
	//
	// member
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The workspace ID.
	//
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
