// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListContextDatabaseSyncRunLogsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncRunLogsRequest
	GetPageSize() *int32
	SetRunId(v string) *ListContextDatabaseSyncRunLogsRequest
	GetRunId() *string
	SetWorkspaceId(v string) *ListContextDatabaseSyncRunLogsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncRunLogsRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseSyncRunLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunLogsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunLogsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncRunLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncRunLogsRequest) GetRunId() *string {
	return s.RunId
}

func (s *ListContextDatabaseSyncRunLogsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncRunLogsRequest) SetPage(v int32) *ListContextDatabaseSyncRunLogsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsRequest) SetPageSize(v int32) *ListContextDatabaseSyncRunLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsRequest) SetRunId(v string) *ListContextDatabaseSyncRunLogsRequest {
	s.RunId = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncRunLogsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsRequest) Validate() error {
	return dara.Validate(s)
}
