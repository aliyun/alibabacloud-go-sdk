// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRulePriorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyRulePriorityRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *ModifyRulePriorityRequest
	GetConsoleRuleId() *int64
	SetPriority(v int32) *ModifyRulePriorityRequest
	GetPriority() *int32
	SetRegId(v string) *ModifyRulePriorityRequest
	GetRegId() *string
	SetRuleId(v int64) *ModifyRulePriorityRequest
	GetRuleId() *int64
}

type ModifyRulePriorityRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID of the policy
	//
	// example:
	//
	// 3581
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Policy priority, the higher the number, the higher the priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101796
	RuleId *int64 `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s ModifyRulePriorityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRulePriorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyRulePriorityRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyRulePriorityRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *ModifyRulePriorityRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyRulePriorityRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyRulePriorityRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyRulePriorityRequest) SetLang(v string) *ModifyRulePriorityRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRulePriorityRequest) SetConsoleRuleId(v int64) *ModifyRulePriorityRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *ModifyRulePriorityRequest) SetPriority(v int32) *ModifyRulePriorityRequest {
	s.Priority = &v
	return s
}

func (s *ModifyRulePriorityRequest) SetRegId(v string) *ModifyRulePriorityRequest {
	s.RegId = &v
	return s
}

func (s *ModifyRulePriorityRequest) SetRuleId(v int64) *ModifyRulePriorityRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyRulePriorityRequest) Validate() error {
	return dara.Validate(s)
}
