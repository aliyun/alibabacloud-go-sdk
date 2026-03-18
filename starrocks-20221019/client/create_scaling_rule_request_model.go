// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *CreateScalingRuleRequest
	GetNodeGroupId() *string
	SetRule(v string) *CreateScalingRuleRequest
	GetRule() *string
	SetTriggerType(v string) *CreateScalingRuleRequest
	GetTriggerType() *string
}

type CreateScalingRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "nodeNumber": 1,
	//
	//   "scalingRuleName": "test1",
	//
	//   "scalingOutRule": {
	//
	//     "year": 2026,
	//
	//     "month": 3,
	//
	//     "day": 2,
	//
	//     "hour": 5,
	//
	//     "minute": 50,
	//
	//     "second": 0,
	//
	//     "recurrenceInterval": 1,
	//
	//     "recurrenceType": "DAILY",
	//
	//     "recurrenceValues": null
	//
	//   },
	//
	//   "scalingInRule": {
	//
	//     "year": 2026,
	//
	//     "month": 3,
	//
	//     "day": 2,
	//
	//     "hour": 6,
	//
	//     "minute": 50,
	//
	//     "second": 0,
	//
	//     "recurrenceInterval": 1,
	//
	//     "recurrenceType": "DAILY",
	//
	//     "recurrenceValues": null
	//
	//   }
	//
	// }
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TIME_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s CreateScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *CreateScalingRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateScalingRuleRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateScalingRuleRequest) SetNodeGroupId(v string) *CreateScalingRuleRequest {
	s.NodeGroupId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetRule(v string) *CreateScalingRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateScalingRuleRequest) SetTriggerType(v string) *CreateScalingRuleRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
