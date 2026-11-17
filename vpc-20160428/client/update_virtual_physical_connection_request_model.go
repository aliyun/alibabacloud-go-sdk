// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *UpdateVirtualPhysicalConnectionRequest
	GetDryRun() *bool
	SetExpectSpec(v string) *UpdateVirtualPhysicalConnectionRequest
	GetExpectSpec() *string
	SetInstanceId(v string) *UpdateVirtualPhysicalConnectionRequest
	GetInstanceId() *string
	SetRegionId(v string) *UpdateVirtualPhysicalConnectionRequest
	GetRegionId() *string
	SetToken(v string) *UpdateVirtualPhysicalConnectionRequest
	GetToken() *string
	SetVlanId(v int64) *UpdateVirtualPhysicalConnectionRequest
	GetVlanId() *int64
}

type UpdateVirtualPhysicalConnectionRequest struct {
	// Specifies whether to only perform a dry run for the request. Valid values:
	//
	// - **true**: The system performs a dry run without modifying the VLAN ID of the shared Express Connect circuit. The system checks whether the required parameters are specified, the request format, and the instance status. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the corresponding request ID is returned.
	//
	// - **false*	- (default): The system sends a normal request. After the request passes the dry run, the VLAN ID of the shared Express Connect circuit is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The expected bandwidth of the shared Express Connect circuit. The bandwidth takes effect only after the payment is completed.
	//
	// Valid values: **50M**, **100M**, **200M**, **300M**, **400M**, **500M**, **1G**, **2G**, **5G**, **8G**, and **10G**.
	//
	// <props="china">
	//
	// > The bandwidth values of **2G**, **5G**, **8G**, and **10G*	- are not available by default. To use these values, contact your account manager.
	//
	// <props="intl">
	//
	// > The bandwidth values of **2G**, **5G**, **8G**, and **10G*	- are not available by default. To use these values, contact your account manager.
	//
	// Unit: **M*	- indicates Mbps, and **G*	- indicates Gbps.
	//
	// example:
	//
	// 50M
	ExpectSpec *string `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	// The instance ID of the shared Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the shared Express Connect circuit.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The VLAN ID of the shared Express Connect circuit. Valid values: **0*	- to **2999**.
	//
	// - A value of **0*	- specifies that the physical vSwitch port of the Virtual Border Router (VBR), which acts as a vRouter, does not use the VLAN mode, but uses the Layer 3 routing interface mode. In the Layer 3 routing interface mode, each Express Connect circuit corresponds to one VBR.
	//
	// - A value of **1*	- to **2999*	- specifies that the physical vSwitch port of the VBR uses a VLAN-based Layer 3 subinterface. In the Layer 3 subinterface mode, each VLAN ID corresponds to one VBR. In this case, the Express Connect circuit of the VBR can be used to connect VPCs that belong to different Alibaba Cloud accounts. Layer 2 network isolation is implemented for VBRs in different VLANs, and they cannot communicate with each other.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	VlanId *int64 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s UpdateVirtualPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetExpectSpec() *string {
	return s.ExpectSpec
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetToken() *string {
	return s.Token
}

func (s *UpdateVirtualPhysicalConnectionRequest) GetVlanId() *int64 {
	return s.VlanId
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetDryRun(v bool) *UpdateVirtualPhysicalConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetExpectSpec(v string) *UpdateVirtualPhysicalConnectionRequest {
	s.ExpectSpec = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetInstanceId(v string) *UpdateVirtualPhysicalConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetRegionId(v string) *UpdateVirtualPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetToken(v string) *UpdateVirtualPhysicalConnectionRequest {
	s.Token = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) SetVlanId(v int64) *UpdateVirtualPhysicalConnectionRequest {
	s.VlanId = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
