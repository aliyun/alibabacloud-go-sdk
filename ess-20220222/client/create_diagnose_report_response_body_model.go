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
	// The unique ID of the diagnostic report.
	//
	// example:
	//
	// dr-uf6enpbnri1xhcy9qc7s
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0189C6CB-07BA-5AFE-B533-D93892324774
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
