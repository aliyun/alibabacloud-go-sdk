// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventBaseInfoByEventCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventBaseInfoByEventCodeResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) *DescribeEventBaseInfoByEventCodeResponseBody
	GetResultObject() *DescribeEventBaseInfoByEventCodeResponseBodyResultObject
}

type DescribeEventBaseInfoByEventCodeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeEventBaseInfoByEventCodeResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeEventBaseInfoByEventCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventBaseInfoByEventCodeResponseBody) GetResultObject() *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventBaseInfoByEventCodeResponseBody) SetRequestId(v string) *DescribeEventBaseInfoByEventCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBody) SetResultObject(v *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) *DescribeEventBaseInfoByEventCodeResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventBaseInfoByEventCodeResponseBodyResultObject struct {
	// Business version number
	//
	// example:
	//
	// 1
	BizVersion *int32 `json:"bizVersion,omitempty" xml:"bizVersion,omitempty"`
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
	// Event status.
	//
	// example:
	//
	// ONLINE
	EventStauts *string `json:"eventStauts,omitempty" xml:"eventStauts,omitempty"`
	// Field list.
	InputFields []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields `json:"inputFields,omitempty" xml:"inputFields,omitempty" type:"Repeated"`
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Policy Information
	RuleDetails []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails `json:"ruleDetails,omitempty" xml:"ruleDetails,omitempty" type:"Repeated"`
	// Operation template code
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Template name
	//
	// example:
	//
	// 注册事件模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Template type.
	//
	// example:
	//
	// TEMPLATE
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetBizVersion() *int32 {
	return s.BizVersion
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetEventStauts() *string {
	return s.EventStauts
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetInputFields() []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	return s.InputFields
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetRuleDetails() []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	return s.RuleDetails
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetBizVersion(v int32) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.BizVersion = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetEventCode(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetEventName(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetEventStauts(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.EventStauts = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetInputFields(v []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.InputFields = v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetMemo(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetRuleDetails(v []*DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.RuleDetails = v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetTemplateCode(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetTemplateName(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.TemplateName = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) SetTemplateType(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObject {
	s.TemplateType = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields struct {
	// Field description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Field code
	//
	// example:
	//
	// age
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// Field ranking
	//
	// example:
	//
	// 1
	FieldRank *string `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	// Field source.
	//
	// example:
	//
	// DEFAULT
	FieldSource *string `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	// Field type.
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Field name.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetFieldCode() *string {
	return s.FieldCode
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetFieldRank() *string {
	return s.FieldRank
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetDescription(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.Description = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetFieldCode(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.FieldCode = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetFieldRank(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetFieldSource(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetFieldType(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.FieldType = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) SetTitle(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields {
	s.Title = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectInputFields) Validate() error {
	return dara.Validate(s)
}

type DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails struct {
	// Policy Execution Logic
	//
	// example:
	//
	// 3&((1&2&4)
	LogicExpression *string `json:"logicExpression,omitempty" xml:"logicExpression,omitempty"`
	// Memo
	//
	// example:
	//
	// 描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Rule Actions
	//
	// example:
	//
	// [{\\"inputs\\":[\\"auto_accselist\\"],\\"name\\":\\"__addDeTags__\\",\\"actionType\\":\\"TAG\\",\\"outputType\\":\\"const\\"}]
	RuleActions *string `json:"ruleActions,omitempty" xml:"ruleActions,omitempty"`
	// Policy Type
	//
	// example:
	//
	// DEFAULT
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Event Expressions.
	//
	// example:
	//
	// [{\\"expressionName\\":\\"同一设备同一IP上的注册用户数\\",\\"itemId\\":1,\\"left\\":{\\"name\\":\\"dK7EXHr490f\\"},\\"operatorCode\\":\\"gte\\",\\"operatorName\\":\\"大于等于\\",\\"right\\":{\\"fieldValue\\":\\"2\\"}}]
	RuleExpressions *string `json:"ruleExpressions,omitempty" xml:"ruleExpressions,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 101544
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy Name
	//
	// example:
	//
	// 手机号MD5命中人脸测试名单
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy Status
	//
	// example:
	//
	// DRAFT
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetLogicExpression() *string {
	return s.LogicExpression
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetMemo() *string {
	return s.Memo
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleActions() *string {
	return s.RuleActions
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleExpressions() *string {
	return s.RuleExpressions
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetLogicExpression(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.LogicExpression = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetMemo(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.Memo = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleActions(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleActions = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleAuthType(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleExpressions(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleExpressions = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleId(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleId = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleName(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleName = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) SetRuleStatus(v string) *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails {
	s.RuleStatus = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponseBodyResultObjectRuleDetails) Validate() error {
	return dara.Validate(s)
}
