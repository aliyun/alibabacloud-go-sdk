// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DisableRuleRequest
	GetEventBusName() *string
	SetRuleName(v string) *DisableRuleRequest
	GetRuleName() *string
}

type DisableRuleRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testacc-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DisableRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableRuleRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DisableRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DisableRuleRequest) SetEventBusName(v string) *DisableRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *DisableRuleRequest) SetRuleName(v string) *DisableRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DisableRuleRequest) Validate() error {
	return dara.Validate(s)
}
