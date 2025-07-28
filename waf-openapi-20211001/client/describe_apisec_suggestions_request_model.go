// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSuggestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApisecSuggestionsRequest
	GetApiId() *string
	SetClusterId(v string) *DescribeApisecSuggestionsRequest
	GetClusterId() *string
	SetInstanceId(v string) *DescribeApisecSuggestionsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecSuggestionsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecSuggestionsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecSuggestionsRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// a60fd7e3021fe371c06dc1dcb883def0
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-pe336n43m04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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

func (s DescribeApisecSuggestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSuggestionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecSuggestionsRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecSuggestionsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecSuggestionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecSuggestionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecSuggestionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecSuggestionsRequest) SetApiId(v string) *DescribeApisecSuggestionsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecSuggestionsRequest) SetClusterId(v string) *DescribeApisecSuggestionsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecSuggestionsRequest) SetInstanceId(v string) *DescribeApisecSuggestionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecSuggestionsRequest) SetRegionId(v string) *DescribeApisecSuggestionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecSuggestionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecSuggestionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecSuggestionsRequest) Validate() error {
	return dara.Validate(s)
}
