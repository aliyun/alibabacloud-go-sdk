// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListListsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListsShrinkRequest
	GetPageSize() *int32
	SetQueryArgsShrink(v string) *ListListsShrinkRequest
	GetQueryArgsShrink() *string
}

type ListListsShrinkRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query arguments in the JSON format, which contain filter conditions.
	//
	// example:
	//
	// ListLists
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s ListListsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListListsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListsShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListListsShrinkRequest) SetPageNumber(v int32) *ListListsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListsShrinkRequest) SetPageSize(v int32) *ListListsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListListsShrinkRequest) SetQueryArgsShrink(v string) *ListListsShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListListsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
