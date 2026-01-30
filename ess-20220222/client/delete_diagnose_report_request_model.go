// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteDiagnoseReportRequest
	GetRegionId() *string
	SetReportId(v string) *DeleteDiagnoseReportRequest
	GetReportId() *string
}

type DeleteDiagnoseReportRequest struct {
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the diagnostic report.
	//
	// This parameter is required.
	//
	// example:
	//
	// dr-j6c7wvdzk5103xxxxx
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DeleteDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiagnoseReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDiagnoseReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DeleteDiagnoseReportRequest) SetRegionId(v string) *DeleteDiagnoseReportRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiagnoseReportRequest) SetReportId(v string) *DeleteDiagnoseReportRequest {
	s.ReportId = &v
	return s
}

func (s *DeleteDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
