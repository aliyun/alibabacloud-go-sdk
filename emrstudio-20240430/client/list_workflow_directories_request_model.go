// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListWorkflowDirectoriesRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListWorkflowDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowDirectoriesRequest
	GetNextToken() *string
	SetParentDirectoryId(v string) *ListWorkflowDirectoriesRequest
	GetParentDirectoryId() *string
	SetWorkspaceId(v string) *ListWorkflowDirectoriesRequest
	GetWorkspaceId() *string
}

type ListWorkflowDirectoriesRequest struct {
	// example:
	//
	// wd-y98v7non5dx****
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// wd-y98v7non5d****
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkflowDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListWorkflowDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowDirectoriesRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *ListWorkflowDirectoriesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkflowDirectoriesRequest) SetDirectoryId(v string) *ListWorkflowDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListWorkflowDirectoriesRequest) SetMaxResults(v int32) *ListWorkflowDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowDirectoriesRequest) SetNextToken(v string) *ListWorkflowDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowDirectoriesRequest) SetParentDirectoryId(v string) *ListWorkflowDirectoriesRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListWorkflowDirectoriesRequest) SetWorkspaceId(v string) *ListWorkflowDirectoriesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkflowDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
