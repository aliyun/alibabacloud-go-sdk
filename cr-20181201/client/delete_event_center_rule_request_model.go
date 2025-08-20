// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteEventCenterRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *DeleteEventCenterRuleRequest
	GetRuleId() *string
}

type DeleteEventCenterRuleRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjx*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteEventCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEventCenterRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteEventCenterRuleRequest) SetInstanceId(v string) *DeleteEventCenterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEventCenterRuleRequest) SetRuleId(v string) *DeleteEventCenterRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteEventCenterRuleRequest) Validate() error {
	return dara.Validate(s)
}
