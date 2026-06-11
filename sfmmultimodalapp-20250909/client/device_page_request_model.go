// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevicePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DevicePageRequest
	GetAppId() *string
	SetDeviceName(v string) *DevicePageRequest
	GetDeviceName() *string
	SetPageNumber(v int32) *DevicePageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DevicePageRequest
	GetPageSize() *int32
}

type DevicePageRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DevicePageRequest) String() string {
	return dara.Prettify(s)
}

func (s DevicePageRequest) GoString() string {
	return s.String()
}

func (s *DevicePageRequest) GetAppId() *string {
	return s.AppId
}

func (s *DevicePageRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DevicePageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DevicePageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DevicePageRequest) SetAppId(v string) *DevicePageRequest {
	s.AppId = &v
	return s
}

func (s *DevicePageRequest) SetDeviceName(v string) *DevicePageRequest {
	s.DeviceName = &v
	return s
}

func (s *DevicePageRequest) SetPageNumber(v int32) *DevicePageRequest {
	s.PageNumber = &v
	return s
}

func (s *DevicePageRequest) SetPageSize(v int32) *DevicePageRequest {
	s.PageSize = &v
	return s
}

func (s *DevicePageRequest) Validate() error {
	return dara.Validate(s)
}
