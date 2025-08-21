// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *ModifyDeviceAlarmRequest
	GetAlarmId() *string
	SetChannelId(v int32) *ModifyDeviceAlarmRequest
	GetChannelId() *int32
	SetId(v string) *ModifyDeviceAlarmRequest
	GetId() *string
	SetOwnerId(v int64) *ModifyDeviceAlarmRequest
	GetOwnerId() *int64
	SetStatus(v int32) *ModifyDeviceAlarmRequest
	GetStatus() *int32
}

type ModifyDeviceAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0dGo7jLwwf1000296232
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	ChannelId *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDeviceAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAlarmRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmRequest) GetAlarmId() *string {
	return s.AlarmId
}

func (s *ModifyDeviceAlarmRequest) GetChannelId() *int32 {
	return s.ChannelId
}

func (s *ModifyDeviceAlarmRequest) GetId() *string {
	return s.Id
}

func (s *ModifyDeviceAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeviceAlarmRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyDeviceAlarmRequest) SetAlarmId(v string) *ModifyDeviceAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetChannelId(v int32) *ModifyDeviceAlarmRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetId(v string) *ModifyDeviceAlarmRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetOwnerId(v int64) *ModifyDeviceAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) SetStatus(v int32) *ModifyDeviceAlarmRequest {
	s.Status = &v
	return s
}

func (s *ModifyDeviceAlarmRequest) Validate() error {
	return dara.Validate(s)
}
