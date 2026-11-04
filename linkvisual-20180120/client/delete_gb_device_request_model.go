// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGbDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteGbDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteGbDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteGbDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteGbDeviceRequest
	GetProductKey() *string
}

type DeleteGbDeviceRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteGbDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteGbDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteGbDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteGbDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteGbDeviceRequest) SetDeviceName(v string) *DeleteGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetIotId(v string) *DeleteGbDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetIotInstanceId(v string) *DeleteGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteGbDeviceRequest) SetProductKey(v string) *DeleteGbDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteGbDeviceRequest) Validate() error {
	return dara.Validate(s)
}
