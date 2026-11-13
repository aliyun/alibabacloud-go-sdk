// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetMetricDataShrinkRequest
	GetEndTime() *string
	SetGroupByShrink(v string) *GetMetricDataShrinkRequest
	GetGroupByShrink() *string
	SetMetricNamesShrink(v string) *GetMetricDataShrinkRequest
	GetMetricNamesShrink() *string
	SetPeriod(v int32) *GetMetricDataShrinkRequest
	GetPeriod() *int32
	SetSandboxID(v string) *GetMetricDataShrinkRequest
	GetSandboxID() *string
	SetSandboxPlan(v string) *GetMetricDataShrinkRequest
	GetSandboxPlan() *string
	SetStartTime(v string) *GetMetricDataShrinkRequest
	GetStartTime() *string
	SetTeamID(v string) *GetMetricDataShrinkRequest
	GetTeamID() *string
	SetTemplateID(v string) *GetMetricDataShrinkRequest
	GetTemplateID() *string
	SetTime(v string) *GetMetricDataShrinkRequest
	GetTime() *string
}

type GetMetricDataShrinkRequest struct {
	// 时序查询的结束时间，区间不含此时刻。使用 UTC 整秒时间，须晚于 startTime。一次最多查询 7 天。区间须包含至少一个按 period 对齐的时间点，且这些时间点不能晚于当前时间。
	//
	// example:
	//
	// 2026-11-05T00:02:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 例如，按 teamID 分组可分别查看各团队的指标。可组合最多 4 个不同的分组键，可用键因指标而异。不填时返回查询范围的整体结果。
	//
	// example:
	//
	// ["teamID"]
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// 要查询的指标名称。一次可查询 1～16 个不同的指标。可选指标和查询范围见指标目录。
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cpuUsedCores"]
	MetricNamesShrink *string `json:"metricNames,omitempty" xml:"metricNames,omitempty"`
	// 时序查询的时间间隔，单位为秒。可选 60、300、900、1800、3600、86400。60 对应整分钟，86400 对应每日 UTC 零点。部分统计指标的计算时长随间隔变化。
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// 要查询的沙箱 ID。请同时填写所属的 teamID 和 templateID，并选择支持按沙箱查询的指标。
	//
	// example:
	//
	// sbx-62105527-699e-48df-ad3b-7425001ae4f2
	SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
	// 按沙箱版本查询沙箱数量。仅适用于 sandboxCount，可选 eco、std、pro、ent。
	//
	// example:
	//
	// pro
	SandboxPlan *string `json:"sandboxPlan,omitempty" xml:"sandboxPlan,omitempty"`
	// 时序查询的开始时间，包含此时刻。请按示例填写 UTC 时间，精确到秒，并同时填写 endTime。
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 要查询的团队 ID。不填时查询当前账号。
	//
	// example:
	//
	// 70d1c834-0333-58d8-97ac-5336eb9179eb
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// 要查询的模板 ID。请同时填写模板所属的 teamID。
	//
	// example:
	//
	// 4hrtmghh2nzrsdh6yr65
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	// 查询某一时刻时，只填写 time。请按示例填写当前或过去的 UTC 时间，精确到秒。省略所有时间参数时，查询当前时刻。
	//
	// example:
	//
	// 2026-11-05T00:01:00Z
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetMetricDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetricDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMetricDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMetricDataShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *GetMetricDataShrinkRequest) GetMetricNamesShrink() *string {
	return s.MetricNamesShrink
}

func (s *GetMetricDataShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *GetMetricDataShrinkRequest) GetSandboxID() *string {
	return s.SandboxID
}

func (s *GetMetricDataShrinkRequest) GetSandboxPlan() *string {
	return s.SandboxPlan
}

func (s *GetMetricDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMetricDataShrinkRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *GetMetricDataShrinkRequest) GetTemplateID() *string {
	return s.TemplateID
}

func (s *GetMetricDataShrinkRequest) GetTime() *string {
	return s.Time
}

func (s *GetMetricDataShrinkRequest) SetEndTime(v string) *GetMetricDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetGroupByShrink(v string) *GetMetricDataShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetMetricNamesShrink(v string) *GetMetricDataShrinkRequest {
	s.MetricNamesShrink = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetPeriod(v int32) *GetMetricDataShrinkRequest {
	s.Period = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetSandboxID(v string) *GetMetricDataShrinkRequest {
	s.SandboxID = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetSandboxPlan(v string) *GetMetricDataShrinkRequest {
	s.SandboxPlan = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetStartTime(v string) *GetMetricDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetTeamID(v string) *GetMetricDataShrinkRequest {
	s.TeamID = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetTemplateID(v string) *GetMetricDataShrinkRequest {
	s.TemplateID = &v
	return s
}

func (s *GetMetricDataShrinkRequest) SetTime(v string) *GetMetricDataShrinkRequest {
	s.Time = &v
	return s
}

func (s *GetMetricDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
