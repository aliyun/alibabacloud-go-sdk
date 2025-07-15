// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6EgressOnlyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpv6EgressOnlyRuleRequest
	GetClientToken() *string
	SetIpv6EgressOnlyRuleId(v string) *DeleteIpv6EgressOnlyRuleRequest
	GetIpv6EgressOnlyRuleId() *string
	SetOwnerAccount(v string) *DeleteIpv6EgressOnlyRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpv6EgressOnlyRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpv6EgressOnlyRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpv6EgressOnlyRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpv6EgressOnlyRuleRequest
	GetResourceOwnerId() *int64
}

type DeleteIpv6EgressOnlyRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the **client token**. The value of **RequestId*	- is different for each API request.
	//
	// example:
	//
	// 123456
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the egress-only rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6py-hp3w98rmlbqp0xxxxxxxx
	Ipv6EgressOnlyRuleId *string `json:"Ipv6EgressOnlyRuleId,omitempty" xml:"Ipv6EgressOnlyRuleId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to query the most recent region list.
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

func (s DeleteIpv6EgressOnlyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6EgressOnlyRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetIpv6EgressOnlyRuleId() *string {
	return s.Ipv6EgressOnlyRuleId
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpv6EgressOnlyRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetClientToken(v string) *DeleteIpv6EgressOnlyRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetIpv6EgressOnlyRuleId(v string) *DeleteIpv6EgressOnlyRuleRequest {
	s.Ipv6EgressOnlyRuleId = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetOwnerAccount(v string) *DeleteIpv6EgressOnlyRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetOwnerId(v int64) *DeleteIpv6EgressOnlyRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetRegionId(v string) *DeleteIpv6EgressOnlyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetResourceOwnerAccount(v string) *DeleteIpv6EgressOnlyRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) SetResourceOwnerId(v int64) *DeleteIpv6EgressOnlyRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpv6EgressOnlyRuleRequest) Validate() error {
	return dara.Validate(s)
}
