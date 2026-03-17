// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *AssociateACLRequest
	GetAclId() *string
	SetOwnerAccount(v string) *AssociateACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateACLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateACLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateACLRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *AssociateACLRequest
	GetSmartAGId() *string
}

type AssociateACLRequest struct {
	// The ID of ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-ohlexqptfhy******
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the ACL is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance to be associated with the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-4yr0p2xa6o3k*******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s AssociateACLRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateACLRequest) GoString() string {
	return s.String()
}

func (s *AssociateACLRequest) GetAclId() *string {
	return s.AclId
}

func (s *AssociateACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateACLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateACLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateACLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateACLRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *AssociateACLRequest) SetAclId(v string) *AssociateACLRequest {
	s.AclId = &v
	return s
}

func (s *AssociateACLRequest) SetOwnerAccount(v string) *AssociateACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateACLRequest) SetOwnerId(v int64) *AssociateACLRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateACLRequest) SetRegionId(v string) *AssociateACLRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateACLRequest) SetResourceOwnerAccount(v string) *AssociateACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateACLRequest) SetResourceOwnerId(v int64) *AssociateACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateACLRequest) SetSmartAGId(v string) *AssociateACLRequest {
	s.SmartAGId = &v
	return s
}

func (s *AssociateACLRequest) Validate() error {
	return dara.Validate(s)
}
