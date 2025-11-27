// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLinkedWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsName(v string) *DescribeInstanceLinkedWhitelistTemplateRequest
	GetInsName() *string
	SetRegionId(v string) *DescribeInstanceLinkedWhitelistTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstanceLinkedWhitelistTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceLinkedWhitelistTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceLinkedWhitelistTemplateRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceLinkedWhitelistTemplateRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp191w771kd3****
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can leave this parameter empty.
	//
	// example:
	//
	// rg-aek3dbzqbh6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceLinkedWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLinkedWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) GetInsName() *string {
	return s.InsName
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) SetInsName(v string) *DescribeInstanceLinkedWhitelistTemplateRequest {
	s.InsName = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) SetRegionId(v string) *DescribeInstanceLinkedWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) SetResourceGroupId(v string) *DescribeInstanceLinkedWhitelistTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) SetResourceOwnerAccount(v string) *DescribeInstanceLinkedWhitelistTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) SetResourceOwnerId(v int64) *DescribeInstanceLinkedWhitelistTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
