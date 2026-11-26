// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSmartIndexConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifacts(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetArtifacts() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts
	SetBillingId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetBillingId() *string
	SetCode(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetCode() *string
	SetDirectoryId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetDirectoryId() *string
	SetEnabled(v bool) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetEnabled() *bool
	SetLastError(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLastError() *string
	SetLastSuccessTaskId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLastSuccessTaskId() *string
	SetLifecycle(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetLifecycle() *string
	SetMessage(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetRequestId() *string
	SetRevision(v int64) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetRevision() *int64
	SetSchedule(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetSchedule() *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule
	SetTargetDirId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetTargetDirId() *string
	SetTask(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) *GetKnowledgeBaseSmartIndexConfigurationResponseBody
	GetTask() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBody struct {
	// 已发布产物引用
	Artifacts *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Struct"`
	// 当前或最近任务账单标识
	//
	// example:
	//
	// billing-example
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 企业知识库根目录业务 ID
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 是否开启 LLM Wiki 智能索引
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 最近一次配置或扫描错误
	//
	// example:
	//
	// 自动扫描失败
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// 最近成功任务标识
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	LastSuccessTaskId *string `json:"lastSuccessTaskId,omitempty" xml:"lastSuccessTaskId,omitempty"`
	// 配置生命周期：DISABLED、ENABLED 或 DISABLING
	//
	// example:
	//
	// ENABLED
	Lifecycle *string `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	// 错误描述，成功时为 success
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 配置版本；更新时作为 expectedRevision 提交
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// 自动执行计划
	Schedule *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// 智能索引产物目录业务 ID
	//
	// example:
	//
	// kb-artifacts-example
	TargetDirId *string `json:"targetDirId,omitempty" xml:"targetDirId,omitempty"`
	// 当前或最近任务状态
	Task *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetArtifacts() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	return s.Artifacts
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetBillingId() *string {
	return s.BillingId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetLastError() *string {
	return s.LastError
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetLastSuccessTaskId() *string {
	return s.LastSuccessTaskId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetSchedule() *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	return s.Schedule
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetTargetDirId() *string {
	return s.TargetDirId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) GetTask() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	return s.Task
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetArtifacts(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Artifacts = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetBillingId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.BillingId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetCode(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetDirectoryId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetEnabled(v bool) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetLastError(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.LastError = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetLastSuccessTaskId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.LastSuccessTaskId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetLifecycle(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Lifecycle = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetMessage(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetRequestId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetRevision(v int64) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Revision = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetSchedule(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Schedule = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetTargetDirId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.TargetDirId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) SetTask(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	s.Task = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBody) Validate() error {
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

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts struct {
	// 知识洞察
	KnowledgeInsights *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights `json:"knowledgeInsights,omitempty" xml:"knowledgeInsights,omitempty" type:"Struct"`
	// 主题卡片目录
	TopicCards *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards `json:"topicCards,omitempty" xml:"topicCards,omitempty" type:"Struct"`
	// 主题关系图
	TopicGraph *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph `json:"topicGraph,omitempty" xml:"topicGraph,omitempty" type:"Struct"`
	// 主题概览
	TopicOverview *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview `json:"topicOverview,omitempty" xml:"topicOverview,omitempty" type:"Struct"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetKnowledgeInsights() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights {
	return s.KnowledgeInsights
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicCards() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards {
	return s.TopicCards
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicGraph() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph {
	return s.TopicGraph
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) GetTopicOverview() *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview {
	return s.TopicOverview
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetKnowledgeInsights(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.KnowledgeInsights = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicCards(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicCards = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicGraph(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicGraph = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) SetTopicOverview(v *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts {
	s.TopicOverview = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifacts) Validate() error {
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

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights struct {
	// 产物资源业务 ID
	//
	// example:
	//
	// 2001187
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) SetSourceId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsKnowledgeInsights) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards struct {
	// 产物目录业务 ID
	//
	// example:
	//
	// kb-topic-cards-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) SetDirectoryId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicCards) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph struct {
	// 产物资源业务 ID
	//
	// example:
	//
	// 2001085
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) SetSourceId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicGraph) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview struct {
	// 产物资源业务 ID
	//
	// example:
	//
	// 2001186
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) SetSourceId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyArtifactsTopicOverview) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule struct {
	// Cron 表达式；空表示手动执行
	//
	// example:
	//
	// 0 2 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// 下次计划执行时间，ISO 8601 格式
	//
	// example:
	//
	// 2026-11-19T02:00:00+08:00
	NextRunAt *string `json:"nextRunAt,omitempty" xml:"nextRunAt,omitempty"`
	// 计划使用的 IANA 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) GetTimezone() *string {
	return s.Timezone
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetCronExpression(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.CronExpression = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetNextRunAt(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.NextRunAt = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) SetTimezone(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule {
	s.Timezone = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodySchedule) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask struct {
	// 任务账单标识
	//
	// example:
	//
	// billing-example
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// 任务计费状态
	//
	// example:
	//
	// PENDING
	BillingStatus *string `json:"billingStatus,omitempty" xml:"billingStatus,omitempty"`
	// 任务执行状态
	//
	// example:
	//
	// QUEUED
	ExecutionStatus *string `json:"executionStatus,omitempty" xml:"executionStatus,omitempty"`
	// 任务创建时间，ISO 8601 格式
	//
	// example:
	//
	// 2026-11-18T09:00:00+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 任务更新时间，ISO 8601 格式
	//
	// example:
	//
	// 2026-11-18T09:01:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 最近一次任务错误
	//
	// example:
	//
	// 任务执行失败
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// 任务操作类型
	//
	// example:
	//
	// full
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 任务执行进度
	//
	// example:
	//
	// 0
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// 任务结果上报状态
	//
	// example:
	//
	// pending
	ReportStatus *string `json:"reportStatus,omitempty" xml:"reportStatus,omitempty"`
	// 是否复用已有任务
	Reused *bool `json:"reused,omitempty" xml:"reused,omitempty"`
	// 任务业务标识
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	TaskKey *string `json:"taskKey,omitempty" xml:"taskKey,omitempty"`
	// 任务触发类型
	//
	// example:
	//
	// manual
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// 用量统计是否完成
	UsageComplete *bool `json:"usageComplete,omitempty" xml:"usageComplete,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetBillingId() *string {
	return s.BillingId
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetBillingStatus() *string {
	return s.BillingStatus
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetLastError() *string {
	return s.LastError
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetOperation() *string {
	return s.Operation
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetProgress() *int64 {
	return s.Progress
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetReused() *bool {
	return s.Reused
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetTaskKey() *string {
	return s.TaskKey
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) GetUsageComplete() *bool {
	return s.UsageComplete
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetBillingId(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.BillingId = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetBillingStatus(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.BillingStatus = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetExecutionStatus(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.ExecutionStatus = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetGmtCreate(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.GmtCreate = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetGmtModified(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.GmtModified = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetLastError(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.LastError = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetOperation(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Operation = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetProgress(v int64) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Progress = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetReportStatus(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.ReportStatus = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetReused(v bool) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.Reused = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetTaskKey(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.TaskKey = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetTriggerType(v string) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.TriggerType = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) SetUsageComplete(v bool) *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask {
	s.UsageComplete = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
