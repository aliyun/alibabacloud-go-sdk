// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApmGrafanaDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *QueryApmGrafanaDataShrinkRequest
	GetComponentName() *string
	SetDashboardId(v string) *QueryApmGrafanaDataShrinkRequest
	GetDashboardId() *string
	SetEnd(v string) *QueryApmGrafanaDataShrinkRequest
	GetEnd() *string
	SetProvider(v string) *QueryApmGrafanaDataShrinkRequest
	GetProvider() *string
	SetQuery(v string) *QueryApmGrafanaDataShrinkRequest
	GetQuery() *string
	SetQueryParamsShrink(v string) *QueryApmGrafanaDataShrinkRequest
	GetQueryParamsShrink() *string
	SetQueryUrl(v string) *QueryApmGrafanaDataShrinkRequest
	GetQueryUrl() *string
	SetRegionId(v string) *QueryApmGrafanaDataShrinkRequest
	GetRegionId() *string
	SetStart(v string) *QueryApmGrafanaDataShrinkRequest
	GetStart() *string
	SetStep(v string) *QueryApmGrafanaDataShrinkRequest
	GetStep() *string
	SetTime(v string) *QueryApmGrafanaDataShrinkRequest
	GetTime() *string
	SetVariables(v string) *QueryApmGrafanaDataShrinkRequest
	GetVariables() *string
	SetWorkspaceId(v string) *QueryApmGrafanaDataShrinkRequest
	GetWorkspaceId() *string
}

type QueryApmGrafanaDataShrinkRequest struct {
	// The name of the monitoring dashboard.
	//
	// example:
	//
	// spark-all
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// The dashboard ID.
	//
	// example:
	//
	// ex2tTKoNz
	DashboardId *string `json:"dashboardId,omitempty" xml:"dashboardId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1697033783
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// The datasource provider.
	//
	// example:
	//
	// spark
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The PromQL query expression.
	//
	// example:
	//
	// bizType:
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The panel-level query parameters.
	QueryParamsShrink *string `json:"queryParams,omitempty" xml:"queryParams,omitempty"`
	// The Grafana datasource proxy path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /api/datasources/proxy/1/api/v1/query_range
	QueryUrl *string `json:"queryUrl,omitempty" xml:"queryUrl,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1697030183
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
	// The query step, in seconds.
	//
	// example:
	//
	// 15
	Step *string `json:"step,omitempty" xml:"step,omitempty"`
	// The time point for an instant query. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1697033783
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The dashboard variables, as a JSON string.
	//
	// example:
	//
	// {"fenix_job_runId":"jr-b5059689bb50f360"}
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s QueryApmGrafanaDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApmGrafanaDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryApmGrafanaDataShrinkRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryApmGrafanaDataShrinkRequest) GetDashboardId() *string {
	return s.DashboardId
}

func (s *QueryApmGrafanaDataShrinkRequest) GetEnd() *string {
	return s.End
}

func (s *QueryApmGrafanaDataShrinkRequest) GetProvider() *string {
	return s.Provider
}

func (s *QueryApmGrafanaDataShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *QueryApmGrafanaDataShrinkRequest) GetQueryParamsShrink() *string {
	return s.QueryParamsShrink
}

func (s *QueryApmGrafanaDataShrinkRequest) GetQueryUrl() *string {
	return s.QueryUrl
}

func (s *QueryApmGrafanaDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryApmGrafanaDataShrinkRequest) GetStart() *string {
	return s.Start
}

func (s *QueryApmGrafanaDataShrinkRequest) GetStep() *string {
	return s.Step
}

func (s *QueryApmGrafanaDataShrinkRequest) GetTime() *string {
	return s.Time
}

func (s *QueryApmGrafanaDataShrinkRequest) GetVariables() *string {
	return s.Variables
}

func (s *QueryApmGrafanaDataShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryApmGrafanaDataShrinkRequest) SetComponentName(v string) *QueryApmGrafanaDataShrinkRequest {
	s.ComponentName = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetDashboardId(v string) *QueryApmGrafanaDataShrinkRequest {
	s.DashboardId = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetEnd(v string) *QueryApmGrafanaDataShrinkRequest {
	s.End = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetProvider(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Provider = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetQuery(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Query = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetQueryParamsShrink(v string) *QueryApmGrafanaDataShrinkRequest {
	s.QueryParamsShrink = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetQueryUrl(v string) *QueryApmGrafanaDataShrinkRequest {
	s.QueryUrl = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetRegionId(v string) *QueryApmGrafanaDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetStart(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Start = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetStep(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Step = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetTime(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Time = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetVariables(v string) *QueryApmGrafanaDataShrinkRequest {
	s.Variables = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) SetWorkspaceId(v string) *QueryApmGrafanaDataShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryApmGrafanaDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
