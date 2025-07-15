// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *InsertInterveneRuleShrinkRequest
	GetAgentKey() *string
	SetInterveneRuleConfigShrink(v string) *InsertInterveneRuleShrinkRequest
	GetInterveneRuleConfigShrink() *string
}

type InsertInterveneRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey                  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfigShrink *string `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty"`
}

func (s InsertInterveneRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *InsertInterveneRuleShrinkRequest) GetInterveneRuleConfigShrink() *string {
	return s.InterveneRuleConfigShrink
}

func (s *InsertInterveneRuleShrinkRequest) SetAgentKey(v string) *InsertInterveneRuleShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleShrinkRequest) SetInterveneRuleConfigShrink(v string) *InsertInterveneRuleShrinkRequest {
	s.InterveneRuleConfigShrink = &v
	return s
}

func (s *InsertInterveneRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
