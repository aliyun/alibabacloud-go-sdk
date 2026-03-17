// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *ModifyACLRequest
	GetAclId() *string
	SetName(v string) *ModifyACLRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyACLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyACLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyACLRequest
	GetResourceOwnerId() *int64
}

type ModifyACLRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-e30a66to95cs****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The new name of the ACL.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// newname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the ACL is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
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

func (s ModifyACLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRequest) GoString() string {
	return s.String()
}

func (s *ModifyACLRequest) GetAclId() *string {
	return s.AclId
}

func (s *ModifyACLRequest) GetName() *string {
	return s.Name
}

func (s *ModifyACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyACLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyACLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyACLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyACLRequest) SetAclId(v string) *ModifyACLRequest {
	s.AclId = &v
	return s
}

func (s *ModifyACLRequest) SetName(v string) *ModifyACLRequest {
	s.Name = &v
	return s
}

func (s *ModifyACLRequest) SetOwnerAccount(v string) *ModifyACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyACLRequest) SetOwnerId(v int64) *ModifyACLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyACLRequest) SetRegionId(v string) *ModifyACLRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyACLRequest) SetResourceOwnerAccount(v string) *ModifyACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyACLRequest) SetResourceOwnerId(v int64) *ModifyACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyACLRequest) Validate() error {
	return dara.Validate(s)
}
