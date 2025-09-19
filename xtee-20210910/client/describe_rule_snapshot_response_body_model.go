// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleSnapshotResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRuleSnapshotResponseBodyResultObject) *DescribeRuleSnapshotResponseBody
	GetResultObject() *DescribeRuleSnapshotResponseBodyResultObject
}

type DescribeRuleSnapshotResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeRuleSnapshotResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeRuleSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleSnapshotResponseBody) GetResultObject() *DescribeRuleSnapshotResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleSnapshotResponseBody) SetRequestId(v string) *DescribeRuleSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBody) SetResultObject(v *DescribeRuleSnapshotResponseBodyResultObject) *DescribeRuleSnapshotResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleSnapshotResponseBodyResultObject struct {
	// Business version.
	//
	// example:
	//
	// 1
	BizVersion *string `json:"bizVersion,omitempty" xml:"bizVersion,omitempty"`
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
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Expression for analysis results.
	//
	// example:
	//
	// 1&2
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Rule actions.
	//
	// example:
	//
	// [{\\"inputs\\":[\\"LdShop\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// DSL rule expression. This field is required when ruleType is DSL.
	//
	// example:
	//
	// {\\"elseIfStatement\\":[{\\"condition\\":{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"getLbsRegion(longitude, latitude)?.prov\\",\\"description\\":\\"根据经纬度得到省份信息，比如经度：111.878062，纬度：22.585409，则经过运算，输出”广东省“\\",\\"displayType\\":\\"SELF_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"sl_S02sHLFT7818\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"经纬度自定义系统变量\\",\\"type\\":\\"SELF_BIND\\"},\\"operatorCode\\":\\"equals\\",\\"operatorName\\":\\"等于\\",\\"parentId\\":0,\\"right\\":{\\"name\\":\\"cc\\",\\"rightVariableType\\":\\"constant\\"},\\"sequence\\":5}],\\"parentId\\":0,\\"relationship\\":\\"and\\"},\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"332\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]}],\\"elseStatement\\":{\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"321\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]},\\"ifStatement\\":{\\"condition\\":{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"queryPhoneSimulatorInfo(deviceToken)?.brand\\",\\"description\\":\\"设备信息-终端品牌\\",\\"displayType\\":\\"DEVICE\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__device_brand__\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"设备信息-终端品牌-brand\\",\\"type\\":\\"DEVICE\\"},\\"operatorCode\\":\\"deInNameList\\",\\"operatorName\\":\\"在名单中\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"nl_UN8otElLb490\\",\\"description\\":\\"描述11\\",\\"displayType\\":\\"NAME_LIST\\",\\"name\\":\\"nl_UN8otElLb490\\",\\"rightVariableType\\":\\"constant\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"wtz_名单新建测试02\\",\\"type\\":\\"NAME_LIST\\"},\\"sequence\\":1},{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"deFunctionProcess(ip,\\\\\\"isIp\\\\\\")\\",\\"description\\":\\"判断是否符合IPv4标准\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"BOOLEAN\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":false,\\"name\\":\\"__isIpAddressV4__\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP是否符合IPV4格式\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"boolIsTrue\\",\\"operatorName\\":\\"为true\\",\\"parentId\\":0,\\"right\\":{\\"name\\":\\"\\",\\"rightVariableType\\":\\"constant\\"},\\"sequence\\":2},{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"parseIpV2(ip)?.cityId\\",\\"description\\":\\"通过IP地址库解析IP所在的城市Code，例如，输入“42.120.74.211”，经过该变量运算，输出“330100”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__ipLocationCityCode__\\",\\"outputThreshold\\":{},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_城市Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"equals\\",\\"operatorName\\":\\"等于\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"deFunctionProcess(ip,\\\\\\"getCountry\\\\\\")\\",\\"description\\":\\"通过IP地址库解析IP所在的城市名称，例如，输入“42.120.74.211”，经过该变量运算，输出“CN”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"name\\":\\"__ipLocationCountryId__\\",\\"outputThreshold\\":{},\\"rightVariableType\\":\\"variable\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_国家Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"sequence\\":3},{\\"currentId\\":0,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":0,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"parseIpV2(ip)?.cityId\\",\\"description\\":\\"通过IP地址库解析IP所在的城市Code，例如，输入“42.120.74.211”，经过该变量运算，输出“330100”。\\",\\"displayType\\":\\"SYSTEM_BIND\\",\\"fieldType\\":\\"STRING\\",\\"functionCode\\":\\"\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"__ipLocationCityCode__\\",\\"outputThreshold\\":{\\"$ref\\":\\"$.ifStatement.condition.list[2].list[0].left.outputThreshold\\"},\\"sourceType\\":\\"SAF\\",\\"title\\":\\"IP所在地_城市Code\\",\\"type\\":\\"SYSTEM_BIND\\"},\\"operatorCode\\":\\"deInNameList\\",\\"operatorName\\":\\"在名单中\\",\\"parentId\\":0,\\"right\\":{\\"code\\":\\"nl_NsVwBD2s11e0\\",\\"displayType\\":\\"NAME_LIST\\",\\"name\\":\\"nl_NsVwBD2s11e0\\",\\"rightVariableType\\":\\"constant\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"device_block_list\\",\\"type\\":\\"NAME_LIST\\"},\\"sequence\\":4}],\\"parentId\\":0,\\"relationship\\":\\"and\\"}],\\"parentId\\":0,\\"relationship\\":\\"and\\"}],\\"parentId\\":0,\\"relationship\\":\\"and\\"},\\"then\\":[{\\"actionType\\":\\"TAG\\",\\"code\\":\\"addDeTags\\",\\"description\\":\\"打标签\\",\\"displayType\\":\\"ACTION\\",\\"fieldType\\":\\"STRING\\",\\"inputs\\":[\\"123\\"],\\"name\\":\\"__addDeTags__\\",\\"outputType\\":\\"const\\",\\"sourceType\\":\\"SAF\\",\\"title\\":\\"打标签\\",\\"type\\":\\"ACTION\\"}]}}
	RuleBody *string `json:"ruleBody,omitempty" xml:"ruleBody,omitempty"`
	// Expression.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"代下单_记录日志\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"dhcfX2v7670\\"},\\"operatorCode\\":\\"gte\\",\\"operatorName\\":\\"大于等于\\",\\"right\\":{\\"fieldValue\\":\\"2\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101804
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status
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
}

func (s DescribeRuleSnapshotResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleSnapshotResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetBizVersion() *string {
	return s.BizVersion
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleActions() *string {
	return s.RuleActions
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleBody() *string {
	return s.RuleBody
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetBizVersion(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.BizVersion = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetEventCode(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetEventName(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRuleSnapshotResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetGmtModified(v int64) *DescribeRuleSnapshotResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetLogicExpression(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.LogicExpression = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetMemo(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleActions(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleActions = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleBody(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleBody = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleExpressions(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleExpressions = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleId(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleName(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleStatus(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) SetRuleType(v string) *DescribeRuleSnapshotResponseBodyResultObject {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleSnapshotResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
