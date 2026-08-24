// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextDatabaseWorkspacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextDatabaseWorkspacesRequest
	GetNextToken() *string
	SetStatus(v string) *ListContextDatabaseWorkspacesRequest
	GetStatus() *string
	SetWorkspaceId(v string) *ListContextDatabaseWorkspacesRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseWorkspacesRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status used to filter results. Valid values: Active and Locked.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID used to filter results. Only the workspace that matches this ID is returned.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseWorkspacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseWorkspacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseWorkspacesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseWorkspacesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseWorkspacesRequest) SetMaxResults(v int32) *ListContextDatabaseWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseWorkspacesRequest) SetNextToken(v string) *ListContextDatabaseWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseWorkspacesRequest) SetStatus(v string) *ListContextDatabaseWorkspacesRequest {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseWorkspacesRequest) SetWorkspaceId(v string) *ListContextDatabaseWorkspacesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
