// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityEvaluationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRulesShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetDataQualityRulesShrink() *string
	SetDataSourceId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetDataSourceId() *int64
	SetDescription(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetDescription() *string
	SetHooksShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetHooksShrink() *string
	SetId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetName() *string
	SetNotificationsShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetNotificationsShrink() *string
	SetProjectId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetProjectId() *int64
	SetRuntimeConf(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetRuntimeConf() *string
	SetTargetShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetTargetShrink() *string
	SetTriggerShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest
	GetTriggerShrink() *string
}

type UpdateDataQualityEvaluationTaskShrinkRequest struct {
	// The list of monitoring rules that are associated with the monitor.
	DataQualityRulesShrink *string `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty"`
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
	HooksShrink *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
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
	NotificationsShrink *string `json:"Notifications,omitempty" xml:"Notifications,omitempty"`
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
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The trigger configuration of the monitor.
	TriggerShrink *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetDataQualityRulesShrink() *string {
	return s.DataQualityRulesShrink
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetHooksShrink() *string {
	return s.HooksShrink
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetNotificationsShrink() *string {
	return s.NotificationsShrink
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetTargetShrink() *string {
	return s.TargetShrink
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) GetTriggerShrink() *string {
	return s.TriggerShrink
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetDataQualityRulesShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.DataQualityRulesShrink = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetDataSourceId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetDescription(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetHooksShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.HooksShrink = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetName(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetNotificationsShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.NotificationsShrink = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetProjectId(v int64) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetRuntimeConf(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.RuntimeConf = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetTargetShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) SetTriggerShrink(v string) *UpdateDataQualityEvaluationTaskShrinkRequest {
	s.TriggerShrink = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
