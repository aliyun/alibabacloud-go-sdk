// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest
	GetClientToken() *string
	SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest
	GetConnectionId() *string
	SetCredentials(v map[string]interface{}) *UpdateContextDatabaseSyncConnectionCredentialRequest
	GetCredentials() map[string]interface{}
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncConnectionCredentialRequest struct {
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
	Credentials map[string]interface{} `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) GetCredentials() map[string]interface{} {
	return s.Credentials
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) SetClientToken(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest {
	s.ConnectionId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) SetCredentials(v map[string]interface{}) *UpdateContextDatabaseSyncConnectionCredentialRequest {
	s.Credentials = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionCredentialRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialRequest) Validate() error {
	return dara.Validate(s)
}
