// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeUserEventTrendRequest
	GetClusterId() *string
	SetEventScope(v string) *DescribeUserEventTrendRequest
	GetEventScope() *string
	SetInstanceId(v string) *DescribeUserEventTrendRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserEventTrendRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserEventTrendRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserEventTrendRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7mz2797x***
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
	// rg-aekzwwk****cv5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserEventTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTrendRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserEventTrendRequest) GetEventScope() *string {
	return s.EventScope
}

func (s *DescribeUserEventTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserEventTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserEventTrendRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserEventTrendRequest) SetClusterId(v string) *DescribeUserEventTrendRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserEventTrendRequest) SetEventScope(v string) *DescribeUserEventTrendRequest {
	s.EventScope = &v
	return s
}

func (s *DescribeUserEventTrendRequest) SetInstanceId(v string) *DescribeUserEventTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserEventTrendRequest) SetRegionId(v string) *DescribeUserEventTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEventTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserEventTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserEventTrendRequest) Validate() error {
	return dara.Validate(s)
}
