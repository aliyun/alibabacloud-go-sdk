// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeContainerAppStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeEdgeContainerAppStatsRequest
	GetApp() *string
	SetEndTime(v string) *DescribeEdgeContainerAppStatsRequest
	GetEndTime() *string
	SetFields(v string) *DescribeEdgeContainerAppStatsRequest
	GetFields() *string
	SetIsp(v string) *DescribeEdgeContainerAppStatsRequest
	GetIsp() *string
	SetLocate(v string) *DescribeEdgeContainerAppStatsRequest
	GetLocate() *string
	SetStartTime(v string) *DescribeEdgeContainerAppStatsRequest
	GetStartTime() *string
	SetTenant(v string) *DescribeEdgeContainerAppStatsRequest
	GetTenant() *string
}

type DescribeEdgeContainerAppStatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-xxxx
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 2024-09-02T16:04:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pod_ready_rate
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// telecom,unicom,cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// huizhou
	Locate *string `json:"Locate,omitempty" xml:"Locate,omitempty"`
	// example:
	//
	// 2024-09-02T15:04:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// t-xxxx
	Tenant *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
}

func (s DescribeEdgeContainerAppStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeContainerAppStatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeContainerAppStatsRequest) GetApp() *string {
	return s.App
}

func (s *DescribeEdgeContainerAppStatsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEdgeContainerAppStatsRequest) GetFields() *string {
	return s.Fields
}

func (s *DescribeEdgeContainerAppStatsRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEdgeContainerAppStatsRequest) GetLocate() *string {
	return s.Locate
}

func (s *DescribeEdgeContainerAppStatsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEdgeContainerAppStatsRequest) GetTenant() *string {
	return s.Tenant
}

func (s *DescribeEdgeContainerAppStatsRequest) SetApp(v string) *DescribeEdgeContainerAppStatsRequest {
	s.App = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetEndTime(v string) *DescribeEdgeContainerAppStatsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetFields(v string) *DescribeEdgeContainerAppStatsRequest {
	s.Fields = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetIsp(v string) *DescribeEdgeContainerAppStatsRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetLocate(v string) *DescribeEdgeContainerAppStatsRequest {
	s.Locate = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetStartTime(v string) *DescribeEdgeContainerAppStatsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) SetTenant(v string) *DescribeEdgeContainerAppStatsRequest {
	s.Tenant = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsRequest) Validate() error {
	return dara.Validate(s)
}
