// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContextDatabaseSyncConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *VerifyContextDatabaseSyncConnectionRequest
	GetClientToken() *string
	SetConnectionId(v string) *VerifyContextDatabaseSyncConnectionRequest
	GetConnectionId() *string
	SetWorkspaceId(v string) *VerifyContextDatabaseSyncConnectionRequest
	GetWorkspaceId() *string
}

type VerifyContextDatabaseSyncConnectionRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s VerifyContextDatabaseSyncConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyContextDatabaseSyncConnectionRequest) GoString() string {
	return s.String()
}

func (s *VerifyContextDatabaseSyncConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *VerifyContextDatabaseSyncConnectionRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *VerifyContextDatabaseSyncConnectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *VerifyContextDatabaseSyncConnectionRequest) SetClientToken(v string) *VerifyContextDatabaseSyncConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionRequest) SetConnectionId(v string) *VerifyContextDatabaseSyncConnectionRequest {
	s.ConnectionId = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionRequest) SetWorkspaceId(v string) *VerifyContextDatabaseSyncConnectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionRequest) Validate() error {
	return dara.Validate(s)
}
