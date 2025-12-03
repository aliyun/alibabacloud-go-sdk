// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessControlListAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetAccessControlListAttributeRequest
	GetAclId() *string
	SetAclName(v string) *SetAccessControlListAttributeRequest
	GetAclName() *string
	SetOwnerAccount(v string) *SetAccessControlListAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetAccessControlListAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetAccessControlListAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetAccessControlListAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetAccessControlListAttributeRequest
	GetResourceOwnerId() *int64
}

type SetAccessControlListAttributeRequest struct {
	// The ACL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1l0kk4gxce43k******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ACL name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	AclName      *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ACL.
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

func (s SetAccessControlListAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetAccessControlListAttributeRequest) GetAclName() *string {
	return s.AclName
}

func (s *SetAccessControlListAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetAccessControlListAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetAccessControlListAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetAccessControlListAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetAccessControlListAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetAccessControlListAttributeRequest) SetAclId(v string) *SetAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetAclName(v string) *SetAccessControlListAttributeRequest {
	s.AclName = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetOwnerAccount(v string) *SetAccessControlListAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetOwnerId(v int64) *SetAccessControlListAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetRegionId(v string) *SetAccessControlListAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetResourceOwnerAccount(v string) *SetAccessControlListAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetResourceOwnerId(v int64) *SetAccessControlListAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) Validate() error {
	return dara.Validate(s)
}
