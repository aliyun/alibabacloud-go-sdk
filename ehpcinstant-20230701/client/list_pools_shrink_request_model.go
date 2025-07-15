// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoolsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *ListPoolsShrinkRequest
	GetFilterShrink() *string
	SetPageNumber(v int32) *ListPoolsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoolsShrinkRequest
	GetPageSize() *int32
}

type ListPoolsShrinkRequest struct {
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPoolsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPoolsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListPoolsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoolsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoolsShrinkRequest) SetFilterShrink(v string) *ListPoolsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListPoolsShrinkRequest) SetPageNumber(v int32) *ListPoolsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsShrinkRequest) SetPageSize(v int32) *ListPoolsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoolsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
