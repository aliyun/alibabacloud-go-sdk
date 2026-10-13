// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetContextDatabaseSyncJobScheduleRequest
	GetJobId() *string
	SetWorkspaceId(v string) *GetContextDatabaseSyncJobScheduleRequest
	GetWorkspaceId() *string
}

type GetContextDatabaseSyncJobScheduleRequest struct {
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

func (s GetContextDatabaseSyncJobScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobScheduleRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetContextDatabaseSyncJobScheduleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetContextDatabaseSyncJobScheduleRequest) SetJobId(v string) *GetContextDatabaseSyncJobScheduleRequest {
	s.JobId = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleRequest) SetWorkspaceId(v string) *GetContextDatabaseSyncJobScheduleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleRequest) Validate() error {
	return dara.Validate(s)
}
