// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6EgressOnlyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpv6EgressOnlyRuleRequest
	GetClientToken() *string
	SetDescription(v string) *CreateIpv6EgressOnlyRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateIpv6EgressOnlyRuleRequest
	GetInstanceId() *string
	SetInstanceType(v string) *CreateIpv6EgressOnlyRuleRequest
	GetInstanceType() *string
	SetIpv6GatewayId(v string) *CreateIpv6EgressOnlyRuleRequest
	GetIpv6GatewayId() *string
	SetName(v string) *CreateIpv6EgressOnlyRuleRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateIpv6EgressOnlyRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpv6EgressOnlyRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpv6EgressOnlyRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateIpv6EgressOnlyRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpv6EgressOnlyRuleRequest
	GetResourceOwnerId() *int64
}

type CreateIpv6EgressOnlyRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of RequestId as the client token. The value of RequestId is different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the egress-only rule.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ruledescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the IPv6 address for which you want to create an egress-only rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6-hp3nxjkfxn5pnhgl5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance for which you want to create an egress-only rule.
	//
	// Default value: **Ipv6Address**
	//
	// example:
	//
	// Ipv6Address
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the IPv6 gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6gw-hp3c2paq0ywauasza****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the egress-only rule.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// rulename
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the IPv6 gateway is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s CreateIpv6EgressOnlyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6EgressOnlyRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpv6EgressOnlyRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetClientToken(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetDescription(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetInstanceId(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetInstanceType(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetIpv6GatewayId(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetName(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetOwnerAccount(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetOwnerId(v int64) *CreateIpv6EgressOnlyRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetRegionId(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetResourceOwnerAccount(v string) *CreateIpv6EgressOnlyRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) SetResourceOwnerId(v int64) *CreateIpv6EgressOnlyRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleRequest) Validate() error {
	return dara.Validate(s)
}
