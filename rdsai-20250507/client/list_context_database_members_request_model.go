// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextDatabaseMembersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextDatabaseMembersRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListContextDatabaseMembersRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseMembersRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseMembersRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseMembersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseMembersRequest) SetMaxResults(v int32) *ListContextDatabaseMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseMembersRequest) SetNextToken(v string) *ListContextDatabaseMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseMembersRequest) SetWorkspaceId(v string) *ListContextDatabaseMembersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseMembersRequest) Validate() error {
	return dara.Validate(s)
}
