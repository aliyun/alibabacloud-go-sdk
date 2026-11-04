// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGbDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGbDeviceRequest
	GetDescription() *string
	SetDeviceName(v string) *CreateGbDeviceRequest
	GetDeviceName() *string
	SetDeviceType(v int32) *CreateGbDeviceRequest
	GetDeviceType() *int32
	SetGbId(v string) *CreateGbDeviceRequest
	GetGbId() *string
	SetIotInstanceId(v string) *CreateGbDeviceRequest
	GetIotInstanceId() *string
	SetMediaNetProtocol(v string) *CreateGbDeviceRequest
	GetMediaNetProtocol() *string
	SetPassword(v string) *CreateGbDeviceRequest
	GetPassword() *string
	SetProductKey(v string) *CreateGbDeviceRequest
	GetProductKey() *string
	SetSubProductKey(v string) *CreateGbDeviceRequest
	GetSubProductKey() *string
}

type CreateGbDeviceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType  *int32  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// This parameter is required.
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// This parameter is required.
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	MediaNetProtocol *string `json:"MediaNetProtocol,omitempty" xml:"MediaNetProtocol,omitempty"`
	// This parameter is required.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	SubProductKey *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s CreateGbDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGbDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateGbDeviceRequest) GetDeviceType() *int32 {
	return s.DeviceType
}

func (s *CreateGbDeviceRequest) GetGbId() *string {
	return s.GbId
}

func (s *CreateGbDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateGbDeviceRequest) GetMediaNetProtocol() *string {
	return s.MediaNetProtocol
}

func (s *CreateGbDeviceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateGbDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateGbDeviceRequest) GetSubProductKey() *string {
	return s.SubProductKey
}

func (s *CreateGbDeviceRequest) SetDescription(v string) *CreateGbDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateGbDeviceRequest) SetDeviceName(v string) *CreateGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateGbDeviceRequest) SetDeviceType(v int32) *CreateGbDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateGbDeviceRequest) SetGbId(v string) *CreateGbDeviceRequest {
	s.GbId = &v
	return s
}

func (s *CreateGbDeviceRequest) SetIotInstanceId(v string) *CreateGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateGbDeviceRequest) SetMediaNetProtocol(v string) *CreateGbDeviceRequest {
	s.MediaNetProtocol = &v
	return s
}

func (s *CreateGbDeviceRequest) SetPassword(v string) *CreateGbDeviceRequest {
	s.Password = &v
	return s
}

func (s *CreateGbDeviceRequest) SetProductKey(v string) *CreateGbDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateGbDeviceRequest) SetSubProductKey(v string) *CreateGbDeviceRequest {
	s.SubProductKey = &v
	return s
}

func (s *CreateGbDeviceRequest) Validate() error {
	return dara.Validate(s)
}
