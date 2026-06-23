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
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run to check for required parameters, the request format, and the instance status. The VLAN ID of the virtual physical connection is not modified. If the check fails, an error message is returned. If it passes, the request ID is returned.
	//
	// - **false*	- (default): Sends the request. If the check passes, the VLAN ID of the virtual physical connection is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The expected bandwidth of the virtual physical connection. The new bandwidth takes effect only after the payment is complete.
	//
	// Valid values: **50M**, **100M**, **200M**, **300M**, **400M**, **500M**, **1G**, **2G**, **5G**, **8G**, and **10G**.
	//
	// <props="china">
	//
	// > Bandwidth settings of **2G**, **5G**, **8G**, and **10G*	- are not enabled by default. To use these settings, contact your account manager.
	//
	//
	//
	// <props="intl">
	//
	// > Bandwidth settings of **2G**, **5G**, **8G**, and **10G*	- are not enabled by default. To use these settings, contact your account manager.
	//
	//
	//
	// Units: **M*	- indicates Mbps and **G*	- indicates Gbps.
	//
	// example:
	//
	// 50M
	ExpectSpec *string `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	// The ID of the virtual physical connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the virtual physical connection is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that ensures the idempotence of the request.
	//
	// A client-generated value that must be unique across requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- is different for each request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The VLAN ID of the virtual physical connection. Valid values: **0*	- to **2999**.
	//
	// - If you set the VLAN ID to **0**, the physical switch port of the Virtual Border Router (VBR) operates in Layer 3 routed interface mode. In this mode, one physical connection corresponds to one VBR.
	//
	// - If you set the VLAN ID to a value from **1*	- to **2999**, the physical switch port of the VBR uses a VLAN-based Layer 3 subinterface. In this mode, each VLAN ID corresponds to one VBR. The physical connection can be attached to VPCs that belong to different accounts. VBRs in different VLANs are isolated at Layer 2 and cannot communicate with each other.
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
