// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIPv6TranslatorAclListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntryId(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetAclEntryId() *string
	SetAclId(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetAclId() *string
	SetClientToken(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveIPv6TranslatorAclListEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveIPv6TranslatorAclListEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveIPv6TranslatorAclListEntryRequest
	GetResourceOwnerId() *int64
}

type RemoveIPv6TranslatorAclListEntryRequest struct {
	// The ID of the ACL entry to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transaclentry-bp105jrs****
	AclEntryId *string `json:"AclEntryId,omitempty" xml:"AclEntryId,omitempty"`
	// The ID of the ACL to which the ACL entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
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
	// The region of the ACL.
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

func (s RemoveIPv6TranslatorAclListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveIPv6TranslatorAclListEntryRequest) GoString() string {
	return s.String()
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetAclEntryId() *string {
	return s.AclEntryId
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetAclEntryId(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.AclEntryId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetAclId(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.AclId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetClientToken(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetOwnerAccount(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetOwnerId(v int64) *RemoveIPv6TranslatorAclListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetRegionId(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetResourceOwnerAccount(v string) *RemoveIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) SetResourceOwnerId(v int64) *RemoveIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryRequest) Validate() error {
	return dara.Validate(s)
}
