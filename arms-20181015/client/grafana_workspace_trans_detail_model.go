// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceTransDetail interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardAmount(v int32) *GrafanaWorkspaceTransDetail
	GetDashboardAmount() *int32
	SetDataSourceAmount(v int32) *GrafanaWorkspaceTransDetail
	GetDataSourceAmount() *int32
	SetOriginal(v int64) *GrafanaWorkspaceTransDetail
	GetOriginal() *int64
	SetOriginalName(v string) *GrafanaWorkspaceTransDetail
	GetOriginalName() *string
	SetTarget(v int64) *GrafanaWorkspaceTransDetail
	GetTarget() *int64
	SetTargetName(v string) *GrafanaWorkspaceTransDetail
	GetTargetName() *string
}

type GrafanaWorkspaceTransDetail struct {
	DashboardAmount  *int32  `json:"dashboardAmount,omitempty" xml:"dashboardAmount,omitempty"`
	DataSourceAmount *int32  `json:"dataSourceAmount,omitempty" xml:"dataSourceAmount,omitempty"`
	Original         *int64  `json:"original,omitempty" xml:"original,omitempty"`
	OriginalName     *string `json:"originalName,omitempty" xml:"originalName,omitempty"`
	Target           *int64  `json:"target,omitempty" xml:"target,omitempty"`
	TargetName       *string `json:"targetName,omitempty" xml:"targetName,omitempty"`
}

func (s GrafanaWorkspaceTransDetail) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceTransDetail) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceTransDetail) GetDashboardAmount() *int32 {
	return s.DashboardAmount
}

func (s *GrafanaWorkspaceTransDetail) GetDataSourceAmount() *int32 {
	return s.DataSourceAmount
}

func (s *GrafanaWorkspaceTransDetail) GetOriginal() *int64 {
	return s.Original
}

func (s *GrafanaWorkspaceTransDetail) GetOriginalName() *string {
	return s.OriginalName
}

func (s *GrafanaWorkspaceTransDetail) GetTarget() *int64 {
	return s.Target
}

func (s *GrafanaWorkspaceTransDetail) GetTargetName() *string {
	return s.TargetName
}

func (s *GrafanaWorkspaceTransDetail) SetDashboardAmount(v int32) *GrafanaWorkspaceTransDetail {
	s.DashboardAmount = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) SetDataSourceAmount(v int32) *GrafanaWorkspaceTransDetail {
	s.DataSourceAmount = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) SetOriginal(v int64) *GrafanaWorkspaceTransDetail {
	s.Original = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) SetOriginalName(v string) *GrafanaWorkspaceTransDetail {
	s.OriginalName = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) SetTarget(v int64) *GrafanaWorkspaceTransDetail {
	s.Target = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) SetTargetName(v string) *GrafanaWorkspaceTransDetail {
	s.TargetName = &v
	return s
}

func (s *GrafanaWorkspaceTransDetail) Validate() error {
	return dara.Validate(s)
}
