// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachWhitelistTemplateToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsName(v string) *AttachWhitelistTemplateToInstanceRequest
	GetInsName() *string
	SetRegionId(v string) *AttachWhitelistTemplateToInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AttachWhitelistTemplateToInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AttachWhitelistTemplateToInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachWhitelistTemplateToInstanceRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *AttachWhitelistTemplateToInstanceRequest
	GetTemplateId() *int32
}

type AttachWhitelistTemplateToInstanceRequest struct {
	// The name of the instance.
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
	// The ID of the resource group. For more information about resource groups, see Resource groups.
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

func (s AttachWhitelistTemplateToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetInsName() *string {
	return s.InsName
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachWhitelistTemplateToInstanceRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetInsName(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.InsName = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetRegionId(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetResourceGroupId(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetResourceOwnerAccount(v string) *AttachWhitelistTemplateToInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetResourceOwnerId(v int64) *AttachWhitelistTemplateToInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) SetTemplateId(v int32) *AttachWhitelistTemplateToInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
