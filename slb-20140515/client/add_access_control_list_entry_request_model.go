// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccessControlListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v string) *AddAccessControlListEntryRequest
	GetAclEntrys() *string
	SetAclId(v string) *AddAccessControlListEntryRequest
	GetAclId() *string
	SetOwnerAccount(v string) *AddAccessControlListEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddAccessControlListEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddAccessControlListEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddAccessControlListEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAccessControlListEntryRequest
	GetResourceOwnerId() *int64
}

type AddAccessControlListEntryRequest struct {
	// The configuration of the network ACL. Valid values:
	//
	// 	- **entry**: the IP entries that you want to add to the network ACL. You can add CIDR blocks. Separate multiple CIDR blocks with commas (,).
	//
	// 	- **comment**: the comment on the network ACL.
	//
	// > You can add at most 50 IP entries to a network ACL in each call. If the IP entry that you want to add to a network ACL already exists, the IP entry is not added. The IP entries that you add must be CIDR blocks.
	//
	// example:
	//
	// [{"entry":"``10.0.**.**``/24","comment":"privaterule1"},{"entry":"``192.168.**.**``/16","comment":"privaterule2"}]
	AclEntrys *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	// The ID of the network ACL.
	//
	// example:
	//
	// acl-bp1l0kk4gxce43kze*****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
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

func (s AddAccessControlListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryRequest) GetAclEntrys() *string {
	return s.AclEntrys
}

func (s *AddAccessControlListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *AddAccessControlListEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddAccessControlListEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAccessControlListEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAccessControlListEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAccessControlListEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAccessControlListEntryRequest) SetAclEntrys(v string) *AddAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetAclId(v string) *AddAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetOwnerAccount(v string) *AddAccessControlListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetOwnerId(v int64) *AddAccessControlListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetRegionId(v string) *AddAccessControlListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetResourceOwnerAccount(v string) *AddAccessControlListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetResourceOwnerId(v int64) *AddAccessControlListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) Validate() error {
	return dara.Validate(s)
}
