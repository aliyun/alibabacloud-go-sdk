// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv4GatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpv4GatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpv4GatewayRequest
	GetDryRun() *bool
	SetInternetMode(v string) *DeleteIpv4GatewayRequest
	GetInternetMode() *string
	SetIpv4GatewayId(v string) *DeleteIpv4GatewayRequest
	GetIpv4GatewayId() *string
	SetOwnerAccount(v string) *DeleteIpv4GatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpv4GatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpv4GatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpv4GatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpv4GatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteIpv4GatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without deleting the IPv4 gateway. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the HTTP 2xx status code is returned and the IPv4 gateway is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The public network pattern of the VPC after the IPv4 gateway is deleted. Valid values:
	//
	// - **private**: the default value. After the IPv4 gateway is deleted, the VPC becomes a purely private network VPC and does not have public network access capabilities.
	//
	// - **public**: After the IPv4 gateway is deleted, public network access to the VPC is no longer centrally controlled by the IPv4 gateway. By default, after a public IP address is associated with an instance, the instance can directly access the public network.
	//
	// example:
	//
	// public
	InternetMode *string `json:"InternetMode,omitempty" xml:"InternetMode,omitempty"`
	// The instance ID of the IPv4 gateway to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv4 gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpv4GatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv4GatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpv4GatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpv4GatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpv4GatewayRequest) GetInternetMode() *string {
	return s.InternetMode
}

func (s *DeleteIpv4GatewayRequest) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *DeleteIpv4GatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpv4GatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpv4GatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpv4GatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpv4GatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpv4GatewayRequest) SetClientToken(v string) *DeleteIpv4GatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetDryRun(v bool) *DeleteIpv4GatewayRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetInternetMode(v string) *DeleteIpv4GatewayRequest {
	s.InternetMode = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetIpv4GatewayId(v string) *DeleteIpv4GatewayRequest {
	s.Ipv4GatewayId = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetOwnerAccount(v string) *DeleteIpv4GatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetOwnerId(v int64) *DeleteIpv4GatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetRegionId(v string) *DeleteIpv4GatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetResourceOwnerAccount(v string) *DeleteIpv4GatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) SetResourceOwnerId(v int64) *DeleteIpv4GatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpv4GatewayRequest) Validate() error {
	return dara.Validate(s)
}
