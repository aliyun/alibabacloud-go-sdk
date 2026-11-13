// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetResourceUsageShrinkRequest
	GetEndTime() *string
	SetGranularity(v string) *GetResourceUsageShrinkRequest
	GetGranularity() *string
	SetGroupByShrink(v string) *GetResourceUsageShrinkRequest
	GetGroupByShrink() *string
	SetSandboxPlan(v string) *GetResourceUsageShrinkRequest
	GetSandboxPlan() *string
	SetStartTime(v string) *GetResourceUsageShrinkRequest
	GetStartTime() *string
	SetTeamID(v string) *GetResourceUsageShrinkRequest
	GetTeamID() *string
	SetTemplateID(v string) *GetResourceUsageShrinkRequest
	GetTemplateID() *string
	SetUsageNamesShrink(v string) *GetResourceUsageShrinkRequest
	GetUsageNamesShrink() *string
}

type GetResourceUsageShrinkRequest struct {
	// The end time of the query. This time is exclusive. Specify the time on the hour in UTC. The latest allowed end time is the start of the current hour. You can query data for a maximum of 7 days at a time. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T02:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The granularity at which usage data is aggregated. Set this parameter to All to view the cumulative usage over the entire query interval, to Hour to view the hourly usage, or to Day to aggregate the usage every 24 hours from the start time. The last interval can be less than 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The grouping keys. You can group the results by any combination of teamID, templateID, and sandboxPlan. Each key can be specified only once. If you do not specify this parameter, the results are aggregated over the entire query scope.
	//
	// example:
	//
	// ["teamID"]
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The sandbox edition to query. Valid values: eco, std, pro, and ent.
	//
	// example:
	//
	// pro
	SandboxPlan *string `json:"sandboxPlan,omitempty" xml:"sandboxPlan,omitempty"`
	// The start time of the query. This time is inclusive. Specify the time on the hour in UTC, such as 00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-11-05T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the team to query. If you do not specify this parameter, the usage of the current account is queried.
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
	// The usage items to query. You can query 1 to 6 different usage items at a time. CPU usage items include cpuActive and cpuIdle. Memory usage items include memoryActive and memoryIdle. Disk usage items include disk and diskDeepSleep. The disk item includes the usage of disks in deep sleep mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cpuActive"]
	UsageNamesShrink *string `json:"usageNames,omitempty" xml:"usageNames,omitempty"`
}

func (s GetResourceUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceUsageShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetResourceUsageShrinkRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetResourceUsageShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *GetResourceUsageShrinkRequest) GetSandboxPlan() *string {
	return s.SandboxPlan
}

func (s *GetResourceUsageShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetResourceUsageShrinkRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *GetResourceUsageShrinkRequest) GetTemplateID() *string {
	return s.TemplateID
}

func (s *GetResourceUsageShrinkRequest) GetUsageNamesShrink() *string {
	return s.UsageNamesShrink
}

func (s *GetResourceUsageShrinkRequest) SetEndTime(v string) *GetResourceUsageShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetGranularity(v string) *GetResourceUsageShrinkRequest {
	s.Granularity = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetGroupByShrink(v string) *GetResourceUsageShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetSandboxPlan(v string) *GetResourceUsageShrinkRequest {
	s.SandboxPlan = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetStartTime(v string) *GetResourceUsageShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetTeamID(v string) *GetResourceUsageShrinkRequest {
	s.TeamID = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetTemplateID(v string) *GetResourceUsageShrinkRequest {
	s.TemplateID = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) SetUsageNamesShrink(v string) *GetResourceUsageShrinkRequest {
	s.UsageNamesShrink = &v
	return s
}

func (s *GetResourceUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
