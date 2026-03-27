// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *CreateDiagnoseReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *CreateDiagnoseReportResponseBody
	GetRequestId() *string
}

type CreateDiagnoseReportResponseBody struct {
	// example:
	//
	// report-qe2s****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnoseReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *CreateDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnoseReportResponseBody) SetReportId(v string) *CreateDiagnoseReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *CreateDiagnoseReportResponseBody) SetRequestId(v string) *CreateDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnoseReportResponseBody) Validate() error {
	return dara.Validate(s)
}
