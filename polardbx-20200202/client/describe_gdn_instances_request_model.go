// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGdnInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterType(v string) *DescribeGdnInstancesRequest
	GetFilterType() *string
	SetFilterValue(v string) *DescribeGdnInstancesRequest
	GetFilterValue() *string
	SetGDNId(v string) *DescribeGdnInstancesRequest
	GetGDNId() *string
	SetPageNum(v string) *DescribeGdnInstancesRequest
	GetPageNum() *string
	SetPageSize(v string) *DescribeGdnInstancesRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeGdnInstancesRequest
	GetRegionId() *string
}

type DescribeGdnInstancesRequest struct {
	// example:
	//
	// gdn_id、
	//
	// polarx_id
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// example:
	//
	// gdn-***、
	//
	// pxc-***
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// GDN ID。
	//
	// example:
	//
	// gdn-***
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// example:
	//
	// 50
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGdnInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeGdnInstancesRequest) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeGdnInstancesRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *DescribeGdnInstancesRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *DescribeGdnInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGdnInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGdnInstancesRequest) SetFilterType(v string) *DescribeGdnInstancesRequest {
	s.FilterType = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetFilterValue(v string) *DescribeGdnInstancesRequest {
	s.FilterValue = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetGDNId(v string) *DescribeGdnInstancesRequest {
	s.GDNId = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetPageNum(v string) *DescribeGdnInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetPageSize(v string) *DescribeGdnInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGdnInstancesRequest) SetRegionId(v string) *DescribeGdnInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGdnInstancesRequest) Validate() error {
	return dara.Validate(s)
}
