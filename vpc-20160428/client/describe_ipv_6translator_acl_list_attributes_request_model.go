// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeIPv6TranslatorAclListAttributesRequest
	GetAclId() *string
	SetOwnerAccount(v string) *DescribeIPv6TranslatorAclListAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIPv6TranslatorAclListAttributesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIPv6TranslatorAclListAttributesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorAclListAttributesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIPv6TranslatorAclListAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorAclListAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIPv6TranslatorAclListAttributesRequest
	GetResourceOwnerId() *int64
}

type DescribeIPv6TranslatorAclListAttributesRequest struct {
	// The ACL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
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

func (s DescribeIPv6TranslatorAclListAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetAclId(v string) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetOwnerAccount(v string) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetOwnerId(v int64) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetPageNumber(v int32) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetPageSize(v int32) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetRegionId(v string) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) SetResourceOwnerId(v int64) *DescribeIPv6TranslatorAclListAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesRequest) Validate() error {
	return dara.Validate(s)
}
