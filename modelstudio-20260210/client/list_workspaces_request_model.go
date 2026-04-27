// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListWorkspacesRequest
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *ListWorkspacesRequest
	GetWorkspaceName() *string
}

type ListWorkspacesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// uwCwQ5FFCDo=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ws-32klhjk2312334jkh
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// ws_test
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceId(v string) *ListWorkspacesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceName(v string) *ListWorkspacesRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
