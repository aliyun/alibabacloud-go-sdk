// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGbDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateGbDeviceRequest
	GetDescription() *string
	SetDeviceName(v string) *UpdateGbDeviceRequest
	GetDeviceName() *string
	SetGbId(v string) *UpdateGbDeviceRequest
	GetGbId() *string
	SetIotId(v string) *UpdateGbDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *UpdateGbDeviceRequest
	GetIotInstanceId() *string
	SetPassword(v string) *UpdateGbDeviceRequest
	GetPassword() *string
	SetProductKey(v string) *UpdateGbDeviceRequest
	GetProductKey() *string
}

type UpdateGbDeviceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceName  *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	GbId        *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s UpdateGbDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGbDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGbDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *UpdateGbDeviceRequest) GetGbId() *string {
	return s.GbId
}

func (s *UpdateGbDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *UpdateGbDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *UpdateGbDeviceRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateGbDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *UpdateGbDeviceRequest) SetDescription(v string) *UpdateGbDeviceRequest {
	s.Description = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetDeviceName(v string) *UpdateGbDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetGbId(v string) *UpdateGbDeviceRequest {
	s.GbId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetIotId(v string) *UpdateGbDeviceRequest {
	s.IotId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetIotInstanceId(v string) *UpdateGbDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetPassword(v string) *UpdateGbDeviceRequest {
	s.Password = &v
	return s
}

func (s *UpdateGbDeviceRequest) SetProductKey(v string) *UpdateGbDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateGbDeviceRequest) Validate() error {
	return dara.Validate(s)
}
