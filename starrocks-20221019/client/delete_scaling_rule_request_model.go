// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *DeleteScalingRuleRequest
	GetNodeGroupId() *string
	SetScalingRuleId(v string) *DeleteScalingRuleRequest
	GetScalingRuleId() *string
	SetTriggerType(v string) *DeleteScalingRuleRequest
	GetTriggerType() *string
}

type DeleteScalingRuleRequest struct {
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
	// r-d01a1cac4081s****
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TIME_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s DeleteScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DeleteScalingRuleRequest) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *DeleteScalingRuleRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *DeleteScalingRuleRequest) SetNodeGroupId(v string) *DeleteScalingRuleRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetScalingRuleId(v string) *DeleteScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetTriggerType(v string) *DeleteScalingRuleRequest {
	s.TriggerType = &v
	return s
}

func (s *DeleteScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
