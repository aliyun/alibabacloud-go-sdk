// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcludeSystemPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeExcludeSystemPathRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeExcludeSystemPathRequest
	GetPageSize() *int32
}

type DescribeExcludeSystemPathRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExcludeSystemPathRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcludeSystemPathRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcludeSystemPathRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExcludeSystemPathRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExcludeSystemPathRequest) SetCurrentPage(v int32) *DescribeExcludeSystemPathRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExcludeSystemPathRequest) SetPageSize(v int32) *DescribeExcludeSystemPathRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExcludeSystemPathRequest) Validate() error {
	return dara.Validate(s)
}
