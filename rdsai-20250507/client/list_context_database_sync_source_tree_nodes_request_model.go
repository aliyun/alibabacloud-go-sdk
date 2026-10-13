// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourceTreeNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *ListContextDatabaseSyncSourceTreeNodesRequest
	GetSourceId() *string
	SetWorkspaceId(v string) *ListContextDatabaseSyncSourceTreeNodesRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncSourceTreeNodesRequest struct {
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

func (s ListContextDatabaseSyncSourceTreeNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourceTreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourceTreeNodesRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListContextDatabaseSyncSourceTreeNodesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncSourceTreeNodesRequest) SetSourceId(v string) *ListContextDatabaseSyncSourceTreeNodesRequest {
	s.SourceId = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncSourceTreeNodesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesRequest) Validate() error {
	return dara.Validate(s)
}
