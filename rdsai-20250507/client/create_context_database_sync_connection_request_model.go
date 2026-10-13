// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKind(v string) *CreateContextDatabaseSyncConnectionRequest
	GetAuthKind() *string
	SetClientToken(v string) *CreateContextDatabaseSyncConnectionRequest
	GetClientToken() *string
	SetDisplayName(v string) *CreateContextDatabaseSyncConnectionRequest
	GetDisplayName() *string
	SetPluginId(v string) *CreateContextDatabaseSyncConnectionRequest
	GetPluginId() *string
	SetWorkspaceId(v string) *CreateContextDatabaseSyncConnectionRequest
	GetWorkspaceId() *string
}

type CreateContextDatabaseSyncConnectionRequest struct {
	// The authorization method. Valid values: app, pasted-url, static-secret, and none.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	AuthKind *string `json:"AuthKind,omitempty" xml:"AuthKind,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// R&D team DingTalk data source
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The plugin ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateContextDatabaseSyncConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncConnectionRequest) GetAuthKind() *string {
	return s.AuthKind
}

func (s *CreateContextDatabaseSyncConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateContextDatabaseSyncConnectionRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateContextDatabaseSyncConnectionRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *CreateContextDatabaseSyncConnectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseSyncConnectionRequest) SetAuthKind(v string) *CreateContextDatabaseSyncConnectionRequest {
	s.AuthKind = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionRequest) SetClientToken(v string) *CreateContextDatabaseSyncConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionRequest) SetDisplayName(v string) *CreateContextDatabaseSyncConnectionRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionRequest) SetPluginId(v string) *CreateContextDatabaseSyncConnectionRequest {
	s.PluginId = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionRequest) SetWorkspaceId(v string) *CreateContextDatabaseSyncConnectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionRequest) Validate() error {
	return dara.Validate(s)
}
