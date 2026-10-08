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
	// The IP whitelist of the instance. Separate multiple IP addresses with commas (,). IP addresses cannot be duplicated. The following two formats are supported:
	//
	// - IP address format, such as 10.23.XX.XX.
	//
	// - CIDR format, such as 10.23.XX.XX/24 (Classless Inter-Domain Routing, where 24 indicates the prefix length, with a value range of 1 to 32).
	//
	// > Each instance supports a maximum of 1,000 IP addresses or CIDR blocks. The total number of IP addresses or CIDR blocks across all IP whitelist groups cannot exceed 1,000. If you have a large number of IP addresses, merge them into CIDR blocks, such as 10.23.XX.XX/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139.196.X.X,101.132.X.X
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see What is a resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The whitelist template ID.
	//
	// This parameter is required for modify and delete operations. You can call DescribeAllWhitelistTemplate to obtain the template ID.
	//
	// example:
	//
	// 539
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The whitelist template name. Specify this parameter when creating a template. The name cannot be modified after creation, must be unique within the same account, and must start with a letter. You can call DescribeWhitelistTemplate to obtain the template name.
	//
	// example:
	//
	// template_123
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
