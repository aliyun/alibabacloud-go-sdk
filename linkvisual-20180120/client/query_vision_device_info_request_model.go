// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVisionDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryVisionDeviceInfoRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryVisionDeviceInfoRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryVisionDeviceInfoRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryVisionDeviceInfoRequest
	GetProductKey() *string
}

type QueryVisionDeviceInfoRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryVisionDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryVisionDeviceInfoRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryVisionDeviceInfoRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryVisionDeviceInfoRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryVisionDeviceInfoRequest) SetDeviceName(v string) *QueryVisionDeviceInfoRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetIotId(v string) *QueryVisionDeviceInfoRequest {
	s.IotId = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetIotInstanceId(v string) *QueryVisionDeviceInfoRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) SetProductKey(v string) *QueryVisionDeviceInfoRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryVisionDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
