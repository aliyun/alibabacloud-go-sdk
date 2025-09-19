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
	DstTargetList map[string]interface{} `json:"DstTargetList,omitempty" xml:"DstTargetList,omitempty"`
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
