// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetQualityRuleResponseBodyData) *GetQualityRuleResponseBody
	GetData() *GetQualityRuleResponseBodyData
	SetErrorCode(v string) *GetQualityRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetQualityRuleResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleResponseBody
	GetSuccess() *bool
}

type GetQualityRuleResponseBody struct {
	// The quality rule details.
	Data *GetQualityRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 576b9457-2cf5-4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBody) GetData() *GetQualityRuleResponseBodyData {
	return s.Data
}

func (s *GetQualityRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQualityRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleResponseBody) SetData(v *GetQualityRuleResponseBodyData) *GetQualityRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityRuleResponseBody) SetErrorCode(v string) *GetQualityRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetErrorMessage(v string) *GetQualityRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetHttpStatusCode(v int32) *GetQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetRequestId(v string) *GetQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleResponseBody) SetSuccess(v bool) *GetQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleResponseBodyData struct {
	// The severity of the quality rule. Valid values:
	//
	// - `1`: Strong Rule. If a Strong Rule triggers a critical Alert, the associated Scheduling Task is blocked.
	//
	// - `0`: Weak Rule.
	//
	// example:
	//
	// 1
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The checker ID, which is derived from a primary key and used for front-end identification.
	//
	// example:
	//
	// 9
	Checker *int32 `json:"Checker,omitempty" xml:"Checker,omitempty"`
	// The name of the checker.
	//
	// example:
	//
	// compared with a fixed value
	CheckerName *string `json:"CheckerName,omitempty" xml:"CheckerName,omitempty"`
	// The comment for the rule.
	//
	// example:
	//
	// Verify that the primary key is unique
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical Alert, which defines the maximum acceptable deviation from the Expected Value. If this threshold is met for a Strong Rule, the associated Scheduling Task is blocked.
	//
	// example:
	//
	// 20
	CriticalThreshold *string `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The ID of the partition expression.
	//
	// example:
	//
	// 165523
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 30
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// Indicates whether to validate against a fixed value.
	//
	// example:
	//
	// true
	FixCheck *bool `json:"FixCheck,omitempty" xml:"FixCheck,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 123232
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the method used to collect sample data.
	//
	// example:
	//
	// 8
	MethodId *int32 `json:"MethodId,omitempty" xml:"MethodId,omitempty"`
	// The name of the method used to collect sample data. Valid values include `avg`, `count`, `sum`, `min`, `max`, `count_distinct`, `user_defined`, `table_count`, `table_size`, `table_dt_load_count`, `table_dt_refuseload_count`, `null_value`, `null_value/table_count`, `(table_count-count_distinct)/table_count`, and `table_count-count_distinct`.
	//
	// example:
	//
	// table_count
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The ID of the Alibaba Cloud account that configured the rule.
	//
	// example:
	//
	// 1822931****
	OnDuty *string `json:"OnDuty,omitempty" xml:"OnDuty,omitempty"`
	// The name of the Alibaba Cloud account that configured the rule.
	//
	// example:
	//
	// test
	OnDutyAccountName *string `json:"OnDutyAccountName,omitempty" xml:"OnDutyAccountName,omitempty"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	OpenSwitch *bool `json:"OpenSwitch,omitempty" xml:"OpenSwitch,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// >=
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Indicates whether the threshold is dynamic. Valid values:
	//
	// - `0`: The threshold is static.
	//
	// - `1`: The threshold is dynamic.
	//
	// example:
	//
	// 0
	PredictType *int32 `json:"PredictType,omitempty" xml:"PredictType,omitempty"`
	// The column in the data source table to which the rule applies.
	//
	// example:
	//
	// id
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// View table fluctuations
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the quality rule. Valid values:
	//
	// - `0`: A rule based on a system template.
	//
	// - `1`: A rule based on a custom SQL statement.
	//
	// - `4`: A rule based on a custom template.
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The variable settings inserted before a custom rule. The format is x=a,y=b.
	//
	// example:
	//
	// x=a,y=b
	TaskSetting *string `json:"TaskSetting,omitempty" xml:"TaskSetting,omitempty"`
	// The ID of the template used.
	//
	// example:
	//
	// 7
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template used.
	//
	// example:
	//
	// SQL task table rows, 1,7, 30 days fluctuation test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trend of the validation result.
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The threshold for a warning Alert. This value specifies the acceptable deviation from the Expected Value and can be customized.
	//
	// example:
	//
	// 10
	WarningThreshold *string `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
	// The filter condition or custom SQL statement used for validation.
	//
	// example:
	//
	// id>10
	WhereCondition *string `json:"WhereCondition,omitempty" xml:"WhereCondition,omitempty"`
}

func (s GetQualityRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponseBodyData) GetBlockType() *int32 {
	return s.BlockType
}

func (s *GetQualityRuleResponseBodyData) GetChecker() *int32 {
	return s.Checker
}

func (s *GetQualityRuleResponseBodyData) GetCheckerName() *string {
	return s.CheckerName
}

func (s *GetQualityRuleResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetQualityRuleResponseBodyData) GetCriticalThreshold() *string {
	return s.CriticalThreshold
}

func (s *GetQualityRuleResponseBodyData) GetEntityId() *int64 {
	return s.EntityId
}

func (s *GetQualityRuleResponseBodyData) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *GetQualityRuleResponseBodyData) GetFixCheck() *bool {
	return s.FixCheck
}

func (s *GetQualityRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetQualityRuleResponseBodyData) GetMethodId() *int32 {
	return s.MethodId
}

func (s *GetQualityRuleResponseBodyData) GetMethodName() *string {
	return s.MethodName
}

func (s *GetQualityRuleResponseBodyData) GetOnDuty() *string {
	return s.OnDuty
}

func (s *GetQualityRuleResponseBodyData) GetOnDutyAccountName() *string {
	return s.OnDutyAccountName
}

func (s *GetQualityRuleResponseBodyData) GetOpenSwitch() *bool {
	return s.OpenSwitch
}

func (s *GetQualityRuleResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetQualityRuleResponseBodyData) GetPredictType() *int32 {
	return s.PredictType
}

func (s *GetQualityRuleResponseBodyData) GetProperty() *string {
	return s.Property
}

func (s *GetQualityRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetQualityRuleResponseBodyData) GetRuleType() *int32 {
	return s.RuleType
}

func (s *GetQualityRuleResponseBodyData) GetTaskSetting() *string {
	return s.TaskSetting
}

func (s *GetQualityRuleResponseBodyData) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *GetQualityRuleResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetQualityRuleResponseBodyData) GetTrend() *string {
	return s.Trend
}

func (s *GetQualityRuleResponseBodyData) GetWarningThreshold() *string {
	return s.WarningThreshold
}

func (s *GetQualityRuleResponseBodyData) GetWhereCondition() *string {
	return s.WhereCondition
}

func (s *GetQualityRuleResponseBodyData) SetBlockType(v int32) *GetQualityRuleResponseBodyData {
	s.BlockType = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetChecker(v int32) *GetQualityRuleResponseBodyData {
	s.Checker = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetCheckerName(v string) *GetQualityRuleResponseBodyData {
	s.CheckerName = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetComment(v string) *GetQualityRuleResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetCriticalThreshold(v string) *GetQualityRuleResponseBodyData {
	s.CriticalThreshold = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetEntityId(v int64) *GetQualityRuleResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetExpectValue(v string) *GetQualityRuleResponseBodyData {
	s.ExpectValue = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetFixCheck(v bool) *GetQualityRuleResponseBodyData {
	s.FixCheck = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetId(v int64) *GetQualityRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetMethodId(v int32) *GetQualityRuleResponseBodyData {
	s.MethodId = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetMethodName(v string) *GetQualityRuleResponseBodyData {
	s.MethodName = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetOnDuty(v string) *GetQualityRuleResponseBodyData {
	s.OnDuty = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetOnDutyAccountName(v string) *GetQualityRuleResponseBodyData {
	s.OnDutyAccountName = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetOpenSwitch(v bool) *GetQualityRuleResponseBodyData {
	s.OpenSwitch = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetOperator(v string) *GetQualityRuleResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetPredictType(v int32) *GetQualityRuleResponseBodyData {
	s.PredictType = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetProperty(v string) *GetQualityRuleResponseBodyData {
	s.Property = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetRuleName(v string) *GetQualityRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetRuleType(v int32) *GetQualityRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetTaskSetting(v string) *GetQualityRuleResponseBodyData {
	s.TaskSetting = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetTemplateId(v int32) *GetQualityRuleResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetTemplateName(v string) *GetQualityRuleResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetTrend(v string) *GetQualityRuleResponseBodyData {
	s.Trend = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetWarningThreshold(v string) *GetQualityRuleResponseBodyData {
	s.WarningThreshold = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) SetWhereCondition(v string) *GetQualityRuleResponseBodyData {
	s.WhereCondition = &v
	return s
}

func (s *GetQualityRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
