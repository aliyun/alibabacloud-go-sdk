// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListContextDatabaseSyncJobsRequest
	GetKnowledgeBaseId() *string
	SetPage(v int32) *ListContextDatabaseSyncJobsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncJobsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListContextDatabaseSyncJobsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncJobsRequest struct {
	// The ID of the target knowledge base.
	//
	// example:
	//
	// kb-as1llqmkol****
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The page number. The value starts from 1.
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

func (s ListContextDatabaseSyncJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListContextDatabaseSyncJobsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncJobsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncJobsRequest) SetKnowledgeBaseId(v string) *ListContextDatabaseSyncJobsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsRequest) SetPage(v int32) *ListContextDatabaseSyncJobsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncJobsRequest) SetPageSize(v int32) *ListContextDatabaseSyncJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncJobsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsRequest) Validate() error {
	return dara.Validate(s)
}
