// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQualityResultsByRuleResponseBodyData) *ListQualityResultsByRuleResponseBody
	GetData() *ListQualityResultsByRuleResponseBodyData
	SetErrorCode(v string) *ListQualityResultsByRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListQualityResultsByRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListQualityResultsByRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListQualityResultsByRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityResultsByRuleResponseBody
	GetSuccess() *bool
}

type ListQualityResultsByRuleResponseBody struct {
	// The top-level object of the check result.
	Data *ListQualityResultsByRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListQualityResultsByRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponseBody) GetData() *ListQualityResultsByRuleResponseBodyData {
	return s.Data
}

func (s *ListQualityResultsByRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListQualityResultsByRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListQualityResultsByRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityResultsByRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityResultsByRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityResultsByRuleResponseBody) SetData(v *ListQualityResultsByRuleResponseBodyData) *ListQualityResultsByRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) SetErrorCode(v string) *ListQualityResultsByRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) SetErrorMessage(v string) *ListQualityResultsByRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) SetHttpStatusCode(v int32) *ListQualityResultsByRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) SetRequestId(v string) *ListQualityResultsByRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) SetSuccess(v bool) *ListQualityResultsByRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityResultsByRuleResponseBodyData struct {
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
	// A list of check results.
	RuleChecks []*ListQualityResultsByRuleResponseBodyDataRuleChecks `json:"RuleChecks,omitempty" xml:"RuleChecks,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 200
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityResultsByRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityResultsByRuleResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityResultsByRuleResponseBodyData) GetRuleChecks() []*ListQualityResultsByRuleResponseBodyDataRuleChecks {
	return s.RuleChecks
}

func (s *ListQualityResultsByRuleResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityResultsByRuleResponseBodyData) SetPageNumber(v int32) *ListQualityResultsByRuleResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyData) SetPageSize(v int32) *ListQualityResultsByRuleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyData) SetRuleChecks(v []*ListQualityResultsByRuleResponseBodyDataRuleChecks) *ListQualityResultsByRuleResponseBodyData {
	s.RuleChecks = v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyData) SetTotalCount(v int64) *ListQualityResultsByRuleResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyData) Validate() error {
	if s.RuleChecks != nil {
		for _, item := range s.RuleChecks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityResultsByRuleResponseBodyDataRuleChecks struct {
	// The actual data partition that was checked.
	//
	// example:
	//
	// ds=20200925
	ActualExpression *string `json:"ActualExpression,omitempty" xml:"ActualExpression,omitempty"`
	// The start time of the check.
	//
	// example:
	//
	// 1600704000000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The business date. For an offline table, this is typically the day before the check is performed.
	//
	// example:
	//
	// 1600704000000
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The strength of the monitoring rule. A strong rule can block a downstream scheduling task if a critical alert is triggered. Valid values:
	//
	// - `1`: Strong rule.
	//
	// - `0`: Weak rule.
	//
	// example:
	//
	// 1
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The check result.
	//
	// example:
	//
	// 2
	CheckResult *int32 `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The status of the check result.
	//
	// example:
	//
	// 2
	CheckResultStatus *int32 `json:"CheckResultStatus,omitempty" xml:"CheckResultStatus,omitempty"`
	// The ID of the checker.
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
	// The type of the checker.
	//
	// example:
	//
	// 0
	CheckerType *int32 `json:"CheckerType,omitempty" xml:"CheckerType,omitempty"`
	// The description of the monitoring rule.
	//
	// example:
	//
	// The description of the rule.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical alert. For a strong rule, exceeding this threshold blocks the downstream scheduling task.
	//
	// example:
	//
	// 0.6
	CriticalThreshold *float32 `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The scheduling cycle. For example, `YMD` can represent yearly, monthly, and daily tasks.
	//
	// example:
	//
	// YMD
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// Indicates whether a discrete value check is used. Valid values:
	//
	// - `true`: A discrete value check is used.
	//
	// - `false`: A discrete value check is not used.
	//
	// example:
	//
	// true
	DiscreteCheck *bool `json:"DiscreteCheck,omitempty" xml:"DiscreteCheck,omitempty"`
	// The end time of the check.
	//
	// example:
	//
	// 1600704000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the partition filter expression.
	//
	// example:
	//
	// 14534343
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 90
	ExpectValue *float32 `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The node ID of the scheduling task.
	//
	// example:
	//
	// 123112232
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The type of the scheduling system. Currently, only CWF is supported.
	//
	// example:
	//
	// CWF2
	ExternalType *string `json:"ExternalType,omitempty" xml:"ExternalType,omitempty"`
	// Indicates whether a fixed-value check is used. Valid values:
	//
	// - `true`: A fixed-value check is used.
	//
	// - `false`: A fixed-value check is not used.
	//
	// example:
	//
	// false
	FixedCheck *bool `json:"FixedCheck,omitempty" xml:"FixedCheck,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 2231123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the result is a prediction. Valid values:
	//
	// - `true`: The result is a prediction.
	//
	// - `false`: The result is not a prediction.
	//
	// example:
	//
	// false
	IsPrediction *bool `json:"IsPrediction,omitempty" xml:"IsPrediction,omitempty"`
	// The predicted lower limit, which is automatically generated based on the configured threshold.
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
	// The method used to collect sample data. Examples: `avg`, `count`, `sum`, `min`, `max`, `count_distinct`, `user_defined`, `table_count`, `table_size`, `table_dt_load_count`, `table_dt_refuseload_count`, `null_value`, `null_value/table_count`, `(table_count-count_distinct)/table_count`, and `table_count-count_distinct`.
	//
	// example:
	//
	// max
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// >
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The name of the engine or data source used for the quality check.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The column in the source data table that the rule checks.
	//
	// example:
	//
	// type
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The historical sample values.
	ReferenceValue []*ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue `json:"ReferenceValue,omitempty" xml:"ReferenceValue,omitempty" type:"Repeated"`
	// The string representation of the check result.
	//
	// example:
	//
	// ResultString
	ResultString *string `json:"ResultString,omitempty" xml:"ResultString,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 123421
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// The name of the rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The current sample values.
	SampleValue []*ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue `json:"SampleValue,omitempty" xml:"SampleValue,omitempty" type:"Repeated"`
	// The name of the table being checked.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the check task.
	//
	// example:
	//
	// 16008552981681a0d6****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the rule template.
	//
	// example:
	//
	// 5
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the rule template.
	//
	// example:
	//
	// Expected value verification
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The duration of the check.
	//
	// example:
	//
	// 10
	TimeCost *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
	// The trend of the check result.
	//
	// example:
	//
	// up
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The predicted upper limit, which is automatically generated based on the configured threshold.
	//
	// example:
	//
	// 22200
	UpperValue *float32 `json:"UpperValue,omitempty" xml:"UpperValue,omitempty"`
	// The custom threshold for a warning alert. An alert is triggered if the deviation from the expected value exceeds this threshold.
	//
	// example:
	//
	// 0.1
	WarningThreshold *float32 `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
	// The filter condition for the check.
	//
	// example:
	//
	// type!=\\"type2\\"
	WhereCondition *string `json:"WhereCondition,omitempty" xml:"WhereCondition,omitempty"`
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecks) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecks) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetActualExpression() *string {
	return s.ActualExpression
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetBlockType() *int32 {
	return s.BlockType
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCheckResultStatus() *int32 {
	return s.CheckResultStatus
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCheckerId() *int32 {
	return s.CheckerId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCheckerName() *string {
	return s.CheckerName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCheckerType() *int32 {
	return s.CheckerType
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetComment() *string {
	return s.Comment
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetCriticalThreshold() *float32 {
	return s.CriticalThreshold
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetDateType() *string {
	return s.DateType
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetDiscreteCheck() *bool {
	return s.DiscreteCheck
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetExpectValue() *float32 {
	return s.ExpectValue
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetExternalId() *string {
	return s.ExternalId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetExternalType() *string {
	return s.ExternalType
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetFixedCheck() *bool {
	return s.FixedCheck
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetId() *int64 {
	return s.Id
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetIsPrediction() *bool {
	return s.IsPrediction
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetLowerValue() *float32 {
	return s.LowerValue
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetMethodName() *string {
	return s.MethodName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetOp() *string {
	return s.Op
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetOpen() *bool {
	return s.Open
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetProperty() *string {
	return s.Property
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetReferenceValue() []*ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	return s.ReferenceValue
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetResultString() *string {
	return s.ResultString
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetRuleName() *string {
	return s.RuleName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetSampleValue() []*ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue {
	return s.SampleValue
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTableName() *string {
	return s.TableName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTimeCost() *string {
	return s.TimeCost
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetTrend() *string {
	return s.Trend
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetUpperValue() *float32 {
	return s.UpperValue
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetWarningThreshold() *float32 {
	return s.WarningThreshold
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) GetWhereCondition() *string {
	return s.WhereCondition
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetActualExpression(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ActualExpression = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetBeginTime(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.BeginTime = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetBizDate(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetBlockType(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.BlockType = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCheckResult(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CheckResult = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCheckResultStatus(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CheckResultStatus = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCheckerId(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CheckerId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCheckerName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CheckerName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCheckerType(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CheckerType = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetComment(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Comment = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetCriticalThreshold(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.CriticalThreshold = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetDateType(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.DateType = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetDiscreteCheck(v bool) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.DiscreteCheck = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetEndTime(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.EndTime = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetEntityId(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.EntityId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetExpectValue(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ExpectValue = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetExternalId(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ExternalId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetExternalType(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ExternalType = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetFixedCheck(v bool) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.FixedCheck = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetId(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Id = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetIsPrediction(v bool) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.IsPrediction = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetLowerValue(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.LowerValue = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetMatchExpression(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.MatchExpression = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetMethodName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.MethodName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetOp(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Op = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetOpen(v bool) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Open = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetProjectName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ProjectName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetProperty(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Property = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetReferenceValue(v []*ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ReferenceValue = v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetResultString(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.ResultString = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetRuleId(v int64) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.RuleId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetRuleName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.RuleName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetSampleValue(v []*ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.SampleValue = v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTableName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.TableName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTaskId(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.TaskId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTemplateId(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.TemplateId = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTemplateName(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.TemplateName = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTimeCost(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.TimeCost = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetTrend(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.Trend = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetUpperValue(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.UpperValue = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetWarningThreshold(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.WarningThreshold = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) SetWhereCondition(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecks {
	s.WhereCondition = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecks) Validate() error {
	if s.ReferenceValue != nil {
		for _, item := range s.ReferenceValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SampleValue != nil {
		for _, item := range s.SampleValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue struct {
	// The business date. For an offline table, this is typically the day before the check is performed.
	//
	// example:
	//
	// 1600704000000
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The value of the sample column after being grouped by the `GROUP BY` clause. For example, if you group by a gender column, the values of this parameter can be \\"male\\", \\"female\\", or \\"null\\".
	//
	// example:
	//
	// type1
	DiscreteProperty *string `json:"DiscreteProperty,omitempty" xml:"DiscreteProperty,omitempty"`
	// The result of a single check.
	//
	// example:
	//
	// 2
	SingleCheckResult *int32 `json:"SingleCheckResult,omitempty" xml:"SingleCheckResult,omitempty"`
	// The threshold.
	//
	// example:
	//
	// 0.2
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The sample value.
	//
	// example:
	//
	// 20
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GetDiscreteProperty() *string {
	return s.DiscreteProperty
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GetSingleCheckResult() *int32 {
	return s.SingleCheckResult
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) GetValue() *float32 {
	return s.Value
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) SetBizDate(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) SetDiscreteProperty(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	s.DiscreteProperty = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) SetSingleCheckResult(v int32) *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	s.SingleCheckResult = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) SetThreshold(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	s.Threshold = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) SetValue(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue {
	s.Value = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksReferenceValue) Validate() error {
	return dara.Validate(s)
}

type ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue struct {
	// The business date. For an offline table, this is typically the day before the check is performed.
	//
	// example:
	//
	// 1600704000000
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The value of the sample column after being grouped by the `GROUP BY` clause. For example, if you group by a gender column, the values of this parameter can be \\"male\\", \\"female\\", or \\"null\\".
	//
	// example:
	//
	// type2
	DiscreteProperty *string `json:"DiscreteProperty,omitempty" xml:"DiscreteProperty,omitempty"`
	// The sample value.
	//
	// example:
	//
	// 23
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) GetDiscreteProperty() *string {
	return s.DiscreteProperty
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) GetValue() *float32 {
	return s.Value
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) SetBizDate(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue {
	s.BizDate = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) SetDiscreteProperty(v string) *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue {
	s.DiscreteProperty = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) SetValue(v float32) *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue {
	s.Value = &v
	return s
}

func (s *ListQualityResultsByRuleResponseBodyDataRuleChecksSampleValue) Validate() error {
	return dara.Validate(s)
}
