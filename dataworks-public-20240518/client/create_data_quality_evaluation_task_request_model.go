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
	// The list of data quality rules associated with the data quality monitor. If DataQualityRule.Id is specified, the rule corresponding to that ID is associated with the newly created quality monitor. If not specified, a new rule is created from the other fields and associated with the newly created quality monitor.
	DataQualityRules []*CreateDataQualityEvaluationTaskRequestDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// The ID of the data source. You can call [ListDataSources](https://help.aliyun.com/document_detail/211431.html) to obtain the ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the quality monitoring task.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook settings.
	Hooks []*CreateDataQualityEvaluationTaskRequestHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The name of the quality monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification subscription configuration.
	Notifications *CreateDataQualityEvaluationTaskRequestNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace used by this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The extended configuration, a JSON-formatted string. This setting takes effect only for EMR-type data quality monitors.
	//
	// - queue: The YARN queue used when running EMR data quality validation. The default is the queue configured for the current project.
	//
	// - sqlEngine: The SQL engine used when running EMR data validation.
	//
	//     + HIVE_SQL
	//
	//     + SPARK_SQL
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The data quality monitoring object.
	//
	// This parameter is required.
	Target *CreateDataQualityEvaluationTaskRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the data quality validation task.
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

type CreateDataQualityEvaluationTaskRequestDataQualityRules struct {
	// The sample validation settings.
	CheckingConfig *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the data quality rule.
	//
	// example:
	//
	// OpenAPI test rules
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the quality rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of error handlers for issues detected by the quality rule validation.
	ErrorHandlers []*CreateDataQualityEvaluationTaskRequestDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// example:
	//
	// 2176
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data quality rule.
	//
	// example:
	//
	// OpenAPI test rules
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters required when collecting samples.
	SamplingConfig *CreateDataQualityEvaluationTaskRequestDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The business severity level of the rule (corresponding to strong/weak rules in the console). Valid values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The unique identifier of the rule template that the rule references.
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

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfig struct {
	// For some threshold types, reference samples must be queried and aggregated to derive the threshold used for comparison. This field uses an expression to describe how the reference samples are queried.
	//
	// example:
	//
	// {"bizdate": ["-1"]}
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The validation threshold settings.
	Thresholds *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The method used to compute the threshold. Valid values:
	//
	// - Fixed
	//
	// - Fluctation
	//
	// - FluctationDiscreate
	//
	// - Auto
	//
	// - Average
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
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholds struct {
	// The threshold settings for the critical warning level.
	Critical *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold settings.
	Expected *CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for the normal warning level.
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

type CreateDataQualityEvaluationTaskRequestDataQualityRulesCheckingConfigThresholdsCritical struct {
	// The threshold expression.
	//
	// Fluctuation-type rules must use expressions to specify the fluctuation threshold. For example:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use expressions to configure thresholds. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
	//
	// - !=
	//
	// - =
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
	// Fluctuation-type rules must use expressions to specify the fluctuation threshold. For example:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use expressions to configure thresholds. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
	//
	// - !=
	//
	// - =
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
	// Fluctuation-type rules must use expressions to specify the fluctuation threshold. For example:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use expressions to configure thresholds. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
	//
	// - !=
	//
	// - =
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
	// For custom SQL rules, the user must specify a SQL statement to filter the problematic data.
	//
	// example:
	//
	// SELECT 	- FROM ods_api_log WHERE status = \\"Error\\";
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// The type of the handler. Valid values:
	//
	// - SaveErrorData: retains the problematic data.
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
	// The name of the sampling metric. Valid values:
	//
	// - Count: the number of rows in the table.
	//
	// - Min: the minimum value of the field.
	//
	// - Max: the maximum value of the field.
	//
	// - Avg: the average value of the field.
	//
	// - DistinctCount: the number of distinct values of the field.
	//
	// - DistinctPercent: the ratio of the number of distinct values of the field to the number of rows.
	//
	// - DuplicatedCount: the number of duplicate values of the field.
	//
	// - DuplicatedPercent: the ratio of the number of duplicate values of the field to the number of rows.
	//
	// - TableSize: the size of the table.
	//
	// - NullValueCount: the number of rows in which the field is null.
	//
	// - NullValuePercent: the ratio of rows in which the field is null.
	//
	// - GroupCount: after grouping by the field value, the count of rows for each value.
	//
	// - CountNotIn: the number of rows whose enumeration values do not match.
	//
	// - CountDistinctNotIn: the number of distinct values whose enumeration values do not match.
	//
	// - UserDefinedSql: collect samples using a custom SQL statement.
	//
	// example:
	//
	// NullValueCount
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required when collecting samples.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// An additional filter condition applied during sampling to exclude data that is not of interest. The maximum length is 16,777,215 characters.
	//
	// example:
	//
	// status != \\"Succeeded\\"
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// The runtime parameter statements inserted and executed before the sampling statement is executed. The maximum length is 1000 characters. Only MaxCompute is currently supported.
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
	// The trigger condition of the hook. The hook action is triggered when this condition is met. Currently only two forms of expressions are supported:
	//
	// 1. Specify a single combination of rule severity and rule validation status, for example `${severity} == "High" AND ${status} == "Critical"`, which means the condition is met if among the executed rules there exists a rule whose severity is High and whose validation result is Critical.
	//
	// 2. Specify multiple combinations of rule severity and rule validation status, for example `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`, which means the condition is met if among the executed rules there exists a rule whose severity is High and validation result is Critical, or a rule whose severity is Normal and validation result is Critical, or a rule whose severity is Normal and validation result is Error. The enumeration of severity in the expression is the same as severity in DataQualityRule, and the enumeration of status is the same as status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the hook. Currently only one type is supported:
	//
	// - BlockTaskInstance: blocks the scheduling task from continuing to run. If the data quality monitor is triggered by a scheduling task, after the monitor finishes running, Hook.Condition is evaluated to determine whether to block the scheduling task from continuing to run.
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
	// The trigger condition of the notification. The notification is triggered when this condition is met. Currently only two forms of expressions are supported:
	//
	// Specify a single combination of rule severity and rule validation status, for example `${severity} == "High" AND ${status} == "Critical"`, which means the condition is met if among the executed rules there exists a rule whose severity is High and whose validation result is Critical.
	//
	// Specify multiple combinations of rule severity and rule validation status, for example `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`, which means the condition is met if among the executed rules there exists a rule whose severity is High and validation result is Critical, or a rule whose severity is Normal and validation result is Critical, or a rule whose severity is Normal and validation result is Error. The enumeration of severity in the expression is the same as severity in DataQualityRule, and the enumeration of status is the same as status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The notification settings.
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

type CreateDataQualityEvaluationTaskRequestNotificationsNotifications struct {
	// The notification methods.
	NotificationChannels []*CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert recipient settings.
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

type CreateDataQualityEvaluationTaskRequestNotificationsNotificationsNotificationChannels struct {
	// The notification methods.
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
	// Additional parameters used when sending alerts, in JSON format. Supported keys:
	//
	// - atAll: whether to mention all members (@all) in the group when sending a DingTalk alert. This key takes effect when ReceiverType is set to DingdingUrl.
	//
	// example:
	//
	// {  "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid values:
	//
	// - WebhookUrl: a custom webhook URL.
	//
	// - FeishuUrl: a Lark (Feishu) alert URL.
	//
	// - DingdingUrl: a DingTalk alert URL.
	//
	// - WeixinUrl: a WeCom (Enterprise WeChat) alert URL.
	//
	// - AliUid: an Alibaba Cloud user ID.
	//
	// example:
	//
	// DingdingUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
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
	// - maxcompute
	//
	// - hologres
	//
	// - cdh
	//
	// - analyticdb_for_mysql
	//
	// - starrocks
	//
	// - emr
	//
	// - analyticdb_for_postgresql
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The partition settings of the partitioned table.
	//
	// example:
	//
	// pt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The unique ID of the table in Data Map.
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
	// The list of scheduling task IDs. This parameter is valid when Type is set to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger type of the quality monitoring task. Valid values:
	//
	// - ByManual: triggered manually. This is the default value.
	//
	// - ByScheduledTaskInstance: triggered by an associated scheduling task.
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
