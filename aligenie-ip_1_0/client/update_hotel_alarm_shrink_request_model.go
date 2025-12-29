// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmsShrink(v string) *UpdateHotelAlarmShrinkRequest
	GetAlarmsShrink() *string
	SetHotelId(v string) *UpdateHotelAlarmShrinkRequest
	GetHotelId() *string
	SetScheduleInfoShrink(v string) *UpdateHotelAlarmShrinkRequest
	GetScheduleInfoShrink() *string
}

type UpdateHotelAlarmShrinkRequest struct {
	// This parameter is required.
	AlarmsShrink *string `json:"Alarms,omitempty" xml:"Alarms,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a381a668bc485980bed3876a75e013
	HotelId            *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	ScheduleInfoShrink *string `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty"`
}

func (s UpdateHotelAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmShrinkRequest) GetAlarmsShrink() *string {
	return s.AlarmsShrink
}

func (s *UpdateHotelAlarmShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelAlarmShrinkRequest) GetScheduleInfoShrink() *string {
	return s.ScheduleInfoShrink
}

func (s *UpdateHotelAlarmShrinkRequest) SetAlarmsShrink(v string) *UpdateHotelAlarmShrinkRequest {
	s.AlarmsShrink = &v
	return s
}

func (s *UpdateHotelAlarmShrinkRequest) SetHotelId(v string) *UpdateHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelAlarmShrinkRequest) SetScheduleInfoShrink(v string) *UpdateHotelAlarmShrinkRequest {
	s.ScheduleInfoShrink = &v
	return s
}

func (s *UpdateHotelAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
