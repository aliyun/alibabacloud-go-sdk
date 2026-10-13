// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionCredentialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest
	GetClientToken() *string
	SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest
	GetConnectionId() *string
	SetCredentialsShrink(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest
	GetCredentialsShrink() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncConnectionCredentialShrinkRequest struct {
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"appKey":"dingas1llqmkol****","appSecret":"******"}
	CredentialsShrink *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) GetCredentialsShrink() *string {
	return s.CredentialsShrink
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) SetClientToken(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest {
	s.ConnectionId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) SetCredentialsShrink(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest {
	s.CredentialsShrink = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
