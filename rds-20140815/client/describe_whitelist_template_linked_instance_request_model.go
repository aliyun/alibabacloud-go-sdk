// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateLinkedInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeWhitelistTemplateLinkedInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeWhitelistTemplateLinkedInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeWhitelistTemplateLinkedInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeWhitelistTemplateLinkedInstanceRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *DescribeWhitelistTemplateLinkedInstanceRequest
	GetTemplateId() *int32
}

type DescribeWhitelistTemplateLinkedInstanceRequest struct {
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
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the whitelist template. You can call the DescribeAllWhitelistTemplate operation to obtain the ID of the whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 412
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeWhitelistTemplateLinkedInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateLinkedInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) SetRegionId(v string) *DescribeWhitelistTemplateLinkedInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) SetResourceGroupId(v string) *DescribeWhitelistTemplateLinkedInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) SetResourceOwnerAccount(v string) *DescribeWhitelistTemplateLinkedInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) SetResourceOwnerId(v int64) *DescribeWhitelistTemplateLinkedInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) SetTemplateId(v int32) *DescribeWhitelistTemplateLinkedInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceRequest) Validate() error {
	return dara.Validate(s)
}
