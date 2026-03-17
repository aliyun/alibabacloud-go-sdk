// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DisassociateACLRequest
	GetAclId() *string
	SetOwnerAccount(v string) *DisassociateACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisassociateACLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisassociateACLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisassociateACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisassociateACLRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DisassociateACLRequest
	GetSmartAGId() *string
}

type DisassociateACLRequest struct {
	// The ID of ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0****
	AclId        *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
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
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-ke3kq4evpi8****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DisassociateACLRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateACLRequest) GoString() string {
	return s.String()
}

func (s *DisassociateACLRequest) GetAclId() *string {
	return s.AclId
}

func (s *DisassociateACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisassociateACLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisassociateACLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisassociateACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisassociateACLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisassociateACLRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DisassociateACLRequest) SetAclId(v string) *DisassociateACLRequest {
	s.AclId = &v
	return s
}

func (s *DisassociateACLRequest) SetOwnerAccount(v string) *DisassociateACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisassociateACLRequest) SetOwnerId(v int64) *DisassociateACLRequest {
	s.OwnerId = &v
	return s
}

func (s *DisassociateACLRequest) SetRegionId(v string) *DisassociateACLRequest {
	s.RegionId = &v
	return s
}

func (s *DisassociateACLRequest) SetResourceOwnerAccount(v string) *DisassociateACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisassociateACLRequest) SetResourceOwnerId(v int64) *DisassociateACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisassociateACLRequest) SetSmartAGId(v string) *DisassociateACLRequest {
	s.SmartAGId = &v
	return s
}

func (s *DisassociateACLRequest) Validate() error {
	return dara.Validate(s)
}
