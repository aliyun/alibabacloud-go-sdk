// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUserWafRulesetRequest
	GetDescription() *string
	SetExpression(v string) *UpdateUserWafRulesetRequest
	GetExpression() *string
	SetId(v int64) *UpdateUserWafRulesetRequest
	GetId() *int64
	SetInstanceId(v string) *UpdateUserWafRulesetRequest
	GetInstanceId() *string
	SetName(v string) *UpdateUserWafRulesetRequest
	GetName() *string
	SetPosition(v int64) *UpdateUserWafRulesetRequest
	GetPosition() *int64
	SetRules(v []*WafRuleConfig) *UpdateUserWafRulesetRequest
	GetRules() []*WafRuleConfig
	SetShared(v *WafBatchRuleShared) *UpdateUserWafRulesetRequest
	GetShared() *WafBatchRuleShared
	SetStatus(v string) *UpdateUserWafRulesetRequest
	GetStatus() *string
}

type UpdateUserWafRulesetRequest struct {
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
	Rules  []*WafRuleConfig    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUserWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserWafRulesetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserWafRulesetRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdateUserWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateUserWafRulesetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserWafRulesetRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUserWafRulesetRequest) GetPosition() *int64 {
	return s.Position
}

func (s *UpdateUserWafRulesetRequest) GetRules() []*WafRuleConfig {
	return s.Rules
}

func (s *UpdateUserWafRulesetRequest) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *UpdateUserWafRulesetRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserWafRulesetRequest) SetDescription(v string) *UpdateUserWafRulesetRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetExpression(v string) *UpdateUserWafRulesetRequest {
	s.Expression = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetId(v int64) *UpdateUserWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetInstanceId(v string) *UpdateUserWafRulesetRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetName(v string) *UpdateUserWafRulesetRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetPosition(v int64) *UpdateUserWafRulesetRequest {
	s.Position = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetRules(v []*WafRuleConfig) *UpdateUserWafRulesetRequest {
	s.Rules = v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetShared(v *WafBatchRuleShared) *UpdateUserWafRulesetRequest {
	s.Shared = v
	return s
}

func (s *UpdateUserWafRulesetRequest) SetStatus(v string) *UpdateUserWafRulesetRequest {
	s.Status = &v
	return s
}

func (s *UpdateUserWafRulesetRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Shared != nil {
		if err := s.Shared.Validate(); err != nil {
			return err
		}
	}
	return nil
}
