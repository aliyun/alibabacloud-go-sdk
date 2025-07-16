// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateScheduleConferenceShrinkRequest
	GetEndTime() *int64
	SetScheduleConfSettingModelShrink(v string) *CreateScheduleConferenceShrinkRequest
	GetScheduleConfSettingModelShrink() *string
	SetStartTime(v int64) *CreateScheduleConferenceShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *CreateScheduleConferenceShrinkRequest
	GetTenantContextShrink() *string
	SetTitle(v string) *CreateScheduleConferenceShrinkRequest
	GetTitle() *string
}

type CreateScheduleConferenceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000L
	EndTime                        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ScheduleConfSettingModelShrink *string `json:"ScheduleConfSettingModel,omitempty" xml:"ScheduleConfSettingModel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687924800000L
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 预约会议标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateScheduleConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateScheduleConferenceShrinkRequest) GetScheduleConfSettingModelShrink() *string {
	return s.ScheduleConfSettingModelShrink
}

func (s *CreateScheduleConferenceShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateScheduleConferenceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateScheduleConferenceShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateScheduleConferenceShrinkRequest) SetEndTime(v int64) *CreateScheduleConferenceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleConferenceShrinkRequest) SetScheduleConfSettingModelShrink(v string) *CreateScheduleConferenceShrinkRequest {
	s.ScheduleConfSettingModelShrink = &v
	return s
}

func (s *CreateScheduleConferenceShrinkRequest) SetStartTime(v int64) *CreateScheduleConferenceShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateScheduleConferenceShrinkRequest) SetTenantContextShrink(v string) *CreateScheduleConferenceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateScheduleConferenceShrinkRequest) SetTitle(v string) *CreateScheduleConferenceShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateScheduleConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
