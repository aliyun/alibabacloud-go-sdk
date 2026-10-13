// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *GetContextDatabaseSyncSourceRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncSourceRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncSourceRequest struct {
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

func (s GetContextDatabaseSyncSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetContextDatabaseSyncSourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncSourceRequest) SetSourceId(v string) *GetContextDatabaseSyncSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncSourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceRequest) Validate() error {
	return dara.Validate(s)
}
