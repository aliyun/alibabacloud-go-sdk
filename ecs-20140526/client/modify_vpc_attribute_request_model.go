// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *ModifyVpcAttributeRequest
	GetCidrBlock() *string
	SetDescription(v string) *ModifyVpcAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyVpcAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpcAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVpcAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpcAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpcAttributeRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *ModifyVpcAttributeRequest
	GetUserCidr() *string
	SetVpcId(v string) *ModifyVpcAttributeRequest
	GetVpcId() *string
	SetVpcName(v string) *ModifyVpcAttributeRequest
	GetVpcName() *string
}

type ModifyVpcAttributeRequest struct {
	// The primary IPv4 cidr block for the VPC. You can only expand this cidr block, for example, from `192.168.0.0/24` to `192.168.0.0/16`. You cannot modify the primary cidr block if ClassicLink is enabled.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description must be 2 to 256 characters long. It must start with a letter and cannot begin with `http://` or `https://`.
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the VPC\\"s region.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A user cidr block to add to the VPC. You can add up to three user cidr blocks. They cannot overlap with the primary cidr block, each other, or the reserved `100.64.0.0/10` cidr block.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The ID of the VPC to modify.
	//
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name must be 2 to 128 characters long, start with a letter, and can contain letters, digits, underscores (_), and hyphens (-).
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ModifyVpcAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcAttributeRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ModifyVpcAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpcAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpcAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpcAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpcAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpcAttributeRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *ModifyVpcAttributeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpcAttributeRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *ModifyVpcAttributeRequest) SetCidrBlock(v string) *ModifyVpcAttributeRequest {
	s.CidrBlock = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetDescription(v string) *ModifyVpcAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetOwnerAccount(v string) *ModifyVpcAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetOwnerId(v int64) *ModifyVpcAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetRegionId(v string) *ModifyVpcAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpcAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpcAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetUserCidr(v string) *ModifyVpcAttributeRequest {
	s.UserCidr = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetVpcId(v string) *ModifyVpcAttributeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetVpcName(v string) *ModifyVpcAttributeRequest {
	s.VpcName = &v
	return s
}

func (s *ModifyVpcAttributeRequest) Validate() error {
	return dara.Validate(s)
}
