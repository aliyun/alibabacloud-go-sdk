// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeDdosCountRequest
	GetDdosRegionId() *string
	SetInstanceType(v string) *DescribeDdosCountRequest
	GetInstanceType() *string
}

type DescribeDdosCountRequest struct {
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
	// The type of the asset to query. Valid values:
	//
	// 	- **ecs**: Elastic Compute Service (ECS) instances.
	//
	// 	- **slb**: Server Load Balancer (SLB) instances.
	//
	// 	- **eip**: elastic IP addresses (EIPs).
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

func (s DescribeDdosCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeDdosCountRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDdosCountRequest) SetDdosRegionId(v string) *DescribeDdosCountRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosCountRequest) SetInstanceType(v string) *DescribeDdosCountRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDdosCountRequest) Validate() error {
	return dara.Validate(s)
}
