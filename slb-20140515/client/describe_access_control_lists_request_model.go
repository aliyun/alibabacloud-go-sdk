// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *DescribeAccessControlListsRequest
	GetAclName() *string
	SetAddressIPVersion(v string) *DescribeAccessControlListsRequest
	GetAddressIPVersion() *string
	SetOwnerAccount(v string) *DescribeAccessControlListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccessControlListsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccessControlListsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessControlListsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAccessControlListsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAccessControlListsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAccessControlListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccessControlListsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeAccessControlListsRequestTag) *DescribeAccessControlListsRequest
	GetTag() []*DescribeAccessControlListsRequestTag
}

type DescribeAccessControlListsRequest struct {
	// The ACL name. The ACL name. The name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), hyphens (-), forward slashes (/), and underscores (_). The name of each ACL must be unique within a region. Fuzzy match is supported.
	//
	// example:
	//
	// rule1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP version of the Classic Load Balancer (CLB) instance with which the ACL is associated. Valid values:
	//
	// 	- **ipv4**
	//
	// 	- **ipv6**
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The region ID of the ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop4****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeAccessControlListsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsRequest) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListsRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccessControlListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccessControlListsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessControlListsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessControlListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessControlListsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAccessControlListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccessControlListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccessControlListsRequest) GetTag() []*DescribeAccessControlListsRequestTag {
	return s.Tag
}

func (s *DescribeAccessControlListsRequest) SetAclName(v string) *DescribeAccessControlListsRequest {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetAddressIPVersion(v string) *DescribeAccessControlListsRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetOwnerAccount(v string) *DescribeAccessControlListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetOwnerId(v int64) *DescribeAccessControlListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageNumber(v int32) *DescribeAccessControlListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageSize(v int32) *DescribeAccessControlListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetRegionId(v string) *DescribeAccessControlListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceGroupId(v string) *DescribeAccessControlListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceOwnerAccount(v string) *DescribeAccessControlListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceOwnerId(v int64) *DescribeAccessControlListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetTag(v []*DescribeAccessControlListsRequestTag) *DescribeAccessControlListsRequest {
	s.Tag = v
	return s
}

func (s *DescribeAccessControlListsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessControlListsRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value must be 1 to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessControlListsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAccessControlListsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAccessControlListsRequestTag) SetKey(v string) *DescribeAccessControlListsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAccessControlListsRequestTag) SetValue(v string) *DescribeAccessControlListsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAccessControlListsRequestTag) Validate() error {
	return dara.Validate(s)
}
