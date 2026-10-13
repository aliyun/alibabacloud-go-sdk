// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobExclusionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListContextDatabaseSyncJobExclusionsRequest
	GetJobId() *string
	SetWorkspaceId(v string) *ListContextDatabaseSyncJobExclusionsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncJobExclusionsRequest struct {
	// The node ID, which is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseSyncJobExclusionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobExclusionsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobExclusionsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListContextDatabaseSyncJobExclusionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncJobExclusionsRequest) SetJobId(v string) *ListContextDatabaseSyncJobExclusionsRequest {
	s.JobId = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncJobExclusionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsRequest) Validate() error {
	return dara.Validate(s)
}
