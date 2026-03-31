// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportFromTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetReportFromTemplateResponseBody
	GetRequestId() *string
	SetTemplateReport(v *GetReportFromTemplateResponseBodyTemplateReport) *GetReportFromTemplateResponseBody
	GetTemplateReport() *GetReportFromTemplateResponseBodyTemplateReport
}

type GetReportFromTemplateResponseBody struct {
	// example:
	//
	// DE9FFFE5-FCAD-4B24-9546-BF49273C562B
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateReport *GetReportFromTemplateResponseBodyTemplateReport `json:"TemplateReport,omitempty" xml:"TemplateReport,omitempty" type:"Struct"`
}

func (s GetReportFromTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportFromTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReportFromTemplateResponseBody) GetTemplateReport() *GetReportFromTemplateResponseBodyTemplateReport {
	return s.TemplateReport
}

func (s *GetReportFromTemplateResponseBody) SetRequestId(v string) *GetReportFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportFromTemplateResponseBody) SetTemplateReport(v *GetReportFromTemplateResponseBodyTemplateReport) *GetReportFromTemplateResponseBody {
	s.TemplateReport = v
	return s
}

func (s *GetReportFromTemplateResponseBody) Validate() error {
	if s.TemplateReport != nil {
		if err := s.TemplateReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReportFromTemplateResponseBodyTemplateReport struct {
	// example:
	//
	// 1478085326082xxx
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 1763540426815
	ReportCreateEndTimestamp *int64 `json:"ReportCreateEndTimestamp,omitempty" xml:"ReportCreateEndTimestamp,omitempty"`
	// example:
	//
	// 1763540421815
	ReportCreateStartTimestamp *int64 `json:"ReportCreateStartTimestamp,omitempty" xml:"ReportCreateStartTimestamp,omitempty"`
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
	// example:
	//
	// https://xxx
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// crtr-xxx
	TemplateReportId *string `json:"TemplateReportId,omitempty" xml:"TemplateReportId,omitempty"`
}

func (s GetReportFromTemplateResponseBodyTemplateReport) String() string {
	return dara.Prettify(s)
}

func (s GetReportFromTemplateResponseBodyTemplateReport) GoString() string {
	return s.String()
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetReportCreateEndTimestamp() *int64 {
	return s.ReportCreateEndTimestamp
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetReportCreateStartTimestamp() *int64 {
	return s.ReportCreateStartTimestamp
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetStatus() *string {
	return s.Status
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) GetTemplateReportId() *string {
	return s.TemplateReportId
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetAccountId(v int64) *GetReportFromTemplateResponseBodyTemplateReport {
	s.AccountId = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetReportCreateEndTimestamp(v int64) *GetReportFromTemplateResponseBodyTemplateReport {
	s.ReportCreateEndTimestamp = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetReportCreateStartTimestamp(v int64) *GetReportFromTemplateResponseBodyTemplateReport {
	s.ReportCreateStartTimestamp = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetReportTemplateId(v string) *GetReportFromTemplateResponseBodyTemplateReport {
	s.ReportTemplateId = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetReportUrl(v string) *GetReportFromTemplateResponseBodyTemplateReport {
	s.ReportUrl = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetStatus(v string) *GetReportFromTemplateResponseBodyTemplateReport {
	s.Status = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) SetTemplateReportId(v string) *GetReportFromTemplateResponseBodyTemplateReport {
	s.TemplateReportId = &v
	return s
}

func (s *GetReportFromTemplateResponseBodyTemplateReport) Validate() error {
	return dara.Validate(s)
}
