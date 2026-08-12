// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeApisecSlsProjectsRequest
	GetInstanceId() *string
	SetLogRegionId(v string) *DescribeApisecSlsProjectsRequest
	GetLogRegionId() *string
	SetRegionId(v string) *DescribeApisecSlsProjectsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecSlsProjectsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecSlsProjectsRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID for log storage.
	//
	// > You can call [DescribeUserSlsLogRegions](https://help.aliyun.com/document_detail/2712598.html) to query the available log storage regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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

func (s DescribeApisecSlsProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsProjectsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecSlsProjectsRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *DescribeApisecSlsProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecSlsProjectsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecSlsProjectsRequest) SetInstanceId(v string) *DescribeApisecSlsProjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecSlsProjectsRequest) SetLogRegionId(v string) *DescribeApisecSlsProjectsRequest {
	s.LogRegionId = &v
	return s
}

func (s *DescribeApisecSlsProjectsRequest) SetRegionId(v string) *DescribeApisecSlsProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecSlsProjectsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecSlsProjectsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecSlsProjectsRequest) Validate() error {
	return dara.Validate(s)
}
