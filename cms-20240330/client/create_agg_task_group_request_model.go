// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroupConfig(v string) *CreateAggTaskGroupRequest
	GetAggTaskGroupConfig() *string
	SetAggTaskGroupConfigType(v string) *CreateAggTaskGroupRequest
	GetAggTaskGroupConfigType() *string
	SetAggTaskGroupName(v string) *CreateAggTaskGroupRequest
	GetAggTaskGroupName() *string
	SetCronExpr(v string) *CreateAggTaskGroupRequest
	GetCronExpr() *string
	SetDelay(v int32) *CreateAggTaskGroupRequest
	GetDelay() *int32
	SetDescription(v string) *CreateAggTaskGroupRequest
	GetDescription() *string
	SetFromTime(v int64) *CreateAggTaskGroupRequest
	GetFromTime() *int64
	SetMaxRetries(v int32) *CreateAggTaskGroupRequest
	GetMaxRetries() *int32
	SetMaxRunTimeInSeconds(v int32) *CreateAggTaskGroupRequest
	GetMaxRunTimeInSeconds() *int32
	SetPrecheckString(v string) *CreateAggTaskGroupRequest
	GetPrecheckString() *string
	SetScheduleMode(v string) *CreateAggTaskGroupRequest
	GetScheduleMode() *string
	SetScheduleTimeExpr(v string) *CreateAggTaskGroupRequest
	GetScheduleTimeExpr() *string
	SetStatus(v string) *CreateAggTaskGroupRequest
	GetStatus() *string
	SetTags(v []*CreateAggTaskGroupRequestTags) *CreateAggTaskGroupRequest
	GetTags() []*CreateAggTaskGroupRequestTags
	SetTargetPrometheusId(v string) *CreateAggTaskGroupRequest
	GetTargetPrometheusId() *string
	SetToTime(v int64) *CreateAggTaskGroupRequest
	GetToTime() *int64
	SetOverrideIfExists(v bool) *CreateAggTaskGroupRequest
	GetOverrideIfExists() *bool
}

type CreateAggTaskGroupRequest struct {
	// Aggregation task group configuration.
	//
	// Currently, only the “RecordingRuleYaml” format is supported, which must comply with the format requirements of open-source Prometheus RecordingRules.
	//
	// This parameter is required.
	//
	// example:
	//
	// groups:
	//
	// - name: "node.rules"
	//
	//   interval: "60s"
	//
	//   rules:
	//
	//   - record: "node_namespace_pod:kube_pod_info:"
	//
	//     expr: "max(label_replace(kube_pod_info{job=\\"kubernetes-pods-kube-state-metrics\\"\\
	//
	//       }, \\"pod\\", \\"$1\\", \\"pod\\", \\"(.*)\\")) by (node, namespace, pod, cluster)"
	AggTaskGroupConfig *string `json:"aggTaskGroupConfig,omitempty" xml:"aggTaskGroupConfig,omitempty"`
	// Aggregation task group configuration type, default is “RecordingRuleYaml” (open-source Prometheus RecordingRule format).
	//
	// example:
	//
	// RecordingRuleYaml
	AggTaskGroupConfigType *string `json:"aggTaskGroupConfigType,omitempty" xml:"aggTaskGroupConfigType,omitempty"`
	// Aggregation task group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// When the scheduling mode is selected as “Cron”, this is the specific scheduling expression. For example, “0/1 	- 	- 	- *” means starting from 0 minutes and scheduling every 1 minute.
	//
	// example:
	//
	// 0/1 	- 	- 	- *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// Fixed delay time for scheduling, in seconds, default is 30.
	//
	// example:
	//
	// 30
	Delay *int32 `json:"delay,omitempty" xml:"delay,omitempty"`
	// Description of the aggregation task group.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The second-level timestamp corresponding to the start time of the schedule.
	//
	// example:
	//
	// 1724996015
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// Maximum number of retries for executing the aggregation task, default is 20.
	//
	// example:
	//
	// 20
	MaxRetries *int32 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// Maximum retry time for executing the aggregation task, in seconds, default is 600.
	//
	// example:
	//
	// 600
	MaxRunTimeInSeconds *int32 `json:"maxRunTimeInSeconds,omitempty" xml:"maxRunTimeInSeconds,omitempty"`
	// Pre-check configuration, no configuration by default. The input string needs to be correctly parsed as JSON.
	//
	// example:
	//
	// {"policy":"skip","prometheusId":"xxx","query":"scalar(sum(count_over_time(up{job=\\"_arms/kubelet/cadvisor\\"}[15s])) / 21)","threshold":0.5,"timeout":15,"type":"promql"}
	PrecheckString *string `json:"precheckString,omitempty" xml:"precheckString,omitempty"`
	// Scheduling mode, either “Cron” or “FixedRate”, default is “FixedRate”.
	//
	// example:
	//
	// FixedRate
	ScheduleMode *string `json:"scheduleMode,omitempty" xml:"scheduleMode,omitempty"`
	// Scheduling time expression, recommended “@s” or “@m”, indicating the alignment granularity of the scheduling time window, default is “@m”.
	//
	// example:
	//
	// @m
	ScheduleTimeExpr *string `json:"scheduleTimeExpr,omitempty" xml:"scheduleTimeExpr,omitempty"`
	// Status of the aggregation task group, either “Running” or “Stopped”. Default is Running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Resource group tags.
	Tags []*CreateAggTaskGroupRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The target Prometheus instance ID of the aggregation task group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rw-pq4apob9jm
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
	// The second-level timestamp corresponding to the end time of the schedule, 0 indicates that the scheduling does not stop.
	//
	// example:
	//
	// 0
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// Whether to overwrite and update if a resource with the same name exists when creating an aggregation task group.
	//
	// example:
	//
	// true
	OverrideIfExists *bool `json:"overrideIfExists,omitempty" xml:"overrideIfExists,omitempty"`
}

func (s CreateAggTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAggTaskGroupRequest) GetAggTaskGroupConfig() *string {
	return s.AggTaskGroupConfig
}

func (s *CreateAggTaskGroupRequest) GetAggTaskGroupConfigType() *string {
	return s.AggTaskGroupConfigType
}

func (s *CreateAggTaskGroupRequest) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *CreateAggTaskGroupRequest) GetCronExpr() *string {
	return s.CronExpr
}

func (s *CreateAggTaskGroupRequest) GetDelay() *int32 {
	return s.Delay
}

func (s *CreateAggTaskGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggTaskGroupRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *CreateAggTaskGroupRequest) GetMaxRetries() *int32 {
	return s.MaxRetries
}

func (s *CreateAggTaskGroupRequest) GetMaxRunTimeInSeconds() *int32 {
	return s.MaxRunTimeInSeconds
}

func (s *CreateAggTaskGroupRequest) GetPrecheckString() *string {
	return s.PrecheckString
}

func (s *CreateAggTaskGroupRequest) GetScheduleMode() *string {
	return s.ScheduleMode
}

func (s *CreateAggTaskGroupRequest) GetScheduleTimeExpr() *string {
	return s.ScheduleTimeExpr
}

func (s *CreateAggTaskGroupRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateAggTaskGroupRequest) GetTags() []*CreateAggTaskGroupRequestTags {
	return s.Tags
}

func (s *CreateAggTaskGroupRequest) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *CreateAggTaskGroupRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *CreateAggTaskGroupRequest) GetOverrideIfExists() *bool {
	return s.OverrideIfExists
}

func (s *CreateAggTaskGroupRequest) SetAggTaskGroupConfig(v string) *CreateAggTaskGroupRequest {
	s.AggTaskGroupConfig = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetAggTaskGroupConfigType(v string) *CreateAggTaskGroupRequest {
	s.AggTaskGroupConfigType = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetAggTaskGroupName(v string) *CreateAggTaskGroupRequest {
	s.AggTaskGroupName = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetCronExpr(v string) *CreateAggTaskGroupRequest {
	s.CronExpr = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetDelay(v int32) *CreateAggTaskGroupRequest {
	s.Delay = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetDescription(v string) *CreateAggTaskGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetFromTime(v int64) *CreateAggTaskGroupRequest {
	s.FromTime = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetMaxRetries(v int32) *CreateAggTaskGroupRequest {
	s.MaxRetries = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetMaxRunTimeInSeconds(v int32) *CreateAggTaskGroupRequest {
	s.MaxRunTimeInSeconds = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetPrecheckString(v string) *CreateAggTaskGroupRequest {
	s.PrecheckString = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetScheduleMode(v string) *CreateAggTaskGroupRequest {
	s.ScheduleMode = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetScheduleTimeExpr(v string) *CreateAggTaskGroupRequest {
	s.ScheduleTimeExpr = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetStatus(v string) *CreateAggTaskGroupRequest {
	s.Status = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetTags(v []*CreateAggTaskGroupRequestTags) *CreateAggTaskGroupRequest {
	s.Tags = v
	return s
}

func (s *CreateAggTaskGroupRequest) SetTargetPrometheusId(v string) *CreateAggTaskGroupRequest {
	s.TargetPrometheusId = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetToTime(v int64) *CreateAggTaskGroupRequest {
	s.ToTime = &v
	return s
}

func (s *CreateAggTaskGroupRequest) SetOverrideIfExists(v bool) *CreateAggTaskGroupRequest {
	s.OverrideIfExists = &v
	return s
}

func (s *CreateAggTaskGroupRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAggTaskGroupRequestTags struct {
	// Key of the resource group tag.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Value of the resource group tag.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAggTaskGroupRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAggTaskGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAggTaskGroupRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateAggTaskGroupRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateAggTaskGroupRequestTags) SetKey(v string) *CreateAggTaskGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAggTaskGroupRequestTags) SetValue(v string) *CreateAggTaskGroupRequestTags {
	s.Value = &v
	return s
}

func (s *CreateAggTaskGroupRequestTags) Validate() error {
	return dara.Validate(s)
}
