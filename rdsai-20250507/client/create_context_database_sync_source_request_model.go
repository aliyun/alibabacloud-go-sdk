// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateContextDatabaseSyncSourceRequest
	GetClientToken() *string
	SetConnectionId(v string) *CreateContextDatabaseSyncSourceRequest
	GetConnectionId() *string
	SetName(v string) *CreateContextDatabaseSyncSourceRequest
	GetName() *string
	SetSourceConfig(v map[string]interface{}) *CreateContextDatabaseSyncSourceRequest
	GetSourceConfig() map[string]interface{}
	SetWorkspaceId(v string) *CreateContextDatabaseSyncSourceRequest
	GetWorkspaceId() *string
}

type CreateContextDatabaseSyncSourceRequest struct {
	// The idempotence token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection ID, a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// R&D team document scope
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The datasource config. The structure varies depending on the plug-in.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"folderIds":["https://alidocs.dingtalk.com/i/p/p2N9kLmQxYz****"],"pageSize":50,"maxDepth":5}
	SourceConfig map[string]interface{} `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateContextDatabaseSyncSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncSourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateContextDatabaseSyncSourceRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateContextDatabaseSyncSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateContextDatabaseSyncSourceRequest) GetSourceConfig() map[string]interface{} {
	return s.SourceConfig
}

func (s *CreateContextDatabaseSyncSourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseSyncSourceRequest) SetClientToken(v string) *CreateContextDatabaseSyncSourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceRequest) SetConnectionId(v string) *CreateContextDatabaseSyncSourceRequest {
	s.ConnectionId = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceRequest) SetName(v string) *CreateContextDatabaseSyncSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceRequest) SetSourceConfig(v map[string]interface{}) *CreateContextDatabaseSyncSourceRequest {
	s.SourceConfig = v
	return s
}

func (s *CreateContextDatabaseSyncSourceRequest) SetWorkspaceId(v string) *CreateContextDatabaseSyncSourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceRequest) Validate() error {
	return dara.Validate(s)
}
