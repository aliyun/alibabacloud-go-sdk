// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPagesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPagesShrinkRequest
	GetPageSize() *int32
	SetQueryArgsShrink(v string) *ListPagesShrinkRequest
	GetQueryArgsShrink() *string
}

type ListPagesShrinkRequest struct {
	// The page number. Valid values: **1 to 100000**. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListPagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPagesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPagesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPagesShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListPagesShrinkRequest) SetPageNumber(v int32) *ListPagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPagesShrinkRequest) SetPageSize(v int32) *ListPagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPagesShrinkRequest) SetQueryArgsShrink(v string) *ListPagesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListPagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
