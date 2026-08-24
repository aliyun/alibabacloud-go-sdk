// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeContextDatabaseApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v int64) *RevokeContextDatabaseApiKeyRequest
	GetKeyId() *int64
	SetMemberId(v string) *RevokeContextDatabaseApiKeyRequest
	GetMemberId() *string
	SetWorkspaceId(v string) *RevokeContextDatabaseApiKeyRequest
	GetWorkspaceId() *string
}

type RevokeContextDatabaseApiKeyRequest struct {
	// API Key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The member ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RevokeContextDatabaseApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeContextDatabaseApiKeyRequest) GoString() string {
	return s.String()
}

func (s *RevokeContextDatabaseApiKeyRequest) GetKeyId() *int64 {
	return s.KeyId
}

func (s *RevokeContextDatabaseApiKeyRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *RevokeContextDatabaseApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RevokeContextDatabaseApiKeyRequest) SetKeyId(v int64) *RevokeContextDatabaseApiKeyRequest {
	s.KeyId = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyRequest) SetMemberId(v string) *RevokeContextDatabaseApiKeyRequest {
	s.MemberId = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyRequest) SetWorkspaceId(v string) *RevokeContextDatabaseApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
