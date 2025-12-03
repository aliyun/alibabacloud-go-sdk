// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteAccessControlListRequest
	GetAclId() *string
	SetOwnerAccount(v string) *DeleteAccessControlListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAccessControlListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAccessControlListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAccessControlListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAccessControlListRequest
	GetResourceOwnerId() *int64
}

type DeleteAccessControlListRequest struct {
	// The ACL ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1l0kk4gxce43kz******
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
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

func (s DeleteAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListRequest) GetAclId() *string {
	return s.AclId
}

func (s *DeleteAccessControlListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAccessControlListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccessControlListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAccessControlListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccessControlListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAccessControlListRequest) SetAclId(v string) *DeleteAccessControlListRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetOwnerAccount(v string) *DeleteAccessControlListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetOwnerId(v int64) *DeleteAccessControlListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetRegionId(v string) *DeleteAccessControlListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetResourceOwnerAccount(v string) *DeleteAccessControlListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetResourceOwnerId(v int64) *DeleteAccessControlListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccessControlListRequest) Validate() error {
	return dara.Validate(s)
}
