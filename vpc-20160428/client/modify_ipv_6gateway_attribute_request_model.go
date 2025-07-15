// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6GatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyIpv6GatewayAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyIpv6GatewayAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyIpv6GatewayAttributeRequest
	GetDryRun() *bool
	SetIpv6GatewayId(v string) *ModifyIpv6GatewayAttributeRequest
	GetIpv6GatewayId() *string
	SetName(v string) *ModifyIpv6GatewayAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyIpv6GatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIpv6GatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIpv6GatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIpv6GatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIpv6GatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyIpv6GatewayAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IPv6 gateway.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the API request. After the request passes the check, a 2XX HTTP status code is returned and the gateway endpoint is associated with the route table. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv6 gateway that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6gw-hp39kh1ya51yzp2lu****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the IPv6 gateway.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6name
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ModifyIpv6GatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6GatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpv6GatewayAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIpv6GatewayAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyIpv6GatewayAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyIpv6GatewayAttributeRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *ModifyIpv6GatewayAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyIpv6GatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIpv6GatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIpv6GatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIpv6GatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIpv6GatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIpv6GatewayAttributeRequest) SetClientToken(v string) *ModifyIpv6GatewayAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetDescription(v string) *ModifyIpv6GatewayAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetDryRun(v bool) *ModifyIpv6GatewayAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetIpv6GatewayId(v string) *ModifyIpv6GatewayAttributeRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetName(v string) *ModifyIpv6GatewayAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetOwnerAccount(v string) *ModifyIpv6GatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetOwnerId(v int64) *ModifyIpv6GatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetRegionId(v string) *ModifyIpv6GatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIpv6GatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) SetResourceOwnerId(v int64) *ModifyIpv6GatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
