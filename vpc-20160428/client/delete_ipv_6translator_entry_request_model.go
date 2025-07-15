// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIPv6TranslatorEntryRequest
	GetClientToken() *string
	SetIpv6TranslatorEntryId(v string) *DeleteIPv6TranslatorEntryRequest
	GetIpv6TranslatorEntryId() *string
	SetIpv6TranslatorId(v string) *DeleteIPv6TranslatorEntryRequest
	GetIpv6TranslatorId() *string
	SetOwnerAccount(v string) *DeleteIPv6TranslatorEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIPv6TranslatorEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIPv6TranslatorEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIPv6TranslatorEntryRequest
	GetResourceOwnerId() *int64
}

type DeleteIPv6TranslatorEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IPv6 mapping entry to be deleted.
	//
	// example:
	//
	// ipv6transentry-bp1g8bhrde****
	Ipv6TranslatorEntryId *string `json:"Ipv6TranslatorEntryId,omitempty" xml:"Ipv6TranslatorEntryId,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// > If you do not specify **Ipv6TranslatorEntryId**, all mapping entries in the specified instance are deleted.
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

func (s DeleteIPv6TranslatorEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIPv6TranslatorEntryRequest) GetIpv6TranslatorEntryId() *string {
	return s.Ipv6TranslatorEntryId
}

func (s *DeleteIPv6TranslatorEntryRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DeleteIPv6TranslatorEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIPv6TranslatorEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIPv6TranslatorEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIPv6TranslatorEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIPv6TranslatorEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIPv6TranslatorEntryRequest) SetClientToken(v string) *DeleteIPv6TranslatorEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetIpv6TranslatorEntryId(v string) *DeleteIPv6TranslatorEntryRequest {
	s.Ipv6TranslatorEntryId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetIpv6TranslatorId(v string) *DeleteIPv6TranslatorEntryRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetOwnerAccount(v string) *DeleteIPv6TranslatorEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetOwnerId(v int64) *DeleteIPv6TranslatorEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetRegionId(v string) *DeleteIPv6TranslatorEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) SetResourceOwnerId(v int64) *DeleteIPv6TranslatorEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryRequest) Validate() error {
	return dara.Validate(s)
}
