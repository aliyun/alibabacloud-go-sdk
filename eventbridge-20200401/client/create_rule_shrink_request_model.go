// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRuleShrinkRequest
	GetDescription() *string
	SetEventBusName(v string) *CreateRuleShrinkRequest
	GetEventBusName() *string
	SetEventTargetsShrink(v string) *CreateRuleShrinkRequest
	GetEventTargetsShrink() *string
	SetFilterPattern(v string) *CreateRuleShrinkRequest
	GetFilterPattern() *string
	SetRuleName(v string) *CreateRuleShrinkRequest
	GetRuleName() *string
	SetStatus(v string) *CreateRuleShrinkRequest
	GetStatus() *string
}

type CreateRuleShrinkRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event targets.
	EventTargetsShrink *string `json:"EventTargets,omitempty" xml:"EventTargets,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"source\\": [{\\"prefix\\": \\"acs.\\"}],\\"type\\": [{\\"prefix\\":\\"oss:ObjectReplication\\"}],\\"subject\\":[{\\"prefix\\":\\"acs:oss:cn-hangzhou:123456789098****:my-movie-bucket/\\", \\"suffix\\":\\".txt\\"}]}
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// MNSRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: enables the event rule. It is the default status of the event rule. DISABLE: disables the event rule.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRuleShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateRuleShrinkRequest) GetEventTargetsShrink() *string {
	return s.EventTargetsShrink
}

func (s *CreateRuleShrinkRequest) GetFilterPattern() *string {
	return s.FilterPattern
}

func (s *CreateRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRuleShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateRuleShrinkRequest) SetDescription(v string) *CreateRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetEventBusName(v string) *CreateRuleShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetEventTargetsShrink(v string) *CreateRuleShrinkRequest {
	s.EventTargetsShrink = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetFilterPattern(v string) *CreateRuleShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetRuleName(v string) *CreateRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetStatus(v string) *CreateRuleShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
