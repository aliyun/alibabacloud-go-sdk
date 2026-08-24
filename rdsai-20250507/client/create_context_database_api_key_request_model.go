// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberId(v string) *CreateContextDatabaseApiKeyRequest
	GetMemberId() *string
	SetName(v string) *CreateContextDatabaseApiKeyRequest
	GetName() *string
	SetWorkspaceId(v string) *CreateContextDatabaseApiKeyRequest
	GetWorkspaceId() *string
}

type CreateContextDatabaseApiKeyRequest struct {
	// The member ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The API key name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateContextDatabaseApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseApiKeyRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateContextDatabaseApiKeyRequest) GetName() *string {
	return s.Name
}

func (s *CreateContextDatabaseApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseApiKeyRequest) SetMemberId(v string) *CreateContextDatabaseApiKeyRequest {
	s.MemberId = &v
	return s
}

func (s *CreateContextDatabaseApiKeyRequest) SetName(v string) *CreateContextDatabaseApiKeyRequest {
	s.Name = &v
	return s
}

func (s *CreateContextDatabaseApiKeyRequest) SetWorkspaceId(v string) *CreateContextDatabaseApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
