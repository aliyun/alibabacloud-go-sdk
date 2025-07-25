// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeIspFlushCacheTasksRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeIspFlushCacheTasksRequest
	GetDomainName() *string
	SetInstanceId(v string) *DescribeIspFlushCacheTasksRequest
	GetInstanceId() *string
	SetIsp(v string) *DescribeIspFlushCacheTasksRequest
	GetIsp() *string
	SetLang(v string) *DescribeIspFlushCacheTasksRequest
	GetLang() *string
	SetOrderBy(v string) *DescribeIspFlushCacheTasksRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeIspFlushCacheTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIspFlushCacheTasksRequest
	GetPageSize() *int32
}

type DescribeIspFlushCacheTasksRequest struct {
	Direction  *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeIspFlushCacheTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTasksRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeIspFlushCacheTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeIspFlushCacheTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIspFlushCacheTasksRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeIspFlushCacheTasksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIspFlushCacheTasksRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeIspFlushCacheTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIspFlushCacheTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIspFlushCacheTasksRequest) SetDirection(v string) *DescribeIspFlushCacheTasksRequest {
	s.Direction = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetDomainName(v string) *DescribeIspFlushCacheTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetInstanceId(v string) *DescribeIspFlushCacheTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetIsp(v string) *DescribeIspFlushCacheTasksRequest {
	s.Isp = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetLang(v string) *DescribeIspFlushCacheTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetOrderBy(v string) *DescribeIspFlushCacheTasksRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetPageNumber(v int32) *DescribeIspFlushCacheTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) SetPageSize(v int32) *DescribeIspFlushCacheTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIspFlushCacheTasksRequest) Validate() error {
	return dara.Validate(s)
}
