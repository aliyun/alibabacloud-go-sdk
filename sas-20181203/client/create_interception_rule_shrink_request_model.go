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
	// The ID of the container cluster to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// This parameter must be from an ACK cluster. You can call the DescribeClustersV1 operation of Container Service for Kubernetes (ACK) to query existing clusters, or call the CreateCluster operation to create a cluster, and then call the DescribeGroupedContainerInstances operation of Security Center to obtain the ID of a managed cluster.
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
	// The list of destination objects. The metric descriptions are as follows:
	//
	// - targetId: The ID of the destination object. You can invoke the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to obtain this parameter.
	//
	// - ports: The list of destination port ranges.
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
	// The interception mode. Valid values:
	//
	// - **1**: Block Mode.
	//
	// - **2**: Alert mode.
	//
	// - **3**: Allow mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InterceptType *int64 `json:"InterceptType,omitempty" xml:"InterceptType,omitempty"`
	// The priority of the rule. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether to enable the rule. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the rule. Valid values:
	//
	// - customize: user-defined rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The source object. The metric description is as follows:
	//
	// - targetId: The ID of the source object. You can invoke the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to obtain this parameter.
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
