// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInspectionSchedulesResponseBodyData) *DescribeInspectionSchedulesResponseBody
	GetData() *DescribeInspectionSchedulesResponseBodyData
	SetRequestId(v string) *DescribeInspectionSchedulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInspectionSchedulesResponseBody
	GetSuccess() *bool
}

type DescribeInspectionSchedulesResponseBody struct {
	Data *DescribeInspectionSchedulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInspectionSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionSchedulesResponseBody) GetData() *DescribeInspectionSchedulesResponseBodyData {
	return s.Data
}

func (s *DescribeInspectionSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionSchedulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInspectionSchedulesResponseBody) SetData(v *DescribeInspectionSchedulesResponseBodyData) *DescribeInspectionSchedulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInspectionSchedulesResponseBody) SetRequestId(v string) *DescribeInspectionSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBody) SetSuccess(v bool) *DescribeInspectionSchedulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInspectionSchedulesResponseBodyData struct {
	Items []*DescribeInspectionSchedulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInspectionSchedulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionSchedulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInspectionSchedulesResponseBodyData) GetItems() []*DescribeInspectionSchedulesResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInspectionSchedulesResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeInspectionSchedulesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInspectionSchedulesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeInspectionSchedulesResponseBodyData) SetItems(v []*DescribeInspectionSchedulesResponseBodyDataItems) *DescribeInspectionSchedulesResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyData) SetPageNum(v int64) *DescribeInspectionSchedulesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyData) SetPageSize(v int64) *DescribeInspectionSchedulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyData) SetTotal(v int64) *DescribeInspectionSchedulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyData) Validate() error {
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

type DescribeInspectionSchedulesResponseBodyDataItems struct {
	// example:
	//
	// 2026-04-21T02:26:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0 0 3 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// true
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// HOTKEY
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// example:
	//
	// 1h
	InspectionWindow *string `json:"InspectionWindow,omitempty" xml:"InspectionWindow,omitempty"`
	// example:
	//
	// r-2zed6typz5j6djmb2x
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 2026-07-29T10:00:00Z
	NextFireTime *string `json:"NextFireTime,omitempty" xml:"NextFireTime,omitempty"`
	// example:
	//
	// {}
	NotifyConfig *string `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty"`
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// example:
	//
	// sch-b45811bf4bba46c8b6d233551da9xxxx
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// example:
	//
	// sch-test
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// 2026-07-29T06:50:04Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInspectionSchedulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionSchedulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetEnabled() *int64 {
	return s.Enabled
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetInspectionWindow() *string {
	return s.InspectionWindow
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetNextFireTime() *string {
	return s.NextFireTime
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetNotifyConfig() *string {
	return s.NotifyConfig
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetCreateTime(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetCronExpression(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.CronExpression = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetEnabled(v int64) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.Enabled = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetInspectionItems(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.InspectionItems = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetInspectionWindow(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.InspectionWindow = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetInstanceIds(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetNextFireTime(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.NextFireTime = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetNotifyConfig(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.NotifyConfig = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetReportLanguage(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.ReportLanguage = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetScheduleId(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.ScheduleId = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetScheduleName(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.ScheduleName = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetTimezone(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.Timezone = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) SetUpdateTime(v string) *DescribeInspectionSchedulesResponseBodyDataItems {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInspectionSchedulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
