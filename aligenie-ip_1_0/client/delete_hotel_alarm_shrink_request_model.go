// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmsShrink(v string) *DeleteHotelAlarmShrinkRequest
	GetAlarmsShrink() *string
	SetHotelId(v string) *DeleteHotelAlarmShrinkRequest
	GetHotelId() *string
}

type DeleteHotelAlarmShrinkRequest struct {
	// This parameter is required.
	AlarmsShrink *string `json:"Alarms,omitempty" xml:"Alarms,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteHotelAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmShrinkRequest) GetAlarmsShrink() *string {
	return s.AlarmsShrink
}

func (s *DeleteHotelAlarmShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteHotelAlarmShrinkRequest) SetAlarmsShrink(v string) *DeleteHotelAlarmShrinkRequest {
	s.AlarmsShrink = &v
	return s
}

func (s *DeleteHotelAlarmShrinkRequest) SetHotelId(v string) *DeleteHotelAlarmShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
