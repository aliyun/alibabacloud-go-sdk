// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRules(v []*UpdateDataQualityEvaluationTaskRequestDataQualityRules) *UpdateDataQualityEvaluationTaskRequest
	GetDataQualityRules() []*UpdateDataQualityEvaluationTaskRequestDataQualityRules
	SetDataSourceId(v int64) *UpdateDataQualityEvaluationTaskRequest
	GetDataSourceId() *int64
	SetDescription(v string) *UpdateDataQualityEvaluationTaskRequest
	GetDescription() *string
	SetHooks(v []*UpdateDataQualityEvaluationTaskRequestHooks) *UpdateDataQualityEvaluationTaskRequest
	GetHooks() []*UpdateDataQualityEvaluationTaskRequestHooks
	SetId(v int64) *UpdateDataQualityEvaluationTaskRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityEvaluationTaskRequest
	GetName() *string
	SetNotifications(v *UpdateDataQualityEvaluationTaskRequestNotifications) *UpdateDataQualityEvaluationTaskRequest
	GetNotifications() *UpdateDataQualityEvaluationTaskRequestNotifications
	SetProjectId(v int64) *UpdateDataQualityEvaluationTaskRequest
	GetProjectId() *int64
	SetRuntimeConf(v string) *UpdateDataQualityEvaluationTaskRequest
	GetRuntimeConf() *string
	SetTarget(v *UpdateDataQualityEvaluationTaskRequestTarget) *UpdateDataQualityEvaluationTaskRequest
	GetTarget() *UpdateDataQualityEvaluationTaskRequestTarget
	SetTrigger(v *UpdateDataQualityEvaluationTaskRequestTrigger) *UpdateDataQualityEvaluationTaskRequest
	GetTrigger() *UpdateDataQualityEvaluationTaskRequestTrigger
}

type UpdateDataQualityEvaluationTaskRequest struct {
	// The list of monitoring rules that are associated with the monitor.
	DataQualityRules []*UpdateDataQualityEvaluationTaskRequestDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// The data source ID. You can call the [ListDataSources](https://help.aliyun.com/document_detail/211431.html) operation to query the ID.
	//
	// example:
	//
	// 358750
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the monitor.
	//
	// example:
	//
	// OpenAPI data quality monitoring test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook.
	Hooks []*UpdateDataQualityEvaluationTaskRequestHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7227061794
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitor.
	//
	// example:
	//
	// OpenAPI data quality monitoring test.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of alert notifications.
	Notifications *UpdateDataQualityEvaluationTaskRequestNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The ID of the DataWorks workspace.
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
	// The monitored object of the data quality monitoring task.
	Target *UpdateDataQualityEvaluationTaskRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the monitor.
	Trigger *UpdateDataQualityEvaluationTaskRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateDataQualityEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetDataQualityRules() []*UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	return s.DataQualityRules
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetHooks() []*UpdateDataQualityEvaluationTaskRequestHooks {
	return s.Hooks
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetNotifications() *UpdateDataQualityEvaluationTaskRequestNotifications {
	return s.Notifications
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetTarget() *UpdateDataQualityEvaluationTaskRequestTarget {
	return s.Target
}

func (s *UpdateDataQualityEvaluationTaskRequest) GetTrigger() *UpdateDataQualityEvaluationTaskRequestTrigger {
	return s.Trigger
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetDataQualityRules(v []*UpdateDataQualityEvaluationTaskRequestDataQualityRules) *UpdateDataQualityEvaluationTaskRequest {
	s.DataQualityRules = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetDataSourceId(v int64) *UpdateDataQualityEvaluationTaskRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetDescription(v string) *UpdateDataQualityEvaluationTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetHooks(v []*UpdateDataQualityEvaluationTaskRequestHooks) *UpdateDataQualityEvaluationTaskRequest {
	s.Hooks = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetId(v int64) *UpdateDataQualityEvaluationTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetName(v string) *UpdateDataQualityEvaluationTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetNotifications(v *UpdateDataQualityEvaluationTaskRequestNotifications) *UpdateDataQualityEvaluationTaskRequest {
	s.Notifications = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetProjectId(v int64) *UpdateDataQualityEvaluationTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetRuntimeConf(v string) *UpdateDataQualityEvaluationTaskRequest {
	s.RuntimeConf = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetTarget(v *UpdateDataQualityEvaluationTaskRequestTarget) *UpdateDataQualityEvaluationTaskRequest {
	s.Target = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) SetTrigger(v *UpdateDataQualityEvaluationTaskRequestTrigger) *UpdateDataQualityEvaluationTaskRequest {
	s.Trigger = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequest) Validate() error {
	if s.DataQualityRules != nil {
		for _, item := range s.DataQualityRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Hooks != nil {
		for _, item := range s.Hooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notifications != nil {
		if err := s.Notifications.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRules struct {
	// The check settings for sample data.
	CheckingConfig *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the rule.
	//
	// example:
	//
	// OpenAPI test rules
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The operations that you can perform after the rule-based check fails.
	ErrorHandlers []*UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the ID of the monitoring rule.
	//
	// example:
	//
	// 1022171560
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring rule.
	//
	// example:
	//
	// OpenAPI test rules
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters required for sampling.
	SamplingConfig *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The strength of the rule. Valid values:
	//
	// 	- Normal
	//
	// 	- High
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The ID of the template used by the rule.
	//
	// example:
	//
	// SYSTEM:field:null_value:fixed:0
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRules) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetCheckingConfig() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	return s.CheckingConfig
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetErrorHandlers() []*UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	return s.ErrorHandlers
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetSamplingConfig() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	return s.SamplingConfig
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetCheckingConfig(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.CheckingConfig = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetDescription(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetEnabled(v bool) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Enabled = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetErrorHandlers(v []*UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.ErrorHandlers = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetId(v int64) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetName(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetSamplingConfig(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.SamplingConfig = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetSeverity(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.Severity = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) SetTemplateCode(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRules {
	s.TemplateCode = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRules) Validate() error {
	if s.CheckingConfig != nil {
		if err := s.CheckingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorHandlers != nil {
		for _, item := range s.ErrorHandlers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SamplingConfig != nil {
		if err := s.SamplingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain specific types of thresholds, you must query reference values. In this example, an expression is used to specify the query method of referenced samples.
	//
	// example:
	//
	// {"bizdate": ["-1"]}
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The threshold calculation method. Valid values:
	//
	// 	- Fluctation
	//
	// 	- Auto
	//
	// 	- FluctationDiscreate
	//
	// 	- Average
	//
	// 	- Fixed
	//
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetThresholds() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	return s.Thresholds
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetReferencedSamplesFilter(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetThresholds(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) SetType(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
	Warned *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetCritical() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetExpected() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) GetWarned() *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetCritical(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetExpected(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) SetWarned(v *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Expected != nil {
		if err := s.Expected.Validate(); err != nil {
			return err
		}
	}
	if s.Warned != nil {
		if err := s.Warned.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical struct {
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

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetExpression(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetOperator(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) SetValue(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected struct {
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

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetExpression(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetOperator(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) SetValue(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned struct {
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

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetExpression(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetOperator(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) SetValue(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers struct {
	// The SQL statement that is used to filter failed tasks. If you define the rule by using custom SQL statements, you must specify an SQL statement to filter failed tasks.
	//
	// example:
	//
	// SELECT 	- FROM ods_d_openapi_log WHERE status = \\"Error\\"
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

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) SetErrorDataFilter(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) SetType(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig struct {
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
	// CountNotIn
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

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetMetric(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.Metric = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetMetricParameters(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetSamplingFilter(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) SetSettingConfig(v string) *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestHooks struct {
	// The hook trigger condition. When this condition is met, the hook action is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The hook type. Valid values:
	//
	// 	- BlockTaskInstance: Blocks the running of scheduling tasks.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskRequestHooks) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestHooks) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestHooks) GetCondition() *string {
	return s.Condition
}

func (s *UpdateDataQualityEvaluationTaskRequestHooks) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityEvaluationTaskRequestHooks) SetCondition(v string) *UpdateDataQualityEvaluationTaskRequestHooks {
	s.Condition = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestHooks) SetType(v string) *UpdateDataQualityEvaluationTaskRequestHooks {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestHooks) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestNotifications struct {
	// The notification trigger condition. When this condition is met, the alert notification is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations of the alert notification.
	Notifications []*UpdateDataQualityEvaluationTaskRequestNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s UpdateDataQualityEvaluationTaskRequestNotifications) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestNotifications) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestNotifications) GetCondition() *string {
	return s.Condition
}

func (s *UpdateDataQualityEvaluationTaskRequestNotifications) GetNotifications() []*UpdateDataQualityEvaluationTaskRequestNotificationsNotifications {
	return s.Notifications
}

func (s *UpdateDataQualityEvaluationTaskRequestNotifications) SetCondition(v string) *UpdateDataQualityEvaluationTaskRequestNotifications {
	s.Condition = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotifications) SetNotifications(v []*UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) *UpdateDataQualityEvaluationTaskRequestNotifications {
	s.Notifications = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotifications) Validate() error {
	if s.Notifications != nil {
		for _, item := range s.Notifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestNotificationsNotifications struct {
	// The alert notification methods.
	NotificationChannels []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The configurations of alert recipients.
	NotificationReceivers []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) GetNotificationChannels() []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) GetNotificationReceivers() []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) SetNotificationChannels(v []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) SetNotificationReceivers(v []*UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotifications) Validate() error {
	if s.NotificationChannels != nil {
		for _, item := range s.NotificationChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationReceivers != nil {
		for _, item := range s.NotificationReceivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels struct {
	// The alert notification methods.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers struct {
	// The additional parameters that are required when alerts are sent. The parameters are JSON-formatted strings. The following keys are supported:
	//
	// 	- atAll: specifies that all members in a group are mentioned when alerts are sent by using DingTalk. This parameter is valid only if you set ReceiverType to DingdingUrl.
	//
	// example:
	//
	// {  "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient.
	//
	// Valid values:
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

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetExtension(v string) *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestTarget struct {
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
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The configuration of the partitioned table.
	//
	// example:
	//
	// dt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table in Data Map.
	//
	// example:
	//
	// odsp.openapi.ods_d_openapi_log
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) SetDatabaseType(v string) *UpdateDataQualityEvaluationTaskRequestTarget {
	s.DatabaseType = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) SetPartitionSpec(v string) *UpdateDataQualityEvaluationTaskRequestTarget {
	s.PartitionSpec = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) SetTableGuid(v string) *UpdateDataQualityEvaluationTaskRequestTarget {
	s.TableGuid = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestTarget) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityEvaluationTaskRequestTrigger struct {
	// The IDs of scheduling tasks. This parameter is valid only if you set Type to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger type of the monitor. Valid values:
	//
	// 	- ByScheduledTaskInstance: The monitor is triggered by the associated scheduling tasks.
	//
	// 	- ByManual: The monitor is manually triggered.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskRequestTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *UpdateDataQualityEvaluationTaskRequestTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityEvaluationTaskRequestTrigger) SetTaskIds(v []*int64) *UpdateDataQualityEvaluationTaskRequestTrigger {
	s.TaskIds = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestTrigger) SetType(v string) *UpdateDataQualityEvaluationTaskRequestTrigger {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskRequestTrigger) Validate() error {
	return dara.Validate(s)
}
