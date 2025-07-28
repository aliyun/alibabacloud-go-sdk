// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecStatisticsRequest
	GetClusterId() *string
	SetInstanceId(v string) *DescribeApisecStatisticsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecStatisticsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecStatisticsRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *DescribeApisecStatisticsRequest
	GetType() *string
}

type DescribeApisecStatisticsRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >  This parameter is available only in hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uax***b09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
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
	// rg-aek2***uwbs5q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the statistics. Valid values:
	//
	// 	- **risk**: risk-related statistics.
	//
	// 	- **event**: event-related statistics.
	//
	// example:
	//
	// asset_num
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApisecStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecStatisticsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecStatisticsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecStatisticsRequest) SetClusterId(v string) *DescribeApisecStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetInstanceId(v string) *DescribeApisecStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetRegionId(v string) *DescribeApisecStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecStatisticsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) SetType(v string) *DescribeApisecStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
