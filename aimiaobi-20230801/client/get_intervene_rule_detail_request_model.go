// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetInterveneRuleDetailRequest
	GetAgentKey() *string
	SetRuleId(v int64) *GetInterveneRuleDetailRequest
	GetRuleId() *int64
}

type GetInterveneRuleDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetInterveneRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetInterveneRuleDetailRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetInterveneRuleDetailRequest) SetAgentKey(v string) *GetInterveneRuleDetailRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneRuleDetailRequest) SetRuleId(v int64) *GetInterveneRuleDetailRequest {
	s.RuleId = &v
	return s
}

func (s *GetInterveneRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
