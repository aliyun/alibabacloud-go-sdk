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
	// The scrubbing threshold for traffic in Mbps. This value cannot exceed the peak public network traffic of the instance. If you specify Bps, you must also specify Pps. Otherwise, the change does not take effect.
	//
	// Use the monitoring tools of your instance to query its public network traffic:
	//
	// - For an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// - For an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// <props="china">
	//
	// - For an EIP instance, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
	//
	// example:
	//
	// 100
	Bps         *int32  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// The instance ID of the asset that is assigned a public IP address.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/354191.html) to query the IDs of the ECS, SLB, and EIP instances that belong to your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6idy3c57psf7vu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the asset that is assigned a public IP address. Valid values:
	//
	// - **ecs**: Elastic Compute Service (ECS) instance.
	//
	// - **slb**: Server Load Balancer (SLB) instance.
	//
	// - **eip**: Elastic IP Address (EIP) instance.
	//
	// - **ipv6**: IPv6 Gateway instance.
	//
	// - **swas**: simple application server instance.
	//
	// - **waf**: dedicated Web Application Firewall (WAF) instance.
	//
	// - **ga_basic**: basic Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public IP address of the asset.
	//
	// example:
	//
	// 192.0.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Specifies whether to automatically adjust the scrubbing threshold based on the traffic loads of the instance. Valid values:
	//
	// - **true**: The scrubbing threshold is automatically adjusted. You do not need to set the **Bps*	- and **Pps*	- parameters.
	//
	// - **false**: The scrubbing threshold is not automatically adjusted. You must set the **Bps*	- and **Pps*	- parameters.
	//
	// Default value: false
	//
	// example:
	//
	// false
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto,omitempty"`
	// The scrubbing threshold for packets per second (pps). This value cannot exceed the peak packet traffic of the instance. If you specify Pps, you must also specify Bps. Otherwise, the change does not take effect.
	//
	// Use the monitoring tools of your instance to query its packet traffic:
	//
	// - For an ECS instance, see [View instance monitoring information](https://help.aliyun.com/document_detail/25482.html).
	//
	// - For an SLB instance, see [View monitoring data](https://help.aliyun.com/document_detail/85982.html).
	//
	// <props="china">
	//
	// - For an EIP instance, see [View monitoring data](https://help.aliyun.com/document_detail/85354.html).
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
