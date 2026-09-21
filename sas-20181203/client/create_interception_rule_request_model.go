// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateInterceptionRuleRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateInterceptionRuleRequest
	GetClusterName() *string
	SetDstTargetList(v map[string]interface{}) *CreateInterceptionRuleRequest
	GetDstTargetList() map[string]interface{}
	SetInterceptType(v int64) *CreateInterceptionRuleRequest
	GetInterceptType() *int64
	SetOrderIndex(v int64) *CreateInterceptionRuleRequest
	GetOrderIndex() *int64
	SetRuleName(v string) *CreateInterceptionRuleRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *CreateInterceptionRuleRequest
	GetRuleSwitch() *int32
	SetRuleType(v string) *CreateInterceptionRuleRequest
	GetRuleType() *string
	SetSrcTarget(v map[string]interface{}) *CreateInterceptionRuleRequest
	GetSrcTarget() map[string]interface{}
}

type CreateInterceptionRuleRequest struct {
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
	DstTargetList map[string]interface{} `json:"DstTargetList,omitempty" xml:"DstTargetList,omitempty"`
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
	SrcTarget map[string]interface{} `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty"`
}

func (s CreateInterceptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateInterceptionRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInterceptionRuleRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateInterceptionRuleRequest) GetDstTargetList() map[string]interface{} {
	return s.DstTargetList
}

func (s *CreateInterceptionRuleRequest) GetInterceptType() *int64 {
	return s.InterceptType
}

func (s *CreateInterceptionRuleRequest) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *CreateInterceptionRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateInterceptionRuleRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *CreateInterceptionRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateInterceptionRuleRequest) GetSrcTarget() map[string]interface{} {
	return s.SrcTarget
}

func (s *CreateInterceptionRuleRequest) SetClusterId(v string) *CreateInterceptionRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetClusterName(v string) *CreateInterceptionRuleRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetDstTargetList(v map[string]interface{}) *CreateInterceptionRuleRequest {
	s.DstTargetList = v
	return s
}

func (s *CreateInterceptionRuleRequest) SetInterceptType(v int64) *CreateInterceptionRuleRequest {
	s.InterceptType = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetOrderIndex(v int64) *CreateInterceptionRuleRequest {
	s.OrderIndex = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetRuleName(v string) *CreateInterceptionRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetRuleSwitch(v int32) *CreateInterceptionRuleRequest {
	s.RuleSwitch = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetRuleType(v string) *CreateInterceptionRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateInterceptionRuleRequest) SetSrcTarget(v map[string]interface{}) *CreateInterceptionRuleRequest {
	s.SrcTarget = v
	return s
}

func (s *CreateInterceptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
