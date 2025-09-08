// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeWhiteRuleListResponseBody
	GetCode() *int32
	SetData(v *DescribeWhiteRuleListResponseBodyData) *DescribeWhiteRuleListResponseBody
	GetData() *DescribeWhiteRuleListResponseBodyData
	SetMessage(v string) *DescribeWhiteRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWhiteRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeWhiteRuleListResponseBody
	GetSuccess() *bool
}

type DescribeWhiteRuleListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeWhiteRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWhiteRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeWhiteRuleListResponseBody) GetData() *DescribeWhiteRuleListResponseBodyData {
	return s.Data
}

func (s *DescribeWhiteRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWhiteRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeWhiteRuleListResponseBody) SetCode(v int32) *DescribeWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetData(v *DescribeWhiteRuleListResponseBodyData) *DescribeWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetMessage(v string) *DescribeWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetRequestId(v string) *DescribeWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) SetSuccess(v bool) *DescribeWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeWhiteRuleListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeWhiteRuleListResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeWhiteRuleListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyData) GetPageInfo() *DescribeWhiteRuleListResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeWhiteRuleListResponseBodyData) GetResponseData() []*DescribeWhiteRuleListResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeWhiteRuleListResponseBodyData) SetPageInfo(v *DescribeWhiteRuleListResponseBodyDataPageInfo) *DescribeWhiteRuleListResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyData) SetResponseData(v []*DescribeWhiteRuleListResponseBodyDataResponseData) *DescribeWhiteRuleListResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeWhiteRuleListResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataResponseData struct {
	// The alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameId *string `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The ID of the alert type.
	//
	// example:
	//
	// scan
	AlertTypeId *string `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The conditions in the rule. The value is a JSON array.
	//
	// example:
	//
	// [{"conditions":[{"isNot":false,"itemId":0,"left":{"value":"host_uuid.host_name"},"operator":"containsString","right":{"value":"Cloud-MCH"}}]}]
	Expression *DescribeWhiteRuleListResponseBodyDataResponseDataExpression `json:"Expression,omitempty" xml:"Expression,omitempty" type:"Struct"`
	// The time when the whitelist rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the whitelist rule was modified.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the whitelist rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The status of the whitelist rule. Valid values:
	//
	// 	- 1: enabled.
	//
	// 	- 0: disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the whitelist rule.
	//
	// example:
	//
	// 176555323***
	SubAliuid *int64 `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAlertNameId() *string {
	return s.AlertNameId
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAlertTypeId() *string {
	return s.AlertTypeId
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetExpression() *DescribeWhiteRuleListResponseBodyDataResponseDataExpression {
	return s.Expression
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) GetSubAliuid() *int64 {
	return s.SubAliuid
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertName(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertNameId(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertType(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertTypeId(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetAliuid(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetExpression(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Expression = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetGmtModified(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetId(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetStatus(v int32) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) SetSubAliuid(v int64) *DescribeWhiteRuleListResponseBodyDataResponseData {
	s.SubAliuid = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpression struct {
	// The rule conditions.
	Conditions []*DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The logical relationships among the rule conditions.
	//
	// example:
	//
	// (1&2)|(3&4)
	Logic *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpression) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpression) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) GetConditions() []*DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	return s.Conditions
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) GetLogic() *string {
	return s.Logic
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) SetConditions(v []*DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) *DescribeWhiteRuleListResponseBodyDataResponseDataExpression {
	s.Conditions = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) SetLogic(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpression {
	s.Logic = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpression) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions struct {
	// Indicates whether the result is inverted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsNot *bool `json:"IsNot,omitempty" xml:"IsNot,omitempty"`
	// The ID of the rule condition.
	//
	// example:
	//
	// 1
	ItemId *int32 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The left operand of the rule condition.
	Left *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft `json:"Left,omitempty" xml:"Left,omitempty" type:"Struct"`
	// The logical operator of the rule condition. Valid values:
	//
	// 	- `=`: equals to.
	//
	// 	- `<>`: does not equal to.
	//
	// 	- `in`: contains.
	//
	// 	- `not in`: does not contain.
	//
	// 	- `REGEXP`: matches a regular expression.
	//
	// 	- `NOT REGEXP`: does not match a regular expression.
	//
	// example:
	//
	// REGEXP
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The right operand of the rule condition.
	Right *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight `json:"Right,omitempty" xml:"Right,omitempty" type:"Struct"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GetIsNot() *bool {
	return s.IsNot
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GetItemId() *int32 {
	return s.ItemId
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GetLeft() *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	return s.Left
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GetOperator() *string {
	return s.Operator
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) GetRight() *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	return s.Right
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetIsNot(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.IsNot = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetItemId(v int32) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.ItemId = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetLeft(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Left = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetOperator(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Operator = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) SetRight(v *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions {
	s.Right = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft struct {
	// Indicates whether the left operand is a variable. Valid values:
	//
	// 	- true: variable.
	//
	// 	- false: constant.
	//
	// example:
	//
	// true
	IsVar *bool `json:"IsVar,omitempty" xml:"IsVar,omitempty"`
	// The remarks on the right operand.
	//
	// example:
	//
	// length
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The key-value pair information of the remarks.
	ModifierParam map[string]interface{} `json:"ModifierParam,omitempty" xml:"ModifierParam,omitempty"`
	// Indicates whether the left operand is a constant. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable of the left operand.
	//
	// example:
	//
	// ip
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GetIsVar() *bool {
	return s.IsVar
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GetModifierParam() map[string]interface{} {
	return s.ModifierParam
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GetType() *string {
	return s.Type
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) GetValue() *string {
	return s.Value
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetIsVar(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.IsVar = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetModifier(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Modifier = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetModifierParam(v map[string]interface{}) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.ModifierParam = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetType(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Type = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) SetValue(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft {
	s.Value = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsLeft) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight struct {
	// Indicates whether the right operand is a constant or a runtime variable that is obtained from the runtime context. Valid values:
	//
	// 	- true: runtime variable.
	//
	// 	- false: constant.
	//
	// example:
	//
	// false
	IsVar *bool `json:"IsVar,omitempty" xml:"IsVar,omitempty"`
	// The remarks on the right operand.
	//
	// example:
	//
	// length
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The key-value pair information of the remarks.
	ModifierParam map[string]interface{} `json:"ModifierParam,omitempty" xml:"ModifierParam,omitempty"`
	// The data type of the right operand.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The right operand.
	//
	// example:
	//
	// 12345
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GetIsVar() *bool {
	return s.IsVar
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GetModifierParam() map[string]interface{} {
	return s.ModifierParam
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GetType() *string {
	return s.Type
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) GetValue() *string {
	return s.Value
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetIsVar(v bool) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.IsVar = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetModifier(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Modifier = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetModifierParam(v map[string]interface{}) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.ModifierParam = v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetType(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Type = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) SetValue(v string) *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight {
	s.Value = &v
	return s
}

func (s *DescribeWhiteRuleListResponseBodyDataResponseDataExpressionConditionsRight) Validate() error {
	return dara.Validate(s)
}
