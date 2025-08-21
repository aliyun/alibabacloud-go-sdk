// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBegTime(v int64) *DescribeCapRequest
	GetBegTime() *int64
	SetDdosRegionId(v string) *DescribeCapRequest
	GetDdosRegionId() *string
	SetInstanceId(v string) *DescribeCapRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeCapRequest
	GetInstanceType() *string
	SetInternetIp(v string) *DescribeCapRequest
	GetInternetIp() *string
}

type DescribeCapRequest struct {
	// The start time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > You can call the [DescribeDdosEventList](https://help.aliyun.com/document_detail/354236.html) operation to query the start time of each DDoS attack event that occurred on an asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1637812279000
	BegTime *int64 `json:"BegTime,omitempty" xml:"BegTime,omitempty"`
	// The region ID of the asset that is under DDoS attacks. The asset is assigned a public IP address.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The ID of the asset that is under DDoS attacks.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp10bclrt56fblts****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset that is under DDoS attacks. The asset is assigned a public IP address. Valid values:
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
	// The public IP address of the asset that is under DDoS attacks.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeCapRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapRequest) GetBegTime() *int64 {
	return s.BegTime
}

func (s *DescribeCapRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeCapRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCapRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCapRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCapRequest) SetBegTime(v int64) *DescribeCapRequest {
	s.BegTime = &v
	return s
}

func (s *DescribeCapRequest) SetDdosRegionId(v string) *DescribeCapRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceId(v string) *DescribeCapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapRequest) SetInstanceType(v string) *DescribeCapRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapRequest) SetInternetIp(v string) *DescribeCapRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeCapRequest) Validate() error {
	return dara.Validate(s)
}
