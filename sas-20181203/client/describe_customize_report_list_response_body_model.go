// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportList(v []*DescribeCustomizeReportListResponseBodyReportList) *DescribeCustomizeReportListResponseBody
	GetReportList() []*DescribeCustomizeReportListResponseBodyReportList
	SetRequestId(v string) *DescribeCustomizeReportListResponseBody
	GetRequestId() *string
}

type DescribeCustomizeReportListResponseBody struct {
	// The reports.
	ReportList []*DescribeCustomizeReportListResponseBodyReportList `json:"ReportList,omitempty" xml:"ReportList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9FBC6E47-7508-58C9-9E76-528E118CB1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomizeReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportListResponseBody) GetReportList() []*DescribeCustomizeReportListResponseBodyReportList {
	return s.ReportList
}

func (s *DescribeCustomizeReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizeReportListResponseBody) SetReportList(v []*DescribeCustomizeReportListResponseBodyReportList) *DescribeCustomizeReportListResponseBody {
	s.ReportList = v
	return s
}

func (s *DescribeCustomizeReportListResponseBody) SetRequestId(v string) *DescribeCustomizeReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomizeReportListResponseBodyReportList struct {
	// Indicates whether the report is the default report. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The timestamp when the report is pinned. Unit: milliseconds.
	//
	// example:
	//
	// 1721836800000
	PinnedTime *int64 `json:"PinnedTime,omitempty" xml:"PinnedTime,omitempty"`
	// The most recent days for report statistics.
	//
	// example:
	//
	// 7
	ReportDays *int32 `json:"ReportDays,omitempty" xml:"ReportDays,omitempty"`
	// The end date on which the report is sent. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1721923199999
	ReportEndDate *int64 `json:"ReportEndDate,omitempty" xml:"ReportEndDate,omitempty"`
	// The ID of the report.
	//
	// example:
	//
	// 1
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The start date on which the report is sent. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1721836800000
	ReportStartDate *int64 `json:"ReportStartDate,omitempty" xml:"ReportStartDate,omitempty"`
	// The state of the report. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The type of the report. Valid values:
	//
	// 	- **0**: daily report
	//
	// 	- **1**: weekly report
	//
	// 	- **2**: monthly report
	//
	// 	- **3**: report whose statistics are collected in a custom time range
	//
	// example:
	//
	// 0
	ReportType *int32 `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The report version.
	//
	// example:
	//
	// 2.0.0
	ReportVersion *string `json:"ReportVersion,omitempty" xml:"ReportVersion,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeCustomizeReportListResponseBodyReportList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportListResponseBodyReportList) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetPinnedTime() *int64 {
	return s.PinnedTime
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportDays() *int32 {
	return s.ReportDays
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportEndDate() *int64 {
	return s.ReportEndDate
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportStartDate() *int64 {
	return s.ReportStartDate
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportType() *int32 {
	return s.ReportType
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetReportVersion() *string {
	return s.ReportVersion
}

func (s *DescribeCustomizeReportListResponseBodyReportList) GetTitle() *string {
	return s.Title
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetIsDefault(v string) *DescribeCustomizeReportListResponseBodyReportList {
	s.IsDefault = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetPinnedTime(v int64) *DescribeCustomizeReportListResponseBodyReportList {
	s.PinnedTime = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportDays(v int32) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportDays = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportEndDate(v int64) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportEndDate = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportId(v int64) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportId = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportStartDate(v int64) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportStartDate = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportStatus(v string) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportStatus = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportType(v int32) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportType = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetReportVersion(v string) *DescribeCustomizeReportListResponseBodyReportList {
	s.ReportVersion = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) SetTitle(v string) *DescribeCustomizeReportListResponseBodyReportList {
	s.Title = &v
	return s
}

func (s *DescribeCustomizeReportListResponseBodyReportList) Validate() error {
	return dara.Validate(s)
}
