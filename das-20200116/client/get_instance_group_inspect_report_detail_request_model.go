// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetInstanceGroupInspectReportDetailRequest
	GetAgentId() *string
	SetReportId(v string) *GetInstanceGroupInspectReportDetailRequest
	GetReportId() *string
	SetReportType(v string) *GetInstanceGroupInspectReportDetailRequest
	GetReportType() *string
}

type GetInstanceGroupInspectReportDetailRequest struct {
	// The ID of the agent that the user purchased.
	//
	// example:
	//
	// ag-472T0DxtmjIxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The ID of the automated operations report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13f52040-5a6e-42c3-bb84-051f5d6d****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The type of the report.
	//
	// example:
	//
	// daily
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetInstanceGroupInspectReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportDetailRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetInstanceGroupInspectReportDetailRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceGroupInspectReportDetailRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetInstanceGroupInspectReportDetailRequest) SetAgentId(v string) *GetInstanceGroupInspectReportDetailRequest {
	s.AgentId = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailRequest) SetReportId(v string) *GetInstanceGroupInspectReportDetailRequest {
	s.ReportId = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailRequest) SetReportType(v string) *GetInstanceGroupInspectReportDetailRequest {
	s.ReportType = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailRequest) Validate() error {
	return dara.Validate(s)
}
