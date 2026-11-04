// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteRtmpDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteRtmpDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteRtmpDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteRtmpDeviceRequest
	GetProductKey() *string
}

type DeleteRtmpDeviceRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteRtmpDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteRtmpDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteRtmpDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteRtmpDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteRtmpDeviceRequest) SetDeviceName(v string) *DeleteRtmpDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetIotId(v string) *DeleteRtmpDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetIotInstanceId(v string) *DeleteRtmpDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) SetProductKey(v string) *DeleteRtmpDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRtmpDeviceRequest) Validate() error {
	return dara.Validate(s)
}
