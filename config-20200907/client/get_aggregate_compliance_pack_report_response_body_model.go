// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackReport(v *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) *GetAggregateCompliancePackReportResponseBody
	GetCompliancePackReport() *GetAggregateCompliancePackReportResponseBodyCompliancePackReport
	SetRequestId(v string) *GetAggregateCompliancePackReportResponseBody
	GetRequestId() *string
}

type GetAggregateCompliancePackReportResponseBody struct {
	CompliancePackReport *GetAggregateCompliancePackReportResponseBodyCompliancePackReport `json:"CompliancePackReport,omitempty" xml:"CompliancePackReport,omitempty" type:"Struct"`
	RequestId            *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateCompliancePackReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponseBody) GetCompliancePackReport() *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	return s.CompliancePackReport
}

func (s *GetAggregateCompliancePackReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateCompliancePackReportResponseBody) SetCompliancePackReport(v *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) *GetAggregateCompliancePackReportResponseBody {
	s.CompliancePackReport = v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBody) SetRequestId(v string) *GetAggregateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBody) Validate() error {
	if s.CompliancePackReport != nil {
		if err := s.CompliancePackReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateCompliancePackReportResponseBodyCompliancePackReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId      *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetAggregateCompliancePackReportResponseBodyCompliancePackReport) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GetReportCreateTimestamp() *int64 {
	return s.ReportCreateTimestamp
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetAccountId(v int64) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.AccountId = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetCompliancePackId(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportCreateTimestamp(v int64) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportStatus(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportStatus = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportUrl(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportUrl = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) Validate() error {
	return dara.Validate(s)
}
