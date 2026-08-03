// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateInspectionScheduleRequest
	GetCronExpression() *string
	SetInspectionItems(v string) *CreateInspectionScheduleRequest
	GetInspectionItems() *string
	SetInspectionWindow(v string) *CreateInspectionScheduleRequest
	GetInspectionWindow() *string
	SetInstanceId(v string) *CreateInspectionScheduleRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *CreateInspectionScheduleRequest
	GetInstanceIds() *string
	SetReportLanguage(v string) *CreateInspectionScheduleRequest
	GetReportLanguage() *string
	SetScheduleName(v string) *CreateInspectionScheduleRequest
	GetScheduleName() *string
	SetSecurityToken(v string) *CreateInspectionScheduleRequest
	GetSecurityToken() *string
	SetTimezone(v string) *CreateInspectionScheduleRequest
	GetTimezone() *string
}

type CreateInspectionScheduleRequest struct {
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// RESOURCE_USAGE
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// example:
	//
	// 1h
	InspectionWindow *string `json:"InspectionWindow,omitempty" xml:"InspectionWindow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ta-bp11iljddg37xxxx
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
	// test-sche
	ScheduleName  *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateInspectionScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateInspectionScheduleRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateInspectionScheduleRequest) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *CreateInspectionScheduleRequest) GetInspectionWindow() *string {
	return s.InspectionWindow
}

func (s *CreateInspectionScheduleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInspectionScheduleRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateInspectionScheduleRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *CreateInspectionScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *CreateInspectionScheduleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateInspectionScheduleRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateInspectionScheduleRequest) SetCronExpression(v string) *CreateInspectionScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetInspectionItems(v string) *CreateInspectionScheduleRequest {
	s.InspectionItems = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetInspectionWindow(v string) *CreateInspectionScheduleRequest {
	s.InspectionWindow = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetInstanceId(v string) *CreateInspectionScheduleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetInstanceIds(v string) *CreateInspectionScheduleRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetReportLanguage(v string) *CreateInspectionScheduleRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetScheduleName(v string) *CreateInspectionScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetSecurityToken(v string) *CreateInspectionScheduleRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateInspectionScheduleRequest) SetTimezone(v string) *CreateInspectionScheduleRequest {
	s.Timezone = &v
	return s
}

func (s *CreateInspectionScheduleRequest) Validate() error {
	return dara.Validate(s)
}
