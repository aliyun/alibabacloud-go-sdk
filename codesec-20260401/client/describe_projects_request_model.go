// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeProjectsRequest
	GetNextToken() *string
	SetQuery(v string) *DescribeProjectsRequest
	GetQuery() *string
}

type DescribeProjectsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProjectsRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeProjectsRequest) SetMaxResults(v int32) *DescribeProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProjectsRequest) SetNextToken(v string) *DescribeProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProjectsRequest) SetQuery(v string) *DescribeProjectsRequest {
	s.Query = &v
	return s
}

func (s *DescribeProjectsRequest) Validate() error {
	return dara.Validate(s)
}
