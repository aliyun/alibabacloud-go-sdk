// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListContextDatabaseSyncJobDocumentsRequest
	GetJobId() *string
	SetPage(v int32) *ListContextDatabaseSyncJobDocumentsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncJobDocumentsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListContextDatabaseSyncJobDocumentsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncJobDocumentsRequest struct {
	// The task ID. This is a system-generated UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number, starting from 1.
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

func (s ListContextDatabaseSyncJobDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) SetJobId(v string) *ListContextDatabaseSyncJobDocumentsRequest {
	s.JobId = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) SetPage(v int32) *ListContextDatabaseSyncJobDocumentsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) SetPageSize(v int32) *ListContextDatabaseSyncJobDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncJobDocumentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
