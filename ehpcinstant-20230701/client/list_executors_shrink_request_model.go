// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *ListExecutorsShrinkRequest
	GetFilterShrink() *string
	SetPageNumber(v int32) *ListExecutorsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorsShrinkRequest
	GetPageSize() *int32
}

type ListExecutorsShrinkRequest struct {
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExecutorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListExecutorsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorsShrinkRequest) SetFilterShrink(v string) *ListExecutorsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListExecutorsShrinkRequest) SetPageNumber(v int32) *ListExecutorsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsShrinkRequest) SetPageSize(v int32) *ListExecutorsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListExecutorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
