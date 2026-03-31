// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetApi(v string) *DescribeTemplateResourcesRequest
	GetAssetApi() *string
	SetInstanceId(v string) *DescribeTemplateResourcesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeTemplateResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTemplateResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeTemplateResourcesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeTemplateResourcesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *DescribeTemplateResourcesRequest
	GetResourceType() *string
	SetTemplateId(v int64) *DescribeTemplateResourcesRequest
	GetTemplateId() *int64
}

type DescribeTemplateResourcesRequest struct {
	// example:
	//
	// abc.com
	AssetApi *string `json:"AssetApi,omitempty" xml:"AssetApi,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// 	- **single:*	- protected object.
	//
	// 	- **group:*	- protected object group.
	//
	// This parameter is required.
	//
	// example:
	//
	// single
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1020
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesRequest) GetAssetApi() *string {
	return s.AssetApi
}

func (s *DescribeTemplateResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTemplateResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTemplateResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTemplateResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTemplateResourcesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeTemplateResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeTemplateResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTemplateResourcesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeTemplateResourcesRequest) SetAssetApi(v string) *DescribeTemplateResourcesRequest {
	s.AssetApi = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetInstanceId(v string) *DescribeTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetMaxResults(v int32) *DescribeTemplateResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetNextToken(v string) *DescribeTemplateResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetRegionId(v string) *DescribeTemplateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResource(v string) *DescribeTemplateResourcesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResourceType(v string) *DescribeTemplateResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetTemplateId(v int64) *DescribeTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
