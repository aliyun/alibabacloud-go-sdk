// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBorderRouterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedPhysicalConnections(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetAssociatedPhysicalConnections() *string
	SetBandwidth(v int32) *ModifyVirtualBorderRouterAttributeRequest
	GetBandwidth() *int32
	SetCircuitCode(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetCircuitCode() *string
	SetClientToken(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetDescription() *string
	SetDetectMultiplier(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetDetectMultiplier() *int64
	SetEnableIpv6(v bool) *ModifyVirtualBorderRouterAttributeRequest
	GetEnableIpv6() *bool
	SetLocalGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetLocalGatewayIp() *string
	SetLocalIpv6GatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetLocalIpv6GatewayIp() *string
	SetMinRxInterval(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetMinRxInterval() *int64
	SetMinTxInterval(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetMinTxInterval() *int64
	SetMtu(v int32) *ModifyVirtualBorderRouterAttributeRequest
	GetMtu() *int32
	SetName(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetOwnerId() *int64
	SetPeerGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeerGatewayIp() *string
	SetPeerIpv6GatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeerIpv6GatewayIp() *string
	SetPeeringIpv6SubnetMask(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeeringIpv6SubnetMask() *string
	SetPeeringSubnetMask(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeeringSubnetMask() *string
	SetRegionId(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetResourceOwnerId() *int64
	SetSitelinkEnable(v bool) *ModifyVirtualBorderRouterAttributeRequest
	GetSitelinkEnable() *bool
	SetVbrId(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetVbrId() *string
	SetVlanId(v int32) *ModifyVirtualBorderRouterAttributeRequest
	GetVlanId() *int32
}

type ModifyVirtualBorderRouterAttributeRequest struct {
	// The list of Express Connect circuits associated with the VBR, which contains the following parameters:
	//
	// - **VlanId**: The VLAN ID of the VBR instance.
	//
	// - **CircuitCode**: The circuit encoding provided by the carrier for the Express Connect circuit.
	//
	// - **LocalGatewayIp**: The Alibaba Cloud-side IP address of the VBR instance.
	//
	// - **PeerGatewayIp**: The client-side IP address of the VBR instance.
	//
	// - **PeeringSubnetMask**: The subnet mask for the Alibaba Cloud-side and client-side IP addresses of the VBR instance.
	//
	// - **LocalIpv6GatewayIp**: The Alibaba Cloud-side IPv6 address of the VBR instance.
	//
	// - **PeerIpv6GatewayIp**: The client-side IPv6 address of the VBR instance.
	//
	// - **PeeringIpv6SubnetMask**: The subnet mask for the Alibaba Cloud-side and client-side IPv6 addresses of the VBR instance.
	//
	// - **ipv6Enable**: Enables IPv6.
	//
	// - **PhysicalConnectionId**: The Express Connect circuit instance ID.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "VlanId": "5",
	//
	//     "CircuitCode": "longtel001",
	//
	//     "LocalGatewayIp": "192.168.XX.XX",
	//
	//     "PeerGatewayIp": "192.168.XX.XX",
	//
	//     "PeeringSubnetMask": "255.255.255.252",
	//
	//     "LocalIpv6GatewayIp": "fd00:XXXX:1030:3f41::1",
	//
	//     "PeerIpv6GatewayIp": "fd00:XXXX:1030:3f41::2",
	//
	//     "PeeringIpv6SubnetMask": "fd00:0:1030:3f41::/126",
	//
	//     "ipv6Enable": true,
	//
	//     "PhysicalConnectionId": "pc-kojok19****"
	//
	//   }
	//
	// ]
	AssociatedPhysicalConnections *string `json:"AssociatedPhysicalConnections,omitempty" xml:"AssociatedPhysicalConnections,omitempty"`
	// The bandwidth value. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The circuit code provided by the carrier for the Express Connect circuit.
	//
	//
	//
	// > Only the owner of the Express Connect circuit can specify this parameter.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VBR.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The detection multiplier, which specifies the maximum number of consecutive packet losses allowed by the receiver from the sender. This parameter is used to detect whether the link is normal.
	//
	// Valid values: **3 to 10**.
	//
	// example:
	//
	// 3
	DetectMultiplier *int64 `json:"DetectMultiplier,omitempty" xml:"DetectMultiplier,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false*	- (default): Disabled.
	//
	// example:
	//
	// false
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The Alibaba Cloud-side IP address of the VBR instance.
	//
	// This property can be specified or modified only by the VBR owner.
	//
	// example:
	//
	// 192.168.XX.XX
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// The Alibaba Cloud-side IPv6 address of the VBR instance.
	//
	// example:
	//
	// 2001:XXXX:3c4d:0015:0000:0000:0000:1a2b
	LocalIpv6GatewayIp *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	// The receive interval of BFD packets. Valid values: **200 to 1000**. Unit: ms.
	//
	// example:
	//
	// 300
	MinRxInterval *int64 `json:"MinRxInterval,omitempty" xml:"MinRxInterval,omitempty"`
	// The alert interval for sending Bidirectional Forwarding Detection (BFD) packets. Valid values: **200 to 1000**. Unit: ms.
	//
	// example:
	//
	// 300
	MinTxInterval *int64 `json:"MinTxInterval,omitempty" xml:"MinTxInterval,omitempty"`
	// The MTU value supported by the VBR. Valid values: 1500 and 8500.
	//
	// This value can be set only when the VBR is attached to an Express Connect Router (ECR). This value also affects all other VBRs and VPCs within the same ECR.
	//
	// example:
	//
	// 1500
	Mtu *int32 `json:"Mtu,omitempty" xml:"Mtu,omitempty"`
	// The name of the VBR.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// VBR
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The client-side IP address of the VBR instance.
	//
	// This property can be specified or modified only by the VBR owner.
	//
	// example:
	//
	// 192.168.XX.X
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The client-side IPv6 address of the VBR instance.
	//
	// - This property can be specified or modified only by the VBR owner.
	//
	// - This parameter is required when the Express Connect circuit owner creates a VBR instance. It is not required when creating a VBR instance for another account.
	//
	// example:
	//
	// 2001:XXXX:3c4d:0015:0000:0000:0000:2a2b
	PeerIpv6GatewayIp *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	// The subnet mask for the Alibaba Cloud-side and client-side IPv6 addresses of the VBR instance.
	//
	// The two IPv6 addresses must be in the same subnet.
	//
	// example:
	//
	// 2408:4004:cc:400::/56
	PeeringIpv6SubnetMask *string `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	// The subnet mask for the Alibaba Cloud-side and client-side IP addresses of the VBR instance. This property can be specified or modified only by the VBR owner.
	//
	// The two IP addresses must be in the same subnet.
	//
	// example:
	//
	// 255.255.255.252
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// The region ID of the VBR.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to allow inter-IDC service access. Valid values:
	//
	// - true: Allowed.
	//
	// - false (default): Not allowed.
	//
	// example:
	//
	// false
	SitelinkEnable *bool `json:"SitelinkEnable,omitempty" xml:"SitelinkEnable,omitempty"`
	// The VBR instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp1lhl0taikrte****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The VLAN ID of the VBR. Valid values: **0 to 2999**.
	//
	// > Only the owner of the Express Connect circuit can specify this parameter. The VLAN IDs of two VBRs on the same Express Connect circuit must be different.
	//
	// example:
	//
	// 0
	VlanId *int32 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s ModifyVirtualBorderRouterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBorderRouterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetAssociatedPhysicalConnections() *string {
	return s.AssociatedPhysicalConnections
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetDetectMultiplier() *int64 {
	return s.DetectMultiplier
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetMinRxInterval() *int64 {
	return s.MinRxInterval
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetMinTxInterval() *int64 {
	return s.MinTxInterval
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetMtu() *int32 {
	return s.Mtu
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetSitelinkEnable() *bool {
	return s.SitelinkEnable
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetVlanId() *int32 {
	return s.VlanId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetAssociatedPhysicalConnections(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.AssociatedPhysicalConnections = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetBandwidth(v int32) *ModifyVirtualBorderRouterAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetCircuitCode(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.CircuitCode = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetClientToken(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetDescription(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetDetectMultiplier(v int64) *ModifyVirtualBorderRouterAttributeRequest {
	s.DetectMultiplier = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetEnableIpv6(v bool) *ModifyVirtualBorderRouterAttributeRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetLocalGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.LocalGatewayIp = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetLocalIpv6GatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetMinRxInterval(v int64) *ModifyVirtualBorderRouterAttributeRequest {
	s.MinRxInterval = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetMinTxInterval(v int64) *ModifyVirtualBorderRouterAttributeRequest {
	s.MinTxInterval = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetMtu(v int32) *ModifyVirtualBorderRouterAttributeRequest {
	s.Mtu = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetName(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetPeerGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.PeerGatewayIp = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetPeerIpv6GatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetPeeringIpv6SubnetMask(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetPeeringSubnetMask(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.PeeringSubnetMask = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetRegionId(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetResourceOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetSitelinkEnable(v bool) *ModifyVirtualBorderRouterAttributeRequest {
	s.SitelinkEnable = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetVbrId(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.VbrId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) SetVlanId(v int32) *ModifyVirtualBorderRouterAttributeRequest {
	s.VlanId = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
