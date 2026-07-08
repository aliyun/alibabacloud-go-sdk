// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarm(v int32) *CreateDeviceAlarmRequest
	GetAlarm() *int32
	SetChannelId(v int32) *CreateDeviceAlarmRequest
	GetChannelId() *int32
	SetEndTime(v int64) *CreateDeviceAlarmRequest
	GetEndTime() *int64
	SetExpire(v int64) *CreateDeviceAlarmRequest
	GetExpire() *int64
	SetId(v string) *CreateDeviceAlarmRequest
	GetId() *string
	SetObjectType(v int32) *CreateDeviceAlarmRequest
	GetObjectType() *int32
	SetOwnerId(v int64) *CreateDeviceAlarmRequest
	GetOwnerId() *int64
	SetStartTime(v int64) *CreateDeviceAlarmRequest
	GetStartTime() *int64
	SetSubAlarm(v int32) *CreateDeviceAlarmRequest
	GetSubAlarm() *int32
}

type CreateDeviceAlarmRequest struct {
	// The Alarm Metric. Valid values:
	//
	// - 0: other
	//
	// - 1: motion detection alerting
	//
	// - 2: intelligent alerting
	//
	// - 3: switch alerting
	//
	// - 4: video loss alerting.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Alarm *int32 `json:"Alarm,omitempty" xml:"Alarm,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ChannelId *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The end time. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1632314789000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expiration time of the media upload URL. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The device ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The media object type. Valid values:
	//
	// - 0: none
	//
	// - 1: JPEG image
	//
	// - 2: video
	//
	// - 3: GIF image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ObjectType *int32 `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1632121707000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The alert subtype.
	//
	// This parameter is required.
	SubAlarm *int32 `json:"SubAlarm,omitempty" xml:"SubAlarm,omitempty"`
}

func (s CreateDeviceAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceAlarmRequest) GetAlarm() *int32 {
	return s.Alarm
}

func (s *CreateDeviceAlarmRequest) GetChannelId() *int32 {
	return s.ChannelId
}

func (s *CreateDeviceAlarmRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateDeviceAlarmRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *CreateDeviceAlarmRequest) GetId() *string {
	return s.Id
}

func (s *CreateDeviceAlarmRequest) GetObjectType() *int32 {
	return s.ObjectType
}

func (s *CreateDeviceAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDeviceAlarmRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateDeviceAlarmRequest) GetSubAlarm() *int32 {
	return s.SubAlarm
}

func (s *CreateDeviceAlarmRequest) SetAlarm(v int32) *CreateDeviceAlarmRequest {
	s.Alarm = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetChannelId(v int32) *CreateDeviceAlarmRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetEndTime(v int64) *CreateDeviceAlarmRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetExpire(v int64) *CreateDeviceAlarmRequest {
	s.Expire = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetId(v string) *CreateDeviceAlarmRequest {
	s.Id = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetObjectType(v int32) *CreateDeviceAlarmRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetOwnerId(v int64) *CreateDeviceAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetStartTime(v int64) *CreateDeviceAlarmRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDeviceAlarmRequest) SetSubAlarm(v int32) *CreateDeviceAlarmRequest {
	s.SubAlarm = &v
	return s
}

func (s *CreateDeviceAlarmRequest) Validate() error {
	return dara.Validate(s)
}
