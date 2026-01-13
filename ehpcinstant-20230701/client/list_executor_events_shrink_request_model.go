// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *ListExecutorEventsShrinkRequest
	GetFilterShrink() *string
	SetPageNumber(v int32) *ListExecutorEventsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorEventsShrinkRequest
	GetPageSize() *int32
}

type ListExecutorEventsShrinkRequest struct {
	// Queries the Executor filter conditions.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The current page number.\\
	//
	// Starting value: 1\\
	//
	// Default value: 1
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
}

func (s ListExecutorEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListExecutorEventsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorEventsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorEventsShrinkRequest) SetFilterShrink(v string) *ListExecutorEventsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListExecutorEventsShrinkRequest) SetPageNumber(v int32) *ListExecutorEventsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorEventsShrinkRequest) SetPageSize(v int32) *ListExecutorEventsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListExecutorEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
