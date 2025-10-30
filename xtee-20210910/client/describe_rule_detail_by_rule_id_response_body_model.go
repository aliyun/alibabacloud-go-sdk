// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleDetailByRuleIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleDetailByRuleIdResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRuleDetailByRuleIdResponseBodyResultObject) *DescribeRuleDetailByRuleIdResponseBody
	GetResultObject() *DescribeRuleDetailByRuleIdResponseBodyResultObject
}

type DescribeRuleDetailByRuleIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject *DescribeRuleDetailByRuleIdResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRuleDetailByRuleIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleDetailByRuleIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleDetailByRuleIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleDetailByRuleIdResponseBody) GetResultObject() *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleDetailByRuleIdResponseBody) SetRequestId(v string) *DescribeRuleDetailByRuleIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBody) SetResultObject(v *DescribeRuleDetailByRuleIdResponseBodyResultObject) *DescribeRuleDetailByRuleIdResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRuleDetailByRuleIdResponseBodyResultObject struct {
	// Service authorization type.
	//
	// example:
	//
	// all
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// Version.
	//
	// example:
	//
	// 1
	BizVersion *string `json:"bizVersion,omitempty" xml:"bizVersion,omitempty"`
	// Primary key ID of the policy.
	//
	// example:
	//
	// 6633
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Creation type.
	//
	// example:
	//
	// MORMAL
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
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Log expression.
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
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Rule action structure.
	//
	// example:
	//
	// {
	//
	//                 "SCORE": [
	//
	//                     {
	//
	//                         "actionType": "SCORE",
	//
	//                         "displayType": "ACTION",
	//
	//                         "code": "deAddScore",
	//
	//                         "sourceType": "SAF",
	//
	//                         "inputs": [
	//
	//                             "123"
	//
	//                         ],
	//
	//                         "name": "__addDeScore__",
	//
	//                         "description": "打分",
	//
	//                         "inputTitle": "123",
	//
	//                         "outputType": "const",
	//
	//                         "type": "ACTION",
	//
	//                         "title": "打分",
	//
	//                         "fieldType": "DOUBLE"
	//
	//                     }
	//
	//                 ],
	//
	//                 "VARIABLE": [
	//
	//                     {
	//
	//                         "actionType": "VARIABLE",
	//
	//                         "displayType": "MIDDLE",
	//
	//                         "code": "mid1",
	//
	//                         "inputs": [
	//
	//                             "gg"
	//
	//                         ],
	//
	//                         "name": "mid1",
	//
	//                         "description": "中间变量,mid1",
	//
	//                         "type": "MIDDLE",
	//
	//                         "title": "mid1",
	//
	//                         "fieldType": "STRING"
	//
	//                     }
	//
	//                 ],
	//
	//                 "TAG": [
	//
	//                     {
	//
	//                         "actionType": "TAG",
	//
	//                         "displayType": "ACTION",
	//
	//                         "code": "addDeTags",
	//
	//                         "sourceType": "SAF",
	//
	//                         "inputs": [
	//
	//                             "123"
	//
	//                         ],
	//
	//                         "name": "__addDeTags__",
	//
	//                         "description": "打标签",
	//
	//                         "outputType": "const",
	//
	//                         "type": "ACTION",
	//
	//                         "title": "打标签",
	//
	//                         "fieldType": "STRING"
	//
	//                     }
	//
	//                 ],
	//
	//                 "MIDDLE_VARIABLE": [
	//
	//                     {
	//
	//                         "actionType": "MIDDLE_VARIABLE",
	//
	//                         "inputs": [
	//
	//                             "mid1"
	//
	//                         ],
	//
	//                         "fieldValue": "123",
	//
	//                         "fieldType": "STRING"
	//
	//                     }
	//
	//                 ]
	//
	//             }
	RuleActionMap map[string]interface{} `json:"ruleActionMap,omitempty" xml:"ruleActionMap,omitempty"`
	// Rule actions.
	//
	// example:
	//
	// [{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"123\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"},{\\"actionType\\":\\"SCORE\\",\\"code\\":\\"deAddScore\\",\\"description\\":\\"打分\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"DOUBLE\\",\\"inputTitle\\":\\"123\\",\\"inputs\\":[\\"123\\"],\\"name\\":\\"__addDeScore__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打分\\",\\"type\\":\\"ACTION\\"},{\\"actionType\\":\\"MIDDLE_VARIABLE\\",\\"fieldType\\":\\"STRING\\",\\"fieldValue\\":\\"123\\",\\"inputs\\":[\\"mid1\\"]},{\\"actionType\\":\\"VARIABLE\\",\\"code\\":\\"mid1\\",\\"description\\":\\"中间变量,mid1\\",\\"displayType\\":\\"MIDDLE\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"gg\\"],\\"name\\":\\"mid1\\",\\"title\\":\\"mid1\\",\\"type\\":\\"MIDDLE\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Rule authorization type.
	//
	// example:
	//
	// NOMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// DSL policy expression.
	//
	// example:
	//
	// {\\"elseIfStatement\\":[{\\"condition\\":{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"getLbsRegion(longitude, latitude)?.prov\\",\\"description\\":\\"根据经纬度得到省份信息，比如经度：111.878062，纬度：22.585409，则经过运算，输出”广东省“\\",\\"displayType\\":\\"SELF_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"sl_S02sHLFT7818\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"经纬度自定义系统变量\\",\\"type\\":\\"SELF_BIND\\"},\\"operatorCode\\":\\"equals\\",\\"operatorName\\":\\"等于\\",\\"parentId\\":0,\\"right\\":{\\"name\\":\\"cc\\",\\"rightVariableType\\":\\"constant\\"},\\"sequence\\":5}],\\"parentId\\":0,\\"relationship\\":\\"and\\"},\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"332\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]}],\\"elseStatement\\":{\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"321\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]},\\"ifStatement\\":{\\"condition\\":{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"queryPhoneSimulatorInfo(deviceToken)?.brand\\",\\"description\\":\\"设备信息-终端品牌\\",\\"displayType\\":\\"DEVICE\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__device_brand__\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"设备信息-终端品牌-brand\\",\\"type\\":\\"DEVICE\\"},\\"operatorCode\\":\\"deInNameList\\",\\"operatorName\\":\\"在名单中\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"nl_UN8otElLb490\\",\\"description\\":\\"描述11\\",\\"displayType\\":\\"NAME_LIST\\",\\"name\\":\\"nl_UN8otElLb490\\",\\"rightVariableType\\":\\"constant\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"wtz_名单新建测试02\\",\\"type\\":\\"NAME_LIST\\"},\\"sequence\\":1},{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"deFunctionProcess(ip,\\\\\\"isIp\\\\\\")\\",\\"description\\":\\"判断是否符合IPv4标准\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"BOOLEAN\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":false,\\"name\\":\\"__isIpAddressV4__\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP是否符合IPV4格式\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"boolIsTrue\\",\\"operatorName\\":\\"为true\\",\\"parentId\\":0,\\"right\\":{\\"name\\":\\"\\",\\"rightVariableType\\":\\"constant\\"},\\"sequence\\":2},{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"parseIpV2(ip)?.cityId\\",\\"description\\":\\"通过IP地址库解析IP所在的城市Code，例如，输入“42.120.74.211”，经过该变量运算，输出“330100”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__ipLocationCityCode__\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_城市Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"equals\\",\\"operatorName\\":\\"等于\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"deFunctionProcess(ip,\\\\\\"getCountry\\\\\\")\\",\\"description\\":\\"通过IP地址库解析IP所在的城市名称，例如，输入“42.120.74.211”，经过该变量运算，输出“CN”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"name\\":\\"__ipLocationCountryId__\\",\\"outputThreshold\\":{},\\"rightVariableType\\":\\"variable\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_国家Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"sequence\\":3},{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"parseIpV2(ip)?.cityId\\",\\"description\\":\\"通过IP地址库解析IP所在的城市Code，例如，输入“42.120.74.211”，经过该变量运算，输出“330100”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__ipLocationCityCode__\\",\\"outputThreshold\\":{\\"$ref\\":\\"$.ifStatement.condition.list[2].list[0].left.outputThreshold\\"},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_城市Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"deInNameList\\",\\"operatorName\\":\\"在名单中\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"nl_NsVwBD2s11e0\\",\\"displayType\\":\\"NAME_LIST\\",\\"name\\":\\"nl_NsVwBD2s11e0\\",\\"rightVariableType\\":\\"constant\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"device_block_list\\",\\"type\\":\\"NAME_LIST\\"},\\"sequence\\":4}],\\"parentId\\":0,\\"relationship\\":\\"and\\"}],\\"parentId\\":0,\\"relationship\\":\\"and\\"}],\\"parentId\\":0,\\"relationship\\":\\"and\\"},\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"123\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]}}
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Rule expressions.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"营销风险识别评分\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"sl_rjtsDXK124a5\\"},\\"operatorCode\\":\\"between\\",\\"operatorName\\":\\"数字在[a,b]之间\\",\\"right\\":{\\"fieldValue\\":\\"[65,100]\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name.
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status.
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Rule type.
	//
	// example:
	//
	// DSL
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	// Primary key ID of the policy version.
	//
	// example:
	//
	// 3823
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s DescribeRuleDetailByRuleIdResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleDetailByRuleIdResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetBizVersion() *string {
	return s.BizVersion
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetMainEventCode() *string {
	return s.MainEventCode
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleActionMap() map[string]interface{} {
	return s.RuleActionMap
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleActions() *string {
	return s.RuleActions
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleBody() *string {
	return s.RuleBody
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetAuthType(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.AuthType = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetBizVersion(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.BizVersion = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetConsoleRuleId(v int64) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.ConsoleRuleId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetCreateType(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.CreateType = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetEventCode(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetEventName(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetGmtModified(v int64) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetLogicExpression(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.LogicExpression = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetMainEventCode(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.MainEventCode = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetMemo(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleActionMap(v map[string]interface{}) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleActionMap = v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleActions(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleActions = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleAuthType(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleBody(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleBody = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleExpressions(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleExpressions = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleId(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleName(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleStatus(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleType(v string) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) SetRuleVersionId(v int64) *DescribeRuleDetailByRuleIdResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
