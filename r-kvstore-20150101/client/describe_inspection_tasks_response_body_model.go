// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInspectionTasksResponseBodyData) *DescribeInspectionTasksResponseBody
	GetData() *DescribeInspectionTasksResponseBodyData
	SetRequestId(v string) *DescribeInspectionTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInspectionTasksResponseBody
	GetSuccess() *bool
}

type DescribeInspectionTasksResponseBody struct {
	Data *DescribeInspectionTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2D9F3768-EDA9-4811-943E-42C8006E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInspectionTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTasksResponseBody) GetData() *DescribeInspectionTasksResponseBodyData {
	return s.Data
}

func (s *DescribeInspectionTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInspectionTasksResponseBody) SetData(v *DescribeInspectionTasksResponseBodyData) *DescribeInspectionTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInspectionTasksResponseBody) SetRequestId(v string) *DescribeInspectionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionTasksResponseBody) SetSuccess(v bool) *DescribeInspectionTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInspectionTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInspectionTasksResponseBodyData struct {
	Items []*DescribeInspectionTasksResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInspectionTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTasksResponseBodyData) GetItems() []*DescribeInspectionTasksResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInspectionTasksResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeInspectionTasksResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionTasksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeInspectionTasksResponseBodyData) SetItems(v []*DescribeInspectionTasksResponseBodyDataItems) *DescribeInspectionTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInspectionTasksResponseBodyData) SetPageNum(v int64) *DescribeInspectionTasksResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyData) SetPageSize(v int64) *DescribeInspectionTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyData) SetTotal(v int64) *DescribeInspectionTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyData) Validate() error {
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

type DescribeInspectionTasksResponseBodyDataItems struct {
	// example:
	//
	// 2024-07-01T02:06:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-09-23T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2026-01-09T02:13:01Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// PERFORMANCE_METRICS
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// example:
	//
	// r-uf6ns8txov3mp9jxxx
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// example:
	//
	// sch-4dfb08ddf9f84855bacca35axxx
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// example:
	//
	// scheduler|
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2026-05-30T02:11:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tit-dca42f85c73644e0ab5c80ef641xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTasksResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeInspectionTasksResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetCreateTime(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetEndTime(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetFinishTime(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetInspectionItems(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.InspectionItems = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetInstanceIds(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetReportLanguage(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.ReportLanguage = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetScheduleId(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.ScheduleId = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetSource(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetStartTime(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetStatus(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) SetTaskId(v string) *DescribeInspectionTasksResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTasksResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
