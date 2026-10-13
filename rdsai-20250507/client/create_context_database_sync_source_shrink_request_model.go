// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateContextDatabaseSyncSourceShrinkRequest
	GetClientToken() *string
	SetConnectionId(v string) *CreateContextDatabaseSyncSourceShrinkRequest
	GetConnectionId() *string
	SetName(v string) *CreateContextDatabaseSyncSourceShrinkRequest
	GetName() *string
	SetSourceConfigShrink(v string) *CreateContextDatabaseSyncSourceShrinkRequest
	GetSourceConfigShrink() *string
	SetWorkspaceId(v string) *CreateContextDatabaseSyncSourceShrinkRequest
	GetWorkspaceId() *string
}

type CreateContextDatabaseSyncSourceShrinkRequest struct {
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
	SourceConfigShrink *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateContextDatabaseSyncSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) GetSourceConfigShrink() *string {
	return s.SourceConfigShrink
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) SetClientToken(v string) *CreateContextDatabaseSyncSourceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) SetConnectionId(v string) *CreateContextDatabaseSyncSourceShrinkRequest {
	s.ConnectionId = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) SetName(v string) *CreateContextDatabaseSyncSourceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) SetSourceConfigShrink(v string) *CreateContextDatabaseSyncSourceShrinkRequest {
	s.SourceConfigShrink = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) SetWorkspaceId(v string) *CreateContextDatabaseSyncSourceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
