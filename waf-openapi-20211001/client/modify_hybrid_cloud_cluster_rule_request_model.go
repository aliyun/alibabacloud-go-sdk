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
	// [Deprecated] The hybrid cloud cluster ID.
	//
	// example:
	//
	// 10*
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster rule resource ID.
	//
	// example:
	//
	// hdbc-clusterrule-*******ym0w
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// The Web Application Firewall (WAF) instance ID.
	//
	// > Call [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) to query the current WAF instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-n6w***x52m
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: The Chinese mainland.
	//
	// - **ap-southeast-1**: Outside the Chinese mainland.
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
	// The traffic routing rule configuration:
	//
	// 	Notice:
	//
	// The mode cannot be changed after it is selected.
	//
	//
	//
	// - **check_mode*	- Defines the traffic scope for the routing rule. Valid values:
	//
	//   - **all**: Routes all traffic.
	//
	//   - **part**: Routes a specified portion of traffic.
	//
	// - **type**: The rule\\"s match type. Valid values:
	//
	//   - **exact**: Exact match
	//
	//   - **regex**: Regular expression match.
	//
	// - **substance**: The value of the rule.
	//
	// example:
	//
	// {\\"check_mode\\": \\"all\\", \\"type\\": \\"exact\\", \\"substance\\": \\"122\\"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The status of the rule. Valid values:
	//
	// - **on**: Enabled
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// [Deprecated] The rule type. Valid values:
	//
	// - **pullin**: Traffic routing configuration.
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
