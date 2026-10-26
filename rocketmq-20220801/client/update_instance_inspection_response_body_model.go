// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateInstanceInspectionResponseBodyConfig) *UpdateInstanceInspectionResponseBody
	GetConfig() *UpdateInstanceInspectionResponseBodyConfig
	SetConfigured(v bool) *UpdateInstanceInspectionResponseBody
	GetConfigured() *bool
	SetInstanceId(v string) *UpdateInstanceInspectionResponseBody
	GetInstanceId() *string
	SetItems(v []*UpdateInstanceInspectionResponseBodyItems) *UpdateInstanceInspectionResponseBody
	GetItems() []*UpdateInstanceInspectionResponseBodyItems
	SetLastRun(v *UpdateInstanceInspectionResponseBodyLastRun) *UpdateInstanceInspectionResponseBody
	GetLastRun() *UpdateInstanceInspectionResponseBodyLastRun
	SetNextRunAt(v string) *UpdateInstanceInspectionResponseBody
	GetNextRunAt() *string
	SetRegionId(v string) *UpdateInstanceInspectionResponseBody
	GetRegionId() *string
	SetRevision(v int64) *UpdateInstanceInspectionResponseBody
	GetRevision() *int64
}

type UpdateInstanceInspectionResponseBody struct {
	// example:
	//
	// {}
	Config *UpdateInstanceInspectionResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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
	Items []*UpdateInstanceInspectionResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	LastRun *UpdateInstanceInspectionResponseBodyLastRun `json:"lastRun,omitempty" xml:"lastRun,omitempty" type:"Struct"`
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

func (s UpdateInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBody) GetConfig() *UpdateInstanceInspectionResponseBodyConfig {
	return s.Config
}

func (s *UpdateInstanceInspectionResponseBody) GetConfigured() *bool {
	return s.Configured
}

func (s *UpdateInstanceInspectionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceInspectionResponseBody) GetItems() []*UpdateInstanceInspectionResponseBodyItems {
	return s.Items
}

func (s *UpdateInstanceInspectionResponseBody) GetLastRun() *UpdateInstanceInspectionResponseBodyLastRun {
	return s.LastRun
}

func (s *UpdateInstanceInspectionResponseBody) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *UpdateInstanceInspectionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceInspectionResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *UpdateInstanceInspectionResponseBody) SetConfig(v *UpdateInstanceInspectionResponseBodyConfig) *UpdateInstanceInspectionResponseBody {
	s.Config = v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetConfigured(v bool) *UpdateInstanceInspectionResponseBody {
	s.Configured = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetInstanceId(v string) *UpdateInstanceInspectionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetItems(v []*UpdateInstanceInspectionResponseBodyItems) *UpdateInstanceInspectionResponseBody {
	s.Items = v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetLastRun(v *UpdateInstanceInspectionResponseBodyLastRun) *UpdateInstanceInspectionResponseBody {
	s.LastRun = v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetNextRunAt(v string) *UpdateInstanceInspectionResponseBody {
	s.NextRunAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetRegionId(v string) *UpdateInstanceInspectionResponseBody {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) SetRevision(v int64) *UpdateInstanceInspectionResponseBody {
	s.Revision = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBody) Validate() error {
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

type UpdateInstanceInspectionResponseBodyConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*UpdateInstanceInspectionResponseBodyConfigItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *UpdateInstanceInspectionResponseBodyConfigResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *UpdateInstanceInspectionResponseBodyConfigSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s UpdateInstanceInspectionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyConfig) GetItems() []*UpdateInstanceInspectionResponseBodyConfigItems {
	return s.Items
}

func (s *UpdateInstanceInspectionResponseBodyConfig) GetResourceScope() *UpdateInstanceInspectionResponseBodyConfigResourceScope {
	return s.ResourceScope
}

func (s *UpdateInstanceInspectionResponseBodyConfig) GetSchedule() *UpdateInstanceInspectionResponseBodyConfigSchedule {
	return s.Schedule
}

func (s *UpdateInstanceInspectionResponseBodyConfig) SetItems(v []*UpdateInstanceInspectionResponseBodyConfigItems) *UpdateInstanceInspectionResponseBodyConfig {
	s.Items = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfig) SetResourceScope(v *UpdateInstanceInspectionResponseBodyConfigResourceScope) *UpdateInstanceInspectionResponseBodyConfig {
	s.ResourceScope = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfig) SetSchedule(v *UpdateInstanceInspectionResponseBodyConfigSchedule) *UpdateInstanceInspectionResponseBodyConfig {
	s.Schedule = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfig) Validate() error {
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

type UpdateInstanceInspectionResponseBodyConfigItems struct {
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

func (s UpdateInstanceInspectionResponseBodyConfigItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyConfigItems) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyConfigItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *UpdateInstanceInspectionResponseBodyConfigItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyConfigItems) SetParameters(v interface{}) *UpdateInstanceInspectionResponseBodyConfigItems {
	s.Parameters = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigItems) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyConfigItems {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigItems) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyConfigResourceScope struct {
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

func (s UpdateInstanceInspectionResponseBodyConfigResourceScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyConfigResourceScope) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) GetMode() *string {
	return s.Mode
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) SetConsumerGroups(v []*string) *UpdateInstanceInspectionResponseBodyConfigResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) SetMode(v string) *UpdateInstanceInspectionResponseBodyConfigResourceScope {
	s.Mode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) SetTopics(v []*string) *UpdateInstanceInspectionResponseBodyConfigResourceScope {
	s.Topics = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigResourceScope) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyConfigSchedule struct {
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

func (s UpdateInstanceInspectionResponseBodyConfigSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyConfigSchedule) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyConfigSchedule) GetCron() *string {
	return s.Cron
}

func (s *UpdateInstanceInspectionResponseBodyConfigSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateInstanceInspectionResponseBodyConfigSchedule) SetCron(v string) *UpdateInstanceInspectionResponseBodyConfigSchedule {
	s.Cron = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigSchedule) SetEnabled(v bool) *UpdateInstanceInspectionResponseBodyConfigSchedule {
	s.Enabled = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyConfigSchedule) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyItems struct {
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

func (s UpdateInstanceInspectionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceInspectionResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceInspectionResponseBodyItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *UpdateInstanceInspectionResponseBodyItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *UpdateInstanceInspectionResponseBodyItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyItems) SetDescription(v string) *UpdateInstanceInspectionResponseBodyItems {
	s.Description = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyItems) SetName(v string) *UpdateInstanceInspectionResponseBodyItems {
	s.Name = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyItems) SetResourceCount(v int64) *UpdateInstanceInspectionResponseBodyItems {
	s.ResourceCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyItems) SetResourceTypes(v []*string) *UpdateInstanceInspectionResponseBodyItems {
	s.ResourceTypes = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyItems) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyItems {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRun struct {
	// 每次巡检结果由 LLM 总结；包含生成状态及成功后的原文，未生成时为空。
	//
	// example:
	//
	// {}
	AiSummary *UpdateInstanceInspectionResponseBodyLastRunAiSummary `json:"aiSummary,omitempty" xml:"aiSummary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ConfigSnapshot *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot `json:"configSnapshot,omitempty" xml:"configSnapshot,omitempty" type:"Struct"`
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
	Items []*UpdateInstanceInspectionResponseBodyLastRunItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Notification *UpdateInstanceInspectionResponseBodyLastRunNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {}
	Report *UpdateInstanceInspectionResponseBodyLastRunReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
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
	Summary *UpdateInstanceInspectionResponseBodyLastRunSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TimeRange *UpdateInstanceInspectionResponseBodyLastRunTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// manual
	TriggerSource *string `json:"triggerSource,omitempty" xml:"triggerSource,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRun) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRun) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetAiSummary() *UpdateInstanceInspectionResponseBodyLastRunAiSummary {
	return s.AiSummary
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetConfigSnapshot() *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetItems() []*UpdateInstanceInspectionResponseBodyLastRunItems {
	return s.Items
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetNotification() *UpdateInstanceInspectionResponseBodyLastRunNotification {
	return s.Notification
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetReport() *UpdateInstanceInspectionResponseBodyLastRunReport {
	return s.Report
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetRequestedAt() *string {
	return s.RequestedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetRunId() *string {
	return s.RunId
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetSummary() *UpdateInstanceInspectionResponseBodyLastRunSummary {
	return s.Summary
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetTimeRange() *UpdateInstanceInspectionResponseBodyLastRunTimeRange {
	return s.TimeRange
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetAiSummary(v *UpdateInstanceInspectionResponseBodyLastRunAiSummary) *UpdateInstanceInspectionResponseBodyLastRun {
	s.AiSummary = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetConfigSnapshot(v *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) *UpdateInstanceInspectionResponseBodyLastRun {
	s.ConfigSnapshot = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetFinishedAt(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.FinishedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetInstanceId(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetInstanceName(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetItems(v []*UpdateInstanceInspectionResponseBodyLastRunItems) *UpdateInstanceInspectionResponseBodyLastRun {
	s.Items = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetNotification(v *UpdateInstanceInspectionResponseBodyLastRunNotification) *UpdateInstanceInspectionResponseBodyLastRun {
	s.Notification = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetRegionId(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetReport(v *UpdateInstanceInspectionResponseBodyLastRunReport) *UpdateInstanceInspectionResponseBodyLastRun {
	s.Report = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetRequestedAt(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.RequestedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetRunId(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.RunId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetStatus(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.Status = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetSummary(v *UpdateInstanceInspectionResponseBodyLastRunSummary) *UpdateInstanceInspectionResponseBodyLastRun {
	s.Summary = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetTimeRange(v *UpdateInstanceInspectionResponseBodyLastRunTimeRange) *UpdateInstanceInspectionResponseBodyLastRun {
	s.TimeRange = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) SetTriggerSource(v string) *UpdateInstanceInspectionResponseBodyLastRun {
	s.TriggerSource = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRun) Validate() error {
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

type UpdateInstanceInspectionResponseBodyLastRunAiSummary struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunAiSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunAiSummary) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunAiSummary) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceInspectionResponseBodyLastRunAiSummary) GetText() *string {
	return s.Text
}

func (s *UpdateInstanceInspectionResponseBodyLastRunAiSummary) SetStatus(v string) *UpdateInstanceInspectionResponseBodyLastRunAiSummary {
	s.Status = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunAiSummary) SetText(v string) *UpdateInstanceInspectionResponseBodyLastRunAiSummary {
	s.Text = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunAiSummary) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetItems() []*UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	return s.Items
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetResourceScope() *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	return s.ResourceScope
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) GetSchedule() *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	return s.Schedule
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetItems(v []*UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Items = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetResourceScope(v *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.ResourceScope = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) SetSchedule(v *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Schedule = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshot) Validate() error {
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

type UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetParameters(v interface{}) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.Parameters = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotItems) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetMode() *string {
	return s.Mode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetConsumerGroups(v []*string) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetMode(v string) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Mode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetTopics(v []*string) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Topics = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetCron() *string {
	return s.Cron
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetCron(v string) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Cron = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetEnabled(v bool) *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Enabled = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}}
	EffectiveConfig *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig `json:"effectiveConfig,omitempty" xml:"effectiveConfig,omitempty" type:"Struct"`
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
	Metrics []*UpdateInstanceInspectionResponseBodyLastRunItemsMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
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
	ResourceSummary *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary `json:"resourceSummary,omitempty" xml:"resourceSummary,omitempty" type:"Struct"`
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
	TimeRange *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunItems) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetEffectiveConfig() *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	return s.EffectiveConfig
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetHeadline() *string {
	return s.Headline
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetMetrics() []*UpdateInstanceInspectionResponseBodyLastRunItemsMetrics {
	return s.Metrics
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetResourceSummary() *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	return s.ResourceSummary
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetResult() *string {
	return s.Result
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) GetTimeRange() *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	return s.TimeRange
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetDescription(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Description = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetEffectiveConfig(v *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.EffectiveConfig = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetErrorCode(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetHeadline(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Headline = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetMetrics(v []*UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Metrics = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetName(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Name = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetObservedAt(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.ObservedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetResourceCount(v int64) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetResourceSummary(v *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceSummary = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetResourceTypes(v []*string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.ResourceTypes = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetResult(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Result = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetSeverity(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Severity = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetStatus(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.Status = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) SetTimeRange(v *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) *UpdateInstanceInspectionResponseBodyLastRunItems {
	s.TimeRange = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItems) Validate() error {
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

type UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig struct {
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetParameters() interface{} {
	return s.Parameters
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetParameters(v interface{}) *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.Parameters = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunItemsMetrics struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) GetBad() *bool {
	return s.Bad
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) GetGood() *bool {
	return s.Good
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) GetK() *string {
	return s.K
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) GetV() *string {
	return s.V
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) SetBad(v bool) *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Bad = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) SetGood(v bool) *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Good = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) SetK(v string) *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.K = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) SetV(v string) *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.V = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsMetrics) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetInfoCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.InfoCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetRiskCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.RiskCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsResourceSummary) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) SetEndTime(v string) *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) SetStartTime(v string) *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunItemsTimeRange) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunNotification struct {
	// example:
	//
	// inspection-event-001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Sent
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunNotification) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunNotification) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunNotification) GetEventId() *string {
	return s.EventId
}

func (s *UpdateInstanceInspectionResponseBodyLastRunNotification) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceInspectionResponseBodyLastRunNotification) SetEventId(v string) *UpdateInstanceInspectionResponseBodyLastRunNotification {
	s.EventId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunNotification) SetStatus(v string) *UpdateInstanceInspectionResponseBodyLastRunNotification {
	s.Status = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunNotification) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunReport struct {
	// example:
	//
	// complete
	Completion *string `json:"completion,omitempty" xml:"completion,omitempty"`
	// example:
	//
	// []
	Coverage []*UpdateInstanceInspectionResponseBodyLastRunReportCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Findings []*UpdateInstanceInspectionResponseBodyLastRunReportFindings `json:"findings,omitempty" xml:"findings,omitempty" type:"Repeated"`
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
	ReportSummary *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary `json:"reportSummary,omitempty" xml:"reportSummary,omitempty" type:"Struct"`
	// example:
	//
	// inspection-v2.1
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// example:
	//
	// {}
	TimeRange *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReport) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReport) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetCompletion() *string {
	return s.Completion
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetCoverage() []*UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	return s.Coverage
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetFindings() []*UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	return s.Findings
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetHtml() *string {
	return s.Html
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetInspectionRunId() *string {
	return s.InspectionRunId
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetMarkdown() *string {
	return s.Markdown
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetOverview() *string {
	return s.Overview
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetReportId() *string {
	return s.ReportId
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetReportSummary() *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary {
	return s.ReportSummary
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) GetTimeRange() *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange {
	return s.TimeRange
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetCompletion(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Completion = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetCoverage(v []*UpdateInstanceInspectionResponseBodyLastRunReportCoverage) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Coverage = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetFindings(v []*UpdateInstanceInspectionResponseBodyLastRunReportFindings) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Findings = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetFinishedAt(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.FinishedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetHtml(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Html = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetInspectionRunId(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.InspectionRunId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetInstanceId(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetInstanceName(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetMarkdown(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Markdown = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetOverview(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.Overview = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetRegionId(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetReportId(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.ReportId = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetReportSummary(v *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.ReportSummary = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetTemplateVersion(v string) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) SetTimeRange(v *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) *UpdateInstanceInspectionResponseBodyLastRunReport {
	s.TimeRange = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReport) Validate() error {
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

type UpdateInstanceInspectionResponseBodyLastRunReportCoverage struct {
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
	TimeRange *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportCoverage) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetDiagnosisStatus() *string {
	return s.DiagnosisStatus
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetResult() *string {
	return s.Result
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) GetTimeRange() *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	return s.TimeRange
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetDiagnosisStatus(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.DiagnosisStatus = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetName(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Name = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetObservedAt(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.ObservedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetResult(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Result = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetStatus(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Status = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) SetTimeRange(v *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) *UpdateInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TimeRange = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverage) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetEndTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetStartTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunReportFindings struct {
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
	Resource *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
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
	TimeRange *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// []
	Verification []*string `json:"verification,omitempty" xml:"verification,omitempty" type:"Repeated"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindings) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindings) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetAlternatives() []*string {
	return s.Alternatives
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetBoundary() *string {
	return s.Boundary
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetCause() *string {
	return s.Cause
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetConfidence() *string {
	return s.Confidence
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetEvidence() []*string {
	return s.Evidence
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetImpact() *string {
	return s.Impact
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetLocation() *string {
	return s.Location
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetMissingEvidence() []*string {
	return s.MissingEvidence
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetObservations() []*string {
	return s.Observations
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetRecommendations() []*string {
	return s.Recommendations
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetResource() *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	return s.Resource
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetRiskType() *string {
	return s.RiskType
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetSource() *string {
	return s.Source
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetTimeRange() *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	return s.TimeRange
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) GetVerification() []*string {
	return s.Verification
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetAlternatives(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Alternatives = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetBoundary(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Boundary = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetCause(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Cause = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetConfidence(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Confidence = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetEvidence(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Evidence = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetImpact(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Impact = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetLocation(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Location = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetMissingEvidence(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.MissingEvidence = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetObservations(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Observations = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetObservedAt(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.ObservedAt = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetRecommendations(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Recommendations = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetResource(v *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Resource = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetRiskType(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.RiskType = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetSource(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Source = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetTemplateCode(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetTimeRange(v *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.TimeRange = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) SetVerification(v []*string) *UpdateInstanceInspectionResponseBodyLastRunReportFindings {
	s.Verification = v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindings) Validate() error {
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

type UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetId() *string {
	return s.Id
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) GetType() *string {
	return s.Type
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetId(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Id = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetName(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Name = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) SetType(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Type = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsResource) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetEndTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetStartTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunReportReportSummary struct {
	// example:
	//
	// 4
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) SetInfoCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.InfoCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) SetRiskCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.RiskCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportReportSummary) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunReportTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) SetEndTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) SetStartTime(v string) *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunReportTimeRange) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunSummary struct {
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

func (s UpdateInstanceInspectionResponseBodyLastRunSummary) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunSummary) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) GetResult() *string {
	return s.Result
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) SetInfoCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunSummary {
	s.InfoCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) SetResult(v string) *UpdateInstanceInspectionResponseBodyLastRunSummary {
	s.Result = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) SetRiskCount(v int32) *UpdateInstanceInspectionResponseBodyLastRunSummary {
	s.RiskCount = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunSummary) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionResponseBodyLastRunTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s UpdateInstanceInspectionResponseBodyLastRunTimeRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionResponseBodyLastRunTimeRange) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionResponseBodyLastRunTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateInstanceInspectionResponseBodyLastRunTimeRange) SetEndTime(v string) *UpdateInstanceInspectionResponseBodyLastRunTimeRange {
	s.EndTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunTimeRange) SetStartTime(v string) *UpdateInstanceInspectionResponseBodyLastRunTimeRange {
	s.StartTime = &v
	return s
}

func (s *UpdateInstanceInspectionResponseBodyLastRunTimeRange) Validate() error {
	return dara.Validate(s)
}
