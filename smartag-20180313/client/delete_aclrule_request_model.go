// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteACLRuleRequest
	GetAclId() *string
	SetAcrId(v string) *DeleteACLRuleRequest
	GetAcrId() *string
	SetOwnerAccount(v string) *DeleteACLRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteACLRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteACLRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteACLRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteACLRuleRequest
	GetResourceOwnerId() *int64
}

type DeleteACLRuleRequest struct {
	// The ID of the ACL to which the ACL rule to be deleted belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the ACL rule to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// acr-c1hkd054qywiw******
	AcrId        *string `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the ACL rule to be deleted belongs.
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

func (s DeleteACLRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteACLRuleRequest) GetAclId() *string {
	return s.AclId
}

func (s *DeleteACLRuleRequest) GetAcrId() *string {
	return s.AcrId
}

func (s *DeleteACLRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteACLRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteACLRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteACLRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteACLRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteACLRuleRequest) SetAclId(v string) *DeleteACLRuleRequest {
	s.AclId = &v
	return s
}

func (s *DeleteACLRuleRequest) SetAcrId(v string) *DeleteACLRuleRequest {
	s.AcrId = &v
	return s
}

func (s *DeleteACLRuleRequest) SetOwnerAccount(v string) *DeleteACLRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteACLRuleRequest) SetOwnerId(v int64) *DeleteACLRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteACLRuleRequest) SetRegionId(v string) *DeleteACLRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteACLRuleRequest) SetResourceOwnerAccount(v string) *DeleteACLRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteACLRuleRequest) SetResourceOwnerId(v int64) *DeleteACLRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteACLRuleRequest) Validate() error {
	return dara.Validate(s)
}
