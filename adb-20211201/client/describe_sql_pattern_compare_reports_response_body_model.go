// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSqlPatternCompareReportsResponseBodyItems) *DescribeSqlPatternCompareReportsResponseBody
	GetItems() []*DescribeSqlPatternCompareReportsResponseBodyItems
	SetMaxResults(v int32) *DescribeSqlPatternCompareReportsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSqlPatternCompareReportsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeSqlPatternCompareReportsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSqlPatternCompareReportsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSqlPatternCompareReportsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSqlPatternCompareReportsResponseBody
	GetTotalCount() *int32
}

type DescribeSqlPatternCompareReportsResponseBody struct {
	// The list of reports on the current page. An empty array is returned if no reports match the conditions.
	Items []*DescribeSqlPatternCompareReportsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of rows per page used in this query.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// djE6Mjo1MA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number used in this query. Pages start from 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page used in this query.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1B2C3D-4E5F-6789-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of reports that match the conditions.
	//
	// example:
	//
	// 51
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternCompareReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetItems() []*DescribeSqlPatternCompareReportsResponseBodyItems {
	return s.Items
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlPatternCompareReportsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetItems(v []*DescribeSqlPatternCompareReportsResponseBodyItems) *DescribeSqlPatternCompareReportsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetMaxResults(v int32) *DescribeSqlPatternCompareReportsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetNextToken(v string) *DescribeSqlPatternCompareReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetPageNumber(v int32) *DescribeSqlPatternCompareReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetPageSize(v int32) *DescribeSqlPatternCompareReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetRequestId(v string) *DescribeSqlPatternCompareReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) SetTotalCount(v int32) *DescribeSqlPatternCompareReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlPatternCompareReportsResponseBodyItems struct {
	// Indicates whether the report can be canceled. The value is true when the report is in the PENDING or RUNNING state.
	//
	// example:
	//
	// false
	CancelAvailable *bool `json:"CancelAvailable,omitempty" xml:"CancelAvailable,omitempty"`
	// The end time of time range 2. The time is in the yyyy-MM-ddTHH:mmZ UTC format.
	//
	// example:
	//
	// 2026-09-08T01:00Z
	CompareEndTime *string `json:"CompareEndTime,omitempty" xml:"CompareEndTime,omitempty"`
	// The start time of time range 2. The time is in the yyyy-MM-ddTHH:mmZ UTC format.
	//
	// example:
	//
	// 2026-09-08T00:00Z
	CompareStartTime *string `json:"CompareStartTime,omitempty" xml:"CompareStartTime,omitempty"`
	// The time when the report was created. The time is in the yyyy-MM-ddTHH:mmZ UTC format.
	//
	// example:
	//
	// 2026-09-08T01:05Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether report details can be queried. The value is true when the report is in the SUCCESS state.
	//
	// example:
	//
	// true
	DetailEnabled *bool `json:"DetailEnabled,omitempty" xml:"DetailEnabled,omitempty"`
	// The end time of time range 1. The time is in the yyyy-MM-ddTHH:mmZ UTC format.
	//
	// example:
	//
	// 2026-09-07T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the SQL Pattern comparison report.
	//
	// example:
	//
	// 1001
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The report type. Valid values:
	//
	// - `NEW`: new patterns.
	//
	// - `CHANGED`: patterns with increased metrics.
	//
	// example:
	//
	// CHANGED
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The name of the report type.
	//
	// example:
	//
	// Changed Pattern Comparison Report
	ReportTypeName *string `json:"ReportTypeName,omitempty" xml:"ReportTypeName,omitempty"`
	// The sequence number in the current sorted result. The value starts from 1.
	//
	// example:
	//
	// 1
	RowNumber *int32 `json:"RowNumber,omitempty" xml:"RowNumber,omitempty"`
	// The start time of time range 1. The time is in the yyyy-MM-ddTHH:mmZ UTC format.
	//
	// example:
	//
	// 2026-09-07T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The report status. Valid values:
	//
	// - `PENDING`: waiting to be generated.
	//
	// - `RUNNING`: being generated.
	//
	// - `SUCCESS`: generated.
	//
	// - `FAILED`: failed to be generated.
	//
	// - `CANCELED`: canceled.
	//
	// - `EXPIRED`: expired.
	//
	// > The current list returns only reports in the `PENDING`, `RUNNING`, or `SUCCESS` state.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSqlPatternCompareReportsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetCancelAvailable() *bool {
	return s.CancelAvailable
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetCompareEndTime() *string {
	return s.CompareEndTime
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetCompareStartTime() *string {
	return s.CompareStartTime
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetDetailEnabled() *bool {
	return s.DetailEnabled
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetReportType() *string {
	return s.ReportType
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetReportTypeName() *string {
	return s.ReportTypeName
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetRowNumber() *int32 {
	return s.RowNumber
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetCancelAvailable(v bool) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.CancelAvailable = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetCompareEndTime(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.CompareEndTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetCompareStartTime(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.CompareStartTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetCreatedAt(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetDetailEnabled(v bool) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.DetailEnabled = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetEndTime(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetReportId(v int64) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.ReportId = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetReportType(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.ReportType = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetReportTypeName(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.ReportTypeName = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetRowNumber(v int32) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.RowNumber = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetStartTime(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) SetStatus(v string) *DescribeSqlPatternCompareReportsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
