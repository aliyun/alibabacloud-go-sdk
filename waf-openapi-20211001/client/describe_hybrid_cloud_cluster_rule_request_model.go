// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *DescribeHybridCloudClusterRuleRequest
	GetClusterId() *int64
	SetClusterRuleResourceId(v string) *DescribeHybridCloudClusterRuleRequest
	GetClusterRuleResourceId() *string
	SetInstanceId(v string) *DescribeHybridCloudClusterRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudClusterRuleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeHybridCloudClusterRuleRequest
	GetRuleType() *string
}

type DescribeHybridCloudClusterRuleRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 1
	ClusterId             *int64  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-uqm33n***02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid value:
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
	// The type of the rule. Valid values:
	//
	// 	- **pullin**: The traffic redirection rule of the hybrid cloud cluster.
	//
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeHybridCloudClusterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRuleRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DescribeHybridCloudClusterRuleRequest) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *DescribeHybridCloudClusterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudClusterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudClusterRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudClusterRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHybridCloudClusterRuleRequest) SetClusterId(v int64) *DescribeHybridCloudClusterRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) SetClusterRuleResourceId(v string) *DescribeHybridCloudClusterRuleRequest {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) SetInstanceId(v string) *DescribeHybridCloudClusterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) SetRegionId(v string) *DescribeHybridCloudClusterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) SetRuleType(v string) *DescribeHybridCloudClusterRuleRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleRequest) Validate() error {
	return dara.Validate(s)
}
