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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Whether to perform a dry run of this request, with values:
	//
	// - **true**: Sends a check request without modifying the SNAT entry. The checks include whether the required parameters are filled in, the request format, and business restrictions. If the check fails, the corresponding error is returned. If the check passes, an error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): Sends a normal request. After passing the check, it returns a 2xx HTTP status code and modifies the SNAT entry.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Whether to enable IP affinity. Values:
	//
	// - **0**: Disable IP affinity.
	//
	//  - **1**: Enable IP affinity.
	//
	// > After enabling the IP affinity switch, if an SNAT entry is bound to multiple EIPs or NAT IPs, the same client will use the same EIP or NAT IP for access; otherwise, the client will randomly select from the bound EIPs or NAT IPs for access.
	//
	// example:
	//
	// 1
	EipAffinity *int32 `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// Elastic Network Interface ID. The IPv4 address set of the elastic network interface will be used as the SNAT address.
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
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
	// The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// 	- The elastic IP addresses (EIPs) specified in the SNAT entry when you modify an SNAT entry of an Internet NAT gateway. Separate EIPs with commas (,).
	//
	//     If you select multiple EIPs to create an SNAT address pool, connections are hashed to these EIPs. Network traffic may not be evenly distributed to the EIPs because the amount of traffic passes through each connection varies. We recommend that you associate these EIPs with the same EIP bandwidth plan to prevent service interruptions due to the bandwidth limit of an individual EIP.
	//
	// 	- When you modify an SNAT entry of a VPC NAT gateway, this parameter specifies the NAT IP address in the SNAT entry.
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
