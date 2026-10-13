// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncConnectionRequest
	GetClientToken() *string
	SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionRequest
	GetConnectionId() *string
	SetDisplayName(v string) *UpdateContextDatabaseSyncConnectionRequest
	GetDisplayName() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncConnectionRequest struct {
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
	// 研发团队钉钉数据源
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncConnectionRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateContextDatabaseSyncConnectionRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateContextDatabaseSyncConnectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncConnectionRequest) SetClientToken(v string) *UpdateContextDatabaseSyncConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionRequest) SetConnectionId(v string) *UpdateContextDatabaseSyncConnectionRequest {
	s.ConnectionId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionRequest) SetDisplayName(v string) *UpdateContextDatabaseSyncConnectionRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncConnectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionRequest) Validate() error {
	return dara.Validate(s)
}
