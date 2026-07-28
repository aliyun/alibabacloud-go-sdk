// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6InternetBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *AllocateIpv6InternetBandwidthRequest
	GetBandwidth() *int32
	SetClientToken(v string) *AllocateIpv6InternetBandwidthRequest
	GetClientToken() *string
	SetDryRun(v bool) *AllocateIpv6InternetBandwidthRequest
	GetDryRun() *bool
	SetInternetChargeType(v string) *AllocateIpv6InternetBandwidthRequest
	GetInternetChargeType() *string
	SetIpv6AddressId(v string) *AllocateIpv6InternetBandwidthRequest
	GetIpv6AddressId() *string
	SetIpv6GatewayId(v string) *AllocateIpv6InternetBandwidthRequest
	GetIpv6GatewayId() *string
	SetOwnerAccount(v string) *AllocateIpv6InternetBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateIpv6InternetBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AllocateIpv6InternetBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AllocateIpv6InternetBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateIpv6InternetBandwidthRequest
	GetResourceOwnerId() *int64
}

type AllocateIpv6InternetBandwidthRequest struct {
	// The Internet bandwidth of the IPv6 address. Unit: Mbit/s.
	//
	// <props="china">
	//
	// - If **InternetChargeType*	- is set to **PayByTraffic**, the valid values are **1*	- to **1000**.
	//
	// - If **InternetChargeType*	- is set to **PayByBandwidth**, the valid values are **1*	- to **2000**.
	//
	// - If **InternetChargeType*	- is set to **PayByOld95**, the valid values are **1*	- to **2000**.
	//
	//
	// <props="intl">
	//
	//
	//
	// - If **InternetChargeType*	- is set to **PayByTraffic**, the valid values are **1*	- to **1000**.
	//
	// - If **InternetChargeType*	- is set to **PayByBandwidth**, the valid values are **1*	- to **2000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned. The system does not associate a prefix list with a route table.
	//
	// - **false*	- (default): sends a normal request. After the request passes the check, a 2xx HTTP status code is returned and the system associates the prefix list with a route table. The system checks whether the AccessKey pair is valid, whether the Resource Access Management (RAM) user has the required authorization, and whether the required parameters are specified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billable methods of the IPv6 Internet bandwidth. Valid values:
	//
	// <props="china">
	//
	// - **PayByTraffic**: pay-by-data-transfer.
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// - **PayByOld95**: traditional 95th percentile billing. IPv6 Internet bandwidth does not support traditional 95th percentile billing by default. To use this billing method, contact your account manager.
	//
	//
	// <props="intl">
	//
	// - **PayByTraffic**: pay-by-data-transfer.
	//
	// - **PayByBandwidth*	- (default): pay-by-bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the IPv6 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The ID of the IPv6 gateway.
	//
	// example:
	//
	// ipv6gw-uf6hcyzu65v98v3du****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateIpv6InternetBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6InternetBandwidthRequest) GoString() string {
	return s.String()
}

func (s *AllocateIpv6InternetBandwidthRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *AllocateIpv6InternetBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateIpv6InternetBandwidthRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AllocateIpv6InternetBandwidthRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateIpv6InternetBandwidthRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *AllocateIpv6InternetBandwidthRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *AllocateIpv6InternetBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateIpv6InternetBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateIpv6InternetBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateIpv6InternetBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateIpv6InternetBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateIpv6InternetBandwidthRequest) SetBandwidth(v int32) *AllocateIpv6InternetBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetClientToken(v string) *AllocateIpv6InternetBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetDryRun(v bool) *AllocateIpv6InternetBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetInternetChargeType(v string) *AllocateIpv6InternetBandwidthRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetIpv6AddressId(v string) *AllocateIpv6InternetBandwidthRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetIpv6GatewayId(v string) *AllocateIpv6InternetBandwidthRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetOwnerAccount(v string) *AllocateIpv6InternetBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetOwnerId(v int64) *AllocateIpv6InternetBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetRegionId(v string) *AllocateIpv6InternetBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetResourceOwnerAccount(v string) *AllocateIpv6InternetBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) SetResourceOwnerId(v int64) *AllocateIpv6InternetBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
