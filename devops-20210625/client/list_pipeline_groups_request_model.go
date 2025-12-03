// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListPipelineGroupsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPipelineGroupsRequest
	GetNextToken() *string
}

type ListPipelineGroupsRequest struct {
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aaaaaaaaaa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPipelineGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPipelineGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineGroupsRequest) SetMaxResults(v int64) *ListPipelineGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineGroupsRequest) SetNextToken(v string) *ListPipelineGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupsRequest) Validate() error {
	return dara.Validate(s)
}
