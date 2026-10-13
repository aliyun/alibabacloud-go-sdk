// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetContextDatabaseSyncJobRequest
	GetJobId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncJobRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncJobRequest struct {
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

func (s GetContextDatabaseSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetContextDatabaseSyncJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncJobRequest) SetJobId(v string) *GetContextDatabaseSyncJobRequest {
	s.JobId = &v
	return s
}

func (s *GetContextDatabaseSyncJobRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
