// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceDashboardReport interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v int64) *GrafanaWorkspaceDashboardReport
	GetGmtCreate() *int64
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceDashboardReport
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceDashboardReport
	GetId() *int64
	SetLastSendTime(v int64) *GrafanaWorkspaceDashboardReport
	GetLastSendTime() *int64
	SetMsg(v string) *GrafanaWorkspaceDashboardReport
	GetMsg() *string
	SetName(v string) *GrafanaWorkspaceDashboardReport
	GetName() *string
	SetReportChannelTarget(v string) *GrafanaWorkspaceDashboardReport
	GetReportChannelTarget() *string
	SetReportChannelType(v string) *GrafanaWorkspaceDashboardReport
	GetReportChannelType() *string
	SetReportStyle(v string) *GrafanaWorkspaceDashboardReport
	GetReportStyle() *string
	SetReportType(v string) *GrafanaWorkspaceDashboardReport
	GetReportType() *string
	SetStatus(v string) *GrafanaWorkspaceDashboardReport
	GetStatus() *string
	SetTriggerDay(v string) *GrafanaWorkspaceDashboardReport
	GetTriggerDay() *string
	SetTriggerTime(v string) *GrafanaWorkspaceDashboardReport
	GetTriggerTime() *string
	SetTriggerType(v string) *GrafanaWorkspaceDashboardReport
	GetTriggerType() *string
	SetUrl(v string) *GrafanaWorkspaceDashboardReport
	GetUrl() *string
	SetUserId(v string) *GrafanaWorkspaceDashboardReport
	GetUserId() *string
}

type GrafanaWorkspaceDashboardReport struct {
	// example:
	//
	// 1680861352600
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// grafana-cn-**********
	GrafanaWorkspaceId *string `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1680861352600
	LastSendTime *int64  `json:"lastSendTime,omitempty" xml:"lastSendTime,omitempty"`
	Msg          *string `json:"msg,omitempty" xml:"msg,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ******@aliyun.com
	ReportChannelTarget *string `json:"reportChannelTarget,omitempty" xml:"reportChannelTarget,omitempty"`
	// example:
	//
	// Email
	ReportChannelType *string `json:"reportChannelType,omitempty" xml:"reportChannelType,omitempty"`
	// example:
	//
	// Grid,Simple
	ReportStyle *string `json:"reportStyle,omitempty" xml:"reportStyle,omitempty"`
	// example:
	//
	// Image
	ReportType *string `json:"reportType,omitempty" xml:"reportType,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	TriggerDay *string `json:"triggerDay,omitempty" xml:"triggerDay,omitempty"`
	// example:
	//
	// 12:25
	TriggerTime *string `json:"triggerTime,omitempty" xml:"triggerTime,omitempty"`
	// example:
	//
	// ByWeek
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// example:
	//
	// Dashboard URL
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GrafanaWorkspaceDashboardReport) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceDashboardReport) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceDashboardReport) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GrafanaWorkspaceDashboardReport) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceDashboardReport) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceDashboardReport) GetLastSendTime() *int64 {
	return s.LastSendTime
}

func (s *GrafanaWorkspaceDashboardReport) GetMsg() *string {
	return s.Msg
}

func (s *GrafanaWorkspaceDashboardReport) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceDashboardReport) GetReportChannelTarget() *string {
	return s.ReportChannelTarget
}

func (s *GrafanaWorkspaceDashboardReport) GetReportChannelType() *string {
	return s.ReportChannelType
}

func (s *GrafanaWorkspaceDashboardReport) GetReportStyle() *string {
	return s.ReportStyle
}

func (s *GrafanaWorkspaceDashboardReport) GetReportType() *string {
	return s.ReportType
}

func (s *GrafanaWorkspaceDashboardReport) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspaceDashboardReport) GetTriggerDay() *string {
	return s.TriggerDay
}

func (s *GrafanaWorkspaceDashboardReport) GetTriggerTime() *string {
	return s.TriggerTime
}

func (s *GrafanaWorkspaceDashboardReport) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GrafanaWorkspaceDashboardReport) GetUrl() *string {
	return s.Url
}

func (s *GrafanaWorkspaceDashboardReport) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspaceDashboardReport) SetGmtCreate(v int64) *GrafanaWorkspaceDashboardReport {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceDashboardReport {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetId(v int64) *GrafanaWorkspaceDashboardReport {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetLastSendTime(v int64) *GrafanaWorkspaceDashboardReport {
	s.LastSendTime = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetMsg(v string) *GrafanaWorkspaceDashboardReport {
	s.Msg = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetName(v string) *GrafanaWorkspaceDashboardReport {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetReportChannelTarget(v string) *GrafanaWorkspaceDashboardReport {
	s.ReportChannelTarget = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetReportChannelType(v string) *GrafanaWorkspaceDashboardReport {
	s.ReportChannelType = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetReportStyle(v string) *GrafanaWorkspaceDashboardReport {
	s.ReportStyle = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetReportType(v string) *GrafanaWorkspaceDashboardReport {
	s.ReportType = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetStatus(v string) *GrafanaWorkspaceDashboardReport {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetTriggerDay(v string) *GrafanaWorkspaceDashboardReport {
	s.TriggerDay = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetTriggerTime(v string) *GrafanaWorkspaceDashboardReport {
	s.TriggerTime = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetTriggerType(v string) *GrafanaWorkspaceDashboardReport {
	s.TriggerType = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetUrl(v string) *GrafanaWorkspaceDashboardReport {
	s.Url = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) SetUserId(v string) *GrafanaWorkspaceDashboardReport {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspaceDashboardReport) Validate() error {
	return dara.Validate(s)
}
