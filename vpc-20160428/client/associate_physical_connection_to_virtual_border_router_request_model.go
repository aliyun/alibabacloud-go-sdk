// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePhysicalConnectionToVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCircuitCode(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetCircuitCode() *string
	SetClientToken(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetClientToken() *string
	SetEnableIpv6(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetEnableIpv6() *string
	SetLocalGatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetLocalGatewayIp() *string
	SetLocalIpv6GatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetLocalIpv6GatewayIp() *string
	SetOwnerAccount(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetPeerGatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetPeerGatewayIp() *string
	SetPeerIpv6GatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetPeerIpv6GatewayIp() *string
	SetPeeringIpv6SubnetMask(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetPeeringIpv6SubnetMask() *string
	SetPeeringSubnetMask(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetPeeringSubnetMask() *string
	SetPhysicalConnectionId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetVbrId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetVbrId() *string
	SetVlanId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest
	GetVlanId() *string
}

type AssociatePhysicalConnectionToVirtualBorderRouterRequest struct {
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// >  Only the Express Connect circuit owner can specify this parameter.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	EnableIpv6 *string `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The IP address of the gateway device on the Alibaba Cloud side.
	//
	// example:
	//
	// 192.168.XX.XX
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// The IPv6 address of the gateway device on the Alibaba Cloud side.
	//
	// example:
	//
	// 2001:XXXX:3c4d:0015:0000:0000:0000:1a2b
	LocalIpv6GatewayIp *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IP address of the gateway device on the user side.
	//
	// 	- Only the owner of the VBR can set or modify this parameter.
	//
	// 	- When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The IPv6 address of the gateway device in the data center.
	//
	// 	- Only the owner of the VBR can specify or modify this parameter.
	//
	// 	- When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 2001:XXXX:4:4:4:4:4:4
	PeerIpv6GatewayIp *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	// The subnet mask of the IPv6 addresses of the gateway devices on the user side and Alibaba Cloud side.
	//
	// The two IPv6 addresses must fall within the same subnet.
	//
	// example:
	//
	// 2408:4004:cc:400::/56
	PeeringIpv6SubnetMask *string `json:"PeeringIpv6SubnetMask,omitempty" xml:"PeeringIpv6SubnetMask,omitempty"`
	// The subnet mask of the IP addresses of the VBR and the gateway device in the data center.
	//
	// The two IP addresses must fall within the same subnet.
	//
	// example:
	//
	// 255.255.255.0
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1qrb3044eqixog****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp186tnz6rijyhj******
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The VLAN ID of the VBR. Valid values: **0 to 2999**.
	//
	// >  Only the Express Connect circuit owner can specify this parameter. Two VBRs associated with the same Express Connect circuit cannot use the same VLAN ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetEnableIpv6() *string {
	return s.EnableIpv6
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) GetVlanId() *string {
	return s.VlanId
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetCircuitCode(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.CircuitCode = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetClientToken(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetEnableIpv6(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetLocalGatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.LocalGatewayIp = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetLocalIpv6GatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetOwnerAccount(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetOwnerId(v int64) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetPeerGatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.PeerGatewayIp = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetPeerIpv6GatewayIp(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetPeeringIpv6SubnetMask(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetPeeringSubnetMask(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.PeeringSubnetMask = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetPhysicalConnectionId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetRegionId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetVbrId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) SetVlanId(v string) *AssociatePhysicalConnectionToVirtualBorderRouterRequest {
	s.VlanId = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
