// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInspectReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListClusterInspectReportsResponseBody
	GetNextToken() *string
	SetReports(v []*ListClusterInspectReportsResponseBodyReports) *ListClusterInspectReportsResponseBody
	GetReports() []*ListClusterInspectReportsResponseBodyReports
	SetRequestId(v string) *ListClusterInspectReportsResponseBody
	GetRequestId() *string
}

type ListClusterInspectReportsResponseBody struct {
	// The pagination token.
	//
	// example:
	//
	// 405b99e5411f9a4e7148506e45
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of inspection reports.
	Reports []*ListClusterInspectReportsResponseBodyReports `json:"reports,omitempty" xml:"reports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListClusterInspectReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInspectReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterInspectReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterInspectReportsResponseBody) GetReports() []*ListClusterInspectReportsResponseBodyReports {
	return s.Reports
}

func (s *ListClusterInspectReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterInspectReportsResponseBody) SetNextToken(v string) *ListClusterInspectReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClusterInspectReportsResponseBody) SetReports(v []*ListClusterInspectReportsResponseBodyReports) *ListClusterInspectReportsResponseBody {
	s.Reports = v
	return s
}

func (s *ListClusterInspectReportsResponseBody) SetRequestId(v string) *ListClusterInspectReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterInspectReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterInspectReportsResponseBodyReports struct {
	// The report completion time.
	//
	// example:
	//
	// 2024-12-18T19:40:16.778333+08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// An inspection report ID.
	//
	// example:
	//
	// 782df89346054a0000562063a6****
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The report start time.
	//
	// example:
	//
	// 2024-12-18T19:40:16.778333+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The inspection report status.
	//
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The inspection summary.
	Summary *ListClusterInspectReportsResponseBodyReportsSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
}

func (s ListClusterInspectReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInspectReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *ListClusterInspectReportsResponseBodyReports) GetEndTime() *string {
	return s.EndTime
}

func (s *ListClusterInspectReportsResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *ListClusterInspectReportsResponseBodyReports) GetStartTime() *string {
	return s.StartTime
}

func (s *ListClusterInspectReportsResponseBodyReports) GetStatus() *string {
	return s.Status
}

func (s *ListClusterInspectReportsResponseBodyReports) GetSummary() *ListClusterInspectReportsResponseBodyReportsSummary {
	return s.Summary
}

func (s *ListClusterInspectReportsResponseBodyReports) SetEndTime(v string) *ListClusterInspectReportsResponseBodyReports {
	s.EndTime = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReports) SetReportId(v string) *ListClusterInspectReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReports) SetStartTime(v string) *ListClusterInspectReportsResponseBodyReports {
	s.StartTime = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReports) SetStatus(v string) *ListClusterInspectReportsResponseBodyReports {
	s.Status = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReports) SetSummary(v *ListClusterInspectReportsResponseBodyReportsSummary) *ListClusterInspectReportsResponseBodyReports {
	s.Summary = v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}

type ListClusterInspectReportsResponseBodyReportsSummary struct {
	// The number of items whose result is advice.
	//
	// example:
	//
	// 0
	AdviceCount *int32 `json:"adviceCount,omitempty" xml:"adviceCount,omitempty"`
	// Aggregated inspection task result code.
	//
	// example:
	//
	// warning
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The number of items whose result is error.
	//
	// example:
	//
	// 0
	ErrorCount *int32 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// The number of items whose result is normal.
	//
	// example:
	//
	// 1
	NormalCount *int32 `json:"normalCount,omitempty" xml:"normalCount,omitempty"`
	// The number of items whose result is warning.
	//
	// example:
	//
	// 0
	WarnCount *int32 `json:"warnCount,omitempty" xml:"warnCount,omitempty"`
}

func (s ListClusterInspectReportsResponseBodyReportsSummary) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInspectReportsResponseBodyReportsSummary) GoString() string {
	return s.String()
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) GetAdviceCount() *int32 {
	return s.AdviceCount
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) GetCode() *string {
	return s.Code
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) GetErrorCount() *int32 {
	return s.ErrorCount
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) GetNormalCount() *int32 {
	return s.NormalCount
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) GetWarnCount() *int32 {
	return s.WarnCount
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) SetAdviceCount(v int32) *ListClusterInspectReportsResponseBodyReportsSummary {
	s.AdviceCount = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) SetCode(v string) *ListClusterInspectReportsResponseBodyReportsSummary {
	s.Code = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) SetErrorCount(v int32) *ListClusterInspectReportsResponseBodyReportsSummary {
	s.ErrorCount = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) SetNormalCount(v int32) *ListClusterInspectReportsResponseBodyReportsSummary {
	s.NormalCount = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) SetWarnCount(v int32) *ListClusterInspectReportsResponseBodyReportsSummary {
	s.WarnCount = &v
	return s
}

func (s *ListClusterInspectReportsResponseBodyReportsSummary) Validate() error {
	return dara.Validate(s)
}
