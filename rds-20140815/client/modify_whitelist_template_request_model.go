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
	// The IP addresses in an IP address whitelist. Separate multiple IP addresses with commas (,). Each IP address in the IP address whitelist must be unique. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.23.XX.XX.
	//
	// 	- CIDR blocks, such as 10.23.XX.XX/24. In this example, 24 indicates that the prefix of the CIDR block is 24-bit in length. You can replace 24 with a value that ranges from 1 to 32.
	//
	// > : A maximum of 1,000 IP addresses or CIDR blocks can be added for each instance. If you want to add a large number of IP addresses, we recommend that you merge them into CIDR blocks, such as 10.23.XX.XX/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139.196.X.X,101.132.X.X
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see related documentation.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the whitelist template. This parameter is required when you modify or delete a whitelist. You can call the DescribeAllWhitelistTemplate operation to obtain the ID of the whitelist.
	//
	// example:
	//
	// 539
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the IP whitelist. This parameter is required when you create a whitelist. The value of this parameter cannot be modified after the whitelist is created. The value must be unique to an Alibaba Cloud account and start with a letter. You can call the DescribeWhitelistTemplate operation to obtain the name of the whitelist.
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
