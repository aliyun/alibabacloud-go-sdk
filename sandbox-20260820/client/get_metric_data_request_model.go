// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetMetricDataRequest
	GetEndTime() *string
	SetGroupBy(v []*string) *GetMetricDataRequest
	GetGroupBy() []*string
	SetMetricNames(v []*string) *GetMetricDataRequest
	GetMetricNames() []*string
	SetPeriod(v int32) *GetMetricDataRequest
	GetPeriod() *int32
	SetSandboxID(v string) *GetMetricDataRequest
	GetSandboxID() *string
	SetSandboxPlan(v string) *GetMetricDataRequest
	GetSandboxPlan() *string
	SetStartTime(v string) *GetMetricDataRequest
	GetStartTime() *string
	SetTeamID(v string) *GetMetricDataRequest
	GetTeamID() *string
	SetTemplateID(v string) *GetMetricDataRequest
	GetTemplateID() *string
	SetTime(v string) *GetMetricDataRequest
	GetTime() *string
}

type GetMetricDataRequest struct {
	// The end time of the time series query. This time is exclusive. Specify a UTC time accurate to seconds. The end time must be later than the start time. You can query metrics for a maximum of 7 days at a time. The time range must include at least one time point aligned with the period, and these time points cannot be later than the current time.
	//
	// example:
	//
	// 2026-11-05T00:02:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Groups metrics by a specific key. For example, you can group metrics by teamID to view the metrics of each team separately. You can combine up to four different grouping keys. The available keys vary by metric. If you do not specify this parameter, the overall results within the query scope are returned.
	//
	// example:
	//
	// ["teamID"]
	GroupBy []*string `json:"groupBy" xml:"groupBy" type:"Repeated"`
	// The names of the metrics to query. You can query 1 to 16 different metrics at a time. For available metrics and query scopes, refer to the metric catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cpuUsedCores"]
	MetricNames []*string `json:"metricNames" xml:"metricNames" type:"Repeated"`
	// The time interval for time series queries, in seconds. Valid values: 60, 300, 900, 1800, 3600, and 86400. A value of 60 corresponds to exact minutes, and 86400 corresponds to 00:00:00 UTC every day. The calculation duration of some statistical metrics varies with the interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The ID of the sandbox to query. You must also specify the teamID and templateID to which the sandbox belongs, and select a metric that supports sandbox-level queries.
	//
	// example:
	//
	// sbx-62105527-699e-48df-ad3b-7425001ae4f2
	SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
	// Queries the number of sandboxes by sandbox plan. This parameter applies only to the sandboxCount metric. Valid values: eco, std, pro, and ent.
	//
	// example:
	//
	// pro
	SandboxPlan *string `json:"sandboxPlan,omitempty" xml:"sandboxPlan,omitempty"`
	// The start time of the time series query. This time is inclusive. Specify a UTC time accurate to seconds, and also specify the endTime parameter.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the team to query. If you leave this parameter empty, the metrics of the current account are queried.
	//
	// example:
	//
	// 70d1c834-0333-58d8-97ac-5336eb9179eb
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The ID of the template to query. You must also specify the teamID to which the template belongs.
	//
	// example:
	//
	// 4hrtmghh2nzrsdh6yr65
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	// The specific point in time to query. If you query metrics at a specific point in time, specify only this parameter. Specify a current or past UTC time accurate to seconds. If you omit all time parameters, the metrics at the current time are queried.
	//
	// example:
	//
	// 2026-11-05T00:01:00Z
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetricDataRequest) GoString() string {
	return s.String()
}

func (s *GetMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMetricDataRequest) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *GetMetricDataRequest) GetMetricNames() []*string {
	return s.MetricNames
}

func (s *GetMetricDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *GetMetricDataRequest) GetSandboxID() *string {
	return s.SandboxID
}

func (s *GetMetricDataRequest) GetSandboxPlan() *string {
	return s.SandboxPlan
}

func (s *GetMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMetricDataRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *GetMetricDataRequest) GetTemplateID() *string {
	return s.TemplateID
}

func (s *GetMetricDataRequest) GetTime() *string {
	return s.Time
}

func (s *GetMetricDataRequest) SetEndTime(v string) *GetMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricDataRequest) SetGroupBy(v []*string) *GetMetricDataRequest {
	s.GroupBy = v
	return s
}

func (s *GetMetricDataRequest) SetMetricNames(v []*string) *GetMetricDataRequest {
	s.MetricNames = v
	return s
}

func (s *GetMetricDataRequest) SetPeriod(v int32) *GetMetricDataRequest {
	s.Period = &v
	return s
}

func (s *GetMetricDataRequest) SetSandboxID(v string) *GetMetricDataRequest {
	s.SandboxID = &v
	return s
}

func (s *GetMetricDataRequest) SetSandboxPlan(v string) *GetMetricDataRequest {
	s.SandboxPlan = &v
	return s
}

func (s *GetMetricDataRequest) SetStartTime(v string) *GetMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetMetricDataRequest) SetTeamID(v string) *GetMetricDataRequest {
	s.TeamID = &v
	return s
}

func (s *GetMetricDataRequest) SetTemplateID(v string) *GetMetricDataRequest {
	s.TemplateID = &v
	return s
}

func (s *GetMetricDataRequest) SetTime(v string) *GetMetricDataRequest {
	s.Time = &v
	return s
}

func (s *GetMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
