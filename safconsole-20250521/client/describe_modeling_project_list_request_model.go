// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeModelingProjectListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeModelingProjectListRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeModelingProjectListRequest
	GetStatus() *string
}

type DescribeModelingProjectListRequest struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Pagination parameter: number of items per page, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Project status.
	//
	// example:
	//
	// PREPARING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModelingProjectListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectListRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeModelingProjectListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelingProjectListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelingProjectListRequest) SetCurrentPage(v int32) *DescribeModelingProjectListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeModelingProjectListRequest) SetPageSize(v int32) *DescribeModelingProjectListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModelingProjectListRequest) SetStatus(v string) *DescribeModelingProjectListRequest {
	s.Status = &v
	return s
}

func (s *DescribeModelingProjectListRequest) Validate() error {
	return dara.Validate(s)
}
