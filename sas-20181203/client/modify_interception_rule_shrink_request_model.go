// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInterceptionRuleShrinkRequest
	GetClusterId() *string
	SetDstTargetShrink(v string) *ModifyInterceptionRuleShrinkRequest
	GetDstTargetShrink() *string
	SetInterceptType(v int32) *ModifyInterceptionRuleShrinkRequest
	GetInterceptType() *int32
	SetOrderIndex(v int64) *ModifyInterceptionRuleShrinkRequest
	GetOrderIndex() *int64
	SetRuleId(v int64) *ModifyInterceptionRuleShrinkRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifyInterceptionRuleShrinkRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *ModifyInterceptionRuleShrinkRequest
	GetRuleSwitch() *int32
	SetSrcTargetShrink(v string) *ModifyInterceptionRuleShrinkRequest
	GetSrcTargetShrink() *string
}

type ModifyInterceptionRuleShrinkRequest struct {
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
	DstTargetShrink *string `json:"DstTarget,omitempty" xml:"DstTarget,omitempty"`
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
	SrcTargetShrink *string `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty"`
}

func (s ModifyInterceptionRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInterceptionRuleShrinkRequest) GetDstTargetShrink() *string {
	return s.DstTargetShrink
}

func (s *ModifyInterceptionRuleShrinkRequest) GetInterceptType() *int32 {
	return s.InterceptType
}

func (s *ModifyInterceptionRuleShrinkRequest) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *ModifyInterceptionRuleShrinkRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyInterceptionRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyInterceptionRuleShrinkRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyInterceptionRuleShrinkRequest) GetSrcTargetShrink() *string {
	return s.SrcTargetShrink
}

func (s *ModifyInterceptionRuleShrinkRequest) SetClusterId(v string) *ModifyInterceptionRuleShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetDstTargetShrink(v string) *ModifyInterceptionRuleShrinkRequest {
	s.DstTargetShrink = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetInterceptType(v int32) *ModifyInterceptionRuleShrinkRequest {
	s.InterceptType = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetOrderIndex(v int64) *ModifyInterceptionRuleShrinkRequest {
	s.OrderIndex = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetRuleId(v int64) *ModifyInterceptionRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetRuleName(v string) *ModifyInterceptionRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetRuleSwitch(v int32) *ModifyInterceptionRuleShrinkRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) SetSrcTargetShrink(v string) *ModifyInterceptionRuleShrinkRequest {
	s.SrcTargetShrink = &v
	return s
}

func (s *ModifyInterceptionRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
