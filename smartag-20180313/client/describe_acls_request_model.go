// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v string) *DescribeACLsRequest
	GetAclIds() *string
	SetAclType(v string) *DescribeACLsRequest
	GetAclType() *string
	SetName(v string) *DescribeACLsRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeACLsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeACLsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeACLsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeACLsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeACLsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeACLsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeACLsRequest
	GetResourceOwnerId() *int64
}

type DescribeACLsRequest struct {
	// The ID of the ACL.
	//
	// 	- If you want to query multiple ACLs at the same time, separate the ACL IDs with commas (,).
	//
	// 	- If you do not set this parameter, all ACLs in the specified region are queried.
	//
	// example:
	//
	// acl-xhwhyuo43l*******
	AclIds *string `json:"AclIds,omitempty" xml:"AclIds,omitempty"`
	// The type of the SAG instance associated with the ACL. Valid values:
	//
	// 	- **acl-hardware**: SAG CPE instance
	//
	// 	- **acl-software**: SAG app instance
	//
	// example:
	//
	// acl-hardware
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The name of the ACL.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The ID of the region where the ACL is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeACLsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLsRequest) GoString() string {
	return s.String()
}

func (s *DescribeACLsRequest) GetAclIds() *string {
	return s.AclIds
}

func (s *DescribeACLsRequest) GetAclType() *string {
	return s.AclType
}

func (s *DescribeACLsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeACLsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeACLsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeACLsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeACLsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeACLsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeACLsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeACLsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeACLsRequest) SetAclIds(v string) *DescribeACLsRequest {
	s.AclIds = &v
	return s
}

func (s *DescribeACLsRequest) SetAclType(v string) *DescribeACLsRequest {
	s.AclType = &v
	return s
}

func (s *DescribeACLsRequest) SetName(v string) *DescribeACLsRequest {
	s.Name = &v
	return s
}

func (s *DescribeACLsRequest) SetOwnerAccount(v string) *DescribeACLsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeACLsRequest) SetOwnerId(v int64) *DescribeACLsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeACLsRequest) SetPageNumber(v int32) *DescribeACLsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeACLsRequest) SetPageSize(v int32) *DescribeACLsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeACLsRequest) SetRegionId(v string) *DescribeACLsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeACLsRequest) SetResourceOwnerAccount(v string) *DescribeACLsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeACLsRequest) SetResourceOwnerId(v int64) *DescribeACLsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeACLsRequest) Validate() error {
	return dara.Validate(s)
}
