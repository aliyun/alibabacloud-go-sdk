// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSDGShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGShrinkRequest
	GetPageSize() *int32
	SetSDGIdsShrink(v string) *DescribeSDGShrinkRequest
	GetSDGIdsShrink() *string
}

type DescribeSDGShrinkRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of SDGs that you want to query. By default, all SDGs are queried.
	SDGIdsShrink *string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty"`
}

func (s DescribeSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGShrinkRequest) GetSDGIdsShrink() *string {
	return s.SDGIdsShrink
}

func (s *DescribeSDGShrinkRequest) SetPageNumber(v int32) *DescribeSDGShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGShrinkRequest) SetPageSize(v int32) *DescribeSDGShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGShrinkRequest) SetSDGIdsShrink(v string) *DescribeSDGShrinkRequest {
	s.SDGIdsShrink = &v
	return s
}

func (s *DescribeSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
