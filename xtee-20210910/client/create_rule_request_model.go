// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateRuleRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *CreateRuleRequest
	GetConsoleRuleId() *int64
	SetCreateType(v string) *CreateRuleRequest
	GetCreateType() *string
	SetEventCode(v string) *CreateRuleRequest
	GetEventCode() *string
	SetEventName(v string) *CreateRuleRequest
	GetEventName() *string
	SetLogicExpression(v string) *CreateRuleRequest
	GetLogicExpression() *string
	SetMemo(v string) *CreateRuleRequest
	GetMemo() *string
	SetRegId(v string) *CreateRuleRequest
	GetRegId() *string
	SetRuleActions(v string) *CreateRuleRequest
	GetRuleActions() *string
	SetRuleBody(v string) *CreateRuleRequest
	GetRuleBody() *string
	SetRuleExpressions(v string) *CreateRuleRequest
	GetRuleExpressions() *string
	SetRuleName(v string) *CreateRuleRequest
	GetRuleName() *string
	SetRuleStatus(v string) *CreateRuleRequest
	GetRuleStatus() *string
	SetRuleType(v string) *CreateRuleRequest
	GetRuleType() *string
}

type CreateRuleRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 5178
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_acytyt7036
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 登录事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Policy expression execution logic
	//
	// example:
	//
	// 1&2
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Memo
	//
	// example:
	//
	// 描述信息
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy execution output action
	//
	// example:
	//
	// [{"inputs":["123"],"name":"__addDeTags__","actionType":"TAG","outputType":"const"},{"inputs":["123"],"name":"__addDeScore__","actionType":"SCORE","outputType":"const","inputTitle":"123"},{"actionType":"MIDDLE_VARIABLE","fieldValue":"123","inputs":["mid1"]},{"actionType":"VARIABLE","inputs":["gg"],"name":"mid1"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// DSL policy execution logic
	//
	// example:
	//
	// {"elseIfStatement":[],"elseStatement":{},"ifStatement":{"condition":{"currentId":0,"deepCount":1,"list":[{"currentId":0,"deepCount":1,"left":{"code":"deFunctionProcess(ip,\\"isIp\\")","description":"判断是否符合IPv4标准","displayType":"SYSTEM_BIND","fieldType":"BOOLEAN","functionCode":"","functionName":"","hasRightVariable":false,"name":"__isIpAddressV4__","outputThreshold":{},"sourceType":"SAF","title":"IP是否符合IPV4格式","type":"SYSTEM_BIND"},"operatorCode":"boolIsFalse","operatorName":"为false","parentId":0,"sequence":1}],"parentId":0,"relationship":"and"},"then":[{"inputs":["123"],"name":"__addDeTags__","actionType":"TAG","outputType":"const"}]}}
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Policy expression
	//
	// example:
	//
	// [{"expressionName":"cc","itemId":1,"left":{"name":"__ipLocationCityCode__"},"operatorCode":"equals","operatorName":"等于","right":{"fieldValue":"a"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy name
	//
	// example:
	//
	// 注册手机号是11位数字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status
	//
	// example:
	//
	// DRAFT
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Policy type
	//
	// example:
	//
	// DRAFT
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateRuleRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *CreateRuleRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateRuleRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *CreateRuleRequest) GetEventName() *string {
	return s.EventName
}

func (s *CreateRuleRequest) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *CreateRuleRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateRuleRequest) GetRuleActions() *string {
	return s.RuleActions
}

func (s *CreateRuleRequest) GetRuleBody() *string {
	return s.RuleBody
}

func (s *CreateRuleRequest) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *CreateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRuleRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CreateRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateRuleRequest) SetLang(v string) *CreateRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateRuleRequest) SetConsoleRuleId(v int64) *CreateRuleRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *CreateRuleRequest) SetCreateType(v string) *CreateRuleRequest {
	s.CreateType = &v
	return s
}

func (s *CreateRuleRequest) SetEventCode(v string) *CreateRuleRequest {
	s.EventCode = &v
	return s
}

func (s *CreateRuleRequest) SetEventName(v string) *CreateRuleRequest {
	s.EventName = &v
	return s
}

func (s *CreateRuleRequest) SetLogicExpression(v string) *CreateRuleRequest {
	s.LogicExpression = &v
	return s
}

func (s *CreateRuleRequest) SetMemo(v string) *CreateRuleRequest {
	s.Memo = &v
	return s
}

func (s *CreateRuleRequest) SetRegId(v string) *CreateRuleRequest {
	s.RegId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleActions(v string) *CreateRuleRequest {
	s.RuleActions = &v
	return s
}

func (s *CreateRuleRequest) SetRuleBody(v string) *CreateRuleRequest {
	s.RuleBody = &v
	return s
}

func (s *CreateRuleRequest) SetRuleExpressions(v string) *CreateRuleRequest {
	s.RuleExpressions = &v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleRequest) SetRuleStatus(v string) *CreateRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *CreateRuleRequest) SetRuleType(v string) *CreateRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateRuleRequest) Validate() error {
	return dara.Validate(s)
}
