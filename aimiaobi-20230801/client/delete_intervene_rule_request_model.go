// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterveneRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteInterveneRuleRequest
	GetAgentKey() *string
	SetRuleId(v int64) *DeleteInterveneRuleRequest
	GetRuleId() *int64
}

type DeleteInterveneRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteInterveneRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteInterveneRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteInterveneRuleRequest) SetAgentKey(v string) *DeleteInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteInterveneRuleRequest) SetRuleId(v int64) *DeleteInterveneRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteInterveneRuleRequest) Validate() error {
	return dara.Validate(s)
}
