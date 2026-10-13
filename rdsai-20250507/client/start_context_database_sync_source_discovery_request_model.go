// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartContextDatabaseSyncSourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartContextDatabaseSyncSourceDiscoveryRequest
	GetClientToken() *string
	SetSourceId(v string) *StartContextDatabaseSyncSourceDiscoveryRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *StartContextDatabaseSyncSourceDiscoveryRequest
	GetWorkspaceId() *string
}

type StartContextDatabaseSyncSourceDiscoveryRequest struct {
	// The idempotence token.
	//
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data source ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartContextDatabaseSyncSourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s StartContextDatabaseSyncSourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) SetClientToken(v string) *StartContextDatabaseSyncSourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) SetSourceId(v string) *StartContextDatabaseSyncSourceDiscoveryRequest {
	s.SourceId = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) SetWorkspaceId(v string) *StartContextDatabaseSyncSourceDiscoveryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryRequest) Validate() error {
	return dara.Validate(s)
}
