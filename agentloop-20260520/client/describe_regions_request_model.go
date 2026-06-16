// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeRegionsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeRegionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRegionsRequest
	GetNextToken() *string
}

type DescribeRegionsRequest struct {
	// This parameter is required.
	Language   *string `json:"language,omitempty" xml:"language,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeRegionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRegionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRegionsRequest) SetLanguage(v string) *DescribeRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeRegionsRequest) SetMaxResults(v int32) *DescribeRegionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRegionsRequest) SetNextToken(v string) *DescribeRegionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
