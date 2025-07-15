// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyIPv6TranslatorBandwidthRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *ModifyIPv6TranslatorBandwidthRequest
	GetBandwidth() *int32
	SetClientToken(v string) *ModifyIPv6TranslatorBandwidthRequest
	GetClientToken() *string
	SetIpv6TranslatorId(v string) *ModifyIPv6TranslatorBandwidthRequest
	GetIpv6TranslatorId() *string
	SetOwnerAccount(v string) *ModifyIPv6TranslatorBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIPv6TranslatorBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIPv6TranslatorBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIPv6TranslatorBandwidthRequest
	GetResourceOwnerId() *int64
}

type ModifyIPv6TranslatorBandwidthRequest struct {
	// Specifies whether to enable auto-payment for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the IPv6 Translation Service instance. Valid values: **1*	- to **200**. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6trans-bp1858ys****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the IPv6 Translation Service instance. You can call the **DescribeRegions*	- operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyIPv6TranslatorBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIPv6TranslatorBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetAutoPay(v bool) *ModifyIPv6TranslatorBandwidthRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetBandwidth(v int32) *ModifyIPv6TranslatorBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetClientToken(v string) *ModifyIPv6TranslatorBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetIpv6TranslatorId(v string) *ModifyIPv6TranslatorBandwidthRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetOwnerAccount(v string) *ModifyIPv6TranslatorBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetOwnerId(v int64) *ModifyIPv6TranslatorBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetRegionId(v string) *ModifyIPv6TranslatorBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) SetResourceOwnerId(v int64) *ModifyIPv6TranslatorBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
