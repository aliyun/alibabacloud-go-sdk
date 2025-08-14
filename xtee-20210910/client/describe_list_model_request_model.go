// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *DescribeListModelRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeListModelRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeListModelRequest
	GetRegId() *string
}

type DescribeListModelRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s DescribeListModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListModelRequest) GoString() string {
	return s.String()
}

func (s *DescribeListModelRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeListModelRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeListModelRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeListModelRequest) SetCurrentPage(v int64) *DescribeListModelRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListModelRequest) SetPageSize(v int64) *DescribeListModelRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeListModelRequest) SetRegId(v string) *DescribeListModelRequest {
	s.RegId = &v
	return s
}

func (s *DescribeListModelRequest) Validate() error {
	return dara.Validate(s)
}
