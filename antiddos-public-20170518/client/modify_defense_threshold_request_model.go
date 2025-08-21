// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBps(v int32) *ModifyDefenseThresholdRequest
	GetBps() *int32
	SetClientToken(v string) *ModifyDefenseThresholdRequest
	GetClientToken() *string
	SetDdosRegionId(v string) *ModifyDefenseThresholdRequest
	GetDdosRegionId() *string
	SetInstanceId(v string) *ModifyDefenseThresholdRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyDefenseThresholdRequest
	GetInstanceType() *string
	SetInternetIp(v string) *ModifyDefenseThresholdRequest
	GetInternetIp() *string
	SetIsAuto(v bool) *ModifyDefenseThresholdRequest
	GetIsAuto() *bool
	SetPps(v int32) *ModifyDefenseThresholdRequest
	GetPps() *int32
}

type ModifyDefenseThresholdRequest struct {
	// The traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset. When you modify Bps, Pps is required. Otherwise, Bps does not take effect.
	//
	// You can use the monitoring tool that is provided by the asset to query the Internet traffic of the asset:
	//
	// 	- If the asset is an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// 	- If the asset is an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// example:
	//
	// 100
	Bps         *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the asset for which you want to change the scrubbing thresholds.
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
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Specifies whether to automatically adjust the scrubbing threshold based on the traffic load on the asset. Valid values:
	//
	// 	- **true**: automatically adjusts the scrubbing thresholds. You do not need to configure the **Bps*	- and **Pps*	- parameters.
	//
	// 	- **false**: The scrubbing threshold is not automatically adjusted. You must configure the **Bps*	- and **Pps*	- parameters.
	//
	// Default value: false.
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
	// example:
	//
	// 70000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s ModifyDefenseThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseThresholdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdRequest) GetBps() *int32 {
	return s.Bps
}

func (s *ModifyDefenseThresholdRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDefenseThresholdRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *ModifyDefenseThresholdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseThresholdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyDefenseThresholdRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ModifyDefenseThresholdRequest) GetIsAuto() *bool {
	return s.IsAuto
}

func (s *ModifyDefenseThresholdRequest) GetPps() *int32 {
	return s.Pps
}

func (s *ModifyDefenseThresholdRequest) SetBps(v int32) *ModifyDefenseThresholdRequest {
	s.Bps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetClientToken(v string) *ModifyDefenseThresholdRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetDdosRegionId(v string) *ModifyDefenseThresholdRequest {
	s.DdosRegionId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceId(v string) *ModifyDefenseThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInstanceType(v string) *ModifyDefenseThresholdRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetInternetIp(v string) *ModifyDefenseThresholdRequest {
	s.InternetIp = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetIsAuto(v bool) *ModifyDefenseThresholdRequest {
	s.IsAuto = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) SetPps(v int32) *ModifyDefenseThresholdRequest {
	s.Pps = &v
	return s
}

func (s *ModifyDefenseThresholdRequest) Validate() error {
	return dara.Validate(s)
}
