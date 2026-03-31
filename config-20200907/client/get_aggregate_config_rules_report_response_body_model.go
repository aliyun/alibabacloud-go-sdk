// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRulesReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRulesReport(v *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) *GetAggregateConfigRulesReportResponseBody
	GetConfigRulesReport() *GetAggregateConfigRulesReportResponseBodyConfigRulesReport
	SetRequestId(v string) *GetAggregateConfigRulesReportResponseBody
	GetRequestId() *string
}

type GetAggregateConfigRulesReportResponseBody struct {
	// The compliance evaluation report.
	ConfigRulesReport *GetAggregateConfigRulesReportResponseBodyConfigRulesReport `json:"ConfigRulesReport,omitempty" xml:"ConfigRulesReport,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F0BCC7B2-D0E4-49B0-95D2-6689CFB08D31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRulesReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponseBody) GetConfigRulesReport() *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	return s.ConfigRulesReport
}

func (s *GetAggregateConfigRulesReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateConfigRulesReportResponseBody) SetConfigRulesReport(v *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) *GetAggregateConfigRulesReportResponseBody {
	s.ConfigRulesReport = v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBody) SetRequestId(v string) *GetAggregateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBody) Validate() error {
	if s.ConfigRulesReport != nil {
		if err := s.ConfigRulesReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateConfigRulesReportResponseBodyConfigRulesReport struct {
	// The ID of the management account to which the rules belong.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The timestamp when the compliance evaluation report was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1624332329593
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
	// 	- COMPLETE: The compliance evaluation report was generated.
	//
	// example:
	//
	// CREATING
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The URL used to download the compliance evaluation report.
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetAggregateConfigRulesReportResponseBodyConfigRulesReport) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetReportCreateTimestamp() *int64 {
	return s.ReportCreateTimestamp
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetReportId() *string {
	return s.ReportId
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetAccountId(v int64) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetAggregatorId(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportCreateTimestamp(v int64) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportId(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportStatus(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportStatus = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportUrl(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportUrl = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) Validate() error {
	return dara.Validate(s)
}
