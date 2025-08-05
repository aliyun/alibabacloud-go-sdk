// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DeleteRuleRequest
	GetEventBusName() *string
	SetRuleName(v string) *DeleteRuleRequest
	GetRuleName() *string
}

type DeleteRuleRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testacc-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ramrolechange-mns
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DeleteRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteRuleRequest) SetEventBusName(v string) *DeleteRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleName(v string) *DeleteRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
