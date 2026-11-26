// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSmartIndexConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifacts(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetArtifacts() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts
	SetBillingId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetBillingId() *string
	SetCode(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetCode() *string
	SetDirectoryId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetDirectoryId() *string
	SetEnabled(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetEnabled() *bool
	SetLastError(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLastError() *string
	SetLastSuccessTaskId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLastSuccessTaskId() *string
	SetLifecycle(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLifecycle() *string
	SetMessage(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetRequestId() *string
	SetRevision(v int64) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetRevision() *int64
	SetSchedule(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetSchedule() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule
	SetTargetDirId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetTargetDirId() *string
	SetTask(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
	GetTask() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBody struct {
	// The references to published smart index artifacts. This value is empty if no artifacts are available or if the feature is disabled. Internal storage credentials are not included.
	Artifacts *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Struct"`
	// The bill ID of the current or most recent task. This value is empty if no corresponding task exists. The existence of a bill ID does not indicate that the deduction is complete.
	//
	// example:
	//
	// billing-example
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// The business status code. A value of 200 indicates success. If the request fails, a specific business error code is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business ID of the root directory for the enterprise knowledge base. Subsequent configuration operations use this ID.
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// Indicates whether the LLM Wiki smart index is currently enabled. This parameter indicates the configuration switch status and does not indicate whether the build task is complete.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The description of the most recent configuration or scan error. This value is empty if no error occurred.
	//
	// example:
	//
	// Automatic scan failed
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// The business ID of the most recent successful build task. This value is empty if no successful task exists.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	LastSuccessTaskId *string `json:"lastSuccessTaskId,omitempty" xml:"lastSuccessTaskId,omitempty"`
	// The configuration lifecycle. Valid values:
	//
	// - DISABLED: Disabled.
	//
	// - ENABLED: Enabled.
	//
	// - DISABLING: Artifacts are being cleaned up after disabling.
	//
	// example:
	//
	// ENABLED
	Lifecycle *string `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	// The response message. The value is success if the request is successful. If the request fails, this parameter describes the cause of the error.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID. You can use this ID to troubleshoot issues related to this call.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The current configuration version. Submit this value as expectedRevision during the next update. The value is 0 when not configured, and increments by 1 after each successful save.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The automatic execution schedule. An empty cronExpression indicates manual execution.
	Schedule *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// The business ID of the smart index artifact directory. This value is empty if no artifact directory exists or if the feature is disabled.
	//
	// example:
	//
	// kb-artifacts-example
	TargetDirId *string `json:"targetDirId,omitempty" xml:"targetDirId,omitempty"`
	// The status of the task generated by this save operation, or the current or most recent task status returned by the read operation. This value is empty if no corresponding task exists.
	Task *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetArtifacts() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	return s.Artifacts
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetBillingId() *string {
	return s.BillingId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetLastError() *string {
	return s.LastError
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetLastSuccessTaskId() *string {
	return s.LastSuccessTaskId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetSchedule() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	return s.Schedule
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetTargetDirId() *string {
	return s.TargetDirId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) GetTask() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	return s.Task
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetArtifacts(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Artifacts = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetBillingId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.BillingId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetCode(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetDirectoryId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetEnabled(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Enabled = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetLastError(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.LastError = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetLastSuccessTaskId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.LastSuccessTaskId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetLifecycle(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Lifecycle = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetMessage(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetRevision(v int64) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Revision = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetSchedule(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Schedule = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetTargetDirId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.TargetDirId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) SetTask(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Task = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) Validate() error {
	if s.Artifacts != nil {
		if err := s.Artifacts.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts struct {
	// The resource reference for knowledge insights. This value is empty if the artifact is not yet available.
	KnowledgeInsights *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights `json:"knowledgeInsights,omitempty" xml:"knowledgeInsights,omitempty" type:"Struct"`
	// The reference to the directory where topic cards are located. This value is empty if the artifact is not yet available.
	TopicCards *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards `json:"topicCards,omitempty" xml:"topicCards,omitempty" type:"Struct"`
	// The resource reference for the topic graph. This value is empty if the artifact is not yet available.
	TopicGraph *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph `json:"topicGraph,omitempty" xml:"topicGraph,omitempty" type:"Struct"`
	// The resource reference for the topic overview. This value is empty if the artifact is not yet available.
	TopicOverview *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview `json:"topicOverview,omitempty" xml:"topicOverview,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetKnowledgeInsights() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights {
	return s.KnowledgeInsights
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicCards() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards {
	return s.TopicCards
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicGraph() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph {
	return s.TopicGraph
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicOverview() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview {
	return s.TopicOverview
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetKnowledgeInsights(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.KnowledgeInsights = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicCards(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicCards = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicGraph(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicGraph = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicOverview(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicOverview = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) Validate() error {
	if s.KnowledgeInsights != nil {
		if err := s.KnowledgeInsights.Validate(); err != nil {
			return err
		}
	}
	if s.TopicCards != nil {
		if err := s.TopicCards.Validate(); err != nil {
			return err
		}
	}
	if s.TopicGraph != nil {
		if err := s.TopicGraph.Validate(); err != nil {
			return err
		}
	}
	if s.TopicOverview != nil {
		if err := s.TopicOverview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights struct {
	// The business ID of the knowledge insights resource.
	//
	// example:
	//
	// 2001187
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) SetSourceId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards struct {
	// The business ID of the topic card artifact directory.
	//
	// example:
	//
	// kb-topic-cards-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) SetDirectoryId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph struct {
	// The business ID of the topic graph resource.
	//
	// example:
	//
	// 2001085
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) SetSourceId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview struct {
	// The business ID of the topic overview resource.
	//
	// example:
	//
	// 2001186
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) SetSourceId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule struct {
	// The current five-part cron automatic execution schedule. An empty value indicates manual execution.
	//
	// example:
	//
	// 0 2 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// The next scheduled execution time in ISO 8601 format. This value is empty when no automatic schedule is configured.
	//
	// example:
	//
	// 2026-11-19T02:00:00+08:00
	NextRunAt *string `json:"nextRunAt,omitempty" xml:"nextRunAt,omitempty"`
	// The IANA time zone name for the execution schedule. Default value: Asia/Shanghai when not configured.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetCronExpression(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.CronExpression = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetNextRunAt(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.NextRunAt = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetTimezone(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.Timezone = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask struct {
	// The bill ID corresponding to the task. The existence of a bill ID does not indicate that the deduction is complete.
	//
	// example:
	//
	// billing-example
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// The billing status of the task. For example, PENDING indicates that settlement is not complete, and ENDED indicates that the billing process has ended. This status is returned separately from the task execution status.
	//
	// example:
	//
	// PENDING
	BillingStatus *string `json:"billingStatus,omitempty" xml:"billingStatus,omitempty"`
	// The execution status of the task. For example, QUEUED indicates that the task is queued, RUNNING indicates that the task is running, SUCCEEDED indicates that the task succeeded, and FAILED indicates that the task failed. Evaluate the build result based on the task status. A successful configuration save does not indicate a successful task.
	//
	// example:
	//
	// QUEUED
	ExecutionStatus *string `json:"executionStatus,omitempty" xml:"executionStatus,omitempty"`
	// The time when the task was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-18T09:00:00+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the task was updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-18T09:01:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The description of the most recent task error. This value is empty if no error occurred.
	//
	// example:
	//
	// Task execution failed
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// The operation type of the smart index task. Valid values:
	//
	// - full: Full build.
	//
	// - incremental: Incremental update.
	//
	// example:
	//
	// full
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The execution progress reported by the task. This value is empty if the task has not yet reported progress.
	//
	// example:
	//
	// 0
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The reporting status of the task result. For example, pending indicates that the result is waiting to be reported, and not_required indicates that reporting is not required.
	//
	// example:
	//
	// pending
	ReportStatus *string `json:"reportStatus,omitempty" xml:"reportStatus,omitempty"`
	// Indicates whether an existing task was reused. The value is false if no task was reused.
	Reused *bool `json:"reused,omitempty" xml:"reused,omitempty"`
	// The business ID of the smart index task.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	TaskKey *string `json:"taskKey,omitempty" xml:"taskKey,omitempty"`
	// The trigger type of the task. Valid values:
	//
	// - manual: Manually triggered.
	//
	// - scheduled: Triggered by an automatic schedule.
	//
	// The build task generated upon the first enabling uses manual.
	//
	// example:
	//
	// manual
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// Indicates whether the task usage statistics are complete. This status is independent of the task execution status and billing status.
	UsageComplete *bool `json:"usageComplete,omitempty" xml:"usageComplete,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetBillingId() *string {
	return s.BillingId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetBillingStatus() *string {
	return s.BillingStatus
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetLastError() *string {
	return s.LastError
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetOperation() *string {
	return s.Operation
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetProgress() *int64 {
	return s.Progress
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetReused() *bool {
	return s.Reused
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetTaskKey() *string {
	return s.TaskKey
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetUsageComplete() *bool {
	return s.UsageComplete
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetBillingId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.BillingId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetBillingStatus(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.BillingStatus = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetExecutionStatus(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.ExecutionStatus = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetGmtCreate(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.GmtCreate = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetGmtModified(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.GmtModified = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetLastError(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.LastError = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetOperation(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Operation = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetProgress(v int64) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Progress = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetReportStatus(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.ReportStatus = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetReused(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Reused = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetTaskKey(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.TaskKey = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetTriggerType(v string) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.TriggerType = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetUsageComplete(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.UsageComplete = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
