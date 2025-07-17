// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRules(v []*CreateDataQualityEvaluationTaskRequestDataQualityRules) *CreateDataQualityEvaluationTaskRequest
	GetDataQualityRules() []*CreateDataQualityEvaluationTaskRequestDataQualityRules
	SetDataSourceId(v int64) *CreateDataQualityEvaluationTaskRequest
	GetDataSourceId() *int64
	SetDescription(v string) *CreateDataQualityEvaluationTaskRequest
	GetDescription() *string
	SetHooks(v []*CreateDataQualityEvaluationTaskRequestHooks) *CreateDataQualityEvaluationTaskRequest
	GetHooks() []*CreateDataQualityEvaluationTaskRequestHooks
	SetName(v string) *CreateDataQualityEvaluationTaskRequest
	GetName() *string
	SetNotifications(v *CreateDataQualityEvaluationTaskRequestNotifications) *CreateDataQualityEvaluationTaskRequest
	GetNotifications() *CreateDataQualityEvaluationTaskRequestNotifications
	SetProjectId(v int64) *CreateDataQualityEvaluationTaskRequest
	GetProjectId() *int64
	SetRuntimeConf(v string) *CreateDataQualityEvaluationTaskRequest
	GetRuntimeConf() *string
	SetTarget(v *CreateDataQualityEvaluationTaskRequestTarget) *CreateDataQualityEvaluationTaskRequest
	GetTarget() *CreateDataQualityEvaluationTaskRequestTarget
	SetTrigger(v *CreateDataQualityEvaluationTaskRequestTrigger) *CreateDataQualityEvaluationTaskRequest
	GetTrigger() *CreateDataQualityEvaluationTaskRequestTrigger
}

type CreateDataQualityEvaluationTaskRequest struct {
	// The list of monitoring rules that are associated with the monitor. If you configure the ID of a monitoring rule by using the DataQualityRule.Id parameter, the system associates the rule with a created monitor. If you do not configure the ID of a monitoring rule, the system creates a new monitoring rule by using other fields and associates the rule with a created monitor.
	DataQualityRules []*CreateDataQualityEvaluationTaskRequestDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// The data source ID. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the monitor.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook.
	Hooks []*CreateDataQualityEvaluationTaskRequestHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The name of the monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of alert notifications.
	Notifications *CreateDataQualityEvaluationTaskRequestNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You can use this parameter to specify the DataWorks workspace on which you want to perform the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The extended configurations in JSON-formatted strings. You can use this parameter only for monitors that are used to monitor the quality of E-MapReduce (EMR) data.
	//
	// 	- queue: The Yarn queue used when a monitor checks the quality of EMR data. By default, the queue configured for the current workspace is used.
	//
	// 	- sqlEngine: The SQL engine used when a monitor checks the quality of EMR data.
	//
	//     	- HIVE_SQL
	//
	//     	- SPARK_SQL
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the monitor.
	//
	// This parameter is required.
	Target *CreateDataQualityEvaluationTaskRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the monitor.
	Trigger *CreateDataQualityEvaluationTaskRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s CreateDataQualityEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequest) GetDataQualityRules() []*CreateDataQualityEvaluationTaskRequestDataQualityRules {
	return s.DataQualityRules
}

func (s *CreateDataQualityEvaluationTaskRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateDataQualityEvaluationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityEvaluationTaskRequest) GetHooks() []*CreateDataQualityEvaluationTaskRequestHooks {
	return s.Hooks
}

func (s *CreateDataQualityEvaluationTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityEvaluationTaskRequest) GetNotifications() *CreateDataQualityEvaluationTaskRequestNotifications {
	return s.Notifications
}

func (s *CreateDataQualityEvaluationTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityEvaluationTaskRequest) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *CreateDataQualityEvaluationTaskRequest) GetTarget() *CreateDataQualityEvaluationTaskRequestTarget {
	return s.Target
}

func (s *CreateDataQualityEvaluationTaskRequest) GetTrigger() *CreateDataQualityEvaluationTaskRequestTrigger {
	return s.Trigger
}

func (s *CreateDataQualityEvaluationTaskRequest) SetDataQualityRules(v []*CreateDataQualityEvaluationTaskRequestDataQualityRules) *CreateDataQualityEvaluationTaskRequest {
	s.DataQualityRules = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetDataSourceId(v int64) *CreateDataQualityEvaluationTaskRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetDescription(v string) *CreateDataQualityEvaluationTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetHooks(v []*CreateDataQualityEvaluationTaskRequestHooks) *CreateDataQualityEvaluationTaskRequest {
	s.Hooks = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetName(v string) *CreateDataQualityEvaluationTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetNotifications(v *CreateDataQualityEvaluationTaskRequestNotifications) *CreateDataQualityEvaluationTaskRequest {
	s.Notifications = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetProjectId(v int64) *CreateDataQualityEvaluationTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetRuntimeConf(v string) *CreateDataQualityEvaluationTaskRequest {
	s.RuntimeConf = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetTarget(v *CreateDataQualityEvaluationTaskRequestTarget) *CreateDataQualityEvaluationTaskRequest {
	s.Target = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) SetTrigger(v *CreateDataQualityEvaluationTaskRequestTrigger) *CreateDataQualityEvaluationTaskRequest {
	s.Trigger = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRules struct {
	// The check settings for sample data.
	CheckingConfig *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the monitoring rule.
	//
	// example:
	//
	// OpenAPI test rules
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the monitoring rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The operations that you can perform after the rule-based check fails.
	ErrorHandlers []*CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 2176
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring rule.
	//
	// example:
	//
	// OpenAPI test rules
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters required for sampling.
	SamplingConfig *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The strength of the monitoring rule. Valid values:
	//
	// 	- Normal
	//
	// 	- High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The ID of the template used by the monitoring rule.
	//
	// example:
	//
	// SYSTEM:field:null_value:fixed:0
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRules) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetCheckingConfig() *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	return s.CheckingConfig
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetErrorHandlers() []*CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	return s.ErrorHandlers
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetSamplingConfig() *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	return s.SamplingConfig
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetSeverity() *string {
	return s.Severity
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetCheckingConfig(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.CheckingConfig = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetDescription(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Description = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetEnabled(v bool) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Enabled = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetErrorHandlers(v []*CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.ErrorHandlers = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetId(v int64) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Id = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetName(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Name = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetSamplingConfig(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.SamplingConfig = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetSeverity(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Severity = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) SetTemplateCode(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRules {
	s.TemplateCode = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRules) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain specific types of thresholds, you must query reference values. In this example, an expression is used to specify the query method of referenced samples.
	//
	// example:
	//
	// {"bizdate": ["-1"]}
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The threshold calculation method. Valid values:
	//
	// 	- Fixed
	//
	// 	- Fluctation
	//
	// 	- FluctationDiscreate
	//
	// 	- Auto
	//
	// 	- Average
	//
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetThresholds() *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	return s.Thresholds
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetReferencedSamplesFilter(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetThresholds(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetType(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.Type = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
	Warned *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetCritical() *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetExpected() *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetWarned() *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetCritical(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetExpected(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetWarned(v *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical struct {
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
	// example:
	//
	// 0.01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetExpression(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetOperator(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetValue(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected struct {
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetExpression(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetOperator(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetValue(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned struct {
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
	// example:
	//
	// 0.001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetExpression(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetOperator(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetValue(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers struct {
	// The SQL statement that is used to filter failed tasks. If you define the rule by using custom SQL statements, you must specify an SQL statement to filter failed tasks.
	//
	// example:
	//
	// SELECT 	- FROM ods_api_log WHERE status = \\"Error\\";
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- SaveErrorData
	//
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) SetErrorDataFilter(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) SetType(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	s.Type = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig struct {
	// The metrics used for sampling. Valid values:
	//
	// 	- Count: the number of rows in the table.
	//
	// 	- Min: the minimum value of the field.
	//
	// 	- Max: the maximum value of the field.
	//
	// 	- Avg: the average value of the field.
	//
	// 	- DistinctCount: the number of unique values of the field after deduplication.
	//
	// 	- DistinctPercent: the proportion of the number of unique values of the field after deduplication to the number of rows in the table.
	//
	// 	- DuplicatedCount: the number of duplicated values of the field.
	//
	// 	- DuplicatedPercent: the proportion of the number of duplicated values of the field to the number of rows in the table.
	//
	// 	- TableSize: the table size.
	//
	// 	- NullValueCount: the number of rows in which the field value is null.
	//
	// 	- NullValuePercent: the proportion of the number of rows in which the field value is null to the number of rows in the table.
	//
	// 	- GroupCount: the field value and the number of rows for each field value.
	//
	// 	- CountNotIn: the number of rows in which the field values are different from the referenced values that you specified in the rule.
	//
	// 	- CountDistinctNotIn: the number of unique values that are different from the referenced values that you specified in the rule after deduplication.
	//
	// 	- UserDefinedSql: specifies that data is sampled by executing custom SQL statements.
	//
	// example:
	//
	// NullValueCount
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sampling.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The statements that are used to filter unnecessary data during sampling. The statements can be up to 16,777,215 characters in length.
	//
	// example:
	//
	// status != \\"Succeeded\\"
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// The statements that are used to configure the parameters required for sampling before you execute the sampling statements. The statements can be up to 1,000 characters in length. Only the MaxCompute database is supported.
	//
	// example:
	//
	// odps.sql.type.system.odps2=True,odps.sql.hive.compatible=True
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetMetric(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.Metric = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetMetricParameters(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetSamplingFilter(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetSettingConfig(v string) *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestHooks struct {
	// The hook trigger condition. When this condition is met, the hook action is triggered. Only two conditional expressions are supported:
	//
	// 1.  Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 2.  Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The hook type. Only one hook type is supported.
	//
	// 	- BlockTaskInstance: Blocks the running of scheduling tasks. A monitor is triggered by scheduling tasks. After a monitor finishes running, the monitor determines whether to block the running of scheduling tasks based on the hook condition.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestHooks) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestHooks) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestHooks) GetCondition() *string {
	return s.Condition
}

func (s *CreateDataQualityEvaluationTaskRequestHooks) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityEvaluationTaskRequestHooks) SetCondition(v string) *CreateDataQualityEvaluationTaskRequestHooks {
	s.Condition = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestHooks) SetType(v string) *CreateDataQualityEvaluationTaskRequestHooks {
	s.Type = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestHooks) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestNotifications struct {
	// The notification trigger condition. When this condition is met, the alert notification is triggered. Only two conditional expressions are supported:
	//
	// Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical. Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations of the alert notification.
	Notifications []*CreateDataQualityEvaluationTaskRequestNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s CreateDataQualityEvaluationTaskRequestNotifications) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestNotifications) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestNotifications) GetCondition() *string {
	return s.Condition
}

func (s *CreateDataQualityEvaluationTaskRequestNotifications) GetNotifications() []*CreateDataQualityEvaluationTaskRequestNotificationsNotifications {
	return s.Notifications
}

func (s *CreateDataQualityEvaluationTaskRequestNotifications) SetCondition(v string) *CreateDataQualityEvaluationTaskRequestNotifications {
	s.Condition = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotifications) SetNotifications(v []*CreateDataQualityEvaluationTaskRequestNotificationsNotifications) *CreateDataQualityEvaluationTaskRequestNotifications {
	s.Notifications = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotifications) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestNotificationsNotifications struct {
	// The alert notification methods.
	NotificationChannels []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The configurations of alert recipients.
	NotificationReceivers []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotifications) GetNotificationChannels() []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotifications) GetNotificationReceivers() []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotifications) SetNotificationChannels(v []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) *CreateDataQualityEvaluationTaskRequestNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotifications) SetNotificationReceivers(v []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) *CreateDataQualityEvaluationTaskRequestNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotifications) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels struct {
	// The alert notification methods.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers struct {
	// The additional parameters that are required when alerts are sent. The parameters are JSON-formatted strings. The following keys are supported:
	//
	// 	- atAll: specifies that all members in a group are mentioned when alerts are sent by using DingTalk. This parameter is valid only if you set ReceiverType to DingdingUrl.
	//
	// example:
	//
	// {  "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid values:
	//
	// 	- WebhookUrl
	//
	// 	- FeishuUrl
	//
	// 	- DingdingUrl
	//
	// 	- WeixinUrl
	//
	// 	- AliUid
	//
	// example:
	//
	// DingdingUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipient.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetExtension(v string) *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestTarget struct {
	// The type of the database to which the table belongs. Valid values:
	//
	// 	- maxcompute
	//
	// 	- hologres
	//
	// 	- cdh
	//
	// 	- analyticdb_for_mysql
	//
	// 	- starrocks
	//
	// 	- emr
	//
	// 	- analyticdb_for_postgresql
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The configuration of the partitioned table.
	//
	// example:
	//
	// pt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table in Data Map.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.api_test.ods_openapi_log_d
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestTarget) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) SetDatabaseType(v string) *CreateDataQualityEvaluationTaskRequestTarget {
	s.DatabaseType = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) SetPartitionSpec(v string) *CreateDataQualityEvaluationTaskRequestTarget {
	s.PartitionSpec = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) SetTableGuid(v string) *CreateDataQualityEvaluationTaskRequestTarget {
	s.TableGuid = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestTarget) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityEvaluationTaskRequestTrigger struct {
	// The IDs of scheduling tasks. This parameter is valid only if you set Type to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger type of the monitor. Valid values:
	//
	// 	- ByManual (default): The monitor is manually triggered.
	//
	// 	- ByScheduledTaskInstance: The monitor is triggered by the associated scheduling tasks.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityEvaluationTaskRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskRequestTrigger) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskRequestTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *CreateDataQualityEvaluationTaskRequestTrigger) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityEvaluationTaskRequestTrigger) SetTaskIds(v []*int64) *CreateDataQualityEvaluationTaskRequestTrigger {
	s.TaskIds = v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestTrigger) SetType(v string) *CreateDataQualityEvaluationTaskRequestTrigger {
	s.Type = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskRequestTrigger) Validate() error {
	return dara.Validate(s)
}
