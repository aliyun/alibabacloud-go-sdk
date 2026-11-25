// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetDefenseScene() *string
	SetInstanceId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseTemplateValidResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseTemplateValidResourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *DescribeDefenseTemplateValidResourcesRequest
	GetTemplateId() *int64
}

type DescribeDefenseTemplateValidResourcesRequest struct {
	// The protection scene to query. For more information, see the valid values of the **DefenseScene*	- parameter in [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html) when the protection rule type **DefenseType*	- is set to **template**.
	//
	// This parameter is required.
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-27a3jyw0v04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the current paging in a paged query. Default value: **1**.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The Settings for the number of protected objects to return on each page of the paging in a paged query. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object to query.
	//
	// example:
	//
	// mgw.realperson.antdigital.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm2thcppfv6ay
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection template.
	//
	//
	// > If you do not specify this parameter, you can query the list of protected objects to be associated when you create a template for the protection scene (**DefenseScene**).
	//
	// example:
	//
	// 21202
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateValidResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetDefenseScene(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetInstanceId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetPageNumber(v int32) *DescribeDefenseTemplateValidResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetPageSize(v int32) *DescribeDefenseTemplateValidResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetRegionId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetResource(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetTemplateId(v int64) *DescribeDefenseTemplateValidResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) Validate() error {
	return dara.Validate(s)
}
