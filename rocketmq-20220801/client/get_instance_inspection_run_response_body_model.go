// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiSummary(v *GetInstanceInspectionRunResponseBodyAiSummary) *GetInstanceInspectionRunResponseBody
	GetAiSummary() *GetInstanceInspectionRunResponseBodyAiSummary
	SetConfigSnapshot(v *GetInstanceInspectionRunResponseBodyConfigSnapshot) *GetInstanceInspectionRunResponseBody
	GetConfigSnapshot() *GetInstanceInspectionRunResponseBodyConfigSnapshot
	SetFinishedAt(v string) *GetInstanceInspectionRunResponseBody
	GetFinishedAt() *string
	SetInstanceId(v string) *GetInstanceInspectionRunResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *GetInstanceInspectionRunResponseBody
	GetInstanceName() *string
	SetItems(v []*GetInstanceInspectionRunResponseBodyItems) *GetInstanceInspectionRunResponseBody
	GetItems() []*GetInstanceInspectionRunResponseBodyItems
	SetNotification(v *GetInstanceInspectionRunResponseBodyNotification) *GetInstanceInspectionRunResponseBody
	GetNotification() *GetInstanceInspectionRunResponseBodyNotification
	SetRegionId(v string) *GetInstanceInspectionRunResponseBody
	GetRegionId() *string
	SetReport(v *GetInstanceInspectionRunResponseBodyReport) *GetInstanceInspectionRunResponseBody
	GetReport() *GetInstanceInspectionRunResponseBodyReport
	SetRequestedAt(v string) *GetInstanceInspectionRunResponseBody
	GetRequestedAt() *string
	SetRunId(v string) *GetInstanceInspectionRunResponseBody
	GetRunId() *string
	SetStatus(v string) *GetInstanceInspectionRunResponseBody
	GetStatus() *string
	SetSummary(v *GetInstanceInspectionRunResponseBodySummary) *GetInstanceInspectionRunResponseBody
	GetSummary() *GetInstanceInspectionRunResponseBodySummary
	SetTimeRange(v *GetInstanceInspectionRunResponseBodyTimeRange) *GetInstanceInspectionRunResponseBody
	GetTimeRange() *GetInstanceInspectionRunResponseBodyTimeRange
	SetTriggerSource(v string) *GetInstanceInspectionRunResponseBody
	GetTriggerSource() *string
}

type GetInstanceInspectionRunResponseBody struct {
	// 每次巡检结果由 LLM 总结；包含生成状态及成功后的原文，未生成时为空。
	//
	// example:
	//
	// {}
	AiSummary *GetInstanceInspectionRunResponseBodyAiSummary `json:"aiSummary,omitempty" xml:"aiSummary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ConfigSnapshot *GetInstanceInspectionRunResponseBodyConfigSnapshot `json:"configSnapshot,omitempty" xml:"configSnapshot,omitempty" type:"Struct"`
	// example:
	//
	// 2026-09-08T00:01:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// rmq-cn-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 订单实例
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// []
	Items []*GetInstanceInspectionRunResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Notification *GetInstanceInspectionRunResponseBodyNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {}
	Report *GetInstanceInspectionRunResponseBodyReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
	// example:
	//
	// 2026-09-08T00:00:00Z
	RequestedAt *string `json:"requestedAt,omitempty" xml:"requestedAt,omitempty"`
	// example:
	//
	// inspection-run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// {}
	Summary *GetInstanceInspectionRunResponseBodySummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionRunResponseBodyTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// manual
	TriggerSource *string `json:"triggerSource,omitempty" xml:"triggerSource,omitempty"`
}

func (s GetInstanceInspectionRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBody) GetAiSummary() *GetInstanceInspectionRunResponseBodyAiSummary {
	return s.AiSummary
}

func (s *GetInstanceInspectionRunResponseBody) GetConfigSnapshot() *GetInstanceInspectionRunResponseBodyConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *GetInstanceInspectionRunResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetInstanceInspectionRunResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionRunResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceInspectionRunResponseBody) GetItems() []*GetInstanceInspectionRunResponseBodyItems {
	return s.Items
}

func (s *GetInstanceInspectionRunResponseBody) GetNotification() *GetInstanceInspectionRunResponseBodyNotification {
	return s.Notification
}

func (s *GetInstanceInspectionRunResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionRunResponseBody) GetReport() *GetInstanceInspectionRunResponseBodyReport {
	return s.Report
}

func (s *GetInstanceInspectionRunResponseBody) GetRequestedAt() *string {
	return s.RequestedAt
}

func (s *GetInstanceInspectionRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetInstanceInspectionRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionRunResponseBody) GetSummary() *GetInstanceInspectionRunResponseBodySummary {
	return s.Summary
}

func (s *GetInstanceInspectionRunResponseBody) GetTimeRange() *GetInstanceInspectionRunResponseBodyTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionRunResponseBody) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetInstanceInspectionRunResponseBody) SetAiSummary(v *GetInstanceInspectionRunResponseBodyAiSummary) *GetInstanceInspectionRunResponseBody {
	s.AiSummary = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetConfigSnapshot(v *GetInstanceInspectionRunResponseBodyConfigSnapshot) *GetInstanceInspectionRunResponseBody {
	s.ConfigSnapshot = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetFinishedAt(v string) *GetInstanceInspectionRunResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetInstanceId(v string) *GetInstanceInspectionRunResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetInstanceName(v string) *GetInstanceInspectionRunResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetItems(v []*GetInstanceInspectionRunResponseBodyItems) *GetInstanceInspectionRunResponseBody {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetNotification(v *GetInstanceInspectionRunResponseBodyNotification) *GetInstanceInspectionRunResponseBody {
	s.Notification = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetRegionId(v string) *GetInstanceInspectionRunResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetReport(v *GetInstanceInspectionRunResponseBodyReport) *GetInstanceInspectionRunResponseBody {
	s.Report = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetRequestedAt(v string) *GetInstanceInspectionRunResponseBody {
	s.RequestedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetRunId(v string) *GetInstanceInspectionRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetStatus(v string) *GetInstanceInspectionRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetSummary(v *GetInstanceInspectionRunResponseBodySummary) *GetInstanceInspectionRunResponseBody {
	s.Summary = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetTimeRange(v *GetInstanceInspectionRunResponseBodyTimeRange) *GetInstanceInspectionRunResponseBody {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) SetTriggerSource(v string) *GetInstanceInspectionRunResponseBody {
	s.TriggerSource = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBody) Validate() error {
	if s.AiSummary != nil {
		if err := s.AiSummary.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigSnapshot != nil {
		if err := s.ConfigSnapshot.Validate(); err != nil {
			return err
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyAiSummary struct {
	// Pending 等待、Running 生成中、Succeeded 已生成、Failed 生成失败；不改变巡检结果。
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 仅生成成功时返回，最多 200 字，不由前端拼接或截取。
	//
	// example:
	//
	// 本次巡检未发现风险，实例运行正常。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyAiSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyAiSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyAiSummary) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionRunResponseBodyAiSummary) GetText() *string {
	return s.Text
}

func (s *GetInstanceInspectionRunResponseBodyAiSummary) SetStatus(v string) *GetInstanceInspectionRunResponseBodyAiSummary {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyAiSummary) SetText(v string) *GetInstanceInspectionRunResponseBodyAiSummary {
	s.Text = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyAiSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyConfigSnapshot struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*GetInstanceInspectionRunResponseBodyConfigSnapshotItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshot) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) GetItems() []*GetInstanceInspectionRunResponseBodyConfigSnapshotItems {
	return s.Items
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) GetResourceScope() *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope {
	return s.ResourceScope
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) GetSchedule() *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule {
	return s.Schedule
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) SetItems(v []*GetInstanceInspectionRunResponseBodyConfigSnapshotItems) *GetInstanceInspectionRunResponseBodyConfigSnapshot {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) SetResourceScope(v *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) *GetInstanceInspectionRunResponseBodyConfigSnapshot {
	s.ResourceScope = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) SetSchedule(v *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) *GetInstanceInspectionRunResponseBodyConfigSnapshot {
	s.Schedule = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshot) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceScope != nil {
		if err := s.ResourceScope.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyConfigSnapshotItems struct {
	// This parameter is required.
	//
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotItems) SetParameters(v interface{}) *GetInstanceInspectionRunResponseBodyConfigSnapshotItems {
	s.Parameters = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotItems) SetTemplateCode(v string) *GetInstanceInspectionRunResponseBodyConfigSnapshotItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotItems) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope struct {
	// example:
	//
	// ["orders-g1"]
	ConsumerGroups []*string `json:"consumerGroups,omitempty" xml:"consumerGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// All
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// ["orders"]
	Topics []*string `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) GetMode() *string {
	return s.Mode
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) SetConsumerGroups(v []*string) *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) SetMode(v string) *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope {
	s.Mode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) SetTopics(v []*string) *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope {
	s.Topics = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotResourceScope) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule struct {
	// example:
	//
	// 0 3 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) GetCron() *string {
	return s.Cron
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) SetCron(v string) *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule {
	s.Cron = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) SetEnabled(v bool) *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule {
	s.Enabled = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyConfigSnapshotSchedule) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}}
	EffectiveConfig *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig `json:"effectiveConfig,omitempty" xml:"effectiveConfig,omitempty" type:"Struct"`
	// example:
	//
	// INSPECTION_EXECUTION_INTERRUPTED
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 消费延迟超过阈值
	Headline *string `json:"headline,omitempty" xml:"headline,omitempty"`
	// example:
	//
	// []
	Metrics []*GetInstanceInspectionRunResponseBodyItemsMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 消费健康巡检
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2026-09-01T01:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// example:
	//
	// 3
	ResourceCount *int64 `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
	// example:
	//
	// {"infoCount":2,"riskCount":1}
	ResourceSummary *GetInstanceInspectionRunResponseBodyItemsResourceSummary `json:"resourceSummary,omitempty" xml:"resourceSummary,omitempty" type:"Struct"`
	// example:
	//
	// ["ConsumerGroup"]
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// Risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// warning
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// {"startTime":"2026-09-01T00:00:00Z","endTime":"2026-09-01T01:00:00Z"}
	TimeRange *GetInstanceInspectionRunResponseBodyItemsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionRunResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetEffectiveConfig() *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig {
	return s.EffectiveConfig
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetHeadline() *string {
	return s.Headline
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetMetrics() []*GetInstanceInspectionRunResponseBodyItemsMetrics {
	return s.Metrics
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetResourceSummary() *GetInstanceInspectionRunResponseBodyItemsResourceSummary {
	return s.ResourceSummary
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetSeverity() *string {
	return s.Severity
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionRunResponseBodyItems) GetTimeRange() *GetInstanceInspectionRunResponseBodyItemsTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetDescription(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Description = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetEffectiveConfig(v *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) *GetInstanceInspectionRunResponseBodyItems {
	s.EffectiveConfig = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetErrorCode(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetHeadline(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Headline = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetMetrics(v []*GetInstanceInspectionRunResponseBodyItemsMetrics) *GetInstanceInspectionRunResponseBodyItems {
	s.Metrics = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetName(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetObservedAt(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetResourceCount(v int64) *GetInstanceInspectionRunResponseBodyItems {
	s.ResourceCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetResourceSummary(v *GetInstanceInspectionRunResponseBodyItemsResourceSummary) *GetInstanceInspectionRunResponseBodyItems {
	s.ResourceSummary = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetResourceTypes(v []*string) *GetInstanceInspectionRunResponseBodyItems {
	s.ResourceTypes = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetResult(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetSeverity(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Severity = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetStatus(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetTemplateCode(v string) *GetInstanceInspectionRunResponseBodyItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) SetTimeRange(v *GetInstanceInspectionRunResponseBodyItemsTimeRange) *GetInstanceInspectionRunResponseBodyItems {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItems) Validate() error {
	if s.EffectiveConfig != nil {
		if err := s.EffectiveConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceSummary != nil {
		if err := s.ResourceSummary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyItemsEffectiveConfig struct {
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) GetParameters() interface{} {
	return s.Parameters
}

func (s *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) SetParameters(v interface{}) *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig {
	s.Parameters = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) SetTemplateCode(v string) *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsEffectiveConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyItemsMetrics struct {
	// example:
	//
	// true
	Bad *bool `json:"bad,omitempty" xml:"bad,omitempty"`
	// example:
	//
	// false
	Good *bool `json:"good,omitempty" xml:"good,omitempty"`
	// example:
	//
	// 消费延迟
	K *string `json:"k,omitempty" xml:"k,omitempty"`
	// example:
	//
	// 320 秒
	V *string `json:"v,omitempty" xml:"v,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyItemsMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyItemsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) GetBad() *bool {
	return s.Bad
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) GetGood() *bool {
	return s.Good
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) GetK() *string {
	return s.K
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) GetV() *string {
	return s.V
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) SetBad(v bool) *GetInstanceInspectionRunResponseBodyItemsMetrics {
	s.Bad = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) SetGood(v bool) *GetInstanceInspectionRunResponseBodyItemsMetrics {
	s.Good = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) SetK(v string) *GetInstanceInspectionRunResponseBodyItemsMetrics {
	s.K = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) SetV(v string) *GetInstanceInspectionRunResponseBodyItemsMetrics {
	s.V = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsMetrics) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyItemsResourceSummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyItemsResourceSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyItemsResourceSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyItemsResourceSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionRunResponseBodyItemsResourceSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionRunResponseBodyItemsResourceSummary) SetInfoCount(v int32) *GetInstanceInspectionRunResponseBodyItemsResourceSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsResourceSummary) SetRiskCount(v int32) *GetInstanceInspectionRunResponseBodyItemsResourceSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsResourceSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyItemsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyItemsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyItemsTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyItemsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionRunResponseBodyItemsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionRunResponseBodyItemsTimeRange) SetEndTime(v string) *GetInstanceInspectionRunResponseBodyItemsTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsTimeRange) SetStartTime(v string) *GetInstanceInspectionRunResponseBodyItemsTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyItemsTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyNotification struct {
	// example:
	//
	// inspection-event-001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Sent
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyNotification) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyNotification) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyNotification) GetEventId() *string {
	return s.EventId
}

func (s *GetInstanceInspectionRunResponseBodyNotification) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionRunResponseBodyNotification) SetEventId(v string) *GetInstanceInspectionRunResponseBodyNotification {
	s.EventId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyNotification) SetStatus(v string) *GetInstanceInspectionRunResponseBodyNotification {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyNotification) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyReport struct {
	// example:
	//
	// complete
	Completion *string `json:"completion,omitempty" xml:"completion,omitempty"`
	// example:
	//
	// []
	Coverage []*GetInstanceInspectionRunResponseBodyReportCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Findings []*GetInstanceInspectionRunResponseBodyReportFindings `json:"findings,omitempty" xml:"findings,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-18T08:02:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// <h1>实例巡检诊断报告</h1>
	Html *string `json:"html,omitempty" xml:"html,omitempty"`
	// example:
	//
	// irun-example
	InspectionRunId *string `json:"inspectionRunId,omitempty" xml:"inspectionRunId,omitempty"`
	// example:
	//
	// rmq-cn-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 订单实例
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// # 实例巡检诊断报告
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// example:
	//
	// 消费延迟与容量限制可能相关
	Overview *string `json:"overview,omitempty" xml:"overview,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// irpt-example
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// {}
	ReportSummary *GetInstanceInspectionRunResponseBodyReportReportSummary `json:"reportSummary,omitempty" xml:"reportSummary,omitempty" type:"Struct"`
	// example:
	//
	// inspection-v2.1
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionRunResponseBodyReportTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionRunResponseBodyReport) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReport) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetCompletion() *string {
	return s.Completion
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetCoverage() []*GetInstanceInspectionRunResponseBodyReportCoverage {
	return s.Coverage
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetFindings() []*GetInstanceInspectionRunResponseBodyReportFindings {
	return s.Findings
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetHtml() *string {
	return s.Html
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetInspectionRunId() *string {
	return s.InspectionRunId
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetMarkdown() *string {
	return s.Markdown
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetOverview() *string {
	return s.Overview
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetReportSummary() *GetInstanceInspectionRunResponseBodyReportReportSummary {
	return s.ReportSummary
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetInstanceInspectionRunResponseBodyReport) GetTimeRange() *GetInstanceInspectionRunResponseBodyReportTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetCompletion(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.Completion = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetCoverage(v []*GetInstanceInspectionRunResponseBodyReportCoverage) *GetInstanceInspectionRunResponseBodyReport {
	s.Coverage = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetFindings(v []*GetInstanceInspectionRunResponseBodyReportFindings) *GetInstanceInspectionRunResponseBodyReport {
	s.Findings = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetFinishedAt(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.FinishedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetHtml(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.Html = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetInspectionRunId(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.InspectionRunId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetInstanceId(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetInstanceName(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetMarkdown(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.Markdown = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetOverview(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.Overview = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetRegionId(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetReportId(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.ReportId = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetReportSummary(v *GetInstanceInspectionRunResponseBodyReportReportSummary) *GetInstanceInspectionRunResponseBodyReport {
	s.ReportSummary = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetTemplateVersion(v string) *GetInstanceInspectionRunResponseBodyReport {
	s.TemplateVersion = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) SetTimeRange(v *GetInstanceInspectionRunResponseBodyReportTimeRange) *GetInstanceInspectionRunResponseBodyReport {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReport) Validate() error {
	if s.Coverage != nil {
		for _, item := range s.Coverage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Findings != nil {
		for _, item := range s.Findings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReportSummary != nil {
		if err := s.ReportSummary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyReportCoverage struct {
	// example:
	//
	// Succeeded
	DiagnosisStatus *string `json:"diagnosisStatus,omitempty" xml:"diagnosisStatus,omitempty"`
	// example:
	//
	// 消费健康
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2026-09-01T01:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// example:
	//
	// risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// {"startTime":"2026-09-01T00:00:00Z","endTime":"2026-09-01T01:00:00Z"}
	TimeRange *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionRunResponseBodyReportCoverage) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportCoverage) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetDiagnosisStatus() *string {
	return s.DiagnosisStatus
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) GetTimeRange() *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetDiagnosisStatus(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.DiagnosisStatus = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetName(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetObservedAt(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetResult(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetStatus(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetTemplateCode(v string) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) SetTimeRange(v *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) *GetInstanceInspectionRunResponseBodyReportCoverage {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverage) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyReportCoverageTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) SetEndTime(v string) *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) SetStartTime(v string) *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportCoverageTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyReportFindings struct {
	// example:
	//
	// []
	Alternatives []*string `json:"alternatives,omitempty" xml:"alternatives,omitempty" type:"Repeated"`
	// example:
	//
	// 消费环节
	Boundary *string `json:"boundary,omitempty" xml:"boundary,omitempty"`
	// example:
	//
	// 消费者处理速度可能不足
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// likely
	Confidence *string `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// example:
	//
	// ["消费延迟 60 秒"]
	Evidence []*string `json:"evidence,omitempty" xml:"evidence,omitempty" type:"Repeated"`
	// example:
	//
	// 消息处理延迟升高
	Impact *string `json:"impact,omitempty" xml:"impact,omitempty"`
	// example:
	//
	// 消费组监控
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// []
	MissingEvidence []*string `json:"missingEvidence,omitempty" xml:"missingEvidence,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Observations []*string `json:"observations,omitempty" xml:"observations,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-01T01:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// example:
	//
	// ["检查消费者处理耗时"]
	Recommendations []*string `json:"recommendations,omitempty" xml:"recommendations,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Resource *GetInstanceInspectionRunResponseBodyReportFindingsResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	// example:
	//
	// 消费延迟
	RiskType *string `json:"riskType,omitempty" xml:"riskType,omitempty"`
	// example:
	//
	// script
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// []
	Verification []*string `json:"verification,omitempty" xml:"verification,omitempty" type:"Repeated"`
}

func (s GetInstanceInspectionRunResponseBodyReportFindings) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportFindings) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetAlternatives() []*string {
	return s.Alternatives
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetBoundary() *string {
	return s.Boundary
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetCause() *string {
	return s.Cause
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetConfidence() *string {
	return s.Confidence
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetEvidence() []*string {
	return s.Evidence
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetImpact() *string {
	return s.Impact
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetLocation() *string {
	return s.Location
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetMissingEvidence() []*string {
	return s.MissingEvidence
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetObservations() []*string {
	return s.Observations
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetRecommendations() []*string {
	return s.Recommendations
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetResource() *GetInstanceInspectionRunResponseBodyReportFindingsResource {
	return s.Resource
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetRiskType() *string {
	return s.RiskType
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetSource() *string {
	return s.Source
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetTimeRange() *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) GetVerification() []*string {
	return s.Verification
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetAlternatives(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Alternatives = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetBoundary(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Boundary = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetCause(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Cause = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetConfidence(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Confidence = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetEvidence(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Evidence = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetImpact(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Impact = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetLocation(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Location = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetMissingEvidence(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.MissingEvidence = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetObservations(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Observations = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetObservedAt(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetRecommendations(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Recommendations = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetResource(v *GetInstanceInspectionRunResponseBodyReportFindingsResource) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Resource = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetRiskType(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.RiskType = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetSource(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Source = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetTemplateCode(v string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetTimeRange(v *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) SetVerification(v []*string) *GetInstanceInspectionRunResponseBodyReportFindings {
	s.Verification = v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindings) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionRunResponseBodyReportFindingsResource struct {
	// example:
	//
	// GID-orders
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 订单消费组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// group
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyReportFindingsResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportFindingsResource) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) GetId() *string {
	return s.Id
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) GetType() *string {
	return s.Type
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) SetId(v string) *GetInstanceInspectionRunResponseBodyReportFindingsResource {
	s.Id = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) SetName(v string) *GetInstanceInspectionRunResponseBodyReportFindingsResource {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) SetType(v string) *GetInstanceInspectionRunResponseBodyReportFindingsResource {
	s.Type = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyReportFindingsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) SetEndTime(v string) *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) SetStartTime(v string) *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportFindingsTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyReportReportSummary struct {
	// example:
	//
	// 4
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyReportReportSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportReportSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportReportSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionRunResponseBodyReportReportSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionRunResponseBodyReportReportSummary) SetInfoCount(v int32) *GetInstanceInspectionRunResponseBodyReportReportSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportReportSummary) SetRiskCount(v int32) *GetInstanceInspectionRunResponseBodyReportReportSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportReportSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyReportTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyReportTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyReportTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyReportTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionRunResponseBodyReportTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionRunResponseBodyReportTimeRange) SetEndTime(v string) *GetInstanceInspectionRunResponseBodyReportTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportTimeRange) SetStartTime(v string) *GetInstanceInspectionRunResponseBodyReportTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyReportTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodySummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodySummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionRunResponseBodySummary) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionRunResponseBodySummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionRunResponseBodySummary) SetInfoCount(v int32) *GetInstanceInspectionRunResponseBodySummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodySummary) SetResult(v string) *GetInstanceInspectionRunResponseBodySummary {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodySummary) SetRiskCount(v int32) *GetInstanceInspectionRunResponseBodySummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodySummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionRunResponseBodyTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionRunResponseBodyTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponseBodyTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponseBodyTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionRunResponseBodyTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionRunResponseBodyTimeRange) SetEndTime(v string) *GetInstanceInspectionRunResponseBodyTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyTimeRange) SetStartTime(v string) *GetInstanceInspectionRunResponseBodyTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionRunResponseBodyTimeRange) Validate() error {
	return dara.Validate(s)
}
