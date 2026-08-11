// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetApiKey() *string
	SetApiKeyName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetApiKeyName() *string
	SetCreatedAt(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetMemberId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetMemberId() *string
	SetMemberName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetMemberName() *string
	SetRequestId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetRole(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetRole() *string
	SetStatus(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetWorkspaceId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type CreateContextDatabaseWorkspaceResponseBody struct {
	// example:
	//
	// ctxdb-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// default
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// my-member
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// owner
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetRole() *string {
	return s.Role
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetApiKey(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetApiKeyName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ApiKeyName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetMemberId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.MemberId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetMemberName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetRole(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Role = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetStatus(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetType(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
