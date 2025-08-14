// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthRuleDetailByRuleIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryAuthRuleDetailByRuleIdResponseBody
	GetRequestId() *string
	SetResultObject(v *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) *QueryAuthRuleDetailByRuleIdResponseBody
	GetResultObject() *QueryAuthRuleDetailByRuleIdResponseBodyResultObject
}

type QueryAuthRuleDetailByRuleIdResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *QueryAuthRuleDetailByRuleIdResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s QueryAuthRuleDetailByRuleIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthRuleDetailByRuleIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuthRuleDetailByRuleIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAuthRuleDetailByRuleIdResponseBody) GetResultObject() *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	return s.ResultObject
}

func (s *QueryAuthRuleDetailByRuleIdResponseBody) SetRequestId(v string) *QueryAuthRuleDetailByRuleIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBody) SetResultObject(v *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) *QueryAuthRuleDetailByRuleIdResponseBody {
	s.ResultObject = v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAuthRuleDetailByRuleIdResponseBodyResultObject struct {
	// Audit ID
	//
	// example:
	//
	// 225
	AuditId *int64 `json:"auditId,omitempty" xml:"auditId,omitempty"`
	// Authorization type
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Authorized user UID
	//
	// example:
	//
	// *
	AuthUsers *string `json:"authUsers,omitempty" xml:"authUsers,omitempty"`
	// Primary key ID of the strategy
	//
	// example:
	//
	// 6843
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Execution logic
	//
	// example:
	//
	// 1&2
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Description
	//
	// example:
	//
	// 描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Rule priority, the higher the number, the higher the priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Returned rule action structure.
	RuleActionMap map[string]*string `json:"ruleActionMap,omitempty" xml:"ruleActionMap,omitempty"`
	// Output actions
	//
	// example:
	//
	// [{\\"inputs\\":[\\"unusualBrand\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Rule authorization type
	//
	// example:
	//
	// WHITE_BOX
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Rule expressions.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"同一设备同一IP上的注册用户数\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"dK7EXHr490f\\"},\\"operatorCode\\":\\"gte\\",\\"operatorName\\":\\"大于等于\\",\\"right\\":{\\"fieldValue\\":\\"2\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Strategy ID
	//
	// example:
	//
	// 102224
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Strategy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Strategy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule type
	//
	// example:
	//
	// DEFAULT
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the strategy version
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// Template type
	//
	// example:
	//
	// PUB_SERVICE
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// Version number
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryAuthRuleDetailByRuleIdResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetAuditId() *int64 {
	return s.AuditId
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetAuthType() *string {
	return s.AuthType
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetAuthUsers() *string {
	return s.AuthUsers
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetPriority() *int64 {
	return s.Priority
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleActionMap() map[string]*string {
	return s.RuleActionMap
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleActions() *string {
	return s.RuleActions
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetAuditId(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.AuditId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetAuthType(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.AuthType = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetAuthUsers(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.AuthUsers = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetConsoleRuleId(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.ConsoleRuleId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetEventCode(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetEventName(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetGmtCreate(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetGmtModified(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetLogicExpression(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.LogicExpression = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetMemo(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetPriority(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.Priority = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleActionMap(v map[string]*string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleActionMap = v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleActions(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleActions = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleAuthType(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleAuthType = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleExpressions(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleExpressions = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleId(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleName(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleStatus(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleType(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetRuleVersionId(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetTemplateType(v string) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.TemplateType = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) SetVersion(v int64) *QueryAuthRuleDetailByRuleIdResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
