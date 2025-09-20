// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListProjectsShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListProjectsShrinkRequest
	GetNextToken() *string
	SetPrefix(v string) *ListProjectsShrinkRequest
	GetPrefix() *string
	SetTagShrink(v string) *ListProjectsShrinkRequest
	GetTagShrink() *string
}

type ListProjectsShrinkRequest struct {
	// The maximum number of entries to return. Valid values: 0 to 200. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. The operation returns the projects in lexicographical order starting from the location specified by NextToken.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDAx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix used by the projects that you want to query. The prefix must be up to 128 characters in length.
	//
	// example:
	//
	// immtest
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListProjectsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsShrinkRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListProjectsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListProjectsShrinkRequest) SetMaxResults(v int64) *ListProjectsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetNextToken(v string) *ListProjectsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPrefix(v string) *ListProjectsShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetTagShrink(v string) *ListProjectsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
