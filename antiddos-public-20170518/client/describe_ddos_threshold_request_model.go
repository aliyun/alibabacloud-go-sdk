// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeDdosThresholdRequest
	GetDdosRegionId() *string
	SetDdosType(v string) *DescribeDdosThresholdRequest
	GetDdosType() *string
	SetInstanceIds(v []*string) *DescribeDdosThresholdRequest
	GetInstanceIds() []*string
	SetInstanceType(v string) *DescribeDdosThresholdRequest
	GetInstanceType() *string
}

type DescribeDdosThresholdRequest struct {
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
	// The ID of asset N to query.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
}

func (s DescribeDdosThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosThresholdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeDdosThresholdRequest) GetDdosType() *string {
	return s.DdosType
}

func (s *DescribeDdosThresholdRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDdosThresholdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDdosThresholdRequest) SetDdosRegionId(v string) *DescribeDdosThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetDdosType(v string) *DescribeDdosThresholdRequest {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceIds(v []*string) *DescribeDdosThresholdRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDdosThresholdRequest) SetInstanceType(v string) *DescribeDdosThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosThresholdRequest) Validate() error {
	return dara.Validate(s)
}
