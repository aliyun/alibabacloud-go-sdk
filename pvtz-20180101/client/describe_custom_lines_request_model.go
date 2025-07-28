// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomLinesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeCustomLinesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomLinesRequest
	GetPageSize() *int32
}

type DescribeCustomLinesRequest struct {
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCustomLinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomLinesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomLinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomLinesRequest) SetLang(v string) *DescribeCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageNumber(v int32) *DescribeCustomLinesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomLinesRequest) SetPageSize(v int32) *DescribeCustomLinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomLinesRequest) Validate() error {
	return dara.Validate(s)
}
