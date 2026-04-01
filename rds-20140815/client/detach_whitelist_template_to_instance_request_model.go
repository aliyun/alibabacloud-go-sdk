// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachWhitelistTemplateToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInsName(v string) *DetachWhitelistTemplateToInstanceRequest
	GetInsName() *string
	SetRegionId(v string) *DetachWhitelistTemplateToInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DetachWhitelistTemplateToInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DetachWhitelistTemplateToInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachWhitelistTemplateToInstanceRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *DetachWhitelistTemplateToInstanceRequest
	GetTemplateId() *int32
}

type DetachWhitelistTemplateToInstanceRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp191w771k******
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see Resource groups.
	//
	// example:
	//
	// rg-acfmz3kjr******
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

func (s DetachWhitelistTemplateToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetInsName() *string {
	return s.InsName
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachWhitelistTemplateToInstanceRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetInsName(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.InsName = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetRegionId(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetResourceGroupId(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetResourceOwnerAccount(v string) *DetachWhitelistTemplateToInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetResourceOwnerId(v int64) *DetachWhitelistTemplateToInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) SetTemplateId(v int32) *DetachWhitelistTemplateToInstanceRequest {
	s.TemplateId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
