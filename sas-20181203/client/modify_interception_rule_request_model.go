// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInterceptionRuleRequest
	GetClusterId() *string
	SetDstTarget(v map[string]interface{}) *ModifyInterceptionRuleRequest
	GetDstTarget() map[string]interface{}
	SetInterceptType(v int32) *ModifyInterceptionRuleRequest
	GetInterceptType() *int32
	SetOrderIndex(v int64) *ModifyInterceptionRuleRequest
	GetOrderIndex() *int64
	SetRuleId(v int64) *ModifyInterceptionRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifyInterceptionRuleRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *ModifyInterceptionRuleRequest
	GetRuleSwitch() *int32
	SetSrcTarget(v map[string]interface{}) *ModifyInterceptionRuleRequest
	GetSrcTarget() map[string]interface{}
}

type ModifyInterceptionRuleRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c17ef568f81884cdab402decd5fcd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The destination objects of the rule. The following parameters are included:
	//
	// 	- targetId: the ID of the destination object. You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the ID.
	//
	// 	- ports: the destination port ranges.
	//
	// example:
	//
	// {"targetId":600069,"ports":["80/8088"]}
	DstTarget map[string]interface{} `json:"DstTarget,omitempty" xml:"DstTarget,omitempty"`
	// The interception mode. Valid values:
	//
	// 	- **1**: block
	//
	// 	- **2**: alert
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	InterceptType *int32 `json:"InterceptType,omitempty" xml:"InterceptType,omitempty"`
	// The priority of the rule. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500018
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// tetsRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The source object of the rule. The following parameters are included:
	//
	// 	- targetId: the ID of the source object. You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the ID.
	//
	// example:
	//
	// {"targetId":400989}
	SrcTarget map[string]interface{} `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty"`
}

func (s ModifyInterceptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInterceptionRuleRequest) GetDstTarget() map[string]interface{} {
	return s.DstTarget
}

func (s *ModifyInterceptionRuleRequest) GetInterceptType() *int32 {
	return s.InterceptType
}

func (s *ModifyInterceptionRuleRequest) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *ModifyInterceptionRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyInterceptionRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyInterceptionRuleRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyInterceptionRuleRequest) GetSrcTarget() map[string]interface{} {
	return s.SrcTarget
}

func (s *ModifyInterceptionRuleRequest) SetClusterId(v string) *ModifyInterceptionRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetDstTarget(v map[string]interface{}) *ModifyInterceptionRuleRequest {
	s.DstTarget = v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetInterceptType(v int32) *ModifyInterceptionRuleRequest {
	s.InterceptType = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetOrderIndex(v int64) *ModifyInterceptionRuleRequest {
	s.OrderIndex = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetRuleId(v int64) *ModifyInterceptionRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetRuleName(v string) *ModifyInterceptionRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetRuleSwitch(v int32) *ModifyInterceptionRuleRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyInterceptionRuleRequest) SetSrcTarget(v map[string]interface{}) *ModifyInterceptionRuleRequest {
	s.SrcTarget = v
	return s
}

func (s *ModifyInterceptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
