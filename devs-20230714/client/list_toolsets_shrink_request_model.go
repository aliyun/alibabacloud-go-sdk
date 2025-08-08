// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListToolsetsShrinkRequest
	GetKeyword() *string
	SetLabelSelectorShrink(v string) *ListToolsetsShrinkRequest
	GetLabelSelectorShrink() *string
	SetPageNumber(v int64) *ListToolsetsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListToolsetsShrinkRequest
	GetPageSize() *int64
}

type ListToolsetsShrinkRequest struct {
	// example:
	//
	// demo
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListToolsetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListToolsetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListToolsetsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListToolsetsShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListToolsetsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListToolsetsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListToolsetsShrinkRequest) SetKeyword(v string) *ListToolsetsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListToolsetsShrinkRequest) SetLabelSelectorShrink(v string) *ListToolsetsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListToolsetsShrinkRequest) SetPageNumber(v int64) *ListToolsetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListToolsetsShrinkRequest) SetPageSize(v int64) *ListToolsetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListToolsetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
