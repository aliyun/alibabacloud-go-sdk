// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspaceUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspaceUsersRequest
	GetNextToken() *string
	SetPageNum(v string) *ListWorkspaceUsersRequest
	GetPageNum() *string
	SetWorkspaceId(v string) *ListWorkspaceUsersRequest
	GetWorkspaceId() *string
}

type ListWorkspaceUsersRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspaceUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspaceUsersRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *ListWorkspaceUsersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceUsersRequest) SetMaxResults(v int32) *ListWorkspaceUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspaceUsersRequest) SetNextToken(v string) *ListWorkspaceUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspaceUsersRequest) SetPageNum(v string) *ListWorkspaceUsersRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkspaceUsersRequest) SetWorkspaceId(v string) *ListWorkspaceUsersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceUsersRequest) Validate() error {
	return dara.Validate(s)
}
