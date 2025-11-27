// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *ListJobsShrinkRequest
	GetFilterShrink() *string
	SetPageNumber(v int32) *ListJobsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsShrinkRequest
	GetPageSize() *int32
	SetSortByShrink(v string) *ListJobsShrinkRequest
	GetSortByShrink() *string
}

type ListJobsShrinkRequest struct {
	// Queries job filter conditions.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page. Default value: 50. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting method.
	SortByShrink *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListJobsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsShrinkRequest) GetSortByShrink() *string {
	return s.SortByShrink
}

func (s *ListJobsShrinkRequest) SetFilterShrink(v string) *ListJobsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v int32) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v int32) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortByShrink(v string) *ListJobsShrinkRequest {
	s.SortByShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
