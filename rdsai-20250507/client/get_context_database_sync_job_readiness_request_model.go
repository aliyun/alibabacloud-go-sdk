// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobReadinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetContextDatabaseSyncJobReadinessRequest
	GetJobId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncJobReadinessRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncJobReadinessRequest struct {
	// The job ID, which is a system-generated UUID.
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

func (s GetContextDatabaseSyncJobReadinessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobReadinessRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobReadinessRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetContextDatabaseSyncJobReadinessRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncJobReadinessRequest) SetJobId(v string) *GetContextDatabaseSyncJobReadinessRequest {
	s.JobId = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncJobReadinessRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessRequest) Validate() error {
	return dara.Validate(s)
}
