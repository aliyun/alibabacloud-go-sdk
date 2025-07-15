// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetAclId() *string
	SetAclName(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetAclName() *string
	SetClientToken(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIPv6TranslatorAclAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAclAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAclAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyIPv6TranslatorAclAttributeRequest struct {
	// The ID of the ACL that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the IPv6 Translation Service instance. You can call the DescribeRegions operation to query the most recent region list.
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

func (s ModifyIPv6TranslatorAclAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetAclName() *string {
	return s.AclName
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetAclId(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetAclName(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.AclName = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetClientToken(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetOwnerAccount(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetOwnerId(v int64) *ModifyIPv6TranslatorAclAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetRegionId(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAclAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAclAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeRequest) Validate() error {
	return dara.Validate(s)
}
