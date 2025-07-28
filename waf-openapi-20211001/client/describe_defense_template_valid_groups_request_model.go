// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeDefenseTemplateValidGroupsRequest
	GetDefenseScene() *string
	SetGroupName(v string) *DescribeDefenseTemplateValidGroupsRequest
	GetGroupName() *string
	SetInstanceId(v string) *DescribeDefenseTemplateValidGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseTemplateValidGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseTemplateValidGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseTemplateValidGroupsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidGroupsRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *DescribeDefenseTemplateValidGroupsRequest
	GetTemplateId() *int64
}

type DescribeDefenseTemplateValidGroupsRequest struct {
	// The scenario in which the protection template is used.
	//
	// 	- **waf_group**: basic protection.
	//
	// 	- **antiscan**: scan protection.
	//
	// 	- **ip_blacklist**: IP address blacklist.
	//
	// 	- **custom_acl**: custom rule.
	//
	// 	- **whitelist**: whitelist.
	//
	// 	- **region_block**: region blacklist.
	//
	// 	- **custom_response**: custom response.
	//
	// 	- **cc**: HTTP flood protection.
	//
	// 	- **tamperproof**: website tamper-proofing.
	//
	// 	- **dlp**: data leakage prevention.
	//
	// This parameter is required.
	//
	// example:
	//
	// region_block
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The name of the protected object group that you want to query.
	//
	// example:
	//
	// group221
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-nwy****pf0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm2th****v6ay
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseTemplateValidGroupsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetDefenseScene(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetGroupName(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetInstanceId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetPageNumber(v int32) *DescribeDefenseTemplateValidGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetPageSize(v int32) *DescribeDefenseTemplateValidGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetRegionId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetTemplateId(v int64) *DescribeDefenseTemplateValidGroupsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) Validate() error {
	return dara.Validate(s)
}
