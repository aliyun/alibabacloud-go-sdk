// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRuleRequest
	GetDescription() *string
	SetEventBusName(v string) *UpdateRuleRequest
	GetEventBusName() *string
	SetFilterPattern(v string) *UpdateRuleRequest
	GetFilterPattern() *string
	SetRuleName(v string) *UpdateRuleRequest
	GetRuleName() *string
	SetStatus(v string) *UpdateRuleRequest
	GetStatus() *string
}

type UpdateRuleRequest struct {
	// The description of the event bus.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// hw-test
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual stringExpression Each field can have a maximum of five expressions in the map data structure.
	//
	// Each field can have a maximum of five expressions in the map data structure.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"source\\":[\\"acs.oss\\"],\\"type\\":[\\"oss:BucketQueried:GetBucketStat\\"]}
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testacc-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: The event rule is enabled. It is the default state of the event rule. DISABLE: The event rule is disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRuleRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateRuleRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *UpdateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRuleRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateRuleRequest) SetDescription(v string) *UpdateRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRuleRequest) SetEventBusName(v string) *UpdateRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateRuleRequest) SetFilterPattern(v string) *UpdateRuleRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleName(v string) *UpdateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleRequest) SetStatus(v string) *UpdateRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateRuleRequest) Validate() error {
	return dara.Validate(s)
}
