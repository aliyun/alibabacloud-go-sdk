// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v string) *ModifyDeviceChannelsRequest
	GetChannels() *string
	SetDeviceStatus(v string) *ModifyDeviceChannelsRequest
	GetDeviceStatus() *string
	SetDsn(v string) *ModifyDeviceChannelsRequest
	GetDsn() *string
	SetId(v string) *ModifyDeviceChannelsRequest
	GetId() *string
	SetOwnerId(v int64) *ModifyDeviceChannelsRequest
	GetOwnerId() *int64
}

type ModifyDeviceChannelsRequest struct {
	// This parameter is required.
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// on
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 210235C3GN32090008286cf17e130d
	Dsn *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyDeviceChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceChannelsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsRequest) GetChannels() *string {
	return s.Channels
}

func (s *ModifyDeviceChannelsRequest) GetDeviceStatus() *string {
	return s.DeviceStatus
}

func (s *ModifyDeviceChannelsRequest) GetDsn() *string {
	return s.Dsn
}

func (s *ModifyDeviceChannelsRequest) GetId() *string {
	return s.Id
}

func (s *ModifyDeviceChannelsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeviceChannelsRequest) SetChannels(v string) *ModifyDeviceChannelsRequest {
	s.Channels = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetDeviceStatus(v string) *ModifyDeviceChannelsRequest {
	s.DeviceStatus = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetDsn(v string) *ModifyDeviceChannelsRequest {
	s.Dsn = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetId(v string) *ModifyDeviceChannelsRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) SetOwnerId(v int64) *ModifyDeviceChannelsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceChannelsRequest) Validate() error {
	return dara.Validate(s)
}
