// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserWafRulesetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserWafRulesetShrinkRequest
	GetDescription() *string
	SetExpression(v string) *UpdateUserWafRulesetShrinkRequest
	GetExpression() *string
	SetId(v int64) *UpdateUserWafRulesetShrinkRequest
	GetId() *int64
	SetInstanceId(v string) *UpdateUserWafRulesetShrinkRequest
	GetInstanceId() *string
	SetName(v string) *UpdateUserWafRulesetShrinkRequest
	GetName() *string
	SetPosition(v int64) *UpdateUserWafRulesetShrinkRequest
	GetPosition() *int64
	SetRulesShrink(v string) *UpdateUserWafRulesetShrinkRequest
	GetRulesShrink() *string
	SetSharedShrink(v string) *UpdateUserWafRulesetShrinkRequest
	GetSharedShrink() *string
	SetStatus(v string) *UpdateUserWafRulesetShrinkRequest
	GetStatus() *string
}

type UpdateUserWafRulesetShrinkRequest struct {
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ip.src == 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esa-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// [
	//
	//   {
	//
	//     "Id": 20000001,
	//
	//     "Name": "rule1",
	//
	//     "Expression": "ip.src eq 1.1.1.1",
	//
	//     "Action": "deny"
	//
	//   }
	RulesShrink  *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUserWafRulesetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserWafRulesetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserWafRulesetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserWafRulesetShrinkRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdateUserWafRulesetShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateUserWafRulesetShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserWafRulesetShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUserWafRulesetShrinkRequest) GetPosition() *int64 {
	return s.Position
}

func (s *UpdateUserWafRulesetShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateUserWafRulesetShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *UpdateUserWafRulesetShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserWafRulesetShrinkRequest) SetDescription(v string) *UpdateUserWafRulesetShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetExpression(v string) *UpdateUserWafRulesetShrinkRequest {
	s.Expression = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetId(v int64) *UpdateUserWafRulesetShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetInstanceId(v string) *UpdateUserWafRulesetShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetName(v string) *UpdateUserWafRulesetShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetPosition(v int64) *UpdateUserWafRulesetShrinkRequest {
	s.Position = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetRulesShrink(v string) *UpdateUserWafRulesetShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetSharedShrink(v string) *UpdateUserWafRulesetShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) SetStatus(v string) *UpdateUserWafRulesetShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateUserWafRulesetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
