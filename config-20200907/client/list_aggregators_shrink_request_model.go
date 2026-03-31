// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregatorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAggregatorsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregatorsShrinkRequest
	GetNextToken() *string
	SetTagShrink(v string) *ListAggregatorsShrinkRequest
	GetTagShrink() *string
}

type ListAggregatorsShrinkRequest struct {
	// The maximum number of entries to return in a request. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU2hhcmVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAggregatorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAggregatorsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregatorsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregatorsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListAggregatorsShrinkRequest) SetMaxResults(v int32) *ListAggregatorsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregatorsShrinkRequest) SetNextToken(v string) *ListAggregatorsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregatorsShrinkRequest) SetTagShrink(v string) *ListAggregatorsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListAggregatorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
