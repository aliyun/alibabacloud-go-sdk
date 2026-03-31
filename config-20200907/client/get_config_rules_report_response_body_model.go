// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRulesReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRulesReport(v *GetConfigRulesReportResponseBodyConfigRulesReport) *GetConfigRulesReportResponseBody
	GetConfigRulesReport() *GetConfigRulesReportResponseBodyConfigRulesReport
	SetRequestId(v string) *GetConfigRulesReportResponseBody
	GetRequestId() *string
}

type GetConfigRulesReportResponseBody struct {
	// The information about the compliance evaluation report.
	ConfigRulesReport *GetConfigRulesReportResponseBodyConfigRulesReport `json:"ConfigRulesReport,omitempty" xml:"ConfigRulesReport,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRulesReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponseBody) GetConfigRulesReport() *GetConfigRulesReportResponseBodyConfigRulesReport {
	return s.ConfigRulesReport
}

func (s *GetConfigRulesReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigRulesReportResponseBody) SetConfigRulesReport(v *GetConfigRulesReportResponseBodyConfigRulesReport) *GetConfigRulesReportResponseBody {
	s.ConfigRulesReport = v
	return s
}

func (s *GetConfigRulesReportResponseBody) SetRequestId(v string) *GetConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRulesReportResponseBody) Validate() error {
	if s.ConfigRulesReport != nil {
		if err := s.ConfigRulesReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigRulesReportResponseBodyConfigRulesReport struct {
	// The ID of the Alibaba Cloud account to which the rules belong.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The timestamp when the compliance evaluation report was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	ReportCreateTimestamp *int64 `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	// The ID of the compliance evaluation report.
	//
	// example:
	//
	// crp-88176457e0d900c9****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
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
	// CREATING
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The URL that is used to download the compliance evaluation report.
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetConfigRulesReportResponseBodyConfigRulesReport) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRulesReportResponseBodyConfigRulesReport) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) GetReportCreateTimestamp() *int64 {
	return s.ReportCreateTimestamp
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) GetReportId() *string {
	return s.ReportId
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetAccountId(v int64) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.AccountId = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportCreateTimestamp(v int64) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportId(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportId = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportStatus(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportStatus = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportUrl(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportUrl = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) Validate() error {
	return dara.Validate(s)
}
