// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserWafRulesetRequest
	GetDescription() *string
	SetExpression(v string) *CreateUserWafRulesetRequest
	GetExpression() *string
	SetInstanceId(v string) *CreateUserWafRulesetRequest
	GetInstanceId() *string
	SetName(v string) *CreateUserWafRulesetRequest
	GetName() *string
	SetPhase(v string) *CreateUserWafRulesetRequest
	GetPhase() *string
	SetRules(v []*WafRuleConfig) *CreateUserWafRulesetRequest
	GetRules() []*WafRuleConfig
	SetShared(v *WafBatchRuleShared) *CreateUserWafRulesetRequest
	GetShared() *WafBatchRuleShared
	SetStatus(v string) *CreateUserWafRulesetRequest
	GetStatus() *string
}

type CreateUserWafRulesetRequest struct {
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
	Phase  *string             `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Rules  []*WafRuleConfig    `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateUserWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserWafRulesetRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateUserWafRulesetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserWafRulesetRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserWafRulesetRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateUserWafRulesetRequest) GetRules() []*WafRuleConfig {
	return s.Rules
}

func (s *CreateUserWafRulesetRequest) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *CreateUserWafRulesetRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserWafRulesetRequest) SetDescription(v string) *CreateUserWafRulesetRequest {
	s.Description = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetExpression(v string) *CreateUserWafRulesetRequest {
	s.Expression = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetInstanceId(v string) *CreateUserWafRulesetRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetName(v string) *CreateUserWafRulesetRequest {
	s.Name = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetPhase(v string) *CreateUserWafRulesetRequest {
	s.Phase = &v
	return s
}

func (s *CreateUserWafRulesetRequest) SetRules(v []*WafRuleConfig) *CreateUserWafRulesetRequest {
	s.Rules = v
	return s
}

func (s *CreateUserWafRulesetRequest) SetShared(v *WafBatchRuleShared) *CreateUserWafRulesetRequest {
	s.Shared = v
	return s
}

func (s *CreateUserWafRulesetRequest) SetStatus(v string) *CreateUserWafRulesetRequest {
	s.Status = &v
	return s
}

func (s *CreateUserWafRulesetRequest) Validate() error {
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
