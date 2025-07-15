// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6InternetBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpv6InternetBandwidthRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpv6InternetBandwidthRequest
	GetDryRun() *bool
	SetIpv6AddressId(v string) *DeleteIpv6InternetBandwidthRequest
	GetIpv6AddressId() *string
	SetIpv6InternetBandwidthId(v string) *DeleteIpv6InternetBandwidthRequest
	GetIpv6InternetBandwidthId() *string
	SetOwnerAccount(v string) *DeleteIpv6InternetBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpv6InternetBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpv6InternetBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpv6InternetBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpv6InternetBandwidthRequest
	GetResourceOwnerId() *int64
}

type DeleteIpv6InternetBandwidthRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
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
	// The region ID of the IPv6 gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteIpv6InternetBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6InternetBandwidthRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpv6InternetBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpv6InternetBandwidthRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpv6InternetBandwidthRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *DeleteIpv6InternetBandwidthRequest) GetIpv6InternetBandwidthId() *string {
	return s.Ipv6InternetBandwidthId
}

func (s *DeleteIpv6InternetBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpv6InternetBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpv6InternetBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpv6InternetBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpv6InternetBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpv6InternetBandwidthRequest) SetClientToken(v string) *DeleteIpv6InternetBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetDryRun(v bool) *DeleteIpv6InternetBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetIpv6AddressId(v string) *DeleteIpv6InternetBandwidthRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetIpv6InternetBandwidthId(v string) *DeleteIpv6InternetBandwidthRequest {
	s.Ipv6InternetBandwidthId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetOwnerAccount(v string) *DeleteIpv6InternetBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetOwnerId(v int64) *DeleteIpv6InternetBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetRegionId(v string) *DeleteIpv6InternetBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetResourceOwnerAccount(v string) *DeleteIpv6InternetBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) SetResourceOwnerId(v int64) *DeleteIpv6InternetBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
