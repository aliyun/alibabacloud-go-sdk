// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListContextDatabaseSyncRunsRequest
	GetJobId() *string
	SetPage(v int32) *ListContextDatabaseSyncRunsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncRunsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListContextDatabaseSyncRunsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncRunsRequest struct {
	// The job ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseSyncRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListContextDatabaseSyncRunsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncRunsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncRunsRequest) SetJobId(v string) *ListContextDatabaseSyncRunsRequest {
	s.JobId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsRequest) SetPage(v int32) *ListContextDatabaseSyncRunsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncRunsRequest) SetPageSize(v int32) *ListContextDatabaseSyncRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncRunsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncRunsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsRequest) Validate() error {
	return dara.Validate(s)
}
