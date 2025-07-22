// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLayoutListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeSystemLayoutListRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeSystemLayoutListRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeSystemLayoutListRequest
	GetPageSize() *int64
}

type DescribeSystemLayoutListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSystemLayoutListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSystemLayoutListRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeSystemLayoutListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSystemLayoutListRequest) SetOwnerId(v int64) *DescribeSystemLayoutListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSystemLayoutListRequest) SetPageNum(v int64) *DescribeSystemLayoutListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSystemLayoutListRequest) SetPageSize(v int64) *DescribeSystemLayoutListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSystemLayoutListRequest) Validate() error {
	return dara.Validate(s)
}
