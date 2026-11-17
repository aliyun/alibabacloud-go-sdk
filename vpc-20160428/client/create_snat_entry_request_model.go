// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSnatEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateSnatEntryRequest
	GetDryRun() *bool
	SetEipAffinity(v int32) *CreateSnatEntryRequest
	GetEipAffinity() *int32
	SetNetworkInterfaceId(v string) *CreateSnatEntryRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *CreateSnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateSnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSnatEntryRequest
	GetResourceOwnerId() *int64
	SetSnatEntryName(v string) *CreateSnatEntryRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *CreateSnatEntryRequest
	GetSnatIp() *string
	SetSnatTableId(v string) *CreateSnatEntryRequest
	GetSnatTableId() *string
	SetSourceCIDR(v string) *CreateSnatEntryRequest
	GetSourceCIDR() *string
	SetSourceVSwitchId(v string) *CreateSnatEntryRequest
	GetSourceVSwitchId() *string
}

type CreateSnatEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The `ClientToken` parameter can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run on the request without creating the SNAT entry. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the SNAT entry is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable EIP affinity. Valid values:
	//
	// - **0*	- (default): disables EIP affinity.
	//
	// - **1**: enables EIP affinity.
	//
	// > After you enable EIP affinity, if the SNAT entry is associated with multiple EIPs or NAT IP addresses, the same EIP or NAT IP address is used when the same client accesses the same destination IP address. Otherwise, the client randomly selects an EIP or NAT IP address from the associated EIPs or NAT IP addresses.
	//
	// example:
	//
	// 1
	EipAffinity *int32 `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The ID of the network interface controller (NIC).
	//
	// > The IPv4 addresses of the network interface controller (NIC) are used as the SNAT addresses.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the NAT gateway is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the SNAT entry.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// When you add an SNAT entry to an Internet NAT gateway:
	//
	// 	- The SnatIp parameter is required.
	//
	// 	- This parameter specifies the EIPs in the SNAT entry. Separate multiple EIPs with commas (,).
	//
	// 	- If SnatIp specifies only one public IP address, the ECS instance uses the specified public IP address to access the Internet.
	//
	// 	- If SnatIp specifies multiple public IP addresses, the ECS instance randomly uses one of the public IP addresses in SnatIp to access the Internet.
	//
	// > When you specify multiple EIPs to configure an SNAT IP address pool, service connections are allocated to the multiple EIPs through a hash algorithm. Because service traffic of each connection is different, service traffic of the multiple EIPs may be uneven. Add all EIPs to the same Internet Shared Bandwidth to prevent service interruptions caused by the bandwidth limit of a single EIP.
	//
	// When you add an SNAT entry to a VPC NAT gateway:
	//
	// 	- This parameter specifies the NAT IP addresses in the SNAT entry. Separate multiple NAT IP addresses with commas (,).
	//
	// 	- You must specify either SnatIp or NetworkInterfaceId.
	//
	// example:
	//
	// 47.98.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The ID of the SNAT table.
	//
	// This parameter is required.
	//
	// example:
	//
	// stb-bp190wu8io1vgev****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	// The CIDR block of the VPC, vSwitch, or ECS instance. You can also specify a custom CIDR block.
	//
	// SNAT entries support the following granularities:
	//
	// - VPC granularity: the CIDR block of the VPC where the NAT gateway is deployed. All ECS instances in the VPC can access the Internet or external networks through SNAT rules.
	//
	// - vSwitch granularity: the CIDR block of a specified vSwitch, such as 192.168.1.0/24. All ECS instances in the vSwitch can access the Internet or external networks through SNAT rules.
	//
	// - ECS instance granularity: the IP address of a specified ECS instance, such as 192.168.1.1/32. The ECS instance can access the Internet or external networks through SNAT rules.
	//
	// - Custom CIDR block: all ECS instances within the specified CIDR block can access the Internet or external networks through SNAT rules.
	//
	// > You must specify either the **SourceCIDR*	- parameter or the **SourceVSwtichId*	- parameter.
	//
	// example:
	//
	// 10.1.1.0/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The ID of the vSwitch.
	//
	// 	- When you add an SNAT entry to an Internet NAT gateway, this parameter specifies that all ECS instances in the vSwitch can access the Internet through SNAT rules. When you specify multiple EIPs to configure an SNAT IP address pool, service connections are allocated to the multiple EIPs through a hash algorithm. Because service traffic of each connection is different, service traffic of the multiple EIPs may be uneven. Add all EIPs to the same Internet Shared Bandwidth to prevent service interruptions caused by the bandwidth limit of a single EIP.
	//
	// 	- When you add an SNAT entry to a VPC NAT gateway, this parameter specifies that all ECS instances in the vSwitch can access external networks through SNAT rules.
	//
	// > You must specify either the **SourceCIDR*	- parameter or the **SourceVSwtichId*	- parameter.
	//
	// example:
	//
	// vsw-bp1nhx2s9ui5o****
	SourceVSwitchId *string `json:"SourceVSwitchId,omitempty" xml:"SourceVSwitchId,omitempty"`
}

func (s CreateSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateSnatEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnatEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSnatEntryRequest) GetEipAffinity() *int32 {
	return s.EipAffinity
}

func (s *CreateSnatEntryRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *CreateSnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSnatEntryRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *CreateSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *CreateSnatEntryRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *CreateSnatEntryRequest) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *CreateSnatEntryRequest) GetSourceVSwitchId() *string {
	return s.SourceVSwitchId
}

func (s *CreateSnatEntryRequest) SetClientToken(v string) *CreateSnatEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnatEntryRequest) SetDryRun(v bool) *CreateSnatEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSnatEntryRequest) SetEipAffinity(v int32) *CreateSnatEntryRequest {
	s.EipAffinity = &v
	return s
}

func (s *CreateSnatEntryRequest) SetNetworkInterfaceId(v string) *CreateSnatEntryRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetOwnerAccount(v string) *CreateSnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSnatEntryRequest) SetOwnerId(v int64) *CreateSnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetRegionId(v string) *CreateSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetResourceOwnerAccount(v string) *CreateSnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSnatEntryRequest) SetResourceOwnerId(v int64) *CreateSnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatEntryName(v string) *CreateSnatEntryRequest {
	s.SnatEntryName = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatIp(v string) *CreateSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatTableId(v string) *CreateSnatEntryRequest {
	s.SnatTableId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceCIDR(v string) *CreateSnatEntryRequest {
	s.SourceCIDR = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceVSwitchId(v string) *CreateSnatEntryRequest {
	s.SourceVSwitchId = &v
	return s
}

func (s *CreateSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
