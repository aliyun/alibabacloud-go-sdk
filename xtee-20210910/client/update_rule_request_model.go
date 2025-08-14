// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateRuleRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *UpdateRuleRequest
	GetConsoleRuleId() *int64
	SetEventCode(v string) *UpdateRuleRequest
	GetEventCode() *string
	SetLogicExpression(v string) *UpdateRuleRequest
	GetLogicExpression() *string
	SetMemo(v string) *UpdateRuleRequest
	GetMemo() *string
	SetRegId(v string) *UpdateRuleRequest
	GetRegId() *string
	SetRuleActions(v string) *UpdateRuleRequest
	GetRuleActions() *string
	SetRuleBody(v string) *UpdateRuleRequest
	GetRuleBody() *string
	SetRuleExpressions(v string) *UpdateRuleRequest
	GetRuleExpressions() *string
	SetRuleId(v string) *UpdateRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateRuleRequest
	GetRuleName() *string
	SetRuleStatus(v string) *UpdateRuleRequest
	GetRuleStatus() *string
	SetRuleType(v string) *UpdateRuleRequest
	GetRuleType() *string
	SetRuleVersionId(v int64) *UpdateRuleRequest
	GetRuleVersionId() *int64
}

type UpdateRuleRequest struct {
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
	// Primary key ID of the policy
	//
	// example:
	//
	// 6843
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Event code
	//
	// example:
	//
	// de_agdxgz0246
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
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
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Output action
	//
	// example:
	//
	// [{\\"inputs\\":[\\"rn0301\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// DSL policy expression
	//
	// example:
	//
	// {
	//
	//         "elseIfStatement": [
	//
	//             {
	//
	//                 "condition": {
	//
	//                     "currentId": 0,
	//
	//                     "deepCount": 1,
	//
	//                     "list": [
	//
	//                         {
	//
	//                             "currentId": 0,
	//
	//                             "deepCount": 1,
	//
	//                             "left": {
	//
	//                                 "code": "getLbsRegion(longitude, latitude)?.prov",
	//
	//                                 "description": "根据经纬度得到省份信息，比如经度：111.878062，纬度：22.585409，则经过运算，输出”广东省“",
	//
	//                                 "displayType": "SELF_BIND",
	//
	//                                 "fieldType": "STRING",
	//
	//                                 "functionCode": "",
	//
	//                                 "functionName": "",
	//
	//                                 "hasRightVariable": true,
	//
	//                                 "name": "sl_S02sHLFT7818",
	//
	//                                 "outputThreshold": {
	//
	//                                 },
	//
	//                                 "sourceType": "SAF",
	//
	//                                 "title": "经纬度自定义系统变量",
	//
	//                                 "type": "SELF_BIND"
	//
	//                             },
	//
	//                             "operatorCode": "equals",
	//
	//                             "operatorName": "等于",
	//
	//                             "parentId": 0,
	//
	//                             "right": {
	//
	//                                 "name": "v",
	//
	//                                 "rightVariableType": "constant"
	//
	//                             },
	//
	//                             "sequence": 2
	//
	//                         },
	//
	//                         {
	//
	//                             "currentId": 0,
	//
	//                             "deepCount": 1,
	//
	//                             "list": [
	//
	//                                 {
	//
	//                                     "currentId": 0,
	//
	//                                     "deepCount": 1,
	//
	//                                     "left": {
	//
	//                                         "code": "deReadVelocity(userId,"hK1DMAp1d97",1,"H",0,true,"COUNT")",
	//
	//                                         "description": "测试系统变量累计",
	//
	//                                         "displayType": "SELF_BIND",
	//
	//                                         "fieldType": "DOUBLE",
	//
	//                                         "functionCode": "",
	//
	//                                         "functionName": "",
	//
	//                                         "hasRightVariable": true,
	//
	//                                         "name": "hK1DMAp1d97",
	//
	//                                         "outputThreshold": {
	//
	//                                         },
	//
	//                                         "sourceType": "SAF_SELF",
	//
	//                                         "title": "测试系统变量累计",
	//
	//                                         "type": "SELF_BIND"
	//
	//                                     },
	//
	//                                     "operatorCode": "equals",
	//
	//                                     "operatorName": "等于",
	//
	//                                     "parentId": 0,
	//
	//                                     "right": {
	//
	//                                         "name": "a",
	//
	//                                         "rightVariableType": "constant"
	//
	//                                     },
	//
	//                                     "sequence": 3
	//
	//                                 }
	//
	//                             ],
	//
	//                             "parentId": 0,
	//
	//                             "relationship": "and"
	//
	//                         }
	//
	//                     ],
	//
	//                     "parentId": 0,
	//
	//                     "relationship": "and"
	//
	//                 },
	//
	//                 "then": [
	//
	//                     {
	//
	//                         "inputs": [
	//
	//                             "123"
	//
	//                         ],
	//
	//                         "name": "__addDeScore__",
	//
	//                         "actionType": "SCORE",
	//
	//                         "outputType": "const",
	//
	//                         "inputTitle": "123"
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         ],
	//
	//         "elseStatement": {
	//
	//         },
	//
	//         "ifStatement": {
	//
	//             "condition": {
	//
	//                 "currentId": 0,
	//
	//                 "deepCount": 1,
	//
	//                 "list": [
	//
	//                     {
	//
	//                         "currentId": 0,
	//
	//                         "deepCount": 1,
	//
	//                         "left": {
	//
	//                             "code": "deFunctionProcess(ip,"isIp")",
	//
	//                             "description": "判断是否符合IPv4标准",
	//
	//                             "displayType": "SYSTEM_BIND",
	//
	//                             "fieldType": "BOOLEAN",
	//
	//                             "functionCode": "",
	//
	//                             "functionName": "",
	//
	//                             "hasRightVariable": true,
	//
	//                             "name": "__isIpAddressV4__",
	//
	//                             "outputThreshold": {
	//
	//                             },
	//
	//                             "sourceType": "SAF",
	//
	//                             "title": "IP是否符合IPV4格式",
	//
	//                             "type": "SYSTEM_BIND"
	//
	//                         },
	//
	//                         "operatorCode": "equals",
	//
	//                         "operatorName": "等于",
	//
	//                         "parentId": 0,
	//
	//                         "right": {
	//
	//                             "name": "c
	//
	// d
	//
	// s",
	//
	//                             "rightVariableType": "constant"
	//
	//                         },
	//
	//                         "sequence": 1
	//
	//                     }
	//
	//                 ],
	//
	//                 "parentId": 0,
	//
	//                 "relationship": "and"
	//
	//             },
	//
	//             "then": [
	//
	//                 {
	//
	//                     "inputs": [
	//
	//                         "22"
	//
	//                     ],
	//
	//                     "name": "__addDeTags__",
	//
	//                     "actionType": "TAG",
	//
	//                     "outputType": "const"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     }
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Policy expression
	//
	// example:
	//
	// [{\\"expressionName\\":\\"手机号MD5命中人脸测试名单\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"mobileMd5\\"},\\"operatorCode\\":\\"deInNameList\\",\\"operatorName\\":\\"在名单中\\",\\"right\\":{\\"fieldValue\\":\\"nl_5tolf69W138c\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
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
	// 分析中心事件测试_策略01
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Policy type
	//
	// example:
	//
	// DEFAULT
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the policy version
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateRuleRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *UpdateRuleRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *UpdateRuleRequest) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *UpdateRuleRequest) GetMemo() *string {
	return s.Memo
}

func (s *UpdateRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateRuleRequest) GetRuleActions() *string {
	return s.RuleActions
}

func (s *UpdateRuleRequest) GetRuleBody() *string {
	return s.RuleBody
}

func (s *UpdateRuleRequest) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *UpdateRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRuleRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *UpdateRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *UpdateRuleRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *UpdateRuleRequest) SetLang(v string) *UpdateRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateRuleRequest) SetConsoleRuleId(v int64) *UpdateRuleRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *UpdateRuleRequest) SetEventCode(v string) *UpdateRuleRequest {
	s.EventCode = &v
	return s
}

func (s *UpdateRuleRequest) SetLogicExpression(v string) *UpdateRuleRequest {
	s.LogicExpression = &v
	return s
}

func (s *UpdateRuleRequest) SetMemo(v string) *UpdateRuleRequest {
	s.Memo = &v
	return s
}

func (s *UpdateRuleRequest) SetRegId(v string) *UpdateRuleRequest {
	s.RegId = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleActions(v string) *UpdateRuleRequest {
	s.RuleActions = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleBody(v string) *UpdateRuleRequest {
	s.RuleBody = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleExpressions(v string) *UpdateRuleRequest {
	s.RuleExpressions = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleId(v string) *UpdateRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleName(v string) *UpdateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleStatus(v string) *UpdateRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleType(v string) *UpdateRuleRequest {
	s.RuleType = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleVersionId(v int64) *UpdateRuleRequest {
	s.RuleVersionId = &v
	return s
}

func (s *UpdateRuleRequest) Validate() error {
	return dara.Validate(s)
}
