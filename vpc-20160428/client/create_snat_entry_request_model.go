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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `client token` can contain only ASCII characters.
	//
	// **
	//
	// **Description*	- If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable EIP affinity. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// **
	//
	// **Description*	- After you enable EIP affinity, if multiple EIPs are associated with an SNAT entry, each client uses one EIP to access the Internet. If EIP affinity is disabled, each client uses a random EIP to access the Internet.
	//
	// example:
	//
	// 1
	EipAffinity        *int32  `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// Valid values:
	//
	// 	- ap-northeast-2-pop
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     ap-northeast-2-pop
	//
	//     <!-- -->
	//
	//     .
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
	// The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// 	- The EIPs in the SNAT entry when you add an SNAT entry to an Internet NAT gateway. Separate EIPs with commas (,).
	//
	// >  If you specify multiple EIPs in the SNAT IP address pool, the service connection is allocated to multiple EIPs by using the hashing algorithm. The traffic of each EIP may be different. Therefore, we recommend that you associate the EIPs with an Internet Shared Bandwidth instance to prevent service interruptions caused by bandwidth exhaustion.
	//
	// 	- When you add SNAT entries for a VPC NAT gateway, this parameter specifies the NAT IP addresses in the SNAT entry. Separate multiple NAT IP addresses with commas (,).
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
	// You can specify the CIDR block of a VPC, a vSwitch, or an ECS instance or enter a custom CIDR block.
	//
	// You can specify an SNAT entry in the following ways:
	//
	// 	- You can specify the CIDR block of the VPC where the NAT gateway is deployed. Then, all ECS instances in the VPC can access the Internet or external networks by using SNAT.
	//
	// 	- You can specify the CIDR block of a vSwitch, for example, 192.168.1.0/24. Then, the ECS instances in the vSwitch can access the Internet or external networks by using SNAT.
	//
	// 	- You can specify the IP address of an ECS instance, for example, 192.168.1.1/32. Then, the ECS instance can access the Internet or external networks by using SNAT.
	//
	// 	- You can specify a custom CIDR block. Then, all ECS instances within the specified CIDR block can access the Internet or external networks by using SNAT.
	//
	// When you add an SNAT entry to an Internet NAT gateway, if **SnatIp*	- is set to an EIP, the ECS instance uses the specified EIP to access the Internet.
	//
	// If **SnatIp*	- is set to multiple EIPs, the ECS instance randomly selects an EIP specified in the **SnatIp*	- parameter to access the Internet.
	//
	// You cannot specify this parameter and **SourceVSwtichId*	- at the same time. If **SourceVSwitchId*	- is specified, you cannot specify **SourceCIDR**. If **SourceCIDR*	- is specified, you cannot specify **SourceVSwitchId**.
	//
	// example:
	//
	// 10.1.1.0/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The ID of the vSwitch.
	//
	// 	- When you add an SNAT entry to an Internet NAT gateway, this parameter specifies that ECS instances in the vSwitch can use the SNAT entry to access the Internet. If you select multiple elastic IP addresses (EIPs) to create an SNAT address pool, connections are hashed to these EIPs. Network traffic may not be evenly distributed to the EIPs because the amount of traffic that passes through each connection varies. We recommend that you associate these EIPs with the same EIP bandwidth plan to prevent service interruptions due to the bandwidth limits on individual EIPs.
	//
	// 	- When you add an SNAT entry to a VPC NAT gateway, this parameter specifies that ECS instances in the vSwitch can use the SNAT entry to access external networks.
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
