// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6GatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpv6GatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpv6GatewayRequest
	GetDryRun() *bool
	SetIpv6GatewayId(v string) *DeleteIpv6GatewayRequest
	GetIpv6GatewayId() *string
	SetOwnerAccount(v string) *DeleteIpv6GatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpv6GatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpv6GatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpv6GatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpv6GatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteIpv6GatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv6 gateway that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6gw-hp3y0l3ln89j8****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteIpv6GatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6GatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpv6GatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpv6GatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpv6GatewayRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DeleteIpv6GatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpv6GatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpv6GatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpv6GatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpv6GatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpv6GatewayRequest) SetClientToken(v string) *DeleteIpv6GatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetDryRun(v bool) *DeleteIpv6GatewayRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetIpv6GatewayId(v string) *DeleteIpv6GatewayRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetOwnerAccount(v string) *DeleteIpv6GatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetOwnerId(v int64) *DeleteIpv6GatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetRegionId(v string) *DeleteIpv6GatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetResourceOwnerAccount(v string) *DeleteIpv6GatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) SetResourceOwnerId(v int64) *DeleteIpv6GatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpv6GatewayRequest) Validate() error {
	return dara.Validate(s)
}
