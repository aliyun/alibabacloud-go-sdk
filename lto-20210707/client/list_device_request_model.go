// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *ListDeviceRequest
	GetDeviceGroupId() *string
	SetName(v string) *ListDeviceRequest
	GetName() *string
	SetNum(v int32) *ListDeviceRequest
	GetNum() *int32
	SetRegionId(v string) *ListDeviceRequest
	GetRegionId() *string
	SetSize(v int32) *ListDeviceRequest
	GetSize() *int32
}

type ListDeviceRequest struct {
	// This parameter is required.
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListDeviceRequest) GetName() *string {
	return s.Name
}

func (s *ListDeviceRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListDeviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDeviceRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDeviceRequest) SetDeviceGroupId(v string) *ListDeviceRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceRequest) SetName(v string) *ListDeviceRequest {
	s.Name = &v
	return s
}

func (s *ListDeviceRequest) SetNum(v int32) *ListDeviceRequest {
	s.Num = &v
	return s
}

func (s *ListDeviceRequest) SetRegionId(v string) *ListDeviceRequest {
	s.RegionId = &v
	return s
}

func (s *ListDeviceRequest) SetSize(v int32) *ListDeviceRequest {
	s.Size = &v
	return s
}

func (s *ListDeviceRequest) Validate() error {
	return dara.Validate(s)
}
