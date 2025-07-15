// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeIPv6TranslatorAclListsRequest
	GetAclId() *string
	SetAclName(v string) *DescribeIPv6TranslatorAclListsRequest
	GetAclName() *string
	SetOwnerAccount(v string) *DescribeIPv6TranslatorAclListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIPv6TranslatorAclListsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIPv6TranslatorAclListsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorAclListsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIPv6TranslatorAclListsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorAclListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIPv6TranslatorAclListsRequest
	GetResourceOwnerId() *int64
}

type DescribeIPv6TranslatorAclListsRequest struct {
	// The ID of the ACL.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// acl1
	AclName      *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s DescribeIPv6TranslatorAclListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetAclName() *string {
	return s.AclName
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIPv6TranslatorAclListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetAclId(v string) *DescribeIPv6TranslatorAclListsRequest {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetAclName(v string) *DescribeIPv6TranslatorAclListsRequest {
	s.AclName = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetOwnerAccount(v string) *DescribeIPv6TranslatorAclListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetOwnerId(v int64) *DescribeIPv6TranslatorAclListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetPageNumber(v int32) *DescribeIPv6TranslatorAclListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetPageSize(v int32) *DescribeIPv6TranslatorAclListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetRegionId(v string) *DescribeIPv6TranslatorAclListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorAclListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) SetResourceOwnerId(v int64) *DescribeIPv6TranslatorAclListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsRequest) Validate() error {
	return dara.Validate(s)
}
