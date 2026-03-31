// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeUserAbnormalTrendRequest
	GetClusterId() *string
	SetInstanceId(v string) *DescribeUserAbnormalTrendRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserAbnormalTrendRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserAbnormalTrendRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserAbnormalTrendRequest struct {
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
	// waf-cn-uqm342yj***
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

func (s DescribeUserAbnormalTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTrendRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserAbnormalTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserAbnormalTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserAbnormalTrendRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserAbnormalTrendRequest) SetClusterId(v string) *DescribeUserAbnormalTrendRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserAbnormalTrendRequest) SetInstanceId(v string) *DescribeUserAbnormalTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserAbnormalTrendRequest) SetRegionId(v string) *DescribeUserAbnormalTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserAbnormalTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserAbnormalTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserAbnormalTrendRequest) Validate() error {
	return dara.Validate(s)
}
