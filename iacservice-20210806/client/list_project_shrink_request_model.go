// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProjectShrinkRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListProjectShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListProjectShrinkRequest
	GetPageSize() *string
	SetTagShrink(v string) *ListProjectShrinkRequest
	GetTagShrink() *string
}

type ListProjectShrinkRequest struct {
	// The search keyword. Name-based search is supported.
	//
	// example:
	//
	// test-case2
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of tags.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListProjectShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListProjectShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListProjectShrinkRequest) SetKeyword(v string) *ListProjectShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectShrinkRequest) SetPageNumber(v string) *ListProjectShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectShrinkRequest) SetPageSize(v string) *ListProjectShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectShrinkRequest) SetTagShrink(v string) *ListProjectShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
