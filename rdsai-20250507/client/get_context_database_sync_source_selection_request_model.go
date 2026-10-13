// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceSelectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *GetContextDatabaseSyncSourceSelectionRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncSourceSelectionRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncSourceSelectionRequest struct {
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

func (s GetContextDatabaseSyncSourceSelectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceSelectionRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceSelectionRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetContextDatabaseSyncSourceSelectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncSourceSelectionRequest) SetSourceId(v string) *GetContextDatabaseSyncSourceSelectionRequest {
	s.SourceId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncSourceSelectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionRequest) Validate() error {
	return dara.Validate(s)
}
