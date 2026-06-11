// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeviceUpdateRequest
	GetAppId() *string
	SetDeviceName(v string) *DeviceUpdateRequest
	GetDeviceName() *string
	SetRemark(v string) *DeviceUpdateRequest
	GetRemark() *string
	SetStatus(v int32) *DeviceUpdateRequest
	GetStatus() *int32
}

type DeviceUpdateRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeviceUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceUpdateRequest) GoString() string {
	return s.String()
}

func (s *DeviceUpdateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeviceUpdateRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeviceUpdateRequest) GetRemark() *string {
	return s.Remark
}

func (s *DeviceUpdateRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DeviceUpdateRequest) SetAppId(v string) *DeviceUpdateRequest {
	s.AppId = &v
	return s
}

func (s *DeviceUpdateRequest) SetDeviceName(v string) *DeviceUpdateRequest {
	s.DeviceName = &v
	return s
}

func (s *DeviceUpdateRequest) SetRemark(v string) *DeviceUpdateRequest {
	s.Remark = &v
	return s
}

func (s *DeviceUpdateRequest) SetStatus(v int32) *DeviceUpdateRequest {
	s.Status = &v
	return s
}

func (s *DeviceUpdateRequest) Validate() error {
	return dara.Validate(s)
}
