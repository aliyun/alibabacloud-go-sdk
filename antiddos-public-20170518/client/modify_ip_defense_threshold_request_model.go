// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpDefenseThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBps(v int32) *ModifyIpDefenseThresholdRequest
	GetBps() *int32
	SetDdosRegionId(v string) *ModifyIpDefenseThresholdRequest
	GetDdosRegionId() *string
	SetInstanceId(v string) *ModifyIpDefenseThresholdRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyIpDefenseThresholdRequest
	GetInstanceType() *string
	SetInternetIp(v string) *ModifyIpDefenseThresholdRequest
	GetInternetIp() *string
	SetIsAuto(v bool) *ModifyIpDefenseThresholdRequest
	GetIsAuto() *bool
	SetPps(v int32) *ModifyIpDefenseThresholdRequest
	GetPps() *int32
}

type ModifyIpDefenseThresholdRequest struct {
	// The traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset. When you modify Bps, Pps is required. Otherwise, Bps does not take effect.
	//
	// You can use the monitoring tool that is provided by the asset to query the Internet traffic of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// 	- If the asset is an EIP, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
	//
	// example:
	//
	// 100
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
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
	// The ID of the asset.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) operation to query the IDs of ECS instances, SLB instances, and EIPs within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6idy3c57psf7vu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset. Valid values:
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
	// The IP address of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Specifies whether to automatically adjust the scrubbing threshold based on the traffic load on the asset. Valid values:
	//
	// 	- **true**: automatically adjusts the scrubbing threshold. You do not need to configure the **Bps*	- and **Pps*	- parameters.
	//
	// 	- **false**: The scrubbing threshold is not automatically adjusted. You must configure the **Bps*	- and **Pps*	- parameters. This is the default value.
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The packet scrubbing threshold. Unit: packets per second (PPS). When you modify Pps, Bps is required. Otherwise, Pps does not take effect.
	//
	// The packet scrubbing threshold cannot exceed the peak number of inbound or outbound packets, whichever is larger, of the asset. You can use the monitoring tool that is provided by the asset to query the number of packets of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// 	- If the asset is an EIP, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
	//
	// example:
	//
	// 70000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s ModifyIpDefenseThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdRequest) GetBps() *int32 {
	return s.Bps
}

func (s *ModifyIpDefenseThresholdRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *ModifyIpDefenseThresholdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyIpDefenseThresholdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyIpDefenseThresholdRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyIpDefenseThresholdRequest) GetIsAuto() *bool {
	return s.IsAuto
}

func (s *ModifyIpDefenseThresholdRequest) GetPps() *int32 {
	return s.Pps
}

func (s *ModifyIpDefenseThresholdRequest) SetBps(v int32) *ModifyIpDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyIpDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInstanceId(v string) *ModifyIpDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInstanceType(v string) *ModifyIpDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetInternetIp(v string) *ModifyIpDefenseThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetIsAuto(v bool) *ModifyIpDefenseThresholdRequest {
	s.IsAuto = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) SetPps(v int32) *ModifyIpDefenseThresholdRequest {
	s.Pps = &v
	return s
}

func (s *ModifyIpDefenseThresholdRequest) Validate() error {
	return dara.Validate(s)
}
