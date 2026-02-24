// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackReport(v *GetCompliancePackReportResponseBodyCompliancePackReport) *GetCompliancePackReportResponseBody
	GetCompliancePackReport() *GetCompliancePackReportResponseBodyCompliancePackReport
	SetRequestId(v string) *GetCompliancePackReportResponseBody
	GetRequestId() *string
}

type GetCompliancePackReportResponseBody struct {
	CompliancePackReport *GetCompliancePackReportResponseBodyCompliancePackReport `json:"CompliancePackReport,omitempty" xml:"CompliancePackReport,omitempty" type:"Struct"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCompliancePackReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponseBody) GetCompliancePackReport() *GetCompliancePackReportResponseBodyCompliancePackReport {
	return s.CompliancePackReport
}

func (s *GetCompliancePackReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompliancePackReportResponseBody) SetCompliancePackReport(v *GetCompliancePackReportResponseBodyCompliancePackReport) *GetCompliancePackReportResponseBody {
	s.CompliancePackReport = v
	return s
}

func (s *GetCompliancePackReportResponseBody) SetRequestId(v string) *GetCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompliancePackReportResponseBody) Validate() error {
	if s.CompliancePackReport != nil {
		if err := s.CompliancePackReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCompliancePackReportResponseBodyCompliancePackReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId      *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetCompliancePackReportResponseBodyCompliancePackReport) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackReportResponseBodyCompliancePackReport) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) GetReportCreateTimestamp() *int64 {
	return s.ReportCreateTimestamp
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetAccountId(v int64) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.AccountId = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetCompliancePackId(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportCreateTimestamp(v int64) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportStatus(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportStatus = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportUrl(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportUrl = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) Validate() error {
	return dara.Validate(s)
}
