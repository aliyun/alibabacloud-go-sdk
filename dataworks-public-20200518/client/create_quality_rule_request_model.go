// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockType(v int32) *CreateQualityRuleRequest
	GetBlockType() *int32
	SetChecker(v int32) *CreateQualityRuleRequest
	GetChecker() *int32
	SetComment(v string) *CreateQualityRuleRequest
	GetComment() *string
	SetCriticalThreshold(v string) *CreateQualityRuleRequest
	GetCriticalThreshold() *string
	SetEntityId(v int64) *CreateQualityRuleRequest
	GetEntityId() *int64
	SetExpectValue(v string) *CreateQualityRuleRequest
	GetExpectValue() *string
	SetMethodName(v string) *CreateQualityRuleRequest
	GetMethodName() *string
	SetOperator(v string) *CreateQualityRuleRequest
	GetOperator() *string
	SetPredictType(v int32) *CreateQualityRuleRequest
	GetPredictType() *int32
	SetProjectId(v int64) *CreateQualityRuleRequest
	GetProjectId() *int64
	SetProjectName(v string) *CreateQualityRuleRequest
	GetProjectName() *string
	SetProperty(v string) *CreateQualityRuleRequest
	GetProperty() *string
	SetPropertyType(v string) *CreateQualityRuleRequest
	GetPropertyType() *string
	SetRuleName(v string) *CreateQualityRuleRequest
	GetRuleName() *string
	SetRuleType(v int32) *CreateQualityRuleRequest
	GetRuleType() *int32
	SetTaskSetting(v string) *CreateQualityRuleRequest
	GetTaskSetting() *string
	SetTemplateId(v int32) *CreateQualityRuleRequest
	GetTemplateId() *int32
	SetTrend(v string) *CreateQualityRuleRequest
	GetTrend() *string
	SetWarningThreshold(v string) *CreateQualityRuleRequest
	GetWarningThreshold() *string
	SetWhereCondition(v string) *CreateQualityRuleRequest
	GetWhereCondition() *string
}

type CreateQualityRuleRequest struct {
	// The strength type of the monitoring rule. Valid values:
	//
	// 	- 0: The monitoring rule is a weak rule.
	//
	// 	- 1: The monitoring rule is a strong rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	BlockType *int32 `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The checker ID.
	//
	// example:
	//
	// 9
	Checker *int32 `json:"Checker,omitempty" xml:"Checker,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Verification
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The threshold for a critical alert. The threshold indicates the deviation of the monitoring result from the expected value. You can customize this threshold based on your business requirements. If a strong rule is used and a critical alert is triggered, nodes are blocked.
	//
	// example:
	//
	// 20
	CriticalThreshold *string `json:"CriticalThreshold,omitempty" xml:"CriticalThreshold,omitempty"`
	// The ID of the partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15224
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 0
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The method used to collect sample data. If you want to use a custom SQL statement as a sampling method, set this parameter to user_defined.
	//
	// example:
	//
	// count/table_count
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The comparison operator, such as >, >=, =, ≠, <, or <=.
	//
	// > If you set the Checker parameter to 9, you must configure the Operator parameter.
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Specifies whether the monitoring rule is a dynamic threshold rule. Valid values: 0 and 2. The value 0 indicates that the monitoring rule is not a dynamic threshold rule. The value 2 indicates that the monitoring rule is a dynamic threshold rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PredictType *int32 `json:"PredictType,omitempty" xml:"PredictType,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The fields that you want to monitor. If you want to monitor all fields in a table and check the table rows, set this parameter to table_count. If you want to monitor all fields in a table and check the table size, set this parameter to table_size.
	//
	// example:
	//
	// table_id
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The data type of the fields that you want to monitor. If you want to monitor all fields in a table, set this parameter to table. If you want to monitor only a specific field, set this parameter to bigint.
	//
	// example:
	//
	// bigint
	PropertyType *string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// The name of the monitoring rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the monitoring rule. Valid values: 0, 1, and 2. The value 0 indicates that the monitoring rule is created by the system. The value 1 indicates that the monitoring rule is created by a user. The value 2 indicates that the monitoring rule is a workspace-level rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The variable settings inserted before the custom rule. Format: x=a,y=b.
	//
	// example:
	//
	// x=a,y=b
	TaskSetting *string `json:"TaskSetting,omitempty" xml:"TaskSetting,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 7
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The trend of the monitoring result. Valid values:
	//
	// 	- up: increasing
	//
	// 	- down: decreasing
	//
	// 	- abs: absolute value
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
	// The threshold for a warning alert. The threshold indicates the deviation of the monitoring result from the expected value. You can customize this threshold based on your business requirements.
	//
	// example:
	//
	// 10
	WarningThreshold *string `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
	// The filter condition or custom SQL statement.
	//
	// example:
	//
	// table_id>1
	WhereCondition *string `json:"WhereCondition,omitempty" xml:"WhereCondition,omitempty"`
}

func (s CreateQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleRequest) GetBlockType() *int32 {
	return s.BlockType
}

func (s *CreateQualityRuleRequest) GetChecker() *int32 {
	return s.Checker
}

func (s *CreateQualityRuleRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateQualityRuleRequest) GetCriticalThreshold() *string {
	return s.CriticalThreshold
}

func (s *CreateQualityRuleRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *CreateQualityRuleRequest) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *CreateQualityRuleRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *CreateQualityRuleRequest) GetOperator() *string {
	return s.Operator
}

func (s *CreateQualityRuleRequest) GetPredictType() *int32 {
	return s.PredictType
}

func (s *CreateQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateQualityRuleRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateQualityRuleRequest) GetProperty() *string {
	return s.Property
}

func (s *CreateQualityRuleRequest) GetPropertyType() *string {
	return s.PropertyType
}

func (s *CreateQualityRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateQualityRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *CreateQualityRuleRequest) GetTaskSetting() *string {
	return s.TaskSetting
}

func (s *CreateQualityRuleRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *CreateQualityRuleRequest) GetTrend() *string {
	return s.Trend
}

func (s *CreateQualityRuleRequest) GetWarningThreshold() *string {
	return s.WarningThreshold
}

func (s *CreateQualityRuleRequest) GetWhereCondition() *string {
	return s.WhereCondition
}

func (s *CreateQualityRuleRequest) SetBlockType(v int32) *CreateQualityRuleRequest {
	s.BlockType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetChecker(v int32) *CreateQualityRuleRequest {
	s.Checker = &v
	return s
}

func (s *CreateQualityRuleRequest) SetComment(v string) *CreateQualityRuleRequest {
	s.Comment = &v
	return s
}

func (s *CreateQualityRuleRequest) SetCriticalThreshold(v string) *CreateQualityRuleRequest {
	s.CriticalThreshold = &v
	return s
}

func (s *CreateQualityRuleRequest) SetEntityId(v int64) *CreateQualityRuleRequest {
	s.EntityId = &v
	return s
}

func (s *CreateQualityRuleRequest) SetExpectValue(v string) *CreateQualityRuleRequest {
	s.ExpectValue = &v
	return s
}

func (s *CreateQualityRuleRequest) SetMethodName(v string) *CreateQualityRuleRequest {
	s.MethodName = &v
	return s
}

func (s *CreateQualityRuleRequest) SetOperator(v string) *CreateQualityRuleRequest {
	s.Operator = &v
	return s
}

func (s *CreateQualityRuleRequest) SetPredictType(v int32) *CreateQualityRuleRequest {
	s.PredictType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetProjectId(v int64) *CreateQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateQualityRuleRequest) SetProjectName(v string) *CreateQualityRuleRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityRuleRequest) SetProperty(v string) *CreateQualityRuleRequest {
	s.Property = &v
	return s
}

func (s *CreateQualityRuleRequest) SetPropertyType(v string) *CreateQualityRuleRequest {
	s.PropertyType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetRuleName(v string) *CreateQualityRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateQualityRuleRequest) SetRuleType(v int32) *CreateQualityRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetTaskSetting(v string) *CreateQualityRuleRequest {
	s.TaskSetting = &v
	return s
}

func (s *CreateQualityRuleRequest) SetTemplateId(v int32) *CreateQualityRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateQualityRuleRequest) SetTrend(v string) *CreateQualityRuleRequest {
	s.Trend = &v
	return s
}

func (s *CreateQualityRuleRequest) SetWarningThreshold(v string) *CreateQualityRuleRequest {
	s.WarningThreshold = &v
	return s
}

func (s *CreateQualityRuleRequest) SetWhereCondition(v string) *CreateQualityRuleRequest {
	s.WhereCondition = &v
	return s
}

func (s *CreateQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
