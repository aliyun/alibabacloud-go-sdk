// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySnatEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifySnatEntryRequest
	GetDryRun() *bool
	SetEipAffinity(v int32) *ModifySnatEntryRequest
	GetEipAffinity() *int32
	SetNetworkInterfaceId(v string) *ModifySnatEntryRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *ModifySnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnatEntryRequest
	GetResourceOwnerId() *int64
	SetSnatEntryId(v string) *ModifySnatEntryRequest
	GetSnatEntryId() *string
	SetSnatEntryName(v string) *ModifySnatEntryRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *ModifySnatEntryRequest
	GetSnatIp() *string
	SetSnatTableId(v string) *ModifySnatEntryRequest
	GetSnatTableId() *string
}

type ModifySnatEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the SNAT entry. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, a 2xx HTTP status code is returned and the SNAT entry is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable EIP affinity. Valid values:
	//
	// - **0**: Disables EIP affinity.
	//
	// - **1**: Enables EIP affinity.
	//
	// > After EIP affinity is enabled, if the SNAT entry is associated with multiple EIPs or NAT IP addresses, the same client uses the same EIP or NAT IP address for access. Otherwise, the client randomly selects an EIP or NAT IP address from the associated ones for access.
	//
	// example:
	//
	// 1
	EipAffinity *int32 `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The ID of the elastic network interface (ENI).
	//
	// > The IPv4 addresses of the ENI are used as the SNAT addresses.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SNAT entry that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-bp1vcgcf8tm0plqcg****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// - When you modify a SNAT entry for an Internet NAT gateway, this parameter specifies the EIP in the SNAT entry. Separate multiple EIPs with commas (,).
	//
	//
	//
	// > When you allocate multiple EIPs to configure a SNAT IP address pool, connections are distributed across the EIPs by using a hash algorithm. Because the traffic volume of each connection varies, service traffic may be unevenly distributed across the EIPs. To prevent service interruptions caused by bandwidth limits on a single EIP, add all EIPs to the same Internet Shared Bandwidth instance.
	//
	// - When you modify a SNAT entry for a VPC NAT gateway, this parameter specifies the NAT IP address in the SNAT entry. Separate multiple NAT IP addresses with commas (,).
	//
	// - The SnatIp and NetworkInterfaceId parameters cannot be specified at the same time.
	//
	// example:
	//
	// 47.98.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The ID of the SNAT table to which the SNAT entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// stb-8vbczigrhop8x5u3t****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
}

func (s ModifySnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnatEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifySnatEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySnatEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifySnatEntryRequest) GetEipAffinity() *int32 {
	return s.EipAffinity
}

func (s *ModifySnatEntryRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ModifySnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *ModifySnatEntryRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *ModifySnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *ModifySnatEntryRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *ModifySnatEntryRequest) SetClientToken(v string) *ModifySnatEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySnatEntryRequest) SetDryRun(v bool) *ModifySnatEntryRequest {
	s.DryRun = &v
	return s
}

func (s *ModifySnatEntryRequest) SetEipAffinity(v int32) *ModifySnatEntryRequest {
	s.EipAffinity = &v
	return s
}

func (s *ModifySnatEntryRequest) SetNetworkInterfaceId(v string) *ModifySnatEntryRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetOwnerAccount(v string) *ModifySnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnatEntryRequest) SetOwnerId(v int64) *ModifySnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetRegionId(v string) *ModifySnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetResourceOwnerAccount(v string) *ModifySnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnatEntryRequest) SetResourceOwnerId(v int64) *ModifySnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatEntryId(v string) *ModifySnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatEntryName(v string) *ModifySnatEntryRequest {
	s.SnatEntryName = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatIp(v string) *ModifySnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *ModifySnatEntryRequest) SetSnatTableId(v string) *ModifySnatEntryRequest {
	s.SnatTableId = &v
	return s
}

func (s *ModifySnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
