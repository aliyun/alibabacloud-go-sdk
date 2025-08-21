// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstanceIpAddressRequest
	GetCurrentPage() *int32
	SetDdosRegionId(v string) *DescribeInstanceIpAddressRequest
	GetDdosRegionId() *string
	SetDdosStatus(v string) *DescribeInstanceIpAddressRequest
	GetDdosStatus() *string
	SetInstanceId(v string) *DescribeInstanceIpAddressRequest
	GetInstanceId() *string
	SetInstanceIp(v string) *DescribeInstanceIpAddressRequest
	GetInstanceIp() *string
	SetInstanceName(v string) *DescribeInstanceIpAddressRequest
	GetInstanceName() *string
	SetInstanceType(v string) *DescribeInstanceIpAddressRequest
	GetInstanceType() *string
	SetPageSize(v int32) *DescribeInstanceIpAddressRequest
	GetPageSize() *int32
}

type DescribeInstanceIpAddressRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **defense**: queries assets for which traffic scrubbing is performed.
	//
	// 	- **blackhole**: queries assets for which blackhole filtering is triggered.
	//
	// example:
	//
	// normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The ID of the instance to which the asset is added.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset that is assigned a public IP address. Valid values:
	//
	// 	- **ecs**: ECS instances.
	//
	// 	- **slb**: SLB instances.
	//
	// 	- **eip**: EIPs.
	//
	// 	- **ipv6**: IPv6 gateways.
	//
	// 	- **swas**: simple application servers.
	//
	// 	- **waf**: Web Application Firewall (WAF) instances of the Exclusive edition.
	//
	// 	- **ga_basic**: Global Accelerator (GA) instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceIpAddressRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeInstanceIpAddressRequest) GetDdosStatus() *string {
	return s.DdosStatus
}

func (s *DescribeInstanceIpAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceIpAddressRequest) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeInstanceIpAddressRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceIpAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceIpAddressRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceIpAddressRequest) SetCurrentPage(v int32) *DescribeInstanceIpAddressRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosRegionId(v string) *DescribeInstanceIpAddressRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetDdosStatus(v string) *DescribeInstanceIpAddressRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceId(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceIp(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceName(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetInstanceType(v string) *DescribeInstanceIpAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) SetPageSize(v int32) *DescribeInstanceIpAddressRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
