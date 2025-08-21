// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstanceRequest
	GetCurrentPage() *int32
	SetDdosRegionId(v string) *DescribeInstanceRequest
	GetDdosRegionId() *string
	SetDdosStatus(v string) *DescribeInstanceRequest
	GetDdosStatus() *string
	SetInstanceId(v string) *DescribeInstanceRequest
	GetInstanceId() *string
	SetInstanceIp(v string) *DescribeInstanceRequest
	GetInstanceIp() *string
	SetInstanceName(v string) *DescribeInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *DescribeInstanceRequest
	GetInstanceType() *string
	SetPageSize(v int32) *DescribeInstanceRequest
	GetPageSize() *int32
}

type DescribeInstanceRequest struct {
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
	// 	- **mitigating**: queries assets for which traffic scrubbing is triggered.
	//
	// 	- **blackholed**: queries assets for which blackhole filtering is triggered.
	//
	// 	- **normal**: queries assets that are not under DDoS attacks.
	//
	// example:
	//
	// blackholed
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The ID of the asset. The formats of asset IDs vary based on the value of the **InstanceType**. parameter.
	//
	// 	- If you set **InstanceType*	- to **ecs**, specify the ID of the ECS instance. For example, you can specify i-bp1cb6x80tfgocid\\*\\*\\*\\*.
	//
	// 	- If you set **InstanceType*	- to **slb**, specify the ID of the SLB instance. For example, you can specify alb-vn2dqg3v31y2vd\\*\\*\\*\\*.
	//
	// 	- If you set **InstanceType*	- to **eip**, specify the ID of the EIP. For example, you can specify eip-j6ce6dcx9epi7rs46\\*\\*\\*\\*.
	//
	// example:
	//
	// i-bp1cb6x80tfgocid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 121.199.XX.XX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// launch-advisor-2022****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset to query. Valid values:
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

func (s DescribeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeInstanceRequest) GetDdosStatus() *string {
	return s.DdosStatus
}

func (s *DescribeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRequest) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceRequest) SetCurrentPage(v int32) *DescribeInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosRegionId(v string) *DescribeInstanceRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetDdosStatus(v string) *DescribeInstanceRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceIp(v string) *DescribeInstanceRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceName(v string) *DescribeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceRequest) SetInstanceType(v string) *DescribeInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageSize(v int32) *DescribeInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
