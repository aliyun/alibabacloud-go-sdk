// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompareRuleResponseBody
	GetRequestId() *string
	SetResultObject(v *CompareRuleResponseBodyResultObject) *CompareRuleResponseBody
	GetResultObject() *CompareRuleResponseBodyResultObject
}

type CompareRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject *CompareRuleResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s CompareRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CompareRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareRuleResponseBody) GetResultObject() *CompareRuleResponseBodyResultObject {
	return s.ResultObject
}

func (s *CompareRuleResponseBody) SetRequestId(v string) *CompareRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareRuleResponseBody) SetResultObject(v *CompareRuleResponseBodyResultObject) *CompareRuleResponseBody {
	s.ResultObject = v
	return s
}

func (s *CompareRuleResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareRuleResponseBodyResultObject struct {
	// New policy object.
	NewRule *CompareRuleResponseBodyResultObjectNewRule `json:"newRule,omitempty" xml:"newRule,omitempty" type:"Struct"`
	// Old policy object.
	OldRule *CompareRuleResponseBodyResultObjectOldRule `json:"oldRule,omitempty" xml:"oldRule,omitempty" type:"Struct"`
}

func (s CompareRuleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CompareRuleResponseBodyResultObject) GetNewRule() *CompareRuleResponseBodyResultObjectNewRule {
	return s.NewRule
}

func (s *CompareRuleResponseBodyResultObject) GetOldRule() *CompareRuleResponseBodyResultObjectOldRule {
	return s.OldRule
}

func (s *CompareRuleResponseBodyResultObject) SetNewRule(v *CompareRuleResponseBodyResultObjectNewRule) *CompareRuleResponseBodyResultObject {
	s.NewRule = v
	return s
}

func (s *CompareRuleResponseBodyResultObject) SetOldRule(v *CompareRuleResponseBodyResultObjectOldRule) *CompareRuleResponseBodyResultObject {
	s.OldRule = v
	return s
}

func (s *CompareRuleResponseBodyResultObject) Validate() error {
	if s.NewRule != nil {
		if err := s.NewRule.Validate(); err != nil {
			return err
		}
	}
	if s.OldRule != nil {
		if err := s.OldRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareRuleResponseBodyResultObjectNewRule struct {
	// Audit ID.
	//
	// example:
	//
	// 258
	AuditId *int64 `json:"auditId,omitempty" xml:"auditId,omitempty"`
	// Authorization type.
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Primary key ID of the rule.
	//
	// example:
	//
	// 6760
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_asssce8122
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册_事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1760670462000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1761196952000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Logic of the rule expression execution.
	//
	// example:
	//
	// 1&2&3
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Main event code.
	//
	// example:
	//
	// de_asssce8122
	MainEventCode *string `json:"mainEventCode,omitempty" xml:"mainEventCode,omitempty"`
	// Memo.
	//
	// example:
	//
	// 鸿蒙元服务_交费业务
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Returned rule action structure. Returned when the policy type is DEFAULT.
	//
	// example:
	//
	// {
	//
	//     "TAG": "[{\\"code\\":\\"addDeTags\\",\\"inputs\\":[\\"test\\"],\\"description\\":\\"打标签\\",\\"type\\":\\"ACTION\\",\\"title\\":\\"打标签\\",\\"actionType\\":\\"TAG\\",\\"displayType\\":\\"ACTION\\",\\"sourceType\\":\\"SAF\\",\\"name\\":\\"__addDeTags__\\",\\"fieldType\\":\\"STRING\\"}]"
	//
	// }
	RuleActionMap map[string]interface{} `json:"ruleActionMap,omitempty" xml:"ruleActionMap,omitempty"`
	// Output actions of the rule execution. Returned when the rule type is DEFAULT.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "actionType": "TAG",
	//
	//         "code": "addDeTags",
	//
	//         "description": "打标签",
	//
	//         "displayType": "ACTION",
	//
	//         "fieldType": "STRING",
	//
	//         "inputs": [
	//
	//             "test"
	//
	//         ],
	//
	//         "name": "__addDeTags__",
	//
	//         "sourceType": "SAF",
	//
	//         "title": "打标签",
	//
	//         "type": "ACTION"
	//
	//     }
	//
	// ]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Authorization type of the rule.
	//
	// example:
	//
	// NORMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// DSL logic for rule execution. Returned when the rule type is DSL.
	//
	// example:
	//
	// {
	//
	//     "elseStatement": {
	//
	//     },
	//
	//     "ifStatement": {
	//
	//         "condition": {
	//
	//             "currentId": 0,
	//
	//             "deepCount": 1,
	//
	//             "list": [
	//
	//                 {
	//
	//                     "sequence": 1,
	//
	//                     "left": {
	//
	//                         "displayType": "NATIVE",
	//
	//                         "code": "ip",
	//
	//                         "functionCode": "",
	//
	//                         "functionName": "",
	//
	//                         "name": "ip",
	//
	//                         "description": "IP地址",
	//
	//                         "hasRightVariable": true,
	//
	//                         "title": "IP地址",
	//
	//                         "type": "NATIVE",
	//
	//                         "fieldType": "STRING"
	//
	//                     },
	//
	//                     "currentId": 0,
	//
	//                     "deepCount": 1,
	//
	//                     "right": {
	//
	//                         "name": "192.168.1.1",
	//
	//                         "rightVariableType": "constant"
	//
	//                     },
	//
	//                     "operatorCode": "equals",
	//
	//                     "operatorName": "等于",
	//
	//                     "parentId": 0
	//
	//                 }
	//
	//             ],
	//
	//             "relationship": "and",
	//
	//             "parentId": 0
	//
	//         },
	//
	//         "then": [
	//
	//             {
	//
	//                 "actionType": "TAG",
	//
	//                 "displayType": "ACTION",
	//
	//                 "code": "addDeTags",
	//
	//                 "sourceType": "SAF",
	//
	//                 "inputs": [
	//
	//                     "10"
	//
	//                 ],
	//
	//                 "name": "__addDeTags__",
	//
	//                 "description": "打标签",
	//
	//                 "outputType": "const",
	//
	//                 "title": "打标签",
	//
	//                 "type": "ACTION",
	//
	//                 "fieldType": "STRING"
	//
	//             }
	//
	//         ],
	//
	//         "expressions": [
	//
	//             {
	//
	//                 "itemId": 1,
	//
	//                 "left": {
	//
	//                     "displayType": "NATIVE",
	//
	//                     "code": "ip",
	//
	//                     "functionCode": "",
	//
	//                     "functionName": "",
	//
	//                     "name": "ip",
	//
	//                     "description": "IP地址",
	//
	//                     "hasRightVariable": true,
	//
	//                     "title": "IP地址",
	//
	//                     "type": "NATIVE",
	//
	//                     "fieldType": "STRING"
	//
	//                 },
	//
	//                 "expressionName": "IP地址 等于 192.168.1.1",
	//
	//                 "rightValue": "192.168.1.1",
	//
	//                 "right": {
	//
	//                     "name": "192.168.1.1",
	//
	//                     "fieldValue": "192.168.1.1"
	//
	//                 },
	//
	//                 "operatorCode": "equals",
	//
	//                 "operatorName": "等于"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "elseIfStatement": [
	//
	//     ]
	//
	// }
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Policy expressions. Returned when the policy type is DEFAULT.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"设备token不为空\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"deviceToken\\"},\\"operatorCode\\":\\"isNotEmptyWrapper\\",\\"operatorName\\":\\"不为空\\"}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 101793
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Rule name.
	//
	// example:
	//
	// 返回设备信息
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status.
	//
	// example:
	//
	// DRAFT
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule type.
	//
	// example:
	//
	// DSL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the rule version.
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// User UID.
	//
	// example:
	//
	// 151222xxxxxxxxx
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version number.
	//
	// example:
	//
	// 3
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CompareRuleResponseBodyResultObjectNewRule) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleResponseBodyResultObjectNewRule) GoString() string {
	return s.String()
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetAuditId() *int64 {
	return s.AuditId
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetAuthType() *string {
	return s.AuthType
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetCreateType() *string {
	return s.CreateType
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetEventCode() *string {
	return s.EventCode
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetEventName() *string {
	return s.EventName
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetMainEventCode() *string {
	return s.MainEventCode
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetMemo() *string {
	return s.Memo
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleActionMap() map[string]interface{} {
	return s.RuleActionMap
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleActions() *string {
	return s.RuleActions
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleBody() *string {
	return s.RuleBody
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleId() *string {
	return s.RuleId
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleName() *string {
	return s.RuleName
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleType() *string {
	return s.RuleType
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetUserId() *int64 {
	return s.UserId
}

func (s *CompareRuleResponseBodyResultObjectNewRule) GetVersion() *int64 {
	return s.Version
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetAuditId(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.AuditId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetAuthType(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.AuthType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetConsoleRuleId(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.ConsoleRuleId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetCreateType(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.CreateType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetEventCode(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.EventCode = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetEventName(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.EventName = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetGmtCreate(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.GmtCreate = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetGmtModified(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.GmtModified = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetLogicExpression(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.LogicExpression = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetMainEventCode(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.MainEventCode = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetMemo(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.Memo = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleActionMap(v map[string]interface{}) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleActionMap = v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleActions(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleActions = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleAuthType(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleAuthType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleBody(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleBody = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleExpressions(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleExpressions = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleId(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleName(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleName = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleStatus(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleStatus = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleType(v string) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetRuleVersionId(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.RuleVersionId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetUserId(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.UserId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) SetVersion(v int64) *CompareRuleResponseBodyResultObjectNewRule {
	s.Version = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectNewRule) Validate() error {
	return dara.Validate(s)
}

type CompareRuleResponseBodyResultObjectOldRule struct {
	// Audit ID.
	//
	// example:
	//
	// 257
	AuditId *int64 `json:"auditId,omitempty" xml:"auditId,omitempty"`
	// Authorization type.
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Primary key ID of the policy.
	//
	// example:
	//
	// 6760
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_asssce8122
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册_事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1760670462000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1760670462000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Execution logic of the policy expression.
	//
	// example:
	//
	// 1&2
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Main event code.
	//
	// example:
	//
	// de_asssce8122
	MainEventCode *string `json:"mainEventCode,omitempty" xml:"mainEventCode,omitempty"`
	// Description.
	//
	// example:
	//
	// 鸿蒙元服务_交费业务
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Returned rule action structure. Returned when the policy type is DEFAULT.
	//
	// example:
	//
	// {
	//
	//     "TAG": "[{\\"code\\":\\"addDeTags\\",\\"inputs\\":[\\"test\\"],\\"description\\":\\"打标签\\",\\"type\\":\\"ACTION\\",\\"title\\":\\"打标签\\",\\"actionType\\":\\"TAG\\",\\"displayType\\":\\"ACTION\\",\\"sourceType\\":\\"SAF\\",\\"name\\":\\"__addDeTags__\\",\\"fieldType\\":\\"STRING\\"}]"
	//
	// }
	RuleActionMap map[string]interface{} `json:"ruleActionMap,omitempty" xml:"ruleActionMap,omitempty"`
	// Policy execution output actions. Returned when the policy type is DEFAULT.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "actionType": "TAG",
	//
	//         "code": "addDeTags",
	//
	//         "description": "打标签",
	//
	//         "displayType": "ACTION",
	//
	//         "fieldType": "STRING",
	//
	//         "inputs": [
	//
	//             "test"
	//
	//         ],
	//
	//         "name": "__addDeTags__",
	//
	//         "sourceType": "SAF",
	//
	//         "title": "打标签",
	//
	//         "type": "ACTION"
	//
	//     }
	//
	// ]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Policy authorization type.
	//
	// example:
	//
	// NORMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// DSL policy execution logic. Returned when the policy type is DSL.
	//
	// example:
	//
	// {
	//
	//     "elseStatement": {
	//
	//     },
	//
	//     "ifStatement": {
	//
	//         "condition": {
	//
	//             "currentId": 0,
	//
	//             "deepCount": 1,
	//
	//             "list": [
	//
	//                 {
	//
	//                     "sequence": 1,
	//
	//                     "left": {
	//
	//                         "displayType": "NATIVE",
	//
	//                         "code": "ip",
	//
	//                         "functionCode": "",
	//
	//                         "functionName": "",
	//
	//                         "name": "ip",
	//
	//                         "description": "IP地址",
	//
	//                         "hasRightVariable": true,
	//
	//                         "title": "IP地址",
	//
	//                         "type": "NATIVE",
	//
	//                         "fieldType": "STRING"
	//
	//                     },
	//
	//                     "currentId": 0,
	//
	//                     "deepCount": 1,
	//
	//                     "right": {
	//
	//                         "name": "192.168.1.1",
	//
	//                         "rightVariableType": "constant"
	//
	//                     },
	//
	//                     "operatorCode": "equals",
	//
	//                     "operatorName": "等于",
	//
	//                     "parentId": 0
	//
	//                 }
	//
	//             ],
	//
	//             "relationship": "and",
	//
	//             "parentId": 0
	//
	//         },
	//
	//         "then": [
	//
	//             {
	//
	//                 "actionType": "TAG",
	//
	//                 "displayType": "ACTION",
	//
	//                 "code": "addDeTags",
	//
	//                 "sourceType": "SAF",
	//
	//                 "inputs": [
	//
	//                     "10"
	//
	//                 ],
	//
	//                 "name": "__addDeTags__",
	//
	//                 "description": "打标签",
	//
	//                 "outputType": "const",
	//
	//                 "title": "打标签",
	//
	//                 "type": "ACTION",
	//
	//                 "fieldType": "STRING"
	//
	//             }
	//
	//         ],
	//
	//         "expressions": [
	//
	//             {
	//
	//                 "itemId": 1,
	//
	//                 "left": {
	//
	//                     "displayType": "NATIVE",
	//
	//                     "code": "ip",
	//
	//                     "functionCode": "",
	//
	//                     "functionName": "",
	//
	//                     "name": "ip",
	//
	//                     "description": "IP地址",
	//
	//                     "hasRightVariable": true,
	//
	//                     "title": "IP地址",
	//
	//                     "type": "NATIVE",
	//
	//                     "fieldType": "STRING"
	//
	//                 },
	//
	//                 "expressionName": "IP地址 等于 192.168.1.1",
	//
	//                 "rightValue": "192.168.1.1",
	//
	//                 "right": {
	//
	//                     "name": "192.168.1.1",
	//
	//                     "fieldValue": "192.168.1.1"
	//
	//                 },
	//
	//                 "operatorCode": "equals",
	//
	//                 "operatorName": "等于"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "elseIfStatement": [
	//
	//     ]
	//
	// }
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Policy expression. Returned when the policy type is DEFAULT.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"设备token不为空\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"deviceToken\\"},\\"operatorCode\\":\\"isNotEmptyWrapper\\",\\"operatorName\\":\\"不为空\\"}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 101793
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name.
	//
	// example:
	//
	// 返回设备信息
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status.
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Policy type.
	//
	// example:
	//
	// DSL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the policy version.
	//
	// example:
	//
	// 11518
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// User UID.
	//
	// example:
	//
	// 151222xxxxxxxxx
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version number.
	//
	// example:
	//
	// 2
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CompareRuleResponseBodyResultObjectOldRule) String() string {
	return dara.Prettify(s)
}

func (s CompareRuleResponseBodyResultObjectOldRule) GoString() string {
	return s.String()
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetAuditId() *int64 {
	return s.AuditId
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetAuthType() *string {
	return s.AuthType
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetCreateType() *string {
	return s.CreateType
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetEventCode() *string {
	return s.EventCode
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetEventName() *string {
	return s.EventName
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetMainEventCode() *string {
	return s.MainEventCode
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetMemo() *string {
	return s.Memo
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleActionMap() map[string]interface{} {
	return s.RuleActionMap
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleActions() *string {
	return s.RuleActions
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleBody() *string {
	return s.RuleBody
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleId() *string {
	return s.RuleId
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleName() *string {
	return s.RuleName
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleType() *string {
	return s.RuleType
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetUserId() *int64 {
	return s.UserId
}

func (s *CompareRuleResponseBodyResultObjectOldRule) GetVersion() *int64 {
	return s.Version
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetAuditId(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.AuditId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetAuthType(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.AuthType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetConsoleRuleId(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.ConsoleRuleId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetCreateType(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.CreateType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetEventCode(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.EventCode = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetEventName(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.EventName = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetGmtCreate(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.GmtCreate = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetGmtModified(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.GmtModified = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetLogicExpression(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.LogicExpression = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetMainEventCode(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.MainEventCode = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetMemo(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.Memo = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleActionMap(v map[string]interface{}) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleActionMap = v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleActions(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleActions = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleAuthType(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleAuthType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleBody(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleBody = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleExpressions(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleExpressions = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleId(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleName(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleName = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleStatus(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleStatus = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleType(v string) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleType = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetRuleVersionId(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.RuleVersionId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetUserId(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.UserId = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) SetVersion(v int64) *CompareRuleResponseBodyResultObjectOldRule {
	s.Version = &v
	return s
}

func (s *CompareRuleResponseBodyResultObjectOldRule) Validate() error {
	return dara.Validate(s)
}
