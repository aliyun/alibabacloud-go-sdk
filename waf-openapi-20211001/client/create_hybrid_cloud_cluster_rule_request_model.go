// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *CreateHybridCloudClusterRuleRequest
	GetClusterId() *int64
	SetInstanceId(v string) *CreateHybridCloudClusterRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateHybridCloudClusterRuleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateHybridCloudClusterRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleConfig(v string) *CreateHybridCloudClusterRuleRequest
	GetRuleConfig() *string
	SetRuleStatus(v string) *CreateHybridCloudClusterRuleRequest
	GetRuleStatus() *string
	SetRuleType(v string) *CreateHybridCloudClusterRuleRequest
	GetRuleType() *string
}

type CreateHybridCloudClusterRuleRequest struct {
	// The ID of the Hybrid Cloud WAF cluster.
	//
	// > This parameter applies only to hybrid cloud scenarios. Call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to obtain information about Hybrid Cloud WAF clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 428
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
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
	// The configuration of the traffic redirection rule, in JSON format.
	//
	// > The modes are mutually exclusive. You can select only one mode. If you change the mode, all traffic redirection rules in the previous mode are cleared.
	//
	// - **check_mode**: the traffic redirection mode. Valid values:
	//
	//   - **all**: redirects all traffic.
	//
	//   - **part**: redirects a specific portion of traffic.
	//
	// - **type**: the match type for the rule. Valid values:
	//
	//   - **exact**: exact match.
	//
	//   - **regex**: regular expression match.
	//
	// - **substance**: the match value of the rule.
	//
	// example:
	//
	// full volume drainage:{\\"check_mode\\": \\"all\\", \\"type\\": \\"exact\\", \\"substance\\": \\"122\\"}
	//
	// Specified partial drainage：{\\"check_mode\\": \\"part\\", \\"type\\": \\"exact\\", \\"substance\\": \\"12222\\"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The status of the traffic redirection rule. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The type of the traffic redirection rule. Valid values:
	//
	// - **pullin**: traffic redirection
	//
	// This parameter is required.
	//
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s CreateHybridCloudClusterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterRuleRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *CreateHybridCloudClusterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHybridCloudClusterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridCloudClusterRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateHybridCloudClusterRuleRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *CreateHybridCloudClusterRuleRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CreateHybridCloudClusterRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateHybridCloudClusterRuleRequest) SetClusterId(v int64) *CreateHybridCloudClusterRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetInstanceId(v string) *CreateHybridCloudClusterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetRegionId(v string) *CreateHybridCloudClusterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetResourceManagerResourceGroupId(v string) *CreateHybridCloudClusterRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetRuleConfig(v string) *CreateHybridCloudClusterRuleRequest {
	s.RuleConfig = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetRuleStatus(v string) *CreateHybridCloudClusterRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) SetRuleType(v string) *CreateHybridCloudClusterRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateHybridCloudClusterRuleRequest) Validate() error {
	return dara.Validate(s)
}
