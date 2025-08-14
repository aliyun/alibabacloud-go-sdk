// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateRuleBaseRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *UpdateRuleBaseRequest
	GetConsoleRuleId() *int64
	SetEventCode(v string) *UpdateRuleBaseRequest
	GetEventCode() *string
	SetMemo(v string) *UpdateRuleBaseRequest
	GetMemo() *string
	SetRegId(v string) *UpdateRuleBaseRequest
	GetRegId() *string
	SetRuleId(v string) *UpdateRuleBaseRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateRuleBaseRequest
	GetRuleName() *string
}

type UpdateRuleBaseRequest struct {
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
	// Policy primary key ID
	//
	// example:
	//
	// 6843
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Event code
	//
	// example:
	//
	// de_arcehq4370
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Description
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
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
	// 102224
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s UpdateRuleBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleBaseRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateRuleBaseRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *UpdateRuleBaseRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *UpdateRuleBaseRequest) GetMemo() *string {
	return s.Memo
}

func (s *UpdateRuleBaseRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateRuleBaseRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateRuleBaseRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRuleBaseRequest) SetLang(v string) *UpdateRuleBaseRequest {
	s.Lang = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetConsoleRuleId(v int64) *UpdateRuleBaseRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetEventCode(v string) *UpdateRuleBaseRequest {
	s.EventCode = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetMemo(v string) *UpdateRuleBaseRequest {
	s.Memo = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetRegId(v string) *UpdateRuleBaseRequest {
	s.RegId = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetRuleId(v string) *UpdateRuleBaseRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleBaseRequest) SetRuleName(v string) *UpdateRuleBaseRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleBaseRequest) Validate() error {
	return dara.Validate(s)
}
