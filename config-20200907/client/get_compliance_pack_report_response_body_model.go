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
	// The information about the compliance evaluation report.
	CompliancePackReport *GetCompliancePackReportResponseBodyCompliancePackReport `json:"CompliancePackReport,omitempty" xml:"CompliancePackReport,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the Alibaba Cloud account to which the compliance package belongs.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The timestamp when the compliance evaluation report was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1624329965857
	ReportCreateTimestamp *int64 `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	// The status of the compliance evaluation report. Valid values:
	//
	// 	- NONE: The compliance evaluation report is not generated.
	//
	// 	- CREATING: The compliance evaluation report is being generated.
	//
	// 	- COMPLETE: The compliance evaluation report is generated.
	//
	// example:
	//
	// COMPLETE
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The URL that is used to download the compliance evaluation report.
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
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
