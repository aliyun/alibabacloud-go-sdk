// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateVirtualBorderRouterRequest
	GetBandwidth() *int64
	SetCircuitCode(v string) *CreateVirtualBorderRouterRequest
	GetCircuitCode() *string
	SetClientToken(v string) *CreateVirtualBorderRouterRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVirtualBorderRouterRequest
	GetDescription() *string
	SetEnableIpv6(v bool) *CreateVirtualBorderRouterRequest
	GetEnableIpv6() *bool
	SetLocalGatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetLocalGatewayIp() *string
	SetLocalIpv6GatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetLocalIpv6GatewayIp() *string
	SetName(v string) *CreateVirtualBorderRouterRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetPeerGatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetPeerGatewayIp() *string
	SetPeerIpv6GatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetPeerIpv6GatewayIp() *string
	SetPeeringIpv6SubnetMask(v string) *CreateVirtualBorderRouterRequest
	GetPeeringIpv6SubnetMask() *string
	SetPeeringSubnetMask(v string) *CreateVirtualBorderRouterRequest
	GetPeeringSubnetMask() *string
	SetPhysicalConnectionId(v string) *CreateVirtualBorderRouterRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *CreateVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVirtualBorderRouterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetTags(v []*CreateVirtualBorderRouterRequestTags) *CreateVirtualBorderRouterRequest
	GetTags() []*CreateVirtualBorderRouterRequestTags
	SetVbrOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetVbrOwnerId() *int64
	SetVlanId(v int32) *CreateVirtualBorderRouterRequest
	GetVlanId() *int32
}

type CreateVirtualBorderRouterRequest struct {
	// The bandwidth of the VBR. Unit: Mbit/s.
	//
	// 	- When you create a VBR for a dedicated connection, valid values are **50**, **100**, **200**, **300**, **400**, **500**, **1000**, **2048**, **5120**, **8192**, **10240**, **20480**, **40960**, **50120**, **61440**, and **102400**.
	//
	// 	- You do not need to set this parameter when you create a VBR for a hosted connection. The bandwidth is already configured when the hosted connection is created.
	//
	// example:
	//
	// 100
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// >  Only the owner of the Express Connect circuit can set this parameter.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// >  If you do not set this parameter, the system automatically sets **ClientToken*	- to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VBR.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **true**: enables IPv6.
	//
	// 	- **false*	- (default): disables IPv6.
	//
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The IP address of the VBR. Only the owner of the VBR can set or modify this parameter.
	//
	// When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	LocalGatewayIp *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	// The IPv6 address of the VBR. Only the owner of the VBR can set or modify this parameter.
	//
	// When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 2001:XXXX:3c4d:0015:0000:0000:0000:1a2b
	LocalIpv6GatewayIp *string `json:"LocalIpv6GatewayIp,omitempty" xml:"LocalIpv6GatewayIp,omitempty"`
	// The name of the VBR.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IP address of the gateway device in the data center. Only the owner of the VBR can set or modify this parameter.
	//
	// When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 116.62.XX.XX
	PeerGatewayIp *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	// The IPv6 address of the gateway device in the data center. Only the owner of the VBR can set or modify this parameter.
	//
	// When you create a VBR for the owner of the Express Connect circuit, this parameter is required.
	//
	// example:
	//
	// 2001:XXXX:4:4:4:4:4:4
	PeerIpv6GatewayIp *string `json:"PeerIpv6GatewayIp,omitempty" xml:"PeerIpv6GatewayIp,omitempty"`
	// The subnet mask of the IPv6 addresses of the VBR and the gateway device in the data center.
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
	// 255.255.255.252
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// You can create a VBR for a dedicated connection or a hosted connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-2zextbehcx****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource group, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tags []*CreateVirtualBorderRouterRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The account ID of the VBR owner.
	//
	// The default value is the ID of the current Alibaba Cloud account.
	//
	// example:
	//
	// 168811111****
	VbrOwnerId *int64 `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
	// The VLAN ID of the VBR. Valid values: **0 to 2999**.
	//
	// >  Only the owner of the Express Connect circuit can set this parameter. The VLAN IDs of two VBRs of the same the Express Connect circuit must be different.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	VlanId *int32 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s CreateVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateVirtualBorderRouterRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *CreateVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVirtualBorderRouterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVirtualBorderRouterRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateVirtualBorderRouterRequest) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetLocalIpv6GatewayIp() *string {
	return s.LocalIpv6GatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetName() *string {
	return s.Name
}

func (s *CreateVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetPeerIpv6GatewayIp() *string {
	return s.PeerIpv6GatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetPeeringIpv6SubnetMask() *string {
	return s.PeeringIpv6SubnetMask
}

func (s *CreateVirtualBorderRouterRequest) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *CreateVirtualBorderRouterRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreateVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirtualBorderRouterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetTags() []*CreateVirtualBorderRouterRequestTags {
	return s.Tags
}

func (s *CreateVirtualBorderRouterRequest) GetVbrOwnerId() *int64 {
	return s.VbrOwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetVlanId() *int32 {
	return s.VlanId
}

func (s *CreateVirtualBorderRouterRequest) SetBandwidth(v int64) *CreateVirtualBorderRouterRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetCircuitCode(v string) *CreateVirtualBorderRouterRequest {
	s.CircuitCode = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetClientToken(v string) *CreateVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetDescription(v string) *CreateVirtualBorderRouterRequest {
	s.Description = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetEnableIpv6(v bool) *CreateVirtualBorderRouterRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetLocalGatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.LocalGatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetLocalIpv6GatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.LocalIpv6GatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetName(v string) *CreateVirtualBorderRouterRequest {
	s.Name = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetOwnerAccount(v string) *CreateVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeerGatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.PeerGatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeerIpv6GatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.PeerIpv6GatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeeringIpv6SubnetMask(v string) *CreateVirtualBorderRouterRequest {
	s.PeeringIpv6SubnetMask = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeeringSubnetMask(v string) *CreateVirtualBorderRouterRequest {
	s.PeeringSubnetMask = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPhysicalConnectionId(v string) *CreateVirtualBorderRouterRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetRegionId(v string) *CreateVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetResourceGroupId(v string) *CreateVirtualBorderRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *CreateVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetTags(v []*CreateVirtualBorderRouterRequestTags) *CreateVirtualBorderRouterRequest {
	s.Tags = v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetVbrOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.VbrOwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetVlanId(v int32) *CreateVirtualBorderRouterRequest {
	s.VlanId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVirtualBorderRouterRequestTags struct {
	// The tag key. You must enter at least one tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualBorderRouterRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterRequestTags) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateVirtualBorderRouterRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateVirtualBorderRouterRequestTags) SetKey(v string) *CreateVirtualBorderRouterRequestTags {
	s.Key = &v
	return s
}

func (s *CreateVirtualBorderRouterRequestTags) SetValue(v string) *CreateVirtualBorderRouterRequestTags {
	s.Value = &v
	return s
}

func (s *CreateVirtualBorderRouterRequestTags) Validate() error {
	return dara.Validate(s)
}
