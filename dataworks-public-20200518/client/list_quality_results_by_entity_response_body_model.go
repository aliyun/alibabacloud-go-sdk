// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQualityResultsByEntityResponseBodyData) *ListQualityResultsByEntityResponseBody
	GetData() *ListQualityResultsByEntityResponseBodyData
	SetErrorCode(v string) *ListQualityResultsByEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListQualityResultsByEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListQualityResultsByEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListQualityResultsByEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityResultsByEntityResponseBody
	GetSuccess() *bool
}

type ListQualityResultsByEntityResponseBody struct {
	// The data structure of the check results.
	Data *ListQualityResultsByEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 401
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
	// The request ID.
	//
	// example:
	//
	// CBA58543-00D4-41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityResultsByEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponseBody) GetData() *ListQualityResultsByEntityResponseBodyData {
	return s.Data
}

func (s *ListQualityResultsByEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListQualityResultsByEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListQualityResultsByEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityResultsByEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityResultsByEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityResultsByEntityResponseBody) SetData(v *ListQualityResultsByEntityResponseBodyData) *ListQualityResultsByEntityResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) SetErrorCode(v string) *ListQualityResultsByEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) SetErrorMessage(v string) *ListQualityResultsByEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) SetHttpStatusCode(v int32) *ListQualityResultsByEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) SetRequestId(v string) *ListQualityResultsByEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) SetSuccess(v bool) *ListQualityResultsByEntityResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQualityResultsByEntityResponseBodyData struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned check results.
	RuleChecks []*ListQualityResultsByEntityResponseBodyDataRuleChecks `json:"RuleChecks,omitempty" xml:"RuleChecks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityResultsByEntityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityResultsByEntityResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityResultsByEntityResponseBodyData) GetRuleChecks() []*ListQualityResultsByEntityResponseBodyDataRuleChecks {
	return s.RuleChecks
}

func (s *ListQualityResultsByEntityResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityResultsByEntityResponseBodyData) SetPageNumber(v int32) *ListQualityResultsByEntityResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyData) SetPageSize(v int32) *ListQualityResultsByEntityResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyData) SetRuleChecks(v []*ListQualityResultsByEntityResponseBodyDataRuleChecks) *ListQualityResultsByEntityResponseBodyData {
	s.RuleChecks = v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyData) SetTotalCount(v int64) *ListQualityResultsByEntityResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListQualityResultsByEntityResponseBodyDataRuleChecks struct {
	// The monitored partition in the data source table.
	//
	// example:
	//
	// ds=20200912
	ActualExpression *string `json:"ActualExpression,omitempty" xml:"ActualExpression,omitempty"`
	// The time when the monitoring started.
	//
	// example:
	//
	// 1600704000000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The data timestamp. In most cases, if the monitored business entity is offline data, the value is one day before the monitoring is performed.
	//
	// example:
	//
	// 1600704000000
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The strength of the monitoring rule. The strength of a monitoring rule indicates the importance of the rule. Valid values:
	//
	// 	- 1: the monitoring rule is a strong rule.
	//
	// 	- 0: the monitoring rule is a weak rule. You can specify whether a monitoring rule is a strong rule based on your business requirements. If a monitoring rule is a strong rule and the critical threshold is exceeded, a critical alert is reported and tasks that are associated with the rule are blocked from running.
	//
	// example:
	//
	// 0
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The check result. The value of this parameter is the same as the value of the CheckResultStatus parameter. Valid values:
	//
	// 	- 0: indicates that the data source table is normal.
	//
	// 	- 1: indicates that a warning alert is reported.
	//
	// 	- 2: indicates that a critical alert is reported.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The check result of the monitoring rule. Valid values:
	//
	// 	- 0: indicates that the data source table is normal.
	//
	// 	- 1: indicates that a warning alert is reported.
	//
	// 	- 2: indicates that a critical alert is reported.
	//
	// example:
	//
	// 0
	CheckResultStatus *int32 `json:"CheckResultStatus,omitempty" xml:"CheckResultStatus,omitempty"`
	// The checker ID.
	//
	// example:
	//
	// 7
	CheckerId *int32 `json:"CheckerId,omitempty" xml:"CheckerId,omitempty"`
	// The name of the checker.
	//
	// example:
	//
	// fulx
	CheckerName *string `json:"CheckerName,omitempty" xml:"CheckerName,omitempty"`
	// The check type. Valid values:
	//
	// 	- 0: indicates that the monitoring is performed based on a fixed value.
	//
	// 	- 1: indicates that the monitoring is performed based on a non-fixed value.
	//
	// 	- 2: indicates that the monitoring is performed based on a dynamic threshold.
	//
	// example:
	//
	// 1
	CheckerType *int32 `json:"CheckerType,omitempty" xml:"CheckerType,omitempty"`
	// The description of the monitoring rule.
	//
	// example:
	//
	// The description of the rule.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical alert. The threshold indicates the deviation of the check result from the expected value. You can specify a value for the threshold based on your business requirements. If a monitoring rule is a strong rule and the critical threshold is exceeded, a critical alert is reported and tasks that are associated with the rule are blocked from running.
	//
	// example:
	//
	// 0.5
	CriticalThreshold *float32 `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The scheduling frequency. In most cases, the value of this parameter is YMD. This value indicates year, month, and day.
	//
	// example:
	//
	// YMD
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// Indicates whether the monitoring is discrete monitoring. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DiscreteCheck *bool `json:"DiscreteCheck,omitempty" xml:"DiscreteCheck,omitempty"`
	// The deadline for querying the check result.
	//
	// example:
	//
	// 1600704000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the partition filter expression.
	//
	// example:
	//
	// 15432322
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 200
	ExpectValue *float32 `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 1112323123
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The type of the scheduling system. Only CWF scheduling systems are supported.
	//
	// example:
	//
	// CWF2
	ExternalType *string `json:"ExternalType,omitempty" xml:"ExternalType,omitempty"`
	// Indicates whether the monitoring is performed based on a fixed value. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	FixedCheck *bool `json:"FixedCheck,omitempty" xml:"FixedCheck,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 121212121
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the check result is the same as the predicted result. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsPrediction *bool `json:"IsPrediction,omitempty" xml:"IsPrediction,omitempty"`
	// The lower limit of the predicted result. The value of this parameter is automatically generated based on the threshold that you specify.
	//
	// example:
	//
	// 2344
	LowerValue *float32 `json:"LowerValue,omitempty" xml:"LowerValue,omitempty"`
	// The partition filter expression.
	//
	// example:
	//
	// ds=$[yyyymmdd]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The method used to collect sample data, such as such as avg, count, sum, min, max, count_distinct, user_defined, table_count, table_size, table_dt_load_count, table_dt_refuseload_count, null_value, null_value/table_count, (table_count-count_distinct)/table_count, or table_count-count_distinct.
	//
	// example:
	//
	// count_distinct
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// >
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// The name of the compute engine or data source for which data quality is monitored.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The field whose data quality is checked based on the monitoring rule. This field is a column in the data source table that is monitored.
	//
	// example:
	//
	// type
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The historical sample values.
	ReferenceValue []*ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue `json:"ReferenceValue,omitempty" xml:"ReferenceValue,omitempty" type:"Repeated"`
	// The string of the check result.
	//
	// example:
	//
	// test
	ResultString *string `json:"ResultString,omitempty" xml:"ResultString,omitempty"`
	// The ID of the monitoring rule.
	//
	// example:
	//
	// 123123232
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the monitoring rule.
	//
	// example:
	//
	// The name of the rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The current sample value.
	SampleValue []*ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue `json:"SampleValue,omitempty" xml:"SampleValue,omitempty" type:"Repeated"`
	// The name of the table that is monitored.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The monitoring task ID.
	//
	// example:
	//
	// 16008552981681a0d6****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the monitoring template.
	//
	// example:
	//
	// 5
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the monitoring template.
	//
	// example:
	//
	// Expected value verification
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time that was taken to run the monitoring task. Unit: seconds.
	//
	// example:
	//
	// 202
	TimeCost *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
	// The trend of the check result.
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The upper limit of the predicted result. The value of this parameter is automatically generated based on the threshold that you specify.
	//
	// example:
	//
	// 25555
	UpperValue *float32 `json:"UpperValue,omitempty" xml:"UpperValue,omitempty"`
	// The threshold for a warning alert. The threshold indicates the deviation of the check result from the expected value. You can customize this threshold based on your business requirements.
	//
	// example:
	//
	// 0.1
	WarningThreshold *float32 `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
	// The filter condition of the monitoring rule.
	//
	// example:
	//
	// id>0
	WhereCondition *string `json:"WhereCondition,omitempty" xml:"WhereCondition,omitempty"`
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecks) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecks) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetActualExpression() *string {
	return s.ActualExpression
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetBlockType() *int32 {
	return s.BlockType
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCheckResultStatus() *int32 {
	return s.CheckResultStatus
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCheckerId() *int32 {
	return s.CheckerId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCheckerName() *string {
	return s.CheckerName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCheckerType() *int32 {
	return s.CheckerType
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetComment() *string {
	return s.Comment
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetCriticalThreshold() *float32 {
	return s.CriticalThreshold
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetDateType() *string {
	return s.DateType
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetDiscreteCheck() *bool {
	return s.DiscreteCheck
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetExpectValue() *float32 {
	return s.ExpectValue
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetExternalId() *string {
	return s.ExternalId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetExternalType() *string {
	return s.ExternalType
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetFixedCheck() *bool {
	return s.FixedCheck
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetId() *int64 {
	return s.Id
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetIsPrediction() *bool {
	return s.IsPrediction
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetLowerValue() *float32 {
	return s.LowerValue
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetMethodName() *string {
	return s.MethodName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetOp() *string {
	return s.Op
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetProperty() *string {
	return s.Property
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetReferenceValue() []*ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	return s.ReferenceValue
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetResultString() *string {
	return s.ResultString
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetRuleName() *string {
	return s.RuleName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetSampleValue() []*ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue {
	return s.SampleValue
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTableName() *string {
	return s.TableName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTimeCost() *string {
	return s.TimeCost
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetTrend() *string {
	return s.Trend
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetUpperValue() *float32 {
	return s.UpperValue
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetWarningThreshold() *float32 {
	return s.WarningThreshold
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) GetWhereCondition() *string {
	return s.WhereCondition
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetActualExpression(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ActualExpression = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetBeginTime(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.BeginTime = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetBizDate(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetBlockType(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.BlockType = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCheckResult(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CheckResult = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCheckResultStatus(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CheckResultStatus = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCheckerId(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CheckerId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCheckerName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CheckerName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCheckerType(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CheckerType = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetComment(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.Comment = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetCriticalThreshold(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.CriticalThreshold = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetDateType(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.DateType = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetDiscreteCheck(v bool) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.DiscreteCheck = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetEndTime(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.EndTime = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetEntityId(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.EntityId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetExpectValue(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ExpectValue = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetExternalId(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ExternalId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetExternalType(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ExternalType = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetFixedCheck(v bool) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.FixedCheck = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetId(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.Id = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetIsPrediction(v bool) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.IsPrediction = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetLowerValue(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.LowerValue = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetMatchExpression(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.MatchExpression = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetMethodName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.MethodName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetOp(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.Op = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetProjectName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ProjectName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetProperty(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.Property = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetReferenceValue(v []*ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ReferenceValue = v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetResultString(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.ResultString = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetRuleId(v int64) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.RuleId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetRuleName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.RuleName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetSampleValue(v []*ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.SampleValue = v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTableName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.TableName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTaskId(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.TaskId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTemplateId(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.TemplateId = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTemplateName(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.TemplateName = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTimeCost(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.TimeCost = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetTrend(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.Trend = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetUpperValue(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.UpperValue = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetWarningThreshold(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.WarningThreshold = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) SetWhereCondition(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecks {
	s.WhereCondition = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecks) Validate() error {
	return dara.Validate(s)
}

type ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue struct {
	// The data timestamp. In most cases, if the monitored business entity is offline data, the value is one day before the monitoring is performed.
	//
	// example:
	//
	// 2020-12-03
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The values of the sample field that are grouped by using the GROUP BY clause. For example, the values of the Gender field are grouped by using the GROUP BY clause. In this case, the values of the DiscreteProperty parameter are Male, Female, and null.
	//
	// example:
	//
	// 0
	DiscreteProperty *string `json:"DiscreteProperty,omitempty" xml:"DiscreteProperty,omitempty"`
	// The check result.
	//
	// example:
	//
	// 0
	SingleCheckResult *int32 `json:"SingleCheckResult,omitempty" xml:"SingleCheckResult,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 0.5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The check value.
	//
	// example:
	//
	// 19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GetDiscreteProperty() *string {
	return s.DiscreteProperty
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GetSingleCheckResult() *int32 {
	return s.SingleCheckResult
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) GetValue() *float32 {
	return s.Value
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) SetBizDate(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) SetDiscreteProperty(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	s.DiscreteProperty = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) SetSingleCheckResult(v int32) *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	s.SingleCheckResult = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) SetThreshold(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	s.Threshold = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) SetValue(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue {
	s.Value = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksReferenceValue) Validate() error {
	return dara.Validate(s)
}

type ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue struct {
	// The data timestamp. In most cases, if the monitored business entity is offline data, the value is one day before the monitoring is performed.
	//
	// example:
	//
	// 2020-12-03
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The values of the sample field that are grouped by using the GROUP BY clause. For example, the values of the Gender field are grouped by using the GROUP BY clause. In this case, the values of the DiscreteProperty parameter are Male, Female, and null.
	//
	// example:
	//
	// 0
	DiscreteProperty *string `json:"DiscreteProperty,omitempty" xml:"DiscreteProperty,omitempty"`
	// The current sample value.
	//
	// example:
	//
	// 19
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) GetDiscreteProperty() *string {
	return s.DiscreteProperty
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) GetValue() *float32 {
	return s.Value
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) SetBizDate(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) SetDiscreteProperty(v string) *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue {
	s.DiscreteProperty = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) SetValue(v float32) *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue {
	s.Value = &v
	return s
}

func (s *ListQualityResultsByEntityResponseBodyDataRuleChecksSampleValue) Validate() error {
	return dara.Validate(s)
}
