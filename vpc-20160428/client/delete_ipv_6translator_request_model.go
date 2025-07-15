// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIPv6TranslatorRequest
	GetClientToken() *string
	SetIpv6TranslatorId(v string) *DeleteIPv6TranslatorRequest
	GetIpv6TranslatorId() *string
	SetOwnerAccount(v string) *DeleteIPv6TranslatorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIPv6TranslatorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIPv6TranslatorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIPv6TranslatorRequest
	GetResourceOwnerId() *int64
}

type DeleteIPv6TranslatorRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// ClientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6trans-bp1i8ahxut1ie****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 Translation Service instance.
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

func (s DeleteIPv6TranslatorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorRequest) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIPv6TranslatorRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DeleteIPv6TranslatorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIPv6TranslatorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIPv6TranslatorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIPv6TranslatorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIPv6TranslatorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIPv6TranslatorRequest) SetClientToken(v string) *DeleteIPv6TranslatorRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetIpv6TranslatorId(v string) *DeleteIPv6TranslatorRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetOwnerAccount(v string) *DeleteIPv6TranslatorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetOwnerId(v int64) *DeleteIPv6TranslatorRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetRegionId(v string) *DeleteIPv6TranslatorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) SetResourceOwnerId(v int64) *DeleteIPv6TranslatorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorRequest) Validate() error {
	return dara.Validate(s)
}
