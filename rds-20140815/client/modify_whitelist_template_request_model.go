// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelist(v string) *ModifyWhitelistTemplateRequest
	GetIpWhitelist() *string
	SetRegionId(v string) *ModifyWhitelistTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyWhitelistTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyWhitelistTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyWhitelistTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *ModifyWhitelistTemplateRequest
	GetTemplateId() *int32
	SetTemplateName(v string) *ModifyWhitelistTemplateRequest
	GetTemplateName() *string
}

type ModifyWhitelistTemplateRequest struct {
	// This parameter is required.
	IpWhitelist          *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId           *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistTemplateRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ModifyWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyWhitelistTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWhitelistTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyWhitelistTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyWhitelistTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *ModifyWhitelistTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyWhitelistTemplateRequest) SetIpWhitelist(v string) *ModifyWhitelistTemplateRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetRegionId(v string) *ModifyWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetResourceGroupId(v string) *ModifyWhitelistTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetResourceOwnerAccount(v string) *ModifyWhitelistTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetResourceOwnerId(v int64) *ModifyWhitelistTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetTemplateId(v int32) *ModifyWhitelistTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) SetTemplateName(v string) *ModifyWhitelistTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
