// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApmGrafanaDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *QueryApmGrafanaDataRequest
	GetComponentName() *string
	SetDashboardId(v string) *QueryApmGrafanaDataRequest
	GetDashboardId() *string
	SetEnd(v string) *QueryApmGrafanaDataRequest
	GetEnd() *string
	SetProvider(v string) *QueryApmGrafanaDataRequest
	GetProvider() *string
	SetQuery(v string) *QueryApmGrafanaDataRequest
	GetQuery() *string
	SetQueryParams(v *QueryApmGrafanaDataRequestQueryParams) *QueryApmGrafanaDataRequest
	GetQueryParams() *QueryApmGrafanaDataRequestQueryParams
	SetQueryUrl(v string) *QueryApmGrafanaDataRequest
	GetQueryUrl() *string
	SetRegionId(v string) *QueryApmGrafanaDataRequest
	GetRegionId() *string
	SetStart(v string) *QueryApmGrafanaDataRequest
	GetStart() *string
	SetStep(v string) *QueryApmGrafanaDataRequest
	GetStep() *string
	SetTime(v string) *QueryApmGrafanaDataRequest
	GetTime() *string
	SetVariables(v string) *QueryApmGrafanaDataRequest
	GetVariables() *string
	SetWorkspaceId(v string) *QueryApmGrafanaDataRequest
	GetWorkspaceId() *string
}

type QueryApmGrafanaDataRequest struct {
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
	QueryParams *QueryApmGrafanaDataRequestQueryParams `json:"queryParams,omitempty" xml:"queryParams,omitempty" type:"Struct"`
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

func (s QueryApmGrafanaDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApmGrafanaDataRequest) GoString() string {
	return s.String()
}

func (s *QueryApmGrafanaDataRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *QueryApmGrafanaDataRequest) GetDashboardId() *string {
	return s.DashboardId
}

func (s *QueryApmGrafanaDataRequest) GetEnd() *string {
	return s.End
}

func (s *QueryApmGrafanaDataRequest) GetProvider() *string {
	return s.Provider
}

func (s *QueryApmGrafanaDataRequest) GetQuery() *string {
	return s.Query
}

func (s *QueryApmGrafanaDataRequest) GetQueryParams() *QueryApmGrafanaDataRequestQueryParams {
	return s.QueryParams
}

func (s *QueryApmGrafanaDataRequest) GetQueryUrl() *string {
	return s.QueryUrl
}

func (s *QueryApmGrafanaDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryApmGrafanaDataRequest) GetStart() *string {
	return s.Start
}

func (s *QueryApmGrafanaDataRequest) GetStep() *string {
	return s.Step
}

func (s *QueryApmGrafanaDataRequest) GetTime() *string {
	return s.Time
}

func (s *QueryApmGrafanaDataRequest) GetVariables() *string {
	return s.Variables
}

func (s *QueryApmGrafanaDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryApmGrafanaDataRequest) SetComponentName(v string) *QueryApmGrafanaDataRequest {
	s.ComponentName = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetDashboardId(v string) *QueryApmGrafanaDataRequest {
	s.DashboardId = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetEnd(v string) *QueryApmGrafanaDataRequest {
	s.End = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetProvider(v string) *QueryApmGrafanaDataRequest {
	s.Provider = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetQuery(v string) *QueryApmGrafanaDataRequest {
	s.Query = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetQueryParams(v *QueryApmGrafanaDataRequestQueryParams) *QueryApmGrafanaDataRequest {
	s.QueryParams = v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetQueryUrl(v string) *QueryApmGrafanaDataRequest {
	s.QueryUrl = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetRegionId(v string) *QueryApmGrafanaDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetStart(v string) *QueryApmGrafanaDataRequest {
	s.Start = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetStep(v string) *QueryApmGrafanaDataRequest {
	s.Step = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetTime(v string) *QueryApmGrafanaDataRequest {
	s.Time = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetVariables(v string) *QueryApmGrafanaDataRequest {
	s.Variables = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) SetWorkspaceId(v string) *QueryApmGrafanaDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryApmGrafanaDataRequest) Validate() error {
	if s.QueryParams != nil {
		if err := s.QueryParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryApmGrafanaDataRequestQueryParams struct {
	// The panel ID.
	//
	// example:
	//
	// 111
	PanelId *int64 `json:"panelId,omitempty" xml:"panelId,omitempty"`
	// The query reference ID in the panel.
	//
	// example:
	//
	// A
	RefId *string `json:"refId,omitempty" xml:"refId,omitempty"`
	// The variable name. Used when querying the dropdown values of dashboard variables.
	//
	// example:
	//
	// job_oss_download_bandwidth
	VariableName *string `json:"variableName,omitempty" xml:"variableName,omitempty"`
}

func (s QueryApmGrafanaDataRequestQueryParams) String() string {
	return dara.Prettify(s)
}

func (s QueryApmGrafanaDataRequestQueryParams) GoString() string {
	return s.String()
}

func (s *QueryApmGrafanaDataRequestQueryParams) GetPanelId() *int64 {
	return s.PanelId
}

func (s *QueryApmGrafanaDataRequestQueryParams) GetRefId() *string {
	return s.RefId
}

func (s *QueryApmGrafanaDataRequestQueryParams) GetVariableName() *string {
	return s.VariableName
}

func (s *QueryApmGrafanaDataRequestQueryParams) SetPanelId(v int64) *QueryApmGrafanaDataRequestQueryParams {
	s.PanelId = &v
	return s
}

func (s *QueryApmGrafanaDataRequestQueryParams) SetRefId(v string) *QueryApmGrafanaDataRequestQueryParams {
	s.RefId = &v
	return s
}

func (s *QueryApmGrafanaDataRequestQueryParams) SetVariableName(v string) *QueryApmGrafanaDataRequestQueryParams {
	s.VariableName = &v
	return s
}

func (s *QueryApmGrafanaDataRequestQueryParams) Validate() error {
	return dara.Validate(s)
}
