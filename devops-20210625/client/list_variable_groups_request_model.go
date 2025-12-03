// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVariableGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVariableGroupsRequest
	GetNextToken() *string
	SetPageOrder(v string) *ListVariableGroupsRequest
	GetPageOrder() *string
	SetPageSort(v string) *ListVariableGroupsRequest
	GetPageSort() *string
}

type ListVariableGroupsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aaaaaa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DESC
	PageOrder *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	// example:
	//
	// ID
	PageSort *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
}

func (s ListVariableGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVariableGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVariableGroupsRequest) GetPageOrder() *string {
	return s.PageOrder
}

func (s *ListVariableGroupsRequest) GetPageSort() *string {
	return s.PageSort
}

func (s *ListVariableGroupsRequest) SetMaxResults(v int32) *ListVariableGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVariableGroupsRequest) SetNextToken(v string) *ListVariableGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVariableGroupsRequest) SetPageOrder(v string) *ListVariableGroupsRequest {
	s.PageOrder = &v
	return s
}

func (s *ListVariableGroupsRequest) SetPageSort(v string) *ListVariableGroupsRequest {
	s.PageSort = &v
	return s
}

func (s *ListVariableGroupsRequest) Validate() error {
	return dara.Validate(s)
}
