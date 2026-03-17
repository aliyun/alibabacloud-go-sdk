// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclType(v string) *CreateACLRequest
	GetAclType() *string
	SetName(v string) *CreateACLRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateACLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateACLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateACLRequest
	GetResourceOwnerId() *int64
}

type CreateACLRequest struct {
	// The type of SAG instance to be associated with the ACL. Valid values:
	//
	// 	- **acl-hardware*	- (default): SAG CPE instance
	//
	// 	- **acl-software**: SAG app instance
	//
	// example:
	//
	// acl-hardware
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The name of the ACL.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// username
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to create the ACL.
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

func (s CreateACLRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateACLRequest) GoString() string {
	return s.String()
}

func (s *CreateACLRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateACLRequest) GetName() *string {
	return s.Name
}

func (s *CreateACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateACLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateACLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateACLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateACLRequest) SetAclType(v string) *CreateACLRequest {
	s.AclType = &v
	return s
}

func (s *CreateACLRequest) SetName(v string) *CreateACLRequest {
	s.Name = &v
	return s
}

func (s *CreateACLRequest) SetOwnerAccount(v string) *CreateACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateACLRequest) SetOwnerId(v int64) *CreateACLRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateACLRequest) SetRegionId(v string) *CreateACLRequest {
	s.RegionId = &v
	return s
}

func (s *CreateACLRequest) SetResourceOwnerAccount(v string) *CreateACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateACLRequest) SetResourceOwnerId(v int64) *CreateACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateACLRequest) Validate() error {
	return dara.Validate(s)
}
