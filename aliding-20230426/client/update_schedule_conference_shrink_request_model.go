// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *UpdateScheduleConferenceShrinkRequest
	GetEndTime() *int64
	SetScheduleConferenceId(v string) *UpdateScheduleConferenceShrinkRequest
	GetScheduleConferenceId() *string
	SetStartTime(v int64) *UpdateScheduleConferenceShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *UpdateScheduleConferenceShrinkRequest
	GetTenantContextShrink() *string
	SetTitle(v string) *UpdateScheduleConferenceShrinkRequest
	GetTitle() *string
}

type UpdateScheduleConferenceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000L
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
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

func (s UpdateScheduleConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateScheduleConferenceShrinkRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *UpdateScheduleConferenceShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateScheduleConferenceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateScheduleConferenceShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateScheduleConferenceShrinkRequest) SetEndTime(v int64) *UpdateScheduleConferenceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkRequest) SetScheduleConferenceId(v string) *UpdateScheduleConferenceShrinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkRequest) SetStartTime(v int64) *UpdateScheduleConferenceShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkRequest) SetTenantContextShrink(v string) *UpdateScheduleConferenceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkRequest) SetTitle(v string) *UpdateScheduleConferenceShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
