// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateInterceptionRuleShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateInterceptionRuleShrinkRequest
	GetClusterName() *string
	SetDstTargetListShrink(v string) *CreateInterceptionRuleShrinkRequest
	GetDstTargetListShrink() *string
	SetInterceptType(v int64) *CreateInterceptionRuleShrinkRequest
	GetInterceptType() *int64
	SetOrderIndex(v int64) *CreateInterceptionRuleShrinkRequest
	GetOrderIndex() *int64
	SetRuleName(v string) *CreateInterceptionRuleShrinkRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *CreateInterceptionRuleShrinkRequest
	GetRuleSwitch() *int32
	SetRuleType(v string) *CreateInterceptionRuleShrinkRequest
	GetRuleType() *string
	SetSrcTargetShrink(v string) *CreateInterceptionRuleShrinkRequest
	GetSrcTargetShrink() *string
}

type CreateInterceptionRuleShrinkRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// c35xxxa416
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The information about the destination network object. The value of this parameter contains the following fields:
	//
	// 	- targetId: the ID of the destination network object. You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the ID.
	//
	// 	- ports: the destination port ranges.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "targetId": 600036,
	//
	//             "ports": [
	//
	//                   "1/65535"
	//
	//             ]
	//
	//       }
	//
	// ]
	DstTargetListShrink *string `json:"DstTargetList,omitempty" xml:"DstTargetList,omitempty"`
	// The action on traffic. Valid values:
	//
	// 	- **1**: blocks traffic.
	//
	// 	- **2**: allows traffic and generates alerts.
	//
	// 	- **3**: allows traffic and does not generate alerts.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InterceptType *int64 `json:"InterceptType,omitempty" xml:"InterceptType,omitempty"`
	// The priority of the defense rule. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The name of the defense rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies the status of the defense rule. Valid values:
	//
	// 	- **0**: disables the rule.
	//
	// 	- **1**: enables the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the defense rule. Valid values:
	//
	// 	- customize: custom rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The source network object. The value of this parameter contains the following field:
	//
	// 	- targetId: the ID of the source network object. You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the ID.
	//
	// example:
	//
	// {"targetId":301940}
	SrcTargetShrink *string `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty"`
}

func (s CreateInterceptionRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInterceptionRuleShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInterceptionRuleShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateInterceptionRuleShrinkRequest) GetDstTargetListShrink() *string {
	return s.DstTargetListShrink
}

func (s *CreateInterceptionRuleShrinkRequest) GetInterceptType() *int64 {
	return s.InterceptType
}

func (s *CreateInterceptionRuleShrinkRequest) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *CreateInterceptionRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateInterceptionRuleShrinkRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *CreateInterceptionRuleShrinkRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateInterceptionRuleShrinkRequest) GetSrcTargetShrink() *string {
	return s.SrcTargetShrink
}

func (s *CreateInterceptionRuleShrinkRequest) SetClusterId(v string) *CreateInterceptionRuleShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetClusterName(v string) *CreateInterceptionRuleShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetDstTargetListShrink(v string) *CreateInterceptionRuleShrinkRequest {
	s.DstTargetListShrink = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetInterceptType(v int64) *CreateInterceptionRuleShrinkRequest {
	s.InterceptType = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetOrderIndex(v int64) *CreateInterceptionRuleShrinkRequest {
	s.OrderIndex = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetRuleName(v string) *CreateInterceptionRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetRuleSwitch(v int32) *CreateInterceptionRuleShrinkRequest {
	s.RuleSwitch = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetRuleType(v string) *CreateInterceptionRuleShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) SetSrcTargetShrink(v string) *CreateInterceptionRuleShrinkRequest {
	s.SrcTargetShrink = &v
	return s
}

func (s *CreateInterceptionRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
