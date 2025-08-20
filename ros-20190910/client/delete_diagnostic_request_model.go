// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *DeleteDiagnosticRequest
	GetReportId() *string
}

type DeleteDiagnosticRequest struct {
	// The report ID. You can troubleshoot issues based on the report.
	//
	// example:
	//
	// dr-56a0e30bf9854b00****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DeleteDiagnosticRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DeleteDiagnosticRequest) SetReportId(v string) *DeleteDiagnosticRequest {
	s.ReportId = &v
	return s
}

func (s *DeleteDiagnosticRequest) Validate() error {
	return dara.Validate(s)
}
