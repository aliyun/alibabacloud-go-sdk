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
	// The list of monitoring rules that are associated with the monitor. If you configure the ID of a monitoring rule by using the DataQualityRule.Id parameter, the system associates the rule with a created monitor. If you do not configure the ID of a monitoring rule, the system creates a new monitoring rule by using other fields and associates the rule with a created monitor.
	DataQualityRulesShrink *string `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty"`
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
	HooksShrink *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The name of the monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of alert notifications.
	NotificationsShrink *string `json:"Notifications,omitempty" xml:"Notifications,omitempty"`
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
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The trigger configuration of the monitor.
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
