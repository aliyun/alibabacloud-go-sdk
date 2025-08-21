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
	// The details of the CDN interaction rule. This parameter is a JSON string. The following list describes the fields in the value of the parameter:
	//
	// 	- **ParamType**: the type of the scheduling rule. This field is required and must be of the string type. Set the value to **cdn**. This indicates that you want to modify a CDN interaction rule.
	//
	// 	- **ParamData**: the values of parameters that you want to modify for the CDN interaction rule. This field is required and must be of the map type. ParamData contains the following parameters:
	//
	//     	- **Domain**: the accelerated domain in CDN. This parameter is required and must be of the string type.
	//
	//     	- **Cname**: the CNAME that is assigned to the accelerated domain. This parameter is required and must be of the string type.
	//
	//     	- **AccessQps**: the queries per second (QPS) threshold that is used to switch service traffic to Anti-DDoS Pro or Anti-DDoS Premium. This parameter is required and must be of the integer type.
	//
	//     	- **UpstreamQps**: the QPS threshold that is used to switch service traffic to CDN. This parameter is optional and must be of the integer type.
	//
	// example:
	//
	// {"ParamType":"cdn","ParamData":{"Domain":"example.aliyundoc.com","Cname":"demo.aliyundoc.com","AccessQps":100,"UpstreamQps":100}}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
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
	// 	- **2**: tiered protection
	//
	// 	- **3**: network acceleration
	//
	// 	- **5**: Alibaba Cloud CDN (CDN) interaction
	//
	// 	- **6**: cloud service interaction
	//
	// 	- **8**: secure acceleration
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The details of the scheduling rule. This parameter is a JSON string. The following list describes the fields in the value of the parameter:
	//
	// 	- **Type**: the address type of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the string type. Valid values:
	//
	//     	- **A**: IP address
	//
	//     	- **CNAME**: domain name
	//
	// 	- **Value**: the address of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the string type.
	//
	// 	- **Priority**: the priority of the scheduling rule. This field is required and must be of the integer type. Valid values: **0*	- to **100**. A larger value indicates a higher priority.
	//
	// 	- **ValueType**: the type of the interaction resource that you want to use in the scheduling rule. This field is required and must be of the integer type. Valid values:
	//
	//     	- **1**: the IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance
	//
	//     	- **2**: the IP address of the interaction resource in the tiered protection scenario
	//
	//     	- **3**: the IP address that is used to accelerate access in the network acceleration scenario
	//
	//     	- **5**: the domain name that is configured in Alibaba Cloud CDN (CDN) in the CDN interaction scenario
	//
	//     	- **6*	- the IP address of the interaction resource in the cloud service interaction scenario
	//
	// 	- **RegionId**: the region where the interaction resource is deployed. This parameter must be specified when **ValueType*	- is set to **2**. The value must be of the string type.
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
