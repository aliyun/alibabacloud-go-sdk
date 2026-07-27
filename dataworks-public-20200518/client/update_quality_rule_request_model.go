// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockType(v int32) *UpdateQualityRuleRequest
	GetBlockType() *int32
	SetChecker(v int32) *UpdateQualityRuleRequest
	GetChecker() *int32
	SetComment(v string) *UpdateQualityRuleRequest
	GetComment() *string
	SetCriticalThreshold(v string) *UpdateQualityRuleRequest
	GetCriticalThreshold() *string
	SetEntityId(v int64) *UpdateQualityRuleRequest
	GetEntityId() *int64
	SetExpectValue(v string) *UpdateQualityRuleRequest
	GetExpectValue() *string
	SetId(v int64) *UpdateQualityRuleRequest
	GetId() *int64
	SetMethodName(v string) *UpdateQualityRuleRequest
	GetMethodName() *string
	SetOpenSwitch(v bool) *UpdateQualityRuleRequest
	GetOpenSwitch() *bool
	SetOperator(v string) *UpdateQualityRuleRequest
	GetOperator() *string
	SetPredictType(v int32) *UpdateQualityRuleRequest
	GetPredictType() *int32
	SetProjectId(v int64) *UpdateQualityRuleRequest
	GetProjectId() *int64
	SetProjectName(v string) *UpdateQualityRuleRequest
	GetProjectName() *string
	SetProperty(v string) *UpdateQualityRuleRequest
	GetProperty() *string
	SetPropertyType(v string) *UpdateQualityRuleRequest
	GetPropertyType() *string
	SetRuleName(v string) *UpdateQualityRuleRequest
	GetRuleName() *string
	SetRuleType(v int32) *UpdateQualityRuleRequest
	GetRuleType() *int32
	SetTaskSetting(v string) *UpdateQualityRuleRequest
	GetTaskSetting() *string
	SetTemplateId(v int32) *UpdateQualityRuleRequest
	GetTemplateId() *int32
	SetTrend(v string) *UpdateQualityRuleRequest
	GetTrend() *string
	SetWarningThreshold(v string) *UpdateQualityRuleRequest
	GetWarningThreshold() *string
	SetWhereCondition(v string) *UpdateQualityRuleRequest
	GetWhereCondition() *string
}

type UpdateQualityRuleRequest struct {
	// The strength of the quality rule. You can specify a rule as a strong or weak rule based on the importance of the rule. Valid values:
	//
	// - 1: strong rule
	//
	// - 0: weak rule
	//
	//   If you specify a rule as a strong rule and a critical alert is triggered for the rule, the scheduling of the associated task is blocked.
	//
	// example:
	//
	// 0
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The checker ID. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the checker ID.
	//
	// example:
	//
	// 9
	Checker *int32 `json:"Checker,omitempty" xml:"Checker,omitempty"`
	// The description of the quality rule.
	//
	// example:
	//
	// Verify the number of table rows
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical alert. The threshold specifies the deviation of a check result from the expected value. You can customize the threshold based on your business requirements. If you use a strong rule and a critical alert is triggered, the scheduling of the associated task is blocked.
	//
	// example:
	//
	// 10
	CriticalThreshold *string `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The ID of the partition filter expression. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the ID of the partition filter expression.
	//
	// example:
	//
	// 123
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 300
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The rule ID. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the method used to collect sample data. Valid values: avg, count, sum, min, max, count_distinct, user_defined, table_count, table_size, table_dt_load_count, table_dt_refuseload_count, null_value, null_value/table_count, (table_count-count_distinct)/table_count, and table_count-count_distinct.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_count
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// Specifies whether to enable or disable the quality rule. This parameter specifies whether to run the quality rule in the production environment.
	//
	// - true: The quality rule is triggered when the scheduling task that is associated with the output table of the rule runs.
	//
	// - false: The quality rule is not triggered when the scheduling task that is associated with the output table of the rule runs.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	OpenSwitch *bool `json:"OpenSwitch,omitempty" xml:"OpenSwitch,omitempty"`
	// The comparison operator. Valid values: >, >=, =, !=, <, and <=.
	//
	// > This parameter is required if you set the Checker parameter to 9.
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Specifies whether to use a dynamic threshold. Valid values:
	//
	// - 0: no
	//
	// - 2: yes
	//
	// example:
	//
	// 0
	PredictType *int32 `json:"PredictType,omitempty" xml:"PredictType,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 26
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the engine or data source. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// id
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The data type of the field.
	//
	// example:
	//
	// bigint
	PropertyType *string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// The name of the quality rule.
	//
	// example:
	//
	// 123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// - 0: system template
	//
	// - 1: custom SQL
	//
	// - 2: custom template
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The variable settings that are inserted before a custom rule. The settings are in the format of x=a,y=b.
	//
	// example:
	//
	// x=a,y=b
	TaskSetting *string `json:"TaskSetting,omitempty" xml:"TaskSetting,omitempty"`
	// The ID of the template that is used for the check. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the template ID.
	//
	// example:
	//
	// 7
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The trend of the check result. Valid values:
	//
	// - up: upward trend
	//
	// - down: downward trend
	//
	// - abs: absolute value
	//
	// example:
	//
	// up
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The threshold for a warning alert. The threshold specifies the deviation of a check result from the expected value. You can customize the threshold based on your business requirements.
	//
	// example:
	//
	// 5
	WarningThreshold *string `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
	// The filter condition or custom SQL statement that is used for the check.
	//
	// example:
	//
	// dt=$[yyyymmdd]
	WhereCondition *string `json:"WhereCondition,omitempty" xml:"WhereCondition,omitempty"`
}

func (s UpdateQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleRequest) GetBlockType() *int32 {
	return s.BlockType
}

func (s *UpdateQualityRuleRequest) GetChecker() *int32 {
	return s.Checker
}

func (s *UpdateQualityRuleRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateQualityRuleRequest) GetCriticalThreshold() *string {
	return s.CriticalThreshold
}

func (s *UpdateQualityRuleRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *UpdateQualityRuleRequest) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *UpdateQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateQualityRuleRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *UpdateQualityRuleRequest) GetOpenSwitch() *bool {
	return s.OpenSwitch
}

func (s *UpdateQualityRuleRequest) GetOperator() *string {
	return s.Operator
}

func (s *UpdateQualityRuleRequest) GetPredictType() *int32 {
	return s.PredictType
}

func (s *UpdateQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateQualityRuleRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateQualityRuleRequest) GetProperty() *string {
	return s.Property
}

func (s *UpdateQualityRuleRequest) GetPropertyType() *string {
	return s.PropertyType
}

func (s *UpdateQualityRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateQualityRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *UpdateQualityRuleRequest) GetTaskSetting() *string {
	return s.TaskSetting
}

func (s *UpdateQualityRuleRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *UpdateQualityRuleRequest) GetTrend() *string {
	return s.Trend
}

func (s *UpdateQualityRuleRequest) GetWarningThreshold() *string {
	return s.WarningThreshold
}

func (s *UpdateQualityRuleRequest) GetWhereCondition() *string {
	return s.WhereCondition
}

func (s *UpdateQualityRuleRequest) SetBlockType(v int32) *UpdateQualityRuleRequest {
	s.BlockType = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetChecker(v int32) *UpdateQualityRuleRequest {
	s.Checker = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetComment(v string) *UpdateQualityRuleRequest {
	s.Comment = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetCriticalThreshold(v string) *UpdateQualityRuleRequest {
	s.CriticalThreshold = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetEntityId(v int64) *UpdateQualityRuleRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetExpectValue(v string) *UpdateQualityRuleRequest {
	s.ExpectValue = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetId(v int64) *UpdateQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetMethodName(v string) *UpdateQualityRuleRequest {
	s.MethodName = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetOpenSwitch(v bool) *UpdateQualityRuleRequest {
	s.OpenSwitch = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetOperator(v string) *UpdateQualityRuleRequest {
	s.Operator = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetPredictType(v int32) *UpdateQualityRuleRequest {
	s.PredictType = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetProjectId(v int64) *UpdateQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetProjectName(v string) *UpdateQualityRuleRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetProperty(v string) *UpdateQualityRuleRequest {
	s.Property = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetPropertyType(v string) *UpdateQualityRuleRequest {
	s.PropertyType = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetRuleName(v string) *UpdateQualityRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetRuleType(v int32) *UpdateQualityRuleRequest {
	s.RuleType = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetTaskSetting(v string) *UpdateQualityRuleRequest {
	s.TaskSetting = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetTemplateId(v int32) *UpdateQualityRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetTrend(v string) *UpdateQualityRuleRequest {
	s.Trend = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetWarningThreshold(v string) *UpdateQualityRuleRequest {
	s.WarningThreshold = &v
	return s
}

func (s *UpdateQualityRuleRequest) SetWhereCondition(v string) *UpdateQualityRuleRequest {
	s.WhereCondition = &v
	return s
}

func (s *UpdateQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
