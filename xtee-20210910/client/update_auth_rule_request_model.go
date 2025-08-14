// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateAuthRuleRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *UpdateAuthRuleRequest
	GetConsoleRuleId() *int64
	SetEventCode(v string) *UpdateAuthRuleRequest
	GetEventCode() *string
	SetRegId(v string) *UpdateAuthRuleRequest
	GetRegId() *string
	SetRuleActions(v string) *UpdateAuthRuleRequest
	GetRuleActions() *string
	SetRuleExpressions(v string) *UpdateAuthRuleRequest
	GetRuleExpressions() *string
	SetRuleId(v string) *UpdateAuthRuleRequest
	GetRuleId() *string
	SetRuleVersionId(v int64) *UpdateAuthRuleRequest
	GetRuleVersionId() *int64
}

type UpdateAuthRuleRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
	// 7088
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Event code
	//
	// example:
	//
	// de_afghcf6411
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy output action
	//
	// example:
	//
	// [{\\"inputs\\":[\\"auto_accselist\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Expression
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"设备token不为空\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"deviceToken\\"},\\"operatorCode\\":\\"isNotEmptyWrapper\\",\\"operatorName\\":\\"不为空\\"}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy version primary key ID
	//
	// example:
	//
	// 5190
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s UpdateAuthRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAuthRuleRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *UpdateAuthRuleRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *UpdateAuthRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateAuthRuleRequest) GetRuleActions() *string {
	return s.RuleActions
}

func (s *UpdateAuthRuleRequest) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *UpdateAuthRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateAuthRuleRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *UpdateAuthRuleRequest) SetLang(v string) *UpdateAuthRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetConsoleRuleId(v int64) *UpdateAuthRuleRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetEventCode(v string) *UpdateAuthRuleRequest {
	s.EventCode = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetRegId(v string) *UpdateAuthRuleRequest {
	s.RegId = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetRuleActions(v string) *UpdateAuthRuleRequest {
	s.RuleActions = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetRuleExpressions(v string) *UpdateAuthRuleRequest {
	s.RuleExpressions = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetRuleId(v string) *UpdateAuthRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateAuthRuleRequest) SetRuleVersionId(v int64) *UpdateAuthRuleRequest {
	s.RuleVersionId = &v
	return s
}

func (s *UpdateAuthRuleRequest) Validate() error {
	return dara.Validate(s)
}
