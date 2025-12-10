// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateReport(v *CreateReportInfo) *CreateReportRequest
	GetCreateReport() *CreateReportInfo
}

type CreateReportRequest struct {
	// The details for creating the migration report.
	CreateReport *CreateReportInfo `json:"CreateReport,omitempty" xml:"CreateReport,omitempty"`
}

func (s CreateReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportRequest) GoString() string {
	return s.String()
}

func (s *CreateReportRequest) GetCreateReport() *CreateReportInfo {
	return s.CreateReport
}

func (s *CreateReportRequest) SetCreateReport(v *CreateReportInfo) *CreateReportRequest {
	s.CreateReport = v
	return s
}

func (s *CreateReportRequest) Validate() error {
	if s.CreateReport != nil {
		if err := s.CreateReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}
