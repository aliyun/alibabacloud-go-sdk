// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6InternetBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *ModifyIpv6InternetBandwidthRequest
	GetBandwidth() *int64
	SetClientToken(v string) *ModifyIpv6InternetBandwidthRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyIpv6InternetBandwidthRequest
	GetDryRun() *bool
	SetIpv6AddressId(v string) *ModifyIpv6InternetBandwidthRequest
	GetIpv6AddressId() *string
	SetIpv6InternetBandwidthId(v string) *ModifyIpv6InternetBandwidthRequest
	GetIpv6InternetBandwidthId() *string
	SetOwnerAccount(v string) *ModifyIpv6InternetBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIpv6InternetBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIpv6InternetBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIpv6InternetBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIpv6InternetBandwidthRequest
	GetResourceOwnerId() *int64
}

type ModifyIpv6InternetBandwidthRequest struct {
	// The Internet bandwidth value of the IPv6 address. Unit: Mbit/s.
	//
	// 	- If the billing method is pay-by-data-transfer, valid values are **1*	- to **1000**.
	//
	// 	- If the billing method is pay-by-bandwidth, valid values are **1*	- to **2000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: pre-checks the request but does not create the IPv4 gateway. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the API request. After the request passes the check, an HTTP 2xx status code is returned and the IPv4 gateway is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv6 address.
	//
	// >  You must specify one of **Ipv6AddressId*	- and **Ipv6InternetBandwidthId**.
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The instance ID of the Internet bandwidth of the IPv6 address.
	//
	// example:
	//
	// ipv6bw-uf6hcyzu65v98v3du****
	Ipv6InternetBandwidthId *string `json:"Ipv6InternetBandwidthId,omitempty" xml:"Ipv6InternetBandwidthId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPv6 gateway is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s ModifyIpv6InternetBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6InternetBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpv6InternetBandwidthRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ModifyIpv6InternetBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIpv6InternetBandwidthRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyIpv6InternetBandwidthRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *ModifyIpv6InternetBandwidthRequest) GetIpv6InternetBandwidthId() *string {
	return s.Ipv6InternetBandwidthId
}

func (s *ModifyIpv6InternetBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIpv6InternetBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIpv6InternetBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIpv6InternetBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIpv6InternetBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIpv6InternetBandwidthRequest) SetBandwidth(v int64) *ModifyIpv6InternetBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetClientToken(v string) *ModifyIpv6InternetBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetDryRun(v bool) *ModifyIpv6InternetBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetIpv6AddressId(v string) *ModifyIpv6InternetBandwidthRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetIpv6InternetBandwidthId(v string) *ModifyIpv6InternetBandwidthRequest {
	s.Ipv6InternetBandwidthId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetOwnerAccount(v string) *ModifyIpv6InternetBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetOwnerId(v int64) *ModifyIpv6InternetBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetRegionId(v string) *ModifyIpv6InternetBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetResourceOwnerAccount(v string) *ModifyIpv6InternetBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) SetResourceOwnerId(v int64) *ModifyIpv6InternetBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
