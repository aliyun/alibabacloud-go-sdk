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
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset to query.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of asset to query.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset to query. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **slb**: a Server Load Balancer (SLB) instance.
	//
	// 	- **eip**: an elastic IP address (EIP).
	//
	// 	- **ipv6**: an IPv6 gateway.
	//
	// 	- **swas**: a simple application server.
	//
	// 	- **waf**: a Web Application Firewall (WAF) instance of the Exclusive edition.
	//
	// 	- **ga_basic**: a Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the asset to query.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
