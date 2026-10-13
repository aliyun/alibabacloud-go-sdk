// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncDiscoveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetContextDatabaseSyncDiscoveryTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncDiscoveryTaskRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncDiscoveryTaskRequest struct {
	// The discovery task ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetContextDatabaseSyncDiscoveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncDiscoveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncDiscoveryTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetContextDatabaseSyncDiscoveryTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncDiscoveryTaskRequest) SetTaskId(v string) *GetContextDatabaseSyncDiscoveryTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncDiscoveryTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
