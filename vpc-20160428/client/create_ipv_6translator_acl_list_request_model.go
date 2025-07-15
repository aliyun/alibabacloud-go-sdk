// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorAclListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *CreateIPv6TranslatorAclListRequest
	GetAclName() *string
	SetClientToken(v string) *CreateIPv6TranslatorAclListRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CreateIPv6TranslatorAclListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIPv6TranslatorAclListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIPv6TranslatorAclListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateIPv6TranslatorAclListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIPv6TranslatorAclListRequest
	GetResourceOwnerId() *int64
}

type CreateIPv6TranslatorAclListRequest struct {
	// The ACL name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// example:
	//
	// sha123456
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

func (s CreateIPv6TranslatorAclListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorAclListRequest) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorAclListRequest) GetAclName() *string {
	return s.AclName
}

func (s *CreateIPv6TranslatorAclListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIPv6TranslatorAclListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIPv6TranslatorAclListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIPv6TranslatorAclListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIPv6TranslatorAclListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIPv6TranslatorAclListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIPv6TranslatorAclListRequest) SetAclName(v string) *CreateIPv6TranslatorAclListRequest {
	s.AclName = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetClientToken(v string) *CreateIPv6TranslatorAclListRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetOwnerAccount(v string) *CreateIPv6TranslatorAclListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetOwnerId(v int64) *CreateIPv6TranslatorAclListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetRegionId(v string) *CreateIPv6TranslatorAclListRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetResourceOwnerAccount(v string) *CreateIPv6TranslatorAclListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) SetResourceOwnerId(v int64) *CreateIPv6TranslatorAclListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorAclListRequest) Validate() error {
	return dara.Validate(s)
}
