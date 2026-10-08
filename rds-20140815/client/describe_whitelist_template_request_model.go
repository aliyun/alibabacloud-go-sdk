// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeWhitelistTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeWhitelistTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeWhitelistTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeWhitelistTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *DescribeWhitelistTemplateRequest
	GetTemplateId() *int32
}

type DescribeWhitelistTemplateRequest struct {
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The whitelist template ID. You can call [DescribeAllWhitelistTemplate](~~2aboralibabacloud~~) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1012
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWhitelistTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWhitelistTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeWhitelistTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeWhitelistTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *DescribeWhitelistTemplateRequest) SetRegionId(v string) *DescribeWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWhitelistTemplateRequest) SetResourceGroupId(v string) *DescribeWhitelistTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWhitelistTemplateRequest) SetResourceOwnerAccount(v string) *DescribeWhitelistTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeWhitelistTemplateRequest) SetResourceOwnerId(v int64) *DescribeWhitelistTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeWhitelistTemplateRequest) SetTemplateId(v int32) *DescribeWhitelistTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
