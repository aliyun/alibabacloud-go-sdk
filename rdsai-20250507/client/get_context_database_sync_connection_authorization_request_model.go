// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncConnectionAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v string) *GetContextDatabaseSyncConnectionAuthorizationRequest
	GetConnectionId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncConnectionAuthorizationRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncConnectionAuthorizationRequest struct {
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

func (s GetContextDatabaseSyncConnectionAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncConnectionAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncConnectionAuthorizationRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetContextDatabaseSyncConnectionAuthorizationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncConnectionAuthorizationRequest) SetConnectionId(v string) *GetContextDatabaseSyncConnectionAuthorizationRequest {
	s.ConnectionId = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncConnectionAuthorizationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
