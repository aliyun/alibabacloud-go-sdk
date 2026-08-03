// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionScheduleReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInspectionScheduleReportsResponseBodyData) *DescribeInspectionScheduleReportsResponseBody
	GetData() *DescribeInspectionScheduleReportsResponseBodyData
	SetRequestId(v string) *DescribeInspectionScheduleReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInspectionScheduleReportsResponseBody
	GetSuccess() *bool
}

type DescribeInspectionScheduleReportsResponseBody struct {
	Data *DescribeInspectionScheduleReportsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A057C066-C3F5-4CC9-9FE4-A8D8B0DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInspectionScheduleReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionScheduleReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionScheduleReportsResponseBody) GetData() *DescribeInspectionScheduleReportsResponseBodyData {
	return s.Data
}

func (s *DescribeInspectionScheduleReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionScheduleReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInspectionScheduleReportsResponseBody) SetData(v *DescribeInspectionScheduleReportsResponseBodyData) *DescribeInspectionScheduleReportsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBody) SetRequestId(v string) *DescribeInspectionScheduleReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBody) SetSuccess(v bool) *DescribeInspectionScheduleReportsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInspectionScheduleReportsResponseBodyData struct {
	Items []*DescribeInspectionScheduleReportsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInspectionScheduleReportsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionScheduleReportsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) GetItems() []*DescribeInspectionScheduleReportsResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) SetItems(v []*DescribeInspectionScheduleReportsResponseBodyDataItems) *DescribeInspectionScheduleReportsResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) SetPageNum(v int64) *DescribeInspectionScheduleReportsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) SetPageSize(v int64) *DescribeInspectionScheduleReportsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) SetTotal(v int64) *DescribeInspectionScheduleReportsResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyData) Validate() error {
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

type DescribeInspectionScheduleReportsResponseBodyDataItems struct {
	// example:
	//
	// 1773211755000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2026-06-29T02:12:02Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2026-06-16T13:52:35+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// {\\"Normal\\":10,\\"Warning\\":0,\\"Error\\":0,\\"Failed\\":0}
	LevelSummary *string `json:"LevelSummary,omitempty" xml:"LevelSummary,omitempty"`
	// example:
	//
	// en-US
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// example:
	//
	// 2025-09-26T21:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-0mqomahp4o4uf3aicu
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionScheduleReportsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionScheduleReportsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetLevelSummary() *string {
	return s.LevelSummary
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetCreateTime(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetEndTime(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetFinishTime(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetLevelSummary(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.LevelSummary = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetReportLanguage(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.ReportLanguage = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetStartTime(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetStatus(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) SetTaskId(v string) *DescribeInspectionScheduleReportsResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
