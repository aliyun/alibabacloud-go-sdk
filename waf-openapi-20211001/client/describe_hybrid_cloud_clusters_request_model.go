// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudClustersRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudClustersRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClustersRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudClustersRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uqm35****02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
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

func (s DescribeHybridCloudClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClustersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudClustersRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudClustersRequest) SetInstanceId(v string) *DescribeHybridCloudClustersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudClustersRequest) SetRegionId(v string) *DescribeHybridCloudClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudClustersRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClustersRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudClustersRequest) Validate() error {
	return dara.Validate(s)
}
