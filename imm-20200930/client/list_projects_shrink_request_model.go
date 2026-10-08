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
	// The maximum number of projects to return. Valid values: 0 to 200. If this parameter is not set or is set to 0, the default value is 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous API call. The list of projects is returned in lexicographical order starting from the NextToken value. Leave this parameter empty when you call this API operation for the first time.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDAx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix used to filter projects. The length is limited to 0 to 128 characters.
	//
	// example:
	//
	// immtest
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The tag list.
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
