// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntryComment(v string) *DescribeAccessControlListAttributeRequest
	GetAclEntryComment() *string
	SetAclId(v string) *DescribeAccessControlListAttributeRequest
	GetAclId() *string
	SetOwnerAccount(v string) *DescribeAccessControlListAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccessControlListAttributeRequest
	GetOwnerId() *int64
	SetPage(v int32) *DescribeAccessControlListAttributeRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeAccessControlListAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAccessControlListAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccessControlListAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccessControlListAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAccessControlListAttributeRequest struct {
	// The remarks of the ACL entry.
	//
	// example:
	//
	// test
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	// The ID of the ACL that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1ut10zzvh1y8dfs****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation to query the most recent region list.
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

func (s DescribeAccessControlListAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeRequest) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *DescribeAccessControlListAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccessControlListAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccessControlListAttributeRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeAccessControlListAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessControlListAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessControlListAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccessControlListAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccessControlListAttributeRequest) SetAclEntryComment(v string) *DescribeAccessControlListAttributeRequest {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetAclId(v string) *DescribeAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetOwnerAccount(v string) *DescribeAccessControlListAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetOwnerId(v int64) *DescribeAccessControlListAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetPage(v int32) *DescribeAccessControlListAttributeRequest {
	s.Page = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetPageSize(v int32) *DescribeAccessControlListAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetRegionId(v string) *DescribeAccessControlListAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAccessControlListAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetResourceOwnerId(v int64) *DescribeAccessControlListAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) Validate() error {
	return dara.Validate(s)
}
