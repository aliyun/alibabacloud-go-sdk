// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateRuleV4Request
	GetBaseMeAgentId() *int64
	SetJsonStrForRule(v string) *UpdateRuleV4Request
	GetJsonStrForRule() *string
	SetRuleId(v int64) *UpdateRuleV4Request
	GetRuleId() *int64
}

type UpdateRuleV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateRuleV4Request) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleV4Request) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateRuleV4Request) GetJsonStrForRule() *string {
	return s.JsonStrForRule
}

func (s *UpdateRuleV4Request) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateRuleV4Request) SetBaseMeAgentId(v int64) *UpdateRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleV4Request) SetJsonStrForRule(v string) *UpdateRuleV4Request {
	s.JsonStrForRule = &v
	return s
}

func (s *UpdateRuleV4Request) SetRuleId(v int64) *UpdateRuleV4Request {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleV4Request) Validate() error {
	return dara.Validate(s)
}
