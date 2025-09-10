// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroupConfig(v string) *UpdateAggTaskGroupRequest
	GetAggTaskGroupConfig() *string
	SetAggTaskGroupConfigType(v string) *UpdateAggTaskGroupRequest
	GetAggTaskGroupConfigType() *string
	SetAggTaskGroupName(v string) *UpdateAggTaskGroupRequest
	GetAggTaskGroupName() *string
	SetCronExpr(v string) *UpdateAggTaskGroupRequest
	GetCronExpr() *string
	SetDelay(v int32) *UpdateAggTaskGroupRequest
	GetDelay() *int32
	SetDescription(v string) *UpdateAggTaskGroupRequest
	GetDescription() *string
	SetFromTime(v int64) *UpdateAggTaskGroupRequest
	GetFromTime() *int64
	SetMaxRetries(v int32) *UpdateAggTaskGroupRequest
	GetMaxRetries() *int32
	SetMaxRunTimeInSeconds(v int32) *UpdateAggTaskGroupRequest
	GetMaxRunTimeInSeconds() *int32
	SetPrecheckString(v string) *UpdateAggTaskGroupRequest
	GetPrecheckString() *string
	SetScheduleMode(v string) *UpdateAggTaskGroupRequest
	GetScheduleMode() *string
	SetScheduleTimeExpr(v string) *UpdateAggTaskGroupRequest
	GetScheduleTimeExpr() *string
	SetStatus(v string) *UpdateAggTaskGroupRequest
	GetStatus() *string
	SetTags(v []*UpdateAggTaskGroupRequestTags) *UpdateAggTaskGroupRequest
	GetTags() []*UpdateAggTaskGroupRequestTags
	SetTargetPrometheusId(v string) *UpdateAggTaskGroupRequest
	GetTargetPrometheusId() *string
	SetToTime(v int64) *UpdateAggTaskGroupRequest
	GetToTime() *int64
}

type UpdateAggTaskGroupRequest struct {
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
	// example:
	//
	// RecordingRuleYaml
	AggTaskGroupConfigType *string `json:"aggTaskGroupConfigType,omitempty" xml:"aggTaskGroupConfigType,omitempty"`
	// example:
	//
	// test-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// example:
	//
	// 0/1 	- 	- 	- *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// example:
	//
	// 30
	Delay *int32 `json:"delay,omitempty" xml:"delay,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1724996015
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// 20
	MaxRetries *int32 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// example:
	//
	// 600
	MaxRunTimeInSeconds *int32 `json:"maxRunTimeInSeconds,omitempty" xml:"maxRunTimeInSeconds,omitempty"`
	// example:
	//
	// {"policy":"skip","prometheusId":"xxx","query":"scalar(sum(count_over_time(up{job=\\"_arms/kubelet/cadvisor\\"}[15s])) / 21)","threshold":0.5,"timeout":15,"type":"promql"}
	PrecheckString *string `json:"precheckString,omitempty" xml:"precheckString,omitempty"`
	// example:
	//
	// FixedRate
	ScheduleMode *string `json:"scheduleMode,omitempty" xml:"scheduleMode,omitempty"`
	// example:
	//
	// @m
	ScheduleTimeExpr *string `json:"scheduleTimeExpr,omitempty" xml:"scheduleTimeExpr,omitempty"`
	// example:
	//
	// Running
	Status *string                          `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*UpdateAggTaskGroupRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rw-pq4apob9jm
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
	// example:
	//
	// 0
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s UpdateAggTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupRequest) GetAggTaskGroupConfig() *string {
	return s.AggTaskGroupConfig
}

func (s *UpdateAggTaskGroupRequest) GetAggTaskGroupConfigType() *string {
	return s.AggTaskGroupConfigType
}

func (s *UpdateAggTaskGroupRequest) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *UpdateAggTaskGroupRequest) GetCronExpr() *string {
	return s.CronExpr
}

func (s *UpdateAggTaskGroupRequest) GetDelay() *int32 {
	return s.Delay
}

func (s *UpdateAggTaskGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggTaskGroupRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *UpdateAggTaskGroupRequest) GetMaxRetries() *int32 {
	return s.MaxRetries
}

func (s *UpdateAggTaskGroupRequest) GetMaxRunTimeInSeconds() *int32 {
	return s.MaxRunTimeInSeconds
}

func (s *UpdateAggTaskGroupRequest) GetPrecheckString() *string {
	return s.PrecheckString
}

func (s *UpdateAggTaskGroupRequest) GetScheduleMode() *string {
	return s.ScheduleMode
}

func (s *UpdateAggTaskGroupRequest) GetScheduleTimeExpr() *string {
	return s.ScheduleTimeExpr
}

func (s *UpdateAggTaskGroupRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAggTaskGroupRequest) GetTags() []*UpdateAggTaskGroupRequestTags {
	return s.Tags
}

func (s *UpdateAggTaskGroupRequest) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *UpdateAggTaskGroupRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *UpdateAggTaskGroupRequest) SetAggTaskGroupConfig(v string) *UpdateAggTaskGroupRequest {
	s.AggTaskGroupConfig = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetAggTaskGroupConfigType(v string) *UpdateAggTaskGroupRequest {
	s.AggTaskGroupConfigType = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetAggTaskGroupName(v string) *UpdateAggTaskGroupRequest {
	s.AggTaskGroupName = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetCronExpr(v string) *UpdateAggTaskGroupRequest {
	s.CronExpr = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetDelay(v int32) *UpdateAggTaskGroupRequest {
	s.Delay = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetDescription(v string) *UpdateAggTaskGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetFromTime(v int64) *UpdateAggTaskGroupRequest {
	s.FromTime = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetMaxRetries(v int32) *UpdateAggTaskGroupRequest {
	s.MaxRetries = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetMaxRunTimeInSeconds(v int32) *UpdateAggTaskGroupRequest {
	s.MaxRunTimeInSeconds = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetPrecheckString(v string) *UpdateAggTaskGroupRequest {
	s.PrecheckString = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetScheduleMode(v string) *UpdateAggTaskGroupRequest {
	s.ScheduleMode = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetScheduleTimeExpr(v string) *UpdateAggTaskGroupRequest {
	s.ScheduleTimeExpr = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetStatus(v string) *UpdateAggTaskGroupRequest {
	s.Status = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetTags(v []*UpdateAggTaskGroupRequestTags) *UpdateAggTaskGroupRequest {
	s.Tags = v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetTargetPrometheusId(v string) *UpdateAggTaskGroupRequest {
	s.TargetPrometheusId = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) SetToTime(v int64) *UpdateAggTaskGroupRequest {
	s.ToTime = &v
	return s
}

func (s *UpdateAggTaskGroupRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateAggTaskGroupRequestTags struct {
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateAggTaskGroupRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateAggTaskGroupRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateAggTaskGroupRequestTags) SetKey(v string) *UpdateAggTaskGroupRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateAggTaskGroupRequestTags) SetValue(v string) *UpdateAggTaskGroupRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateAggTaskGroupRequestTags) Validate() error {
	return dara.Validate(s)
}
