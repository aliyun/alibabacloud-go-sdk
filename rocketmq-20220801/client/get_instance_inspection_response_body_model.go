// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *GetInstanceInspectionResponseBodyConfig) *GetInstanceInspectionResponseBody
	GetConfig() *GetInstanceInspectionResponseBodyConfig
	SetConfigured(v bool) *GetInstanceInspectionResponseBody
	GetConfigured() *bool
	SetInstanceId(v string) *GetInstanceInspectionResponseBody
	GetInstanceId() *string
	SetItems(v []*GetInstanceInspectionResponseBodyItems) *GetInstanceInspectionResponseBody
	GetItems() []*GetInstanceInspectionResponseBodyItems
	SetLastRun(v *GetInstanceInspectionResponseBodyLastRun) *GetInstanceInspectionResponseBody
	GetLastRun() *GetInstanceInspectionResponseBodyLastRun
	SetNextRunAt(v string) *GetInstanceInspectionResponseBody
	GetNextRunAt() *string
	SetRegionId(v string) *GetInstanceInspectionResponseBody
	GetRegionId() *string
	SetRevision(v int64) *GetInstanceInspectionResponseBody
	GetRevision() *int64
}

type GetInstanceInspectionResponseBody struct {
	// example:
	//
	// {}
	Config *GetInstanceInspectionResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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
	Items []*GetInstanceInspectionResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	LastRun *GetInstanceInspectionResponseBodyLastRun `json:"lastRun,omitempty" xml:"lastRun,omitempty" type:"Struct"`
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

func (s GetInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBody) GetConfig() *GetInstanceInspectionResponseBodyConfig {
	return s.Config
}

func (s *GetInstanceInspectionResponseBody) GetConfigured() *bool {
	return s.Configured
}

func (s *GetInstanceInspectionResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionResponseBody) GetItems() []*GetInstanceInspectionResponseBodyItems {
	return s.Items
}

func (s *GetInstanceInspectionResponseBody) GetLastRun() *GetInstanceInspectionResponseBodyLastRun {
	return s.LastRun
}

func (s *GetInstanceInspectionResponseBody) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *GetInstanceInspectionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *GetInstanceInspectionResponseBody) SetConfig(v *GetInstanceInspectionResponseBodyConfig) *GetInstanceInspectionResponseBody {
	s.Config = v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetConfigured(v bool) *GetInstanceInspectionResponseBody {
	s.Configured = &v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetInstanceId(v string) *GetInstanceInspectionResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetItems(v []*GetInstanceInspectionResponseBodyItems) *GetInstanceInspectionResponseBody {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetLastRun(v *GetInstanceInspectionResponseBodyLastRun) *GetInstanceInspectionResponseBody {
	s.LastRun = v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetNextRunAt(v string) *GetInstanceInspectionResponseBody {
	s.NextRunAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetRegionId(v string) *GetInstanceInspectionResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionResponseBody) SetRevision(v int64) *GetInstanceInspectionResponseBody {
	s.Revision = &v
	return s
}

func (s *GetInstanceInspectionResponseBody) Validate() error {
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

type GetInstanceInspectionResponseBodyConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*GetInstanceInspectionResponseBodyConfigItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *GetInstanceInspectionResponseBodyConfigResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *GetInstanceInspectionResponseBodyConfigSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyConfig) GetItems() []*GetInstanceInspectionResponseBodyConfigItems {
	return s.Items
}

func (s *GetInstanceInspectionResponseBodyConfig) GetResourceScope() *GetInstanceInspectionResponseBodyConfigResourceScope {
	return s.ResourceScope
}

func (s *GetInstanceInspectionResponseBodyConfig) GetSchedule() *GetInstanceInspectionResponseBodyConfigSchedule {
	return s.Schedule
}

func (s *GetInstanceInspectionResponseBodyConfig) SetItems(v []*GetInstanceInspectionResponseBodyConfigItems) *GetInstanceInspectionResponseBodyConfig {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfig) SetResourceScope(v *GetInstanceInspectionResponseBodyConfigResourceScope) *GetInstanceInspectionResponseBodyConfig {
	s.ResourceScope = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfig) SetSchedule(v *GetInstanceInspectionResponseBodyConfigSchedule) *GetInstanceInspectionResponseBodyConfig {
	s.Schedule = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfig) Validate() error {
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

type GetInstanceInspectionResponseBodyConfigItems struct {
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

func (s GetInstanceInspectionResponseBodyConfigItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyConfigItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyConfigItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *GetInstanceInspectionResponseBodyConfigItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyConfigItems) SetParameters(v interface{}) *GetInstanceInspectionResponseBodyConfigItems {
	s.Parameters = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigItems) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyConfigItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigItems) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyConfigResourceScope struct {
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

func (s GetInstanceInspectionResponseBodyConfigResourceScope) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyConfigResourceScope) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) GetMode() *string {
	return s.Mode
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) SetConsumerGroups(v []*string) *GetInstanceInspectionResponseBodyConfigResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) SetMode(v string) *GetInstanceInspectionResponseBodyConfigResourceScope {
	s.Mode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) SetTopics(v []*string) *GetInstanceInspectionResponseBodyConfigResourceScope {
	s.Topics = v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigResourceScope) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyConfigSchedule struct {
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

func (s GetInstanceInspectionResponseBodyConfigSchedule) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyConfigSchedule) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyConfigSchedule) GetCron() *string {
	return s.Cron
}

func (s *GetInstanceInspectionResponseBodyConfigSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetInstanceInspectionResponseBodyConfigSchedule) SetCron(v string) *GetInstanceInspectionResponseBodyConfigSchedule {
	s.Cron = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigSchedule) SetEnabled(v bool) *GetInstanceInspectionResponseBodyConfigSchedule {
	s.Enabled = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyConfigSchedule) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyItems struct {
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

func (s GetInstanceInspectionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceInspectionResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionResponseBodyItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetInstanceInspectionResponseBodyItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetInstanceInspectionResponseBodyItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyItems) SetDescription(v string) *GetInstanceInspectionResponseBodyItems {
	s.Description = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyItems) SetName(v string) *GetInstanceInspectionResponseBodyItems {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyItems) SetResourceCount(v int64) *GetInstanceInspectionResponseBodyItems {
	s.ResourceCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyItems) SetResourceTypes(v []*string) *GetInstanceInspectionResponseBodyItems {
	s.ResourceTypes = v
	return s
}

func (s *GetInstanceInspectionResponseBodyItems) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRun struct {
	// 每次巡检结果由 LLM 总结；包含生成状态及成功后的原文，未生成时为空。
	//
	// example:
	//
	// {}
	AiSummary *GetInstanceInspectionResponseBodyLastRunAiSummary `json:"aiSummary,omitempty" xml:"aiSummary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	ConfigSnapshot *GetInstanceInspectionResponseBodyLastRunConfigSnapshot `json:"configSnapshot,omitempty" xml:"configSnapshot,omitempty" type:"Struct"`
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
	Items []*GetInstanceInspectionResponseBodyLastRunItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Notification *GetInstanceInspectionResponseBodyLastRunNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {}
	Report *GetInstanceInspectionResponseBodyLastRunReport `json:"report,omitempty" xml:"report,omitempty" type:"Struct"`
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
	Summary *GetInstanceInspectionResponseBodyLastRunSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionResponseBodyLastRunTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// manual
	TriggerSource *string `json:"triggerSource,omitempty" xml:"triggerSource,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRun) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRun) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetAiSummary() *GetInstanceInspectionResponseBodyLastRunAiSummary {
	return s.AiSummary
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetConfigSnapshot() *GetInstanceInspectionResponseBodyLastRunConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetItems() []*GetInstanceInspectionResponseBodyLastRunItems {
	return s.Items
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetNotification() *GetInstanceInspectionResponseBodyLastRunNotification {
	return s.Notification
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetReport() *GetInstanceInspectionResponseBodyLastRunReport {
	return s.Report
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetRequestedAt() *string {
	return s.RequestedAt
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetRunId() *string {
	return s.RunId
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetSummary() *GetInstanceInspectionResponseBodyLastRunSummary {
	return s.Summary
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetTimeRange() *GetInstanceInspectionResponseBodyLastRunTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionResponseBodyLastRun) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetAiSummary(v *GetInstanceInspectionResponseBodyLastRunAiSummary) *GetInstanceInspectionResponseBodyLastRun {
	s.AiSummary = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetConfigSnapshot(v *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) *GetInstanceInspectionResponseBodyLastRun {
	s.ConfigSnapshot = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetFinishedAt(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.FinishedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetInstanceId(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetInstanceName(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetItems(v []*GetInstanceInspectionResponseBodyLastRunItems) *GetInstanceInspectionResponseBodyLastRun {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetNotification(v *GetInstanceInspectionResponseBodyLastRunNotification) *GetInstanceInspectionResponseBodyLastRun {
	s.Notification = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetRegionId(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetReport(v *GetInstanceInspectionResponseBodyLastRunReport) *GetInstanceInspectionResponseBodyLastRun {
	s.Report = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetRequestedAt(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.RequestedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetRunId(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.RunId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetStatus(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetSummary(v *GetInstanceInspectionResponseBodyLastRunSummary) *GetInstanceInspectionResponseBodyLastRun {
	s.Summary = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetTimeRange(v *GetInstanceInspectionResponseBodyLastRunTimeRange) *GetInstanceInspectionResponseBodyLastRun {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) SetTriggerSource(v string) *GetInstanceInspectionResponseBodyLastRun {
	s.TriggerSource = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRun) Validate() error {
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

type GetInstanceInspectionResponseBodyLastRunAiSummary struct {
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

func (s GetInstanceInspectionResponseBodyLastRunAiSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunAiSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunAiSummary) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionResponseBodyLastRunAiSummary) GetText() *string {
	return s.Text
}

func (s *GetInstanceInspectionResponseBodyLastRunAiSummary) SetStatus(v string) *GetInstanceInspectionResponseBodyLastRunAiSummary {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunAiSummary) SetText(v string) *GetInstanceInspectionResponseBodyLastRunAiSummary {
	s.Text = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunAiSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunConfigSnapshot struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshot) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) GetItems() []*GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	return s.Items
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) GetResourceScope() *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	return s.ResourceScope
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) GetSchedule() *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	return s.Schedule
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) SetItems(v []*GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) *GetInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Items = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) SetResourceScope(v *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) *GetInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.ResourceScope = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) SetSchedule(v *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) *GetInstanceInspectionResponseBodyLastRunConfigSnapshot {
	s.Schedule = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshot) Validate() error {
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

type GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems struct {
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

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetParameters(v interface{}) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.Parameters = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotItems) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope struct {
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

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetMode() *string {
	return s.Mode
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetConsumerGroups(v []*string) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetMode(v string) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Mode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) SetTopics(v []*string) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope {
	s.Topics = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotResourceScope) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule struct {
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

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetCron() *string {
	return s.Cron
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetCron(v string) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Cron = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) SetEnabled(v bool) *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule {
	s.Enabled = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunConfigSnapshotSchedule) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}}
	EffectiveConfig *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig `json:"effectiveConfig,omitempty" xml:"effectiveConfig,omitempty" type:"Struct"`
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
	Metrics []*GetInstanceInspectionResponseBodyLastRunItemsMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
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
	ResourceSummary *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary `json:"resourceSummary,omitempty" xml:"resourceSummary,omitempty" type:"Struct"`
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
	TimeRange *GetInstanceInspectionResponseBodyLastRunItemsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionResponseBodyLastRunItems) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunItems) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetEffectiveConfig() *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	return s.EffectiveConfig
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetHeadline() *string {
	return s.Headline
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetMetrics() []*GetInstanceInspectionResponseBodyLastRunItemsMetrics {
	return s.Metrics
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetResourceSummary() *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	return s.ResourceSummary
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetSeverity() *string {
	return s.Severity
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) GetTimeRange() *GetInstanceInspectionResponseBodyLastRunItemsTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetDescription(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Description = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetEffectiveConfig(v *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) *GetInstanceInspectionResponseBodyLastRunItems {
	s.EffectiveConfig = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetErrorCode(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetHeadline(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Headline = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetMetrics(v []*GetInstanceInspectionResponseBodyLastRunItemsMetrics) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Metrics = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetName(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetObservedAt(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetResourceCount(v int64) *GetInstanceInspectionResponseBodyLastRunItems {
	s.ResourceCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetResourceSummary(v *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) *GetInstanceInspectionResponseBodyLastRunItems {
	s.ResourceSummary = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetResourceTypes(v []*string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.ResourceTypes = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetResult(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetSeverity(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Severity = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetStatus(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyLastRunItems {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) SetTimeRange(v *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) *GetInstanceInspectionResponseBodyLastRunItems {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItems) Validate() error {
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

type GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig struct {
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetParameters() interface{} {
	return s.Parameters
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetParameters(v interface{}) *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.Parameters = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsEffectiveConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunItemsMetrics struct {
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

func (s GetInstanceInspectionResponseBodyLastRunItemsMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunItemsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) GetBad() *bool {
	return s.Bad
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) GetGood() *bool {
	return s.Good
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) GetK() *string {
	return s.K
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) GetV() *string {
	return s.V
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) SetBad(v bool) *GetInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Bad = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) SetGood(v bool) *GetInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.Good = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) SetK(v string) *GetInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.K = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) SetV(v string) *GetInstanceInspectionResponseBodyLastRunItemsMetrics {
	s.V = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsMetrics) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunItemsResourceSummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetInfoCount(v int32) *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) SetRiskCount(v int32) *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsResourceSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunItemsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunItemsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunItemsTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) SetEndTime(v string) *GetInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) SetStartTime(v string) *GetInstanceInspectionResponseBodyLastRunItemsTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunItemsTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunNotification struct {
	// example:
	//
	// inspection-event-001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Sent
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunNotification) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunNotification) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunNotification) GetEventId() *string {
	return s.EventId
}

func (s *GetInstanceInspectionResponseBodyLastRunNotification) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionResponseBodyLastRunNotification) SetEventId(v string) *GetInstanceInspectionResponseBodyLastRunNotification {
	s.EventId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunNotification) SetStatus(v string) *GetInstanceInspectionResponseBodyLastRunNotification {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunNotification) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunReport struct {
	// example:
	//
	// complete
	Completion *string `json:"completion,omitempty" xml:"completion,omitempty"`
	// example:
	//
	// []
	Coverage []*GetInstanceInspectionResponseBodyLastRunReportCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Findings []*GetInstanceInspectionResponseBodyLastRunReportFindings `json:"findings,omitempty" xml:"findings,omitempty" type:"Repeated"`
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
	ReportSummary *GetInstanceInspectionResponseBodyLastRunReportReportSummary `json:"reportSummary,omitempty" xml:"reportSummary,omitempty" type:"Struct"`
	// example:
	//
	// inspection-v2.1
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionResponseBodyLastRunReportTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionResponseBodyLastRunReport) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReport) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetCompletion() *string {
	return s.Completion
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetCoverage() []*GetInstanceInspectionResponseBodyLastRunReportCoverage {
	return s.Coverage
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetFindings() []*GetInstanceInspectionResponseBodyLastRunReportFindings {
	return s.Findings
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetHtml() *string {
	return s.Html
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetInspectionRunId() *string {
	return s.InspectionRunId
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetMarkdown() *string {
	return s.Markdown
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetOverview() *string {
	return s.Overview
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetReportSummary() *GetInstanceInspectionResponseBodyLastRunReportReportSummary {
	return s.ReportSummary
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) GetTimeRange() *GetInstanceInspectionResponseBodyLastRunReportTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetCompletion(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Completion = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetCoverage(v []*GetInstanceInspectionResponseBodyLastRunReportCoverage) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Coverage = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetFindings(v []*GetInstanceInspectionResponseBodyLastRunReportFindings) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Findings = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetFinishedAt(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.FinishedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetHtml(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Html = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetInspectionRunId(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.InspectionRunId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetInstanceId(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetInstanceName(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetMarkdown(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Markdown = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetOverview(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.Overview = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetRegionId(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetReportId(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.ReportId = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetReportSummary(v *GetInstanceInspectionResponseBodyLastRunReportReportSummary) *GetInstanceInspectionResponseBodyLastRunReport {
	s.ReportSummary = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetTemplateVersion(v string) *GetInstanceInspectionResponseBodyLastRunReport {
	s.TemplateVersion = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) SetTimeRange(v *GetInstanceInspectionResponseBodyLastRunReportTimeRange) *GetInstanceInspectionResponseBodyLastRunReport {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReport) Validate() error {
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

type GetInstanceInspectionResponseBodyLastRunReportCoverage struct {
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
	TimeRange *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportCoverage) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportCoverage) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetDiagnosisStatus() *string {
	return s.DiagnosisStatus
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) GetTimeRange() *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetDiagnosisStatus(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.DiagnosisStatus = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetName(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetObservedAt(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetResult(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetStatus(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) SetTimeRange(v *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) *GetInstanceInspectionResponseBodyLastRunReportCoverage {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverage) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetEndTime(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) SetStartTime(v string) *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportCoverageTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunReportFindings struct {
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
	Resource *GetInstanceInspectionResponseBodyLastRunReportFindingsResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
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
	TimeRange *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// []
	Verification []*string `json:"verification,omitempty" xml:"verification,omitempty" type:"Repeated"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportFindings) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportFindings) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetAlternatives() []*string {
	return s.Alternatives
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetBoundary() *string {
	return s.Boundary
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetCause() *string {
	return s.Cause
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetConfidence() *string {
	return s.Confidence
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetEvidence() []*string {
	return s.Evidence
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetImpact() *string {
	return s.Impact
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetLocation() *string {
	return s.Location
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetMissingEvidence() []*string {
	return s.MissingEvidence
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetObservations() []*string {
	return s.Observations
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetRecommendations() []*string {
	return s.Recommendations
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetResource() *GetInstanceInspectionResponseBodyLastRunReportFindingsResource {
	return s.Resource
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetRiskType() *string {
	return s.RiskType
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetSource() *string {
	return s.Source
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetTimeRange() *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) GetVerification() []*string {
	return s.Verification
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetAlternatives(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Alternatives = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetBoundary(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Boundary = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetCause(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Cause = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetConfidence(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Confidence = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetEvidence(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Evidence = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetImpact(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Impact = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetLocation(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Location = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetMissingEvidence(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.MissingEvidence = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetObservations(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Observations = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetObservedAt(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetRecommendations(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Recommendations = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetResource(v *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Resource = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetRiskType(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.RiskType = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetSource(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Source = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetTemplateCode(v string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetTimeRange(v *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) SetVerification(v []*string) *GetInstanceInspectionResponseBodyLastRunReportFindings {
	s.Verification = v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindings) Validate() error {
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

type GetInstanceInspectionResponseBodyLastRunReportFindingsResource struct {
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

func (s GetInstanceInspectionResponseBodyLastRunReportFindingsResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportFindingsResource) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) GetId() *string {
	return s.Id
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) GetType() *string {
	return s.Type
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) SetId(v string) *GetInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Id = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) SetName(v string) *GetInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) SetType(v string) *GetInstanceInspectionResponseBodyLastRunReportFindingsResource {
	s.Type = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetEndTime(v string) *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) SetStartTime(v string) *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportFindingsTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunReportReportSummary struct {
	// example:
	//
	// 4
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportReportSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportReportSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportReportSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionResponseBodyLastRunReportReportSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionResponseBodyLastRunReportReportSummary) SetInfoCount(v int32) *GetInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportReportSummary) SetRiskCount(v int32) *GetInstanceInspectionResponseBodyLastRunReportReportSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportReportSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunReportTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunReportTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunReportTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunReportTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionResponseBodyLastRunReportTimeRange) SetEndTime(v string) *GetInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportTimeRange) SetStartTime(v string) *GetInstanceInspectionResponseBodyLastRunReportTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunReportTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunSummary struct {
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

func (s GetInstanceInspectionResponseBodyLastRunSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) SetInfoCount(v int32) *GetInstanceInspectionResponseBodyLastRunSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) SetResult(v string) *GetInstanceInspectionResponseBodyLastRunSummary {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) SetRiskCount(v int32) *GetInstanceInspectionResponseBodyLastRunSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionResponseBodyLastRunTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionResponseBodyLastRunTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponseBodyLastRunTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponseBodyLastRunTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionResponseBodyLastRunTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionResponseBodyLastRunTimeRange) SetEndTime(v string) *GetInstanceInspectionResponseBodyLastRunTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunTimeRange) SetStartTime(v string) *GetInstanceInspectionResponseBodyLastRunTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionResponseBodyLastRunTimeRange) Validate() error {
	return dara.Validate(s)
}
