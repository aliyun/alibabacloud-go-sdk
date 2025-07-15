// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntryComment(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetAclEntryComment() *string
	SetAclEntryId(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetAclEntryId() *string
	SetAclId(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetAclId() *string
	SetOwnerAccount(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIPv6TranslatorAclListEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAclListEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAclListEntryRequest
	GetResourceOwnerId() *int64
}

type ModifyIPv6TranslatorAclListEntryRequest struct {
	// The remarks of the ACL rule.
	//
	// It must be 2 to 100 characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// client IP
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	// The ID of the ACL rule to which the IP entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transaclentry-bp1jzyn7ra8pyxehd****
	AclEntryId *string `json:"AclEntryId,omitempty" xml:"AclEntryId,omitempty"`
	// The ID of the ACL to which the IP entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1b4z3tleyhq1s50****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
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

func (s ModifyIPv6TranslatorAclListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclListEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetAclEntryId() *string {
	return s.AclEntryId
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetAclEntryComment(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.AclEntryComment = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetAclEntryId(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.AclEntryId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetAclId(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.AclId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetOwnerAccount(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetOwnerId(v int64) *ModifyIPv6TranslatorAclListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetRegionId(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryRequest) Validate() error {
	return dara.Validate(s)
}
