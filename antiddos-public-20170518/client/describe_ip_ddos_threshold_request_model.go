// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpDdosThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeIpDdosThresholdRequest
	GetDdosRegionId() *string
	SetDdosType(v string) *DescribeIpDdosThresholdRequest
	GetDdosType() *string
	SetInstanceId(v string) *DescribeIpDdosThresholdRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeIpDdosThresholdRequest
	GetInstanceType() *string
	SetInternetIp(v string) *DescribeIpDdosThresholdRequest
	GetInternetIp() *string
}

type DescribeIpDdosThresholdRequest struct {
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
	// The type of the threshold. Valid values:
	//
	// 	- **defense**: traffic scrubbing threshold
	//
	// 	- **blackhole**: DDoS mitigation threshold
	//
	// This parameter is required.
	//
	// example:
	//
	// defense
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The ID of the asset.
	//
	// > You can call the [DescribeInstanceIpAddress](https://help.aliyun.com/document_detail/429562.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1i88rqjza51s****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The IP address of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpDdosThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpDdosThresholdRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeIpDdosThresholdRequest) GetDdosType() *string {
	return s.DdosType
}

func (s *DescribeIpDdosThresholdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpDdosThresholdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeIpDdosThresholdRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeIpDdosThresholdRequest) SetDdosRegionId(v string) *DescribeIpDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetDdosType(v string) *DescribeIpDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceId(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInstanceType(v string) *DescribeIpDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) SetInternetIp(v string) *DescribeIpDdosThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpDdosThresholdRequest) Validate() error {
	return dara.Validate(s)
}
