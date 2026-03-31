// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeDefenseTemplatesRequest
	GetDefenseScene() *string
	SetDefenseSubScene(v string) *DescribeDefenseTemplatesRequest
	GetDefenseSubScene() *string
	SetInstanceId(v string) *DescribeDefenseTemplatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseTemplatesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseTemplatesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseTemplatesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplatesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *DescribeDefenseTemplatesRequest
	GetResourceType() *string
	SetTemplateId(v int64) *DescribeDefenseTemplatesRequest
	GetTemplateId() *int64
	SetTemplateIds(v string) *DescribeDefenseTemplatesRequest
	GetTemplateIds() *string
	SetTemplateName(v string) *DescribeDefenseTemplatesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *DescribeDefenseTemplatesRequest
	GetTemplateType() *string
}

type DescribeDefenseTemplatesRequest struct {
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
	// example:
	//
	// region_block
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the protection template is used. Valid values:
	//
	// 	- **web**: bot management for website protection.
	//
	// 	- **app**: bot management for app protection.
	//
	// 	- **basic**: bot management for basic protection.
	//
	// example:
	//
	// basic
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-pe33b****03
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
	// The name of the protected object or protected object group.
	//
	// >  If you specify ResourceType, you must specify this parameter.
	//
	// example:
	//
	// xxxqiu.cc-ecs
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmvyknl****fa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// 	- **single**: protected object. This is the default value.
	//
	// 	- **group**: protected object group.
	//
	// >  If you specify Resource, you must specify this parameter.
	//
	// example:
	//
	// single
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The IDs of the protection templates that you want to query. Separate multiple template IDs with commas (,).
	//
	// example:
	//
	// 189731,189539,189538,189531,189540,189542,189541
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
	// The name of the protection template.
	//
	// example:
	//
	// testTemplateName
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the protection template. Valid values:
	//
	// 	- **user_default**: default template.
	//
	// 	- **user_custom**: custom template.
	//
	// example:
	//
	// user_default
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplatesRequest) GetDefenseSubScene() *string {
	return s.DefenseSubScene
}

func (s *DescribeDefenseTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseTemplatesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseTemplatesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseTemplatesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDefenseTemplatesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplatesRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *DescribeDefenseTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDefenseTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDefenseTemplatesRequest) SetDefenseScene(v string) *DescribeDefenseTemplatesRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetDefenseSubScene(v string) *DescribeDefenseTemplatesRequest {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetInstanceId(v string) *DescribeDefenseTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetPageNumber(v int32) *DescribeDefenseTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetPageSize(v int32) *DescribeDefenseTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetRegionId(v string) *DescribeDefenseTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResource(v string) *DescribeDefenseTemplatesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplatesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResourceType(v string) *DescribeDefenseTemplatesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateId(v int64) *DescribeDefenseTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateIds(v string) *DescribeDefenseTemplatesRequest {
	s.TemplateIds = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateName(v string) *DescribeDefenseTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateType(v string) *DescribeDefenseTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
