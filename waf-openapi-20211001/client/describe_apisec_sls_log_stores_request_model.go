// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsLogStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeApisecSlsLogStoresRequest
	GetInstanceId() *string
	SetLogRegionId(v string) *DescribeApisecSlsLogStoresRequest
	GetLogRegionId() *string
	SetProjectName(v string) *DescribeApisecSlsLogStoresRequest
	GetProjectName() *string
	SetRegionId(v string) *DescribeApisecSlsLogStoresRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecSlsLogStoresRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecSlsLogStoresRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where logs are stored.
	//
	// >  You can call the [DescribeUserSlsLogRegions](https://help.aliyun.com/document_detail/2712598.html) operation to query available log storage regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// apisec-project-14316572********
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeApisecSlsLogStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsLogStoresRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsLogStoresRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecSlsLogStoresRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *DescribeApisecSlsLogStoresRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeApisecSlsLogStoresRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecSlsLogStoresRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecSlsLogStoresRequest) SetInstanceId(v string) *DescribeApisecSlsLogStoresRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecSlsLogStoresRequest) SetLogRegionId(v string) *DescribeApisecSlsLogStoresRequest {
	s.LogRegionId = &v
	return s
}

func (s *DescribeApisecSlsLogStoresRequest) SetProjectName(v string) *DescribeApisecSlsLogStoresRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeApisecSlsLogStoresRequest) SetRegionId(v string) *DescribeApisecSlsLogStoresRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecSlsLogStoresRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecSlsLogStoresRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecSlsLogStoresRequest) Validate() error {
	return dara.Validate(s)
}
