// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GetInstanceGroupInspectReportDetailRequest
	GetReportId() *string
}

type GetInstanceGroupInspectReportDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13f52040-5a6e-42c3-bb84-051f5d6d****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetInstanceGroupInspectReportDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportDetailRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceGroupInspectReportDetailRequest) SetReportId(v string) *GetInstanceGroupInspectReportDetailRequest {
	s.ReportId = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailRequest) Validate() error {
	return dara.Validate(s)
}
