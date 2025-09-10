// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggTaskGroup(v *GetAggTaskGroupResponseBodyAggTaskGroup) *GetAggTaskGroupResponseBody
	GetAggTaskGroup() *GetAggTaskGroupResponseBodyAggTaskGroup
	SetRequestId(v string) *GetAggTaskGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAggTaskGroupResponseBody
	GetSuccess() *bool
}

type GetAggTaskGroupResponseBody struct {
	AggTaskGroup *GetAggTaskGroupResponseBodyAggTaskGroup `json:"aggTaskGroup,omitempty" xml:"aggTaskGroup,omitempty" type:"Struct"`
	// example:
	//
	// 68DAF543-35DF-5762-BE90-F5C00B5DC036
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAggTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggTaskGroupResponseBody) GetAggTaskGroup() *GetAggTaskGroupResponseBodyAggTaskGroup {
	return s.AggTaskGroup
}

func (s *GetAggTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggTaskGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAggTaskGroupResponseBody) SetAggTaskGroup(v *GetAggTaskGroupResponseBodyAggTaskGroup) *GetAggTaskGroupResponseBody {
	s.AggTaskGroup = v
	return s
}

func (s *GetAggTaskGroupResponseBody) SetRequestId(v string) *GetAggTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggTaskGroupResponseBody) SetSuccess(v bool) *GetAggTaskGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetAggTaskGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAggTaskGroupResponseBodyAggTaskGroup struct {
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
	// a54136xxx
	AggTaskGroupConfigHash *string `json:"aggTaskGroupConfigHash,omitempty" xml:"aggTaskGroupConfigHash,omitempty"`
	// example:
	//
	// aggTaskGroup-xx
	AggTaskGroupId *string `json:"aggTaskGroupId,omitempty" xml:"aggTaskGroupId,omitempty"`
	// example:
	//
	// pipeline-aggtask-group
	AggTaskGroupName *string `json:"aggTaskGroupName,omitempty" xml:"aggTaskGroupName,omitempty"`
	// example:
	//
	// 0 1 3 	- 	- ? *
	CronExpr *string `json:"cronExpr,omitempty" xml:"cronExpr,omitempty"`
	// example:
	//
	// 2
	Delay *int32 `json:"delay,omitempty" xml:"delay,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1757409495
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// example:
	//
	// 2
	MaxRetries *int32 `json:"maxRetries,omitempty" xml:"maxRetries,omitempty"`
	// example:
	//
	// 50
	MaxRunTimeInSeconds *int32 `json:"maxRunTimeInSeconds,omitempty" xml:"maxRunTimeInSeconds,omitempty"`
	// example:
	//
	// {"policy":"skip","prometheusId":"rw-xx","query":"noPrecheck","threshold":0.5,"timeout":15,"type":"none"}
	PrecheckString *string `json:"precheckString,omitempty" xml:"precheckString,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
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
	// rw-xxx
	SourcePrometheusId *string `json:"sourcePrometheusId,omitempty" xml:"sourcePrometheusId,omitempty"`
	// example:
	//
	// Running
	Status *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*GetAggTaskGroupResponseBodyAggTaskGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// rw-xxx
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
	// example:
	//
	// 1757409495
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// example:
	//
	// 1757409499000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 123xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAggTaskGroupResponseBodyAggTaskGroup) String() string {
	return dara.Prettify(s)
}

func (s GetAggTaskGroupResponseBodyAggTaskGroup) GoString() string {
	return s.String()
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetAggTaskGroupConfig() *string {
	return s.AggTaskGroupConfig
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetAggTaskGroupConfigHash() *string {
	return s.AggTaskGroupConfigHash
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetAggTaskGroupId() *string {
	return s.AggTaskGroupId
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetAggTaskGroupName() *string {
	return s.AggTaskGroupName
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetCronExpr() *string {
	return s.CronExpr
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetDelay() *int32 {
	return s.Delay
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetDescription() *string {
	return s.Description
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetFromTime() *int64 {
	return s.FromTime
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetMaxRetries() *int32 {
	return s.MaxRetries
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetMaxRunTimeInSeconds() *int32 {
	return s.MaxRunTimeInSeconds
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetPrecheckString() *string {
	return s.PrecheckString
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetScheduleMode() *string {
	return s.ScheduleMode
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetScheduleTimeExpr() *string {
	return s.ScheduleTimeExpr
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetSourcePrometheusId() *string {
	return s.SourcePrometheusId
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetStatus() *string {
	return s.Status
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetTags() []*GetAggTaskGroupResponseBodyAggTaskGroupTags {
	return s.Tags
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetToTime() *int64 {
	return s.ToTime
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) GetUserId() *string {
	return s.UserId
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetAggTaskGroupConfig(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.AggTaskGroupConfig = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetAggTaskGroupConfigHash(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.AggTaskGroupConfigHash = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetAggTaskGroupId(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.AggTaskGroupId = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetAggTaskGroupName(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.AggTaskGroupName = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetCronExpr(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.CronExpr = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetDelay(v int32) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.Delay = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetDescription(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.Description = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetFromTime(v int64) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.FromTime = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetMaxRetries(v int32) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.MaxRetries = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetMaxRunTimeInSeconds(v int32) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.MaxRunTimeInSeconds = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetPrecheckString(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.PrecheckString = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetRegionId(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.RegionId = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetScheduleMode(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.ScheduleMode = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetScheduleTimeExpr(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.ScheduleTimeExpr = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetSourcePrometheusId(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.SourcePrometheusId = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetStatus(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.Status = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetTags(v []*GetAggTaskGroupResponseBodyAggTaskGroupTags) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.Tags = v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetTargetPrometheusId(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.TargetPrometheusId = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetToTime(v int64) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.ToTime = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetUpdateTime(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.UpdateTime = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) SetUserId(v string) *GetAggTaskGroupResponseBodyAggTaskGroup {
	s.UserId = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroup) Validate() error {
	return dara.Validate(s)
}

type GetAggTaskGroupResponseBodyAggTaskGroupTags struct {
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAggTaskGroupResponseBodyAggTaskGroupTags) String() string {
	return dara.Prettify(s)
}

func (s GetAggTaskGroupResponseBodyAggTaskGroupTags) GoString() string {
	return s.String()
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroupTags) GetKey() *string {
	return s.Key
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroupTags) GetValue() *string {
	return s.Value
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroupTags) SetKey(v string) *GetAggTaskGroupResponseBodyAggTaskGroupTags {
	s.Key = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroupTags) SetValue(v string) *GetAggTaskGroupResponseBodyAggTaskGroupTags {
	s.Value = &v
	return s
}

func (s *GetAggTaskGroupResponseBodyAggTaskGroupTags) Validate() error {
	return dara.Validate(s)
}
