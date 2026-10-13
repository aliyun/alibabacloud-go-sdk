// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListContextDatabaseSyncSourcesRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncSourcesRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListContextDatabaseSyncSourcesRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncSourcesRequest struct {
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

func (s ListContextDatabaseSyncSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourcesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncSourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncSourcesRequest) SetPage(v int32) *ListContextDatabaseSyncSourcesRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesRequest) SetPageSize(v int32) *ListContextDatabaseSyncSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncSourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesRequest) Validate() error {
	return dara.Validate(s)
}
