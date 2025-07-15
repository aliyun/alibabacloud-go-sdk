// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIPv6TranslatorAclListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntryComment(v string) *AddIPv6TranslatorAclListEntryRequest
	GetAclEntryComment() *string
	SetAclEntryIp(v string) *AddIPv6TranslatorAclListEntryRequest
	GetAclEntryIp() *string
	SetAclId(v string) *AddIPv6TranslatorAclListEntryRequest
	GetAclId() *string
	SetOwnerAccount(v string) *AddIPv6TranslatorAclListEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddIPv6TranslatorAclListEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddIPv6TranslatorAclListEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddIPv6TranslatorAclListEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddIPv6TranslatorAclListEntryRequest
	GetResourceOwnerId() *int64
}

type AddIPv6TranslatorAclListEntryRequest struct {
	// The remarks of the ACL entry.
	//
	// It must be 2 to 100 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// clientIP1
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	// The IPv6 address or IPv6 CIDR block that you want to add to the ACL entry, for example, 12XX:0:0:XXXX::0102 or 12XX:0:0:XXXX::/60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12XX:0:0:XXXX::0102
	AclEntryIp *string `json:"AclEntryIp,omitempty" xml:"AclEntryIp,omitempty"`
	// The ID of the ACL to which you want to add the IP entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1dcdvfe2****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ACL.
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

func (s AddIPv6TranslatorAclListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIPv6TranslatorAclListEntryRequest) GoString() string {
	return s.String()
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetAclEntryIp() *string {
	return s.AclEntryIp
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddIPv6TranslatorAclListEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetAclEntryComment(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.AclEntryComment = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetAclEntryIp(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.AclEntryIp = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetAclId(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.AclId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetOwnerAccount(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetOwnerId(v int64) *AddIPv6TranslatorAclListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetRegionId(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetResourceOwnerAccount(v string) *AddIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) SetResourceOwnerId(v int64) *AddIPv6TranslatorAclListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryRequest) Validate() error {
	return dara.Validate(s)
}
