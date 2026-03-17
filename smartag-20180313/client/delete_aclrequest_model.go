// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteACLRequest
	GetAclId() *string
	SetOwnerAccount(v string) *DeleteACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteACLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteACLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteACLRequest
	GetResourceOwnerId() *int64
}

type DeleteACLRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-ohlexqptfhy*******
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
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

func (s DeleteACLRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLRequest) GoString() string {
	return s.String()
}

func (s *DeleteACLRequest) GetAclId() *string {
	return s.AclId
}

func (s *DeleteACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteACLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteACLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteACLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteACLRequest) SetAclId(v string) *DeleteACLRequest {
	s.AclId = &v
	return s
}

func (s *DeleteACLRequest) SetOwnerAccount(v string) *DeleteACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteACLRequest) SetOwnerId(v int64) *DeleteACLRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteACLRequest) SetRegionId(v string) *DeleteACLRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteACLRequest) SetResourceOwnerAccount(v string) *DeleteACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteACLRequest) SetResourceOwnerId(v int64) *DeleteACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteACLRequest) Validate() error {
	return dara.Validate(s)
}
