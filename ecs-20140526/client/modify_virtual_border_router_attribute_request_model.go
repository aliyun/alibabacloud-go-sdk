// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBorderRouterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCircuitCode(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetCircuitCode() *string
	SetClientToken(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetDescription() *string
	SetLocalGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetLocalGatewayIp() *string
	SetName(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetOwnerId() *int64
	SetPeerGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeerGatewayIp() *string
	SetPeeringSubnetMask(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetPeeringSubnetMask() *string
	SetRegionId(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVirtualBorderRouterAttributeRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetUserCidr() *string
	SetVbrId(v string) *ModifyVirtualBorderRouterAttributeRequest
	GetVbrId() *string
	SetVlanId(v int32) *ModifyVirtualBorderRouterAttributeRequest
	GetVlanId() *int32
}

type ModifyVirtualBorderRouterAttributeRequest struct {
	// The circuit code of the Express Connect circuit. The circuit code is provided by your connectivity provider.
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// A client-generated token that ensures the request is idempotent. The token must be unique for each request, contain only ASCII characters, and not exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VBR.
	//
	// The description must be 2 to 256 characters long. It must start with a letter and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The local gateway IP address for the BGP session.
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// The name of the VBR.
	//
	// The name must be 2 to 128 characters long. It must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The peer gateway IP address for the BGP session.
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The subnet mask for the `LocalGatewayIp` and `PeerGatewayIp`.
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// The ID of the region where the virtual border router (VBR) is deployed.
	//
	// To get a list of the latest regions, call the `DescribeRegions` operation.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The CIDR block of the on-premises network that is connected to the VBR.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The ID of the VBR to modify.
	//
	// This parameter is required.
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The VLAN ID of the VBR. The VLAN ID must be unique on the physical connection.
	//
	// Valid values: 0 to 4095.
	//
	// A value of 0 indicates that the VBR does not use a VLAN.
	VlanId *int32 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s ModifyVirtualBorderRouterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBorderRouterAttributeRequest) GoString() string {
	return s.String()
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

func (s *ModifyVirtualBorderRouterAttributeRequest) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
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

func (s *ModifyVirtualBorderRouterAttributeRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *ModifyVirtualBorderRouterAttributeRequest) GetVlanId() *int32 {
	return s.VlanId
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

func (s *ModifyVirtualBorderRouterAttributeRequest) SetLocalGatewayIp(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.LocalGatewayIp = &v
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

func (s *ModifyVirtualBorderRouterAttributeRequest) SetUserCidr(v string) *ModifyVirtualBorderRouterAttributeRequest {
	s.UserCidr = &v
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
