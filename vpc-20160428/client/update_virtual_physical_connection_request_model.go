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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values: Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including required parameters, request syntax, and instance status. If the request fails to pass the dry run, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The estimated bandwidth value of the hosted connection. The estimated bandwidth value takes effect only after the payment is completed.
	//
	// Valid values: **50M**, **100M**, **200M**, **300M**, **400M**, **500M**, **1G**, **2G**, **5G**, **8G**, and **10G**.
	//
	// >  **2G**, **5G**, **8G**, and **10G*	- are unavailable by default. If you want to use these bandwidth values, contact your account manager.
	//
	// **M*	- indicates Mbit/s and **G*	- indicates Gbit/s.
	//
	// example:
	//
	// 50M
	ExpectSpec *string `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	// The ID of the hosted connection over Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the hosted connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The VLAN ID of the hosted connection over Express Connect circuit. Valid values: **0*	- to **2999**.
	//
	// 	- If the VLAN ID is set to **0**, it indicates that the switch port of the virtual border router (VBR) is a Layer 3 router interface instead of a VLAN interface. When a Layer 3 router interface is used, each Express Connect circuit corresponds to a VBR.
	//
	// 	- If the VLAN ID is set to a value from **1*	- to **2999**, the switch port of the VBR is a Layer 3 VLAN subinterface. When a Layer 3 VLAN subinterface is used, each VLAN ID corresponds to one VBR. In this case, the Express Connect circuit with which the VBR is associated can be used to connect to virtual private clouds (VPCs) that belong to different Alibaba Cloud accounts. VBRs in different VLANs are isolated from each other at Layer 2.
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
