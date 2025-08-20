// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *GetDiagnosticRequest
	GetReportId() *string
}

type GetDiagnosticRequest struct {
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-d540def087714890****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetDiagnosticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosticRequest) GetReportId() *string {
	return s.ReportId
}

func (s *GetDiagnosticRequest) SetReportId(v string) *GetDiagnosticRequest {
	s.ReportId = &v
	return s
}

func (s *GetDiagnosticRequest) Validate() error {
	return dara.Validate(s)
}
