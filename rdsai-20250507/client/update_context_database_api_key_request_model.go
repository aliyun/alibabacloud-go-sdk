// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateContextDatabaseApiKeyRequest
	GetDescription() *string
	SetKeyId(v int64) *UpdateContextDatabaseApiKeyRequest
	GetKeyId() *int64
	SetMemberId(v string) *UpdateContextDatabaseApiKeyRequest
	GetMemberId() *string
	SetName(v string) *UpdateContextDatabaseApiKeyRequest
	GetName() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseApiKeyRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseApiKeyRequest struct {
	// example:
	//
	// for nightly cron
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1024
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseApiKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateContextDatabaseApiKeyRequest) GetKeyId() *int64 {
	return s.KeyId
}

func (s *UpdateContextDatabaseApiKeyRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateContextDatabaseApiKeyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateContextDatabaseApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseApiKeyRequest) SetDescription(v string) *UpdateContextDatabaseApiKeyRequest {
	s.Description = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyRequest) SetKeyId(v int64) *UpdateContextDatabaseApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyRequest) SetMemberId(v string) *UpdateContextDatabaseApiKeyRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyRequest) SetName(v string) *UpdateContextDatabaseApiKeyRequest {
	s.Name = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyRequest) SetWorkspaceId(v string) *UpdateContextDatabaseApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
