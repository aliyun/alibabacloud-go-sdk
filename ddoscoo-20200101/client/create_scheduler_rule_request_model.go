// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParam(v string) *CreateSchedulerRuleRequest
	GetParam() *string
	SetResourceGroupId(v string) *CreateSchedulerRuleRequest
	GetResourceGroupId() *string
	SetRuleName(v string) *CreateSchedulerRuleRequest
	GetRuleName() *string
	SetRuleType(v int32) *CreateSchedulerRuleRequest
	GetRuleType() *int32
	SetRules(v string) *CreateSchedulerRuleRequest
	GetRules() *string
}

type CreateSchedulerRuleRequest struct {
	// The details of the CDN interaction rule. This parameter is a JSON format character string. The following list describes the structure:
	//
	// - **ParamType**: Required. String type. The type of CDN interaction. Valid values: **cdn**, which indicates CDN interaction.
	//
	// - **ParamData**: Required. Map type. The CDN interaction parameters. The following list describes the structure:
	//
	//     - **Domain**: Required. String type. The CDN-accelerated domain name.
	//
	//     - **Cname**: Required. String type. The CNAME of the accelerated domain name.
	//
	//     - **AccessQps**: Required. Integer type. The access QPS threshold. If the QPS exceeds this threshold, traffic is switched to Anti-DDoS Pro and Anti-DDoS Premium.
	//
	//     - **UpstreamQps**: Optional. Integer type. The back-to-origin QPS threshold. If the QPS is lower than this threshold, traffic is switched to CDN.
	//
	// example:
	//
	// {"ParamType":"cdn","ParamData":{"Domain":"example.aliyundoc.com","Cname":"demo.aliyundoc.com","AccessQps":100,"UpstreamQps":100}}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// - **2**: tiered protection
	//
	// - **3**: acceleration for China-to-global traffic
	//
	// - **5**: CDN interaction
	//
	// - **6**: cloud service interaction
	//
	// - **8**: secure acceleration
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The details of the general interaction rule. This parameter is a JSON string that consists of the following characters. The following list describes the structure:
	//
	// - **Type**: String type. Required. The format of the Resource for Interaction address. Valid values:
	//
	//     - **A**: IP address
	//
	//     - **CNAME**: domain name
	//
	// - **Value**: String type. Required. The address of the Resource for Interaction.
	//
	// - **Priority**: Integer type. Required. The priority of the rule. Valid values: **0*	- to **100**. A larger value indicates a higher priority.
	//
	// - **ValueType**: Integer type. Required. The type of the Resource for Interaction. Valid values:
	//
	//     - **1**: IP Address of Anti-DDoS Proxy instance
	//
	//     - **2**: (Tiered protection) IP addresses of cloud resources
	//
	//     - **3**: (Sec-acceleration for China-to-global traffic) acceleration line IP address
	//
	//     - **5**: (CDN interaction) acceleration domain name
	//
	//     - **6**: (Cloud service interaction) IP addresses of cloud resources
	//
	// - **RegionId**: String type. Optional (required when **ValueType*	- is **2**). The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Type":"A", "Value":"1.1.XX.XX", "Priority":80,"ValueType":2, "RegionId":"cn-hangzhou" },{"Type":"A", "Value":"203.199.XX.XX", "Priority":80,"ValueType":1}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateSchedulerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleRequest) GetParam() *string {
	return s.Param
}

func (s *CreateSchedulerRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSchedulerRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateSchedulerRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *CreateSchedulerRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *CreateSchedulerRuleRequest) SetParam(v string) *CreateSchedulerRuleRequest {
	s.Param = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetResourceGroupId(v string) *CreateSchedulerRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRuleName(v string) *CreateSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRuleType(v int32) *CreateSchedulerRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRules(v string) *CreateSchedulerRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateSchedulerRuleRequest) Validate() error {
	return dara.Validate(s)
}
