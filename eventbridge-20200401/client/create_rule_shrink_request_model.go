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
	// The description of the event rule.
	//
	// example:
	//
	// SMQ filter rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// A list of event targets.
	EventTargetsShrink *string `json:"EventTargets,omitempty" xml:"EventTargets,omitempty"`
	// The event pattern, in JSON format. Supported pattern types are `stringEqual` and `stringExpression`. Each field can contain a maximum of five expressions in a map structure.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "source": [
	//
	//     {
	//
	//       "prefix": "acs."
	//
	//     }
	//
	//   ],
	//
	//   "type": [
	//
	//     {
	//
	//       "prefix": "oss:ObjectReplication"
	//
	//     }
	//
	//   ],
	//
	//   "subject": [
	//
	//     {
	//
	//       "prefix": "acs:oss:cn-hangzhou:123456789098****:my-movie-bucket/",
	//
	//       "suffix": ".txt"
	//
	//     }
	//
	//   ]
	//
	// }
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMQRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: `ENABLE`: The rule is enabled. This is the default value. `DISABLE`: The rule is disabled.
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
