// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeIspFlushCacheInstancesRequest
	GetDirection() *string
	SetIsp(v string) *DescribeIspFlushCacheInstancesRequest
	GetIsp() *string
	SetKeyword(v string) *DescribeIspFlushCacheInstancesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeIspFlushCacheInstancesRequest
	GetLang() *string
	SetOrderBy(v string) *DescribeIspFlushCacheInstancesRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeIspFlushCacheInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIspFlushCacheInstancesRequest
	GetPageSize() *int32
	SetType(v string) *DescribeIspFlushCacheInstancesRequest
	GetType() *string
}

type DescribeIspFlushCacheInstancesRequest struct {
	Direction  *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIspFlushCacheInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheInstancesRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeIspFlushCacheInstancesRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeIspFlushCacheInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeIspFlushCacheInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIspFlushCacheInstancesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeIspFlushCacheInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIspFlushCacheInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIspFlushCacheInstancesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeIspFlushCacheInstancesRequest) SetDirection(v string) *DescribeIspFlushCacheInstancesRequest {
	s.Direction = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetIsp(v string) *DescribeIspFlushCacheInstancesRequest {
	s.Isp = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetKeyword(v string) *DescribeIspFlushCacheInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetLang(v string) *DescribeIspFlushCacheInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetOrderBy(v string) *DescribeIspFlushCacheInstancesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetPageNumber(v int32) *DescribeIspFlushCacheInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetPageSize(v int32) *DescribeIspFlushCacheInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) SetType(v string) *DescribeIspFlushCacheInstancesRequest {
	s.Type = &v
	return s
}

func (s *DescribeIspFlushCacheInstancesRequest) Validate() error {
	return dara.Validate(s)
}
