// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDdosEventListRequest
	GetCurrentPage() *int32
	SetDdosRegionId(v string) *DescribeDdosEventListRequest
	GetDdosRegionId() *string
	SetInstanceId(v string) *DescribeDdosEventListRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeDdosEventListRequest
	GetInstanceType() *string
	SetInternetIp(v string) *DescribeDdosEventListRequest
	GetInternetIp() *string
	SetPageSize(v int32) *DescribeDdosEventListRequest
	GetPageSize() *int32
	SetQueryDays(v int32) *DescribeDdosEventListRequest
	GetQueryDays() *int32
}

type DescribeDdosEventListRequest struct {
	// The number of the page to return for a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset that is assigned a public IP address.
	//
	// > Call [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) to query all region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of the instance for the asset that is assigned a public IP address.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) to query the IDs of the ECS, SLB, and EIP instances that belong to your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the asset that is assigned a public IP address. Valid values:
	//
	// - **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// - **slb**: a Server Load Balancer (SLB) instance.
	//
	// - **eip**: an elastic IP address (EIP) instance.
	//
	// - **ipv6**: an IPv6 Gateway instance.
	//
	// - **swas**: a simple application server instance.
	//
	// - **waf**: a dedicated Web Application Firewall (WAF) instance.
	//
	// - **ga_basic**: a Global Accelerator instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset that is assigned a public IP address.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The number of attack events to return on each page for a paged query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of days to query backwards from the current time. Default value: 7.
	//
	// example:
	//
	// 7
	QueryDays *int32 `json:"QueryDays,omitempty" xml:"QueryDays,omitempty"`
}

func (s DescribeDdosEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDdosEventListRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeDdosEventListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDdosEventListRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDdosEventListRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeDdosEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDdosEventListRequest) GetQueryDays() *int32 {
	return s.QueryDays
}

func (s *DescribeDdosEventListRequest) SetCurrentPage(v int32) *DescribeDdosEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetDdosRegionId(v string) *DescribeDdosEventListRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceId(v string) *DescribeDdosEventListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInstanceType(v string) *DescribeDdosEventListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetInternetIp(v string) *DescribeDdosEventListRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetPageSize(v int32) *DescribeDdosEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventListRequest) SetQueryDays(v int32) *DescribeDdosEventListRequest {
	s.QueryDays = &v
	return s
}

func (s *DescribeDdosEventListRequest) Validate() error {
	return dara.Validate(s)
}
