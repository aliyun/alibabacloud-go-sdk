// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProjectsShrinkRequest
	GetKeyword() *string
	SetLabelSelectorShrink(v string) *ListProjectsShrinkRequest
	GetLabelSelectorShrink() *string
	SetPageNumber(v int64) *ListProjectsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListProjectsShrinkRequest
	GetPageSize() *int64
}

type ListProjectsShrinkRequest struct {
	// example:
	//
	// spring-boot
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListProjectsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListProjectsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectsShrinkRequest) SetKeyword(v string) *ListProjectsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetLabelSelectorShrink(v string) *ListProjectsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int64) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int64) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
