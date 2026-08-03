// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInspectionScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyInspectionScheduleRequest
	GetCronExpression() *string
	SetEnabled(v int64) *ModifyInspectionScheduleRequest
	GetEnabled() *int64
	SetInspectionItems(v string) *ModifyInspectionScheduleRequest
	GetInspectionItems() *string
	SetInspectionWindow(v string) *ModifyInspectionScheduleRequest
	GetInspectionWindow() *string
	SetInstanceId(v string) *ModifyInspectionScheduleRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *ModifyInspectionScheduleRequest
	GetInstanceIds() *string
	SetReportLanguage(v string) *ModifyInspectionScheduleRequest
	GetReportLanguage() *string
	SetScheduleId(v string) *ModifyInspectionScheduleRequest
	GetScheduleId() *string
	SetScheduleName(v string) *ModifyInspectionScheduleRequest
	GetScheduleName() *string
	SetTimezone(v string) *ModifyInspectionScheduleRequest
	GetTimezone() *string
}

type ModifyInspectionScheduleRequest struct {
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// false
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// HOTKEY
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// example:
	//
	// 1h
	InspectionWindow *string `json:"InspectionWindow,omitempty" xml:"InspectionWindow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// r-bp1zxszhcgatnx****
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
	// test-sche
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s ModifyInspectionScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInspectionScheduleRequest) GoString() string {
	return s.String()
}

func (s *ModifyInspectionScheduleRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyInspectionScheduleRequest) GetEnabled() *int64 {
	return s.Enabled
}

func (s *ModifyInspectionScheduleRequest) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *ModifyInspectionScheduleRequest) GetInspectionWindow() *string {
	return s.InspectionWindow
}

func (s *ModifyInspectionScheduleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInspectionScheduleRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyInspectionScheduleRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *ModifyInspectionScheduleRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *ModifyInspectionScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *ModifyInspectionScheduleRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *ModifyInspectionScheduleRequest) SetCronExpression(v string) *ModifyInspectionScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetEnabled(v int64) *ModifyInspectionScheduleRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetInspectionItems(v string) *ModifyInspectionScheduleRequest {
	s.InspectionItems = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetInspectionWindow(v string) *ModifyInspectionScheduleRequest {
	s.InspectionWindow = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetInstanceId(v string) *ModifyInspectionScheduleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetInstanceIds(v string) *ModifyInspectionScheduleRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetReportLanguage(v string) *ModifyInspectionScheduleRequest {
	s.ReportLanguage = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetScheduleId(v string) *ModifyInspectionScheduleRequest {
	s.ScheduleId = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetScheduleName(v string) *ModifyInspectionScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) SetTimezone(v string) *ModifyInspectionScheduleRequest {
	s.Timezone = &v
	return s
}

func (s *ModifyInspectionScheduleRequest) Validate() error {
	return dara.Validate(s)
}
