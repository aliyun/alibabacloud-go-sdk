// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQualityRulesResponseBodyData) *ListQualityRulesResponseBody
	GetData() *ListQualityRulesResponseBodyData
	SetErrorCode(v string) *ListQualityRulesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListQualityRulesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListQualityRulesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListQualityRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityRulesResponseBody
	GetSuccess() *bool
}

type ListQualityRulesResponseBody struct {
	// The list of monitoring rules.
	Data *ListQualityRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can troubleshoot errors based on the ID.
	//
	// example:
	//
	// 38cbdef0-f6cf-49****
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

func (s ListQualityRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBody) GetData() *ListQualityRulesResponseBodyData {
	return s.Data
}

func (s *ListQualityRulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListQualityRulesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListQualityRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityRulesResponseBody) SetData(v *ListQualityRulesResponseBodyData) *ListQualityRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityRulesResponseBody) SetErrorCode(v string) *ListQualityRulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetErrorMessage(v string) *ListQualityRulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetHttpStatusCode(v int32) *ListQualityRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetRequestId(v string) *ListQualityRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityRulesResponseBody) SetSuccess(v bool) *ListQualityRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of the monitoring rules.
	Rules []*ListQualityRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 400
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityRulesResponseBodyData) GetRules() []*ListQualityRulesResponseBodyDataRules {
	return s.Rules
}

func (s *ListQualityRulesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityRulesResponseBodyData) SetPageNumber(v int32) *ListQualityRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListQualityRulesResponseBodyData) SetPageSize(v int32) *ListQualityRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQualityRulesResponseBodyData) SetRules(v []*ListQualityRulesResponseBodyDataRules) *ListQualityRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *ListQualityRulesResponseBodyData) SetTotalCount(v int64) *ListQualityRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListQualityRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListQualityRulesResponseBodyDataRules struct {
	// The strength of the monitoring rule. The strength of a monitoring rule indicates the importance of the rule. Valid values:
	//
	// 	- 1: The monitoring rule is a strong rule.
	//
	// 	- 0: The monitoring rule is a weak rule. You can specify the strength of a monitoring rule based on your business requirements. If a monitoring rule is a strong rule and the critical threshold is exceeded, a critical alert is reported and tasks that are associated with the rule are blocked from running.
	//
	// example:
	//
	// 0
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The checker ID.
	//
	// example:
	//
	// 7
	CheckerId *int32 `json:"CheckerId,omitempty" xml:"CheckerId,omitempty"`
	// The description of the monitoring rule.
	//
	// example:
	//
	// Verify table rules
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical alert. The threshold indicates the deviation of the monitoring result from the expected value. You can specify a custom value for the threshold based on your business requirements. If a monitoring rule is a strong rule and the critical threshold is exceeded, a critical alert is reported and tasks that are associated with the rule are blocked from running.
	//
	// example:
	//
	// 40
	CriticalThreshold *string `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The ID of the partition filter expression.
	//
	// example:
	//
	// 1234
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 1000
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// Indicates whether the monitoring is performed based on a fixed value.
	//
	// example:
	//
	// true
	FixCheck *bool `json:"FixCheck,omitempty" xml:"FixCheck,omitempty"`
	// The historical threshold for a critical alert.
	//
	// example:
	//
	// history max:80%,history min:50%
	HistoryCriticalThreshold *string `json:"HistoryCriticalThreshold,omitempty" xml:"HistoryCriticalThreshold,omitempty"`
	// The historical threshold for a warning alert.
	//
	// example:
	//
	// history max:40%,history min:10%
	HistoryWarningThreshold *string `json:"HistoryWarningThreshold,omitempty" xml:"HistoryWarningThreshold,omitempty"`
	// The monitoring rule ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The partition filter expression.
	//
	// example:
	//
	// dt=$[yyyymmdd]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The ID of the task that is associated with the partition filter expression.
	//
	// example:
	//
	// 21
	MethodId *int32 `json:"MethodId,omitempty" xml:"MethodId,omitempty"`
	// The method that is used to collect sample data, such as avg, count, sum, min, max, count_distinct, user_defined, table_count, table_size, table_dt_load_count, table_dt_refuseload_count, null_value, null_value/table_count, (table_count-count_distinct)/table_count, or table_count-count_distinct.
	//
	// example:
	//
	// count/table_count
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The name of the Alibaba Cloud account that is used to configure the monitoring rule.
	//
	// example:
	//
	// 1912****
	OnDuty *string `json:"OnDuty,omitempty" xml:"OnDuty,omitempty"`
	// The name of the Alibaba Cloud account that is used to configure the monitoring rule.
	//
	// example:
	//
	// test
	OnDutyAccountName *string `json:"OnDutyAccountName,omitempty" xml:"OnDutyAccountName,omitempty"`
	// The name of the compute engine or data source.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// id
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The field that is used to associate with monitoring rules at the frontend. This parameter can be ignored.
	//
	// example:
	//
	// table_count
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The ID of the task that is associated with the partition filter expression.
	//
	// example:
	//
	// 132323
	RuleCheckerRelationId *int64 `json:"RuleCheckerRelationId,omitempty" xml:"RuleCheckerRelationId,omitempty"`
	// The name of the monitoring rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the monitoring rule. Valid values:
	//
	// 	- 0: The monitoring rule is created by the system.
	//
	// 	- 1: The monitoring rule is created by a user.
	//
	// 	- 2: The monitoring rule is a workspace-level rule.
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the monitoring template.
	//
	// example:
	//
	// 7
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the monitoring template.
	//
	// example:
	//
	// Number of SQL task table rows, 1, 7, and 30 days wave detection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trend of the monitoring result.
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The threshold for a warning alert. The threshold specifies the deviation of the monitoring result from the expected value. You can specify a custom value for the threshold based on your business requirements.
	//
	// example:
	//
	// 10
	WarningThreshold *string `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
}

func (s ListQualityRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponseBodyDataRules) GetBlockType() *int32 {
	return s.BlockType
}

func (s *ListQualityRulesResponseBodyDataRules) GetCheckerId() *int32 {
	return s.CheckerId
}

func (s *ListQualityRulesResponseBodyDataRules) GetComment() *string {
	return s.Comment
}

func (s *ListQualityRulesResponseBodyDataRules) GetCriticalThreshold() *string {
	return s.CriticalThreshold
}

func (s *ListQualityRulesResponseBodyDataRules) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListQualityRulesResponseBodyDataRules) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *ListQualityRulesResponseBodyDataRules) GetFixCheck() *bool {
	return s.FixCheck
}

func (s *ListQualityRulesResponseBodyDataRules) GetHistoryCriticalThreshold() *string {
	return s.HistoryCriticalThreshold
}

func (s *ListQualityRulesResponseBodyDataRules) GetHistoryWarningThreshold() *string {
	return s.HistoryWarningThreshold
}

func (s *ListQualityRulesResponseBodyDataRules) GetId() *int64 {
	return s.Id
}

func (s *ListQualityRulesResponseBodyDataRules) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *ListQualityRulesResponseBodyDataRules) GetMethodId() *int32 {
	return s.MethodId
}

func (s *ListQualityRulesResponseBodyDataRules) GetMethodName() *string {
	return s.MethodName
}

func (s *ListQualityRulesResponseBodyDataRules) GetOnDuty() *string {
	return s.OnDuty
}

func (s *ListQualityRulesResponseBodyDataRules) GetOnDutyAccountName() *string {
	return s.OnDutyAccountName
}

func (s *ListQualityRulesResponseBodyDataRules) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityRulesResponseBodyDataRules) GetProperty() *string {
	return s.Property
}

func (s *ListQualityRulesResponseBodyDataRules) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *ListQualityRulesResponseBodyDataRules) GetRuleCheckerRelationId() *int64 {
	return s.RuleCheckerRelationId
}

func (s *ListQualityRulesResponseBodyDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListQualityRulesResponseBodyDataRules) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListQualityRulesResponseBodyDataRules) GetTableName() *string {
	return s.TableName
}

func (s *ListQualityRulesResponseBodyDataRules) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *ListQualityRulesResponseBodyDataRules) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListQualityRulesResponseBodyDataRules) GetTrend() *string {
	return s.Trend
}

func (s *ListQualityRulesResponseBodyDataRules) GetWarningThreshold() *string {
	return s.WarningThreshold
}

func (s *ListQualityRulesResponseBodyDataRules) SetBlockType(v int32) *ListQualityRulesResponseBodyDataRules {
	s.BlockType = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetCheckerId(v int32) *ListQualityRulesResponseBodyDataRules {
	s.CheckerId = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetComment(v string) *ListQualityRulesResponseBodyDataRules {
	s.Comment = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetCriticalThreshold(v string) *ListQualityRulesResponseBodyDataRules {
	s.CriticalThreshold = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetEntityId(v int64) *ListQualityRulesResponseBodyDataRules {
	s.EntityId = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetExpectValue(v string) *ListQualityRulesResponseBodyDataRules {
	s.ExpectValue = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetFixCheck(v bool) *ListQualityRulesResponseBodyDataRules {
	s.FixCheck = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetHistoryCriticalThreshold(v string) *ListQualityRulesResponseBodyDataRules {
	s.HistoryCriticalThreshold = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetHistoryWarningThreshold(v string) *ListQualityRulesResponseBodyDataRules {
	s.HistoryWarningThreshold = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetId(v int64) *ListQualityRulesResponseBodyDataRules {
	s.Id = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetMatchExpression(v string) *ListQualityRulesResponseBodyDataRules {
	s.MatchExpression = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetMethodId(v int32) *ListQualityRulesResponseBodyDataRules {
	s.MethodId = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetMethodName(v string) *ListQualityRulesResponseBodyDataRules {
	s.MethodName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetOnDuty(v string) *ListQualityRulesResponseBodyDataRules {
	s.OnDuty = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetOnDutyAccountName(v string) *ListQualityRulesResponseBodyDataRules {
	s.OnDutyAccountName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetProjectName(v string) *ListQualityRulesResponseBodyDataRules {
	s.ProjectName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetProperty(v string) *ListQualityRulesResponseBodyDataRules {
	s.Property = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetPropertyKey(v string) *ListQualityRulesResponseBodyDataRules {
	s.PropertyKey = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetRuleCheckerRelationId(v int64) *ListQualityRulesResponseBodyDataRules {
	s.RuleCheckerRelationId = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetRuleName(v string) *ListQualityRulesResponseBodyDataRules {
	s.RuleName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetRuleType(v int32) *ListQualityRulesResponseBodyDataRules {
	s.RuleType = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetTableName(v string) *ListQualityRulesResponseBodyDataRules {
	s.TableName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetTemplateId(v int32) *ListQualityRulesResponseBodyDataRules {
	s.TemplateId = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetTemplateName(v string) *ListQualityRulesResponseBodyDataRules {
	s.TemplateName = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetTrend(v string) *ListQualityRulesResponseBodyDataRules {
	s.Trend = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) SetWarningThreshold(v string) *ListQualityRulesResponseBodyDataRules {
	s.WarningThreshold = &v
	return s
}

func (s *ListQualityRulesResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
