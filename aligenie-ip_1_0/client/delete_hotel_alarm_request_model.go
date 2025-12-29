// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v []*DeleteHotelAlarmRequestAlarms) *DeleteHotelAlarmRequest
	GetAlarms() []*DeleteHotelAlarmRequestAlarms
	SetHotelId(v string) *DeleteHotelAlarmRequest
	GetHotelId() *string
}

type DeleteHotelAlarmRequest struct {
	// This parameter is required.
	Alarms []*DeleteHotelAlarmRequestAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s DeleteHotelAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmRequest) GetAlarms() []*DeleteHotelAlarmRequestAlarms {
	return s.Alarms
}

func (s *DeleteHotelAlarmRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteHotelAlarmRequest) SetAlarms(v []*DeleteHotelAlarmRequestAlarms) *DeleteHotelAlarmRequest {
	s.Alarms = v
	return s
}

func (s *DeleteHotelAlarmRequest) SetHotelId(v string) *DeleteHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelAlarmRequest) Validate() error {
	if s.Alarms != nil {
		for _, item := range s.Alarms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteHotelAlarmRequestAlarms struct {
	// This parameter is required.
	//
	// example:
	//
	// 5029
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PvkB***TA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mgw/k***HQd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s DeleteHotelAlarmRequestAlarms) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmRequestAlarms) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmRequestAlarms) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *DeleteHotelAlarmRequestAlarms) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *DeleteHotelAlarmRequestAlarms) GetRoomNo() *string {
	return s.RoomNo
}

func (s *DeleteHotelAlarmRequestAlarms) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *DeleteHotelAlarmRequestAlarms) SetAlarmId(v int64) *DeleteHotelAlarmRequestAlarms {
	s.AlarmId = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetDeviceOpenId(v string) *DeleteHotelAlarmRequestAlarms {
	s.DeviceOpenId = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetRoomNo(v string) *DeleteHotelAlarmRequestAlarms {
	s.RoomNo = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) SetUserOpenId(v string) *DeleteHotelAlarmRequestAlarms {
	s.UserOpenId = &v
	return s
}

func (s *DeleteHotelAlarmRequestAlarms) Validate() error {
	return dara.Validate(s)
}
