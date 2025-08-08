// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListConnectionsShrinkRequest
	GetKeyword() *string
	SetLabelSelectorShrink(v string) *ListConnectionsShrinkRequest
	GetLabelSelectorShrink() *string
	SetPageNumber(v int64) *ListConnectionsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListConnectionsShrinkRequest
	GetPageSize() *int64
}

type ListConnectionsShrinkRequest struct {
	// example:
	//
	// auto-
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

func (s ListConnectionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConnectionsShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListConnectionsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListConnectionsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConnectionsShrinkRequest) SetKeyword(v string) *ListConnectionsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetLabelSelectorShrink(v string) *ListConnectionsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetPageNumber(v int64) *ListConnectionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetPageSize(v int64) *ListConnectionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
