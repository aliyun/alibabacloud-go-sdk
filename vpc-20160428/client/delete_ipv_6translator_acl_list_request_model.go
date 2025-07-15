// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorAclListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteIPv6TranslatorAclListRequest
	GetAclId() *string
	SetClientToken(v string) *DeleteIPv6TranslatorAclListRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteIPv6TranslatorAclListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIPv6TranslatorAclListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIPv6TranslatorAclListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorAclListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIPv6TranslatorAclListRequest
	GetResourceOwnerId() *int64
}

type DeleteIPv6TranslatorAclListRequest struct {
	// The ID of the ACL that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the IPv6 translation service instance.
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

func (s DeleteIPv6TranslatorAclListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorAclListRequest) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorAclListRequest) GetAclId() *string {
	return s.AclId
}

func (s *DeleteIPv6TranslatorAclListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIPv6TranslatorAclListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIPv6TranslatorAclListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIPv6TranslatorAclListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIPv6TranslatorAclListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIPv6TranslatorAclListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIPv6TranslatorAclListRequest) SetAclId(v string) *DeleteIPv6TranslatorAclListRequest {
	s.AclId = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetClientToken(v string) *DeleteIPv6TranslatorAclListRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetOwnerAccount(v string) *DeleteIPv6TranslatorAclListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetOwnerId(v int64) *DeleteIPv6TranslatorAclListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetRegionId(v string) *DeleteIPv6TranslatorAclListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetResourceOwnerAccount(v string) *DeleteIPv6TranslatorAclListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) SetResourceOwnerId(v int64) *DeleteIPv6TranslatorAclListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListRequest) Validate() error {
	return dara.Validate(s)
}
