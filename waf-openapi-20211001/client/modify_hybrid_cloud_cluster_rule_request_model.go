// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *ModifyHybridCloudClusterRuleRequest
	GetClusterId() *int64
	SetClusterRuleResourceId(v string) *ModifyHybridCloudClusterRuleRequest
	GetClusterRuleResourceId() *string
	SetInstanceId(v string) *ModifyHybridCloudClusterRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHybridCloudClusterRuleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleConfig(v string) *ModifyHybridCloudClusterRuleRequest
	GetRuleConfig() *string
	SetRuleStatus(v string) *ModifyHybridCloudClusterRuleRequest
	GetRuleStatus() *string
	SetRuleType(v string) *ModifyHybridCloudClusterRuleRequest
	GetRuleType() *string
}

type ModifyHybridCloudClusterRuleRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 1018
	ClusterId             *int64  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the DescribeInstanceInfo operation to query the ID of the WAF instance.[](~~140857~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-n6w***x52m
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the WAF instance. Valid value:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
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
	// The configuration of the rule.
	//
	// example:
	//
	// {\\"check_mode\\":\\"part\\",\\"include\\":{\\"exact\\":[],\\"regex\\":[]}}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **on**: enables the rule.
	//
	// 	- **off**: disables the rule.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **pullin**: The traffic redirection rule.
	//
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ModifyHybridCloudClusterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterRuleRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *ModifyHybridCloudClusterRuleRequest) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *ModifyHybridCloudClusterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudClusterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudClusterRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudClusterRuleRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyHybridCloudClusterRuleRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ModifyHybridCloudClusterRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyHybridCloudClusterRuleRequest) SetClusterId(v int64) *ModifyHybridCloudClusterRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetClusterRuleResourceId(v string) *ModifyHybridCloudClusterRuleRequest {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetInstanceId(v string) *ModifyHybridCloudClusterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetRegionId(v string) *ModifyHybridCloudClusterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetRuleConfig(v string) *ModifyHybridCloudClusterRuleRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetRuleStatus(v string) *ModifyHybridCloudClusterRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) SetRuleType(v string) *ModifyHybridCloudClusterRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleRequest) Validate() error {
	return dara.Validate(s)
}
