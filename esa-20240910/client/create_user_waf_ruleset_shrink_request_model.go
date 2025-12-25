// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserWafRulesetShrinkRequest
	GetDescription() *string
	SetExpression(v string) *CreateUserWafRulesetShrinkRequest
	GetExpression() *string
	SetInstanceId(v string) *CreateUserWafRulesetShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateUserWafRulesetShrinkRequest
	GetName() *string
	SetPhase(v string) *CreateUserWafRulesetShrinkRequest
	GetPhase() *string
	SetRulesShrink(v string) *CreateUserWafRulesetShrinkRequest
	GetRulesShrink() *string
	SetSharedShrink(v string) *CreateUserWafRulesetShrinkRequest
	GetSharedShrink() *string
	SetStatus(v string) *CreateUserWafRulesetShrinkRequest
	GetStatus() *string
}

type CreateUserWafRulesetShrinkRequest struct {
	// example:
	//
	// this is a test ruleset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ip.src == 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esa-site-ads11w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase        *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	RulesShrink  *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateUserWafRulesetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserWafRulesetShrinkRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateUserWafRulesetShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserWafRulesetShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserWafRulesetShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateUserWafRulesetShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateUserWafRulesetShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *CreateUserWafRulesetShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserWafRulesetShrinkRequest) SetDescription(v string) *CreateUserWafRulesetShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetExpression(v string) *CreateUserWafRulesetShrinkRequest {
	s.Expression = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetInstanceId(v string) *CreateUserWafRulesetShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetName(v string) *CreateUserWafRulesetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetPhase(v string) *CreateUserWafRulesetShrinkRequest {
	s.Phase = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetRulesShrink(v string) *CreateUserWafRulesetShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetSharedShrink(v string) *CreateUserWafRulesetShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) SetStatus(v string) *CreateUserWafRulesetShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateUserWafRulesetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
