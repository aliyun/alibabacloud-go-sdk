// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *CreateInstanceInspectionResponseBodyConfig) *CreateInstanceInspectionResponseBody
	GetConfig() *CreateInstanceInspectionResponseBodyConfig
	SetConfigured(v bool) *CreateInstanceInspectionResponseBody
	GetConfigured() *bool
	SetInstanceId(v string) *CreateInstanceInspectionResponseBody
	GetInstanceId() *string
	SetItems(v []*CreateInstanceInspectionResponseBodyItems) *CreateInstanceInspectionResponseBody
	GetItems() []*CreateInstanceInspectionResponseBodyItems
	SetLastRun(v *CreateInstanceInspectionResponseBodyLastRun) *CreateInstanceInspectionResponseBody
	GetLastRun() *CreateInstanceInspectionResponseBodyLastRun
	SetNextRunAt(v string) *CreateInstanceInspectionResponseBody
	GetNextRunAt() *string
	SetRegionId(v string) *CreateInstanceInspectionResponseBody
	GetRegionId() *string
	SetRevision(v int64) *CreateInstanceInspectionResponseBody
	GetRevision() *int64
}

type CreateInstanceInspectionResponseBody struct {
	// example:
	//
	// {}
	Config *CreateInstanceInspectionResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// true
	Configured *bool `json:"configured,omitempty" xml:"configured,omitempty"`
	// example:
	//
	// rmq-cn-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// []
	Items []*CreateInstanceInspectionResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	LastRun *CreateInstanceInspectionResponseBodyLastRun `json:"lastRun,omitempty" xml:"lastRun,omitempty" type:"Struct"`
	// example:
	//
	// 2026-09-09T03:00:00+08:00
	NextRunAt *string `json:"nextRunAt,omitempty" xml:"nextRunAt,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
}

func (s CreateInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBody) GetConfig() *CreateInstanceInspectionResponseBodyConfig {
	return s.Config
}

func (s *CreateInstanceInspectionResponseBody) GetConfigured() *bool {
	return s.Configured
}

func (s *CreateInstanceInspectionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceInspectionResponseBody) GetItems() []*CreateInstanceInspectionResponseBodyItems {
	return s.Items
}

func (s *CreateInstanceInspectionResponseBody) GetLastRun() *CreateInstanceInspectionResponseBodyLastRun {
	return s.LastRun
}

func (s *CreateInstanceInspectionResponseBody) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *CreateInstanceInspectionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceInspectionResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *CreateInstanceInspectionResponseBody) SetConfig(v *CreateInstanceInspectionResponseBodyConfig) *CreateInstanceInspectionResponseBody {
	s.Config = v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetConfigured(v bool) *CreateInstanceInspectionResponseBody {
	s.Configured = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetInstanceId(v string) *CreateInstanceInspectionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetItems(v []*CreateInstanceInspectionResponseBodyItems) *CreateInstanceInspectionResponseBody {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetLastRun(v *CreateInstanceInspectionResponseBodyLastRun) *CreateInstanceInspectionResponseBody {
	s.LastRun = v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetNextRunAt(v string) *CreateInstanceInspectionResponseBody {
	s.NextRunAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetRegionId(v string) *CreateInstanceInspectionResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) SetRevision(v int64) *CreateInstanceInspectionResponseBody {
	s.Revision = &v
	return s
}

func (s *CreateInstanceInspectionResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
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
	if s.LastRun != nil {
		if err := s.LastRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceInspectionResponseBodyConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*CreateInstanceInspectionResponseBodyConfigItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *CreateInstanceInspectionResponseBodyConfigResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *CreateInstanceInspectionResponseBodyConfigSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s CreateInstanceInspectionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyConfig) GetItems() []*CreateInstanceInspectionResponseBodyConfigItems {
	return s.Items
}

func (s *CreateInstanceInspectionResponseBodyConfig) GetResourceScope() *CreateInstanceInspectionResponseBodyConfigResourceScope {
	return s.ResourceScope
}

func (s *CreateInstanceInspectionResponseBodyConfig) GetSchedule() *CreateInstanceInspectionResponseBodyConfigSchedule {
	return s.Schedule
}

func (s *CreateInstanceInspectionResponseBodyConfig) SetItems(v []*CreateInstanceInspectionResponseBodyConfigItems) *CreateInstanceInspectionResponseBodyConfig {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfig) SetResourceScope(v *CreateInstanceInspectionResponseBodyConfigResourceScope) *CreateInstanceInspectionResponseBodyConfig {
	s.ResourceScope = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfig) SetSchedule(v *CreateInstanceInspectionResponseBodyConfigSchedule) *CreateInstanceInspectionResponseBodyConfig {
	s.Schedule = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfig) Validate() error {
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

type CreateInstanceInspectionResponseBodyConfigItems struct {
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

func (s CreateInstanceInspectionResponseBodyConfigItems) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyConfigItems) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyConfigItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *CreateInstanceInspectionResponseBodyConfigItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyConfigItems) SetParameters(v interface{}) *CreateInstanceInspectionResponseBodyConfigItems {
	s.Parameters = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigItems) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyConfigItems {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigItems) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyConfigResourceScope struct {
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

func (s CreateInstanceInspectionResponseBodyConfigResourceScope) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyConfigResourceScope) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) GetMode() *string {
	return s.Mode
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) SetConsumerGroups(v []*string) *CreateInstanceInspectionResponseBodyConfigResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) SetMode(v string) *CreateInstanceInspectionResponseBodyConfigResourceScope {
	s.Mode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) SetTopics(v []*string) *CreateInstanceInspectionResponseBodyConfigResourceScope {
	s.Topics = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigResourceScope) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyConfigSchedule struct {
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

func (s CreateInstanceInspectionResponseBodyConfigSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyConfigSchedule) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyConfigSchedule) GetCron() *string {
	return s.Cron
}

func (s *CreateInstanceInspectionResponseBodyConfigSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateInstanceInspectionResponseBodyConfigSchedule) SetCron(v string) *CreateInstanceInspectionResponseBodyConfigSchedule {
	s.Cron = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigSchedule) SetEnabled(v bool) *CreateInstanceInspectionResponseBodyConfigSchedule {
	s.Enabled = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyConfigSchedule) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 消费健康巡检
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 3
	ResourceCount *int64 `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
	// example:
	//
	// ["ConsumerGroup"]
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceInspectionResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *CreateInstanceInspectionResponseBodyItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *CreateInstanceInspectionResponseBodyItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *CreateInstanceInspectionResponseBodyItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyItems) SetDescription(v string) *CreateInstanceInspectionResponseBodyItems {
	s.Description = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyItems) SetName(v string) *CreateInstanceInspectionResponseBodyItems {
	s.Name = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyItems) SetResourceCount(v int64) *CreateInstanceInspectionResponseBodyItems {
	s.ResourceCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyItems) SetResourceTypes(v []*string) *CreateInstanceInspectionResponseBodyItems {
	s.ResourceTypes = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyItems) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyItems {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRun struct {
	// 每次巡检结果由 LLM 总结；包含生成状态及成功后的原文，未生成时为空。
	//
	// example:
	//
	// {}
	AiSummary *CreateInstanceInspectionResponseBodyLastRunAiSummary `json:"aiSummary,omitempty" xml:"aiSummary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ConfigSnapshot *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot `json:"configSnapshot,omitempty" xml:"configSnapshot,omitempty" type:"Struct"`
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
	Items []*CreateInstanceInspectionResponseBodyLastRunItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Notification *CreateInstanceInspectionResponseBodyLastRunNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {}
	Report *CreateInstanceInspectionResponseBodyLastRunReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
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
	Summary *CreateInstanceInspectionResponseBodyLastRunSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TimeRange *CreateInstanceInspectionResponseBodyLastRunTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// manual
	TriggerSource *string `json:"triggerSource,omitempty" xml:"triggerSource,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRun) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRun) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetAiSummary() *CreateInstanceInspectionResponseBodyLastRunAiSummary {
	return s.AiSummary
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetConfigSnapshot() *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetItems() []*CreateInstanceInspectionResponseBodyLastRunItems {
	return s.Items
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetNotification() *CreateInstanceInspectionResponseBodyLastRunNotification {
	return s.Notification
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetReport() *CreateInstanceInspectionResponseBodyLastRunReport {
	return s.Report
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetRequestedAt() *string {
	return s.RequestedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetRunId() *string {
	return s.RunId
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetStatus() *string {
	return s.Status
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetSummary() *CreateInstanceInspectionResponseBodyLastRunSummary {
	return s.Summary
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetTimeRange() *CreateInstanceInspectionResponseBodyLastRunTimeRange {
	return s.TimeRange
}

func (s *CreateInstanceInspectionResponseBodyLastRun) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetAiSummary(v *CreateInstanceInspectionResponseBodyLastRunAiSummary) *CreateInstanceInspectionResponseBodyLastRun {
	s.AiSummary = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetConfigSnapshot(v *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) *CreateInstanceInspectionResponseBodyLastRun {
	s.ConfigSnapshot = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetFinishedAt(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.FinishedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetInstanceId(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetInstanceName(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetItems(v []*CreateInstanceInspectionResponseBodyLastRunItems) *CreateInstanceInspectionResponseBodyLastRun {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetNotification(v *CreateInstanceInspectionResponseBodyLastRunNotification) *CreateInstanceInspectionResponseBodyLastRun {
	s.Notification = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetRegionId(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetReport(v *CreateInstanceInspectionResponseBodyLastRunReport) *CreateInstanceInspectionResponseBodyLastRun {
	s.Report = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetRequestedAt(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.RequestedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetRunId(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.RunId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetStatus(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.Status = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetSummary(v *CreateInstanceInspectionResponseBodyLastRunSummary) *CreateInstanceInspectionResponseBodyLastRun {
	s.Summary = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetTimeRange(v *CreateInstanceInspectionResponseBodyLastRunTimeRange) *CreateInstanceInspectionResponseBodyLastRun {
	s.TimeRange = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) SetTriggerSource(v string) *CreateInstanceInspectionResponseBodyLastRun {
	s.TriggerSource = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRun) Validate() error {
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

type CreateInstanceInspectionResponseBodyLastRunAiSummary struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunAiSummary) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunAiSummary) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunAiSummary) GetStatus() *string {
	return s.Status
}

func (s *CreateInstanceInspectionResponseBodyLastRunAiSummary) GetText() *string {
	return s.Text
}

func (s *CreateInstanceInspectionResponseBodyLastRunAiSummary) SetStatus(v string) *CreateInstanceInspectionResponseBodyLastRunAiSummary {
	s.Status = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunAiSummary) SetText(v string) *CreateInstanceInspectionResponseBodyLastRunAiSummary {
	s.Text = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunAiSummary) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunConfigSnapshot struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetItems() []*CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	return s.Items
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetResourceScope() *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	return s.ResourceScope
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetSchedule() *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	return s.Schedule
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetItems(v []*CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetResourceScope(v *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.ResourceScope = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetSchedule(v *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Schedule = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshot) Validate() error {
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

type CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetParameters(v interface{}) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.Parameters = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetMode() *string {
	return s.Mode
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetConsumerGroups(v []*string) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetMode(v string) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Mode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetTopics(v []*string) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Topics = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetCron() *string {
	return s.Cron
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetCron(v string) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Cron = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetEnabled(v bool) *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Enabled = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}}
	EffectiveConfig *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig `json:"effectiveConfig,omitempty" xml:"effectiveConfig,omitempty" type:"Struct"`
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
	Metrics []*CreateInstanceInspectionResponseBodyLastRunItemsMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
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
	ResourceSummary *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary `json:"resourceSummary,omitempty" xml:"resourceSummary,omitempty" type:"Struct"`
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
	TimeRange *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s CreateInstanceInspectionResponseBodyLastRunItems) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunItems) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetEffectiveConfig() *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	return s.EffectiveConfig
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetHeadline() *string {
	return s.Headline
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetMetrics() []*CreateInstanceInspectionResponseBodyLastRunItemsMetrics {
	return s.Metrics
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetName() *string {
	return s.Name
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetResourceSummary() *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	return s.ResourceSummary
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetResult() *string {
	return s.Result
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetSeverity() *string {
	return s.Severity
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetStatus() *string {
	return s.Status
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) GetTimeRange() *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	return s.TimeRange
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetDescription(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Description = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetEffectiveConfig(v *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.EffectiveConfig = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetErrorCode(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetHeadline(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Headline = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetMetrics(v []*CreateInstanceInspectionResponseBodyLastRunItemsMetrics) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Metrics = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetName(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Name = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetObservedAt(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.ObservedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetResourceCount(v int64) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetResourceSummary(v *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceSummary = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetResourceTypes(v []*string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceTypes = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetResult(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Result = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetSeverity(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Severity = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetStatus(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.Status = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) SetTimeRange(v *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) *CreateInstanceInspectionResponseBodyLastRunItems {
	s.TimeRange = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItems) Validate() error {
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

type CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig struct {
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetParameters() interface{} {
	return s.Parameters
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetParameters(v interface{}) *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.Parameters = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunItemsMetrics struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunItemsMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsMetrics) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) GetBad() *bool {
	return s.Bad
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) GetGood() *bool {
	return s.Good
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) GetK() *string {
	return s.K
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) GetV() *string {
	return s.V
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) SetBad(v bool) *CreateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Bad = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) SetGood(v bool) *CreateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Good = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) SetK(v string) *CreateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.K = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) SetV(v string) *CreateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.V = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsMetrics) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetInfoCount(v int32) *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.InfoCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetRiskCount(v int32) *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.RiskCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsResourceSummary) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunItemsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) SetEndTime(v string) *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) SetStartTime(v string) *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunItemsTimeRange) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunNotification struct {
	// example:
	//
	// inspection-event-001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Sent
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunNotification) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunNotification) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunNotification) GetEventId() *string {
	return s.EventId
}

func (s *CreateInstanceInspectionResponseBodyLastRunNotification) GetStatus() *string {
	return s.Status
}

func (s *CreateInstanceInspectionResponseBodyLastRunNotification) SetEventId(v string) *CreateInstanceInspectionResponseBodyLastRunNotification {
	s.EventId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunNotification) SetStatus(v string) *CreateInstanceInspectionResponseBodyLastRunNotification {
	s.Status = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunNotification) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunReport struct {
	// example:
	//
	// complete
	Completion *string `json:"completion,omitempty" xml:"completion,omitempty"`
	// example:
	//
	// []
	Coverage []*CreateInstanceInspectionResponseBodyLastRunReportCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Findings []*CreateInstanceInspectionResponseBodyLastRunReportFindings `json:"findings,omitempty" xml:"findings,omitempty" type:"Repeated"`
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
	ReportSummary *CreateInstanceInspectionResponseBodyLastRunReportReportSummary `json:"reportSummary,omitempty" xml:"reportSummary,omitempty" type:"Struct"`
	// example:
	//
	// inspection-v2.1
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// example:
	//
	// {}
	TimeRange *CreateInstanceInspectionResponseBodyLastRunReportTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReport) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReport) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetCompletion() *string {
	return s.Completion
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetCoverage() []*CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	return s.Coverage
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetFindings() []*CreateInstanceInspectionResponseBodyLastRunReportFindings {
	return s.Findings
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetHtml() *string {
	return s.Html
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetInspectionRunId() *string {
	return s.InspectionRunId
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetMarkdown() *string {
	return s.Markdown
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetOverview() *string {
	return s.Overview
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetReportId() *string {
	return s.ReportId
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetReportSummary() *CreateInstanceInspectionResponseBodyLastRunReportReportSummary {
	return s.ReportSummary
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) GetTimeRange() *CreateInstanceInspectionResponseBodyLastRunReportTimeRange {
	return s.TimeRange
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetCompletion(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Completion = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetCoverage(v []*CreateInstanceInspectionResponseBodyLastRunReportCoverage) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Coverage = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetFindings(v []*CreateInstanceInspectionResponseBodyLastRunReportFindings) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Findings = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetFinishedAt(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.FinishedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetHtml(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Html = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetInspectionRunId(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.InspectionRunId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetInstanceId(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetInstanceName(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetMarkdown(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Markdown = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetOverview(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.Overview = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetRegionId(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetReportId(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.ReportId = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetReportSummary(v *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.ReportSummary = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetTemplateVersion(v string) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.TemplateVersion = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) SetTimeRange(v *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) *CreateInstanceInspectionResponseBodyLastRunReport {
	s.TimeRange = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReport) Validate() error {
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

type CreateInstanceInspectionResponseBodyLastRunReportCoverage struct {
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
	TimeRange *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportCoverage) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportCoverage) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetDiagnosisStatus() *string {
	return s.DiagnosisStatus
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetName() *string {
	return s.Name
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetResult() *string {
	return s.Result
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetStatus() *string {
	return s.Status
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) GetTimeRange() *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	return s.TimeRange
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetDiagnosisStatus(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.DiagnosisStatus = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetName(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Name = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetObservedAt(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.ObservedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetResult(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Result = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetStatus(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Status = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) SetTimeRange(v *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) *CreateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TimeRange = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverage) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetEndTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetStartTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunReportFindings struct {
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
	Resource *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
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
	TimeRange *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// []
	Verification []*string `json:"verification,omitempty" xml:"verification,omitempty" type:"Repeated"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportFindings) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportFindings) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetAlternatives() []*string {
	return s.Alternatives
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetBoundary() *string {
	return s.Boundary
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetCause() *string {
	return s.Cause
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetConfidence() *string {
	return s.Confidence
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetEvidence() []*string {
	return s.Evidence
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetImpact() *string {
	return s.Impact
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetLocation() *string {
	return s.Location
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetMissingEvidence() []*string {
	return s.MissingEvidence
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetObservations() []*string {
	return s.Observations
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetRecommendations() []*string {
	return s.Recommendations
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetResource() *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	return s.Resource
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetRiskType() *string {
	return s.RiskType
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetSource() *string {
	return s.Source
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetTimeRange() *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	return s.TimeRange
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) GetVerification() []*string {
	return s.Verification
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetAlternatives(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Alternatives = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetBoundary(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Boundary = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetCause(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Cause = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetConfidence(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Confidence = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetEvidence(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Evidence = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetImpact(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Impact = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetLocation(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Location = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetMissingEvidence(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.MissingEvidence = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetObservations(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Observations = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetObservedAt(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.ObservedAt = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetRecommendations(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Recommendations = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetResource(v *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Resource = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetRiskType(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.RiskType = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetSource(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Source = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetTemplateCode(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetTimeRange(v *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.TimeRange = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) SetVerification(v []*string) *CreateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Verification = v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindings) Validate() error {
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

type CreateInstanceInspectionResponseBodyLastRunReportFindingsResource struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetId() *string {
	return s.Id
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetName() *string {
	return s.Name
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetType() *string {
	return s.Type
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetId(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Id = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetName(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Name = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetType(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Type = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsResource) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetEndTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetStartTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunReportReportSummary struct {
	// example:
	//
	// 4
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportReportSummary) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportReportSummary) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) SetInfoCount(v int32) *CreateInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.InfoCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) SetRiskCount(v int32) *CreateInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.RiskCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportReportSummary) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunReportTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunReportTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunReportTimeRange) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) SetEndTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) SetStartTime(v string) *CreateInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunReportTimeRange) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunSummary struct {
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

func (s CreateInstanceInspectionResponseBodyLastRunSummary) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunSummary) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) GetResult() *string {
	return s.Result
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) SetInfoCount(v int32) *CreateInstanceInspectionResponseBodyLastRunSummary {
	s.InfoCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) SetResult(v string) *CreateInstanceInspectionResponseBodyLastRunSummary {
	s.Result = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) SetRiskCount(v int32) *CreateInstanceInspectionResponseBodyLastRunSummary {
	s.RiskCount = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunSummary) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionResponseBodyLastRunTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateInstanceInspectionResponseBodyLastRunTimeRange) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponseBodyLastRunTimeRange) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponseBodyLastRunTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInstanceInspectionResponseBodyLastRunTimeRange) SetEndTime(v string) *CreateInstanceInspectionResponseBodyLastRunTimeRange {
	s.EndTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunTimeRange) SetStartTime(v string) *CreateInstanceInspectionResponseBodyLastRunTimeRange {
	s.StartTime = &v
	return s
}

func (s *CreateInstanceInspectionResponseBodyLastRunTimeRange) Validate() error {
	return dara.Validate(s)
}
