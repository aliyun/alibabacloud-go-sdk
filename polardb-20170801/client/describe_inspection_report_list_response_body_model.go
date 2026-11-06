// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeInspectionReportListResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInspectionReportListResponseBody
	GetPageSize() *int64
	SetReportItems(v []*DescribeInspectionReportListResponseBodyReportItems) *DescribeInspectionReportListResponseBody
	GetReportItems() []*DescribeInspectionReportListResponseBodyReportItems
	SetRequestId(v string) *DescribeInspectionReportListResponseBody
	GetRequestId() *string
	SetScore(v int32) *DescribeInspectionReportListResponseBody
	GetScore() *int32
	SetTotalCount(v int64) *DescribeInspectionReportListResponseBody
	GetTotalCount() *int64
}

type DescribeInspectionReportListResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of inspection reports.
	ReportItems []*DescribeInspectionReportListResponseBodyReportItems `json:"ReportItems,omitempty" xml:"ReportItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The score returned by the inspection task.
	//
	// example:
	//
	// 90
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The total number of reports that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInspectionReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionReportListResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInspectionReportListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionReportListResponseBody) GetReportItems() []*DescribeInspectionReportListResponseBodyReportItems {
	return s.ReportItems
}

func (s *DescribeInspectionReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionReportListResponseBody) GetScore() *int32 {
	return s.Score
}

func (s *DescribeInspectionReportListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInspectionReportListResponseBody) SetPageNumber(v int64) *DescribeInspectionReportListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInspectionReportListResponseBody) SetPageSize(v int64) *DescribeInspectionReportListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionReportListResponseBody) SetReportItems(v []*DescribeInspectionReportListResponseBodyReportItems) *DescribeInspectionReportListResponseBody {
	s.ReportItems = v
	return s
}

func (s *DescribeInspectionReportListResponseBody) SetRequestId(v string) *DescribeInspectionReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionReportListResponseBody) SetScore(v int32) *DescribeInspectionReportListResponseBody {
	s.Score = &v
	return s
}

func (s *DescribeInspectionReportListResponseBody) SetTotalCount(v int64) *DescribeInspectionReportListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInspectionReportListResponseBody) Validate() error {
	if s.ReportItems != nil {
		for _, item := range s.ReportItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInspectionReportListResponseBodyReportItems struct {
	// The ID of the inspected PolarDB cluster.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether this is the latest inspection report. The value is returned as a string.
	//
	// example:
	//
	// true
	Latest *string `json:"Latest,omitempty" xml:"Latest,omitempty"`
	// The region ID of the inspected cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The detailed description of the inspection report.
	//
	// example:
	//
	// No critical risks were found in this inspection
	ReportDescription *string `json:"ReportDescription,omitempty" xml:"ReportDescription,omitempty"`
	// The ID of the inspection report.
	//
	// example:
	//
	// report-001
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The score of the inspection report.
	//
	// example:
	//
	// 90
	ReportScore *int64 `json:"ReportScore,omitempty" xml:"ReportScore,omitempty"`
	// The status of the inspection report.
	//
	// example:
	//
	// Completed
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The summary of the inspection report.
	//
	// example:
	//
	// The instance is running normally
	ReportSummary *string `json:"ReportSummary,omitempty" xml:"ReportSummary,omitempty"`
	// The time when the inspection report was generated. The time is in the yyyy-MM-dd HH:mm:ss format and is in the UTC+8 time zone.
	//
	// example:
	//
	// 2026-11-05 16:00:00
	ReportTime *string `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	// The number of check items categorized by inspection result level.
	SubItems []*DescribeInspectionReportListResponseBodyReportItemsSubItems `json:"SubItems,omitempty" xml:"SubItems,omitempty" type:"Repeated"`
	// The ID of the inspection task.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionReportListResponseBodyReportItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionReportListResponseBodyReportItems) GoString() string {
	return s.String()
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetLatest() *string {
	return s.Latest
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportDescription() *string {
	return s.ReportDescription
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportScore() *int64 {
	return s.ReportScore
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportSummary() *string {
	return s.ReportSummary
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetReportTime() *string {
	return s.ReportTime
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetSubItems() []*DescribeInspectionReportListResponseBodyReportItemsSubItems {
	return s.SubItems
}

func (s *DescribeInspectionReportListResponseBodyReportItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetDBClusterId(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetLatest(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.Latest = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetRegionId(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.RegionId = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportDescription(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportDescription = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportId(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportId = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportScore(v int64) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportScore = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportStatus(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportStatus = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportSummary(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportSummary = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetReportTime(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.ReportTime = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetSubItems(v []*DescribeInspectionReportListResponseBodyReportItemsSubItems) *DescribeInspectionReportListResponseBodyReportItems {
	s.SubItems = v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) SetTaskId(v string) *DescribeInspectionReportListResponseBodyReportItems {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItems) Validate() error {
	if s.SubItems != nil {
		for _, item := range s.SubItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInspectionReportListResponseBodyReportItemsSubItems struct {
	// The number of check items at the Critical level.
	//
	// example:
	//
	// 0
	Critical *int64 `json:"Critical,omitempty" xml:"Critical,omitempty"`
	// The number of check items at the Info level.
	//
	// example:
	//
	// 1
	Info *int64 `json:"Info,omitempty" xml:"Info,omitempty"`
	// The number of check items at the Normal level.
	//
	// example:
	//
	// 10
	Normal *int64 `json:"Normal,omitempty" xml:"Normal,omitempty"`
	// The number of check items at the Warning level.
	//
	// example:
	//
	// 0
	Warning *int64 `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s DescribeInspectionReportListResponseBodyReportItemsSubItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionReportListResponseBodyReportItemsSubItems) GoString() string {
	return s.String()
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) GetCritical() *int64 {
	return s.Critical
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) GetInfo() *int64 {
	return s.Info
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) GetNormal() *int64 {
	return s.Normal
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) GetWarning() *int64 {
	return s.Warning
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) SetCritical(v int64) *DescribeInspectionReportListResponseBodyReportItemsSubItems {
	s.Critical = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) SetInfo(v int64) *DescribeInspectionReportListResponseBodyReportItemsSubItems {
	s.Info = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) SetNormal(v int64) *DescribeInspectionReportListResponseBodyReportItemsSubItems {
	s.Normal = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) SetWarning(v int64) *DescribeInspectionReportListResponseBodyReportItemsSubItems {
	s.Warning = &v
	return s
}

func (s *DescribeInspectionReportListResponseBodyReportItemsSubItems) Validate() error {
	return dara.Validate(s)
}
