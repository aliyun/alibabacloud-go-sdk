// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAccessControlListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v string) *RemoveAccessControlListEntryRequest
	GetAclEntrys() *string
	SetAclId(v string) *RemoveAccessControlListEntryRequest
	GetAclId() *string
	SetOwnerAccount(v string) *RemoveAccessControlListEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveAccessControlListEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveAccessControlListEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveAccessControlListEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveAccessControlListEntryRequest
	GetResourceOwnerId() *int64
}

type RemoveAccessControlListEntryRequest struct {
	// The IP entries that you want to remove from the network ACL. Valid values:
	//
	// 	- **entry**: the IP address or CIDR block that you want to remove from the network ACL. Separate multiple IP addresses or CIDR blocks with commas (,).
	//
	// 	- **comment**: the description of the network ACL.
	//
	// example:
	//
	// [{"entry":"10.0.10.0/24","comment":"privaterule1"}]
	AclEntrys *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1l0kk4gxce43k******
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the network ACL is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
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

func (s RemoveAccessControlListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryRequest) GetAclEntrys() *string {
	return s.AclEntrys
}

func (s *RemoveAccessControlListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *RemoveAccessControlListEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveAccessControlListEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveAccessControlListEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveAccessControlListEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveAccessControlListEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveAccessControlListEntryRequest) SetAclEntrys(v string) *RemoveAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetAclId(v string) *RemoveAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetOwnerAccount(v string) *RemoveAccessControlListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetOwnerId(v int64) *RemoveAccessControlListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetRegionId(v string) *RemoveAccessControlListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetResourceOwnerAccount(v string) *RemoveAccessControlListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetResourceOwnerId(v int64) *RemoveAccessControlListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) Validate() error {
	return dara.Validate(s)
}
