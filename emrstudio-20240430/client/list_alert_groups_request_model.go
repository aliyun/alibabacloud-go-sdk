// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAlertGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertGroupsRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListAlertGroupsRequest
	GetSearchVal() *string
	SetWorkspaceId(v string) *ListAlertGroupsRequest
	GetWorkspaceId() *string
}

type ListAlertGroupsRequest struct {
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
	// name
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// w-lxyy60mpgpg****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListAlertGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertGroupsRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListAlertGroupsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAlertGroupsRequest) SetMaxResults(v int32) *ListAlertGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertGroupsRequest) SetNextToken(v string) *ListAlertGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertGroupsRequest) SetSearchVal(v string) *ListAlertGroupsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListAlertGroupsRequest) SetWorkspaceId(v string) *ListAlertGroupsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAlertGroupsRequest) Validate() error {
	return dara.Validate(s)
}
