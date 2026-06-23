// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRulesShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetDataQualityRulesShrink() *string
	SetDataSourceId(v int64) *CreateDataQualityEvaluationTaskShrinkRequest
	GetDataSourceId() *int64
	SetDescription(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetDescription() *string
	SetHooksShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetHooksShrink() *string
	SetName(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetName() *string
	SetNotificationsShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetNotificationsShrink() *string
	SetProjectId(v int64) *CreateDataQualityEvaluationTaskShrinkRequest
	GetProjectId() *int64
	SetRuntimeConf(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetRuntimeConf() *string
	SetTargetShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetTargetShrink() *string
	SetTriggerShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest
	GetTriggerShrink() *string
}

type CreateDataQualityEvaluationTaskShrinkRequest struct {
	// The list of data quality rules associated with the data quality monitor. If DataQualityRule.Id is specified, the rule corresponding to that ID is associated with the newly created quality monitor. If not specified, a new rule is created from the other fields and associated with the newly created quality monitor.
	DataQualityRulesShrink *string `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty"`
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
	HooksShrink *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The name of the quality monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification subscription configuration.
	NotificationsShrink *string `json:"Notifications,omitempty" xml:"Notifications,omitempty"`
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
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The trigger configuration of the data quality validation task.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s CreateDataQualityEvaluationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetDataQualityRulesShrink() *string {
	return s.DataQualityRulesShrink
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetHooksShrink() *string {
	return s.HooksShrink
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetNotificationsShrink() *string {
	return s.NotificationsShrink
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetTargetShrink() *string {
	return s.TargetShrink
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetDataQualityRulesShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.DataQualityRulesShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetDataSourceId(v int64) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetDescription(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetHooksShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.HooksShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetName(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetNotificationsShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.NotificationsShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetProjectId(v int64) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetRuntimeConf(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.RuntimeConf = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetTargetShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) SetTriggerShrink(v string) *CreateDataQualityEvaluationTaskShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
