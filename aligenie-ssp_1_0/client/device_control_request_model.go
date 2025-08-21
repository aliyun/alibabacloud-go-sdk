// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlRequest(v *DeviceControlRequestControlRequest) *DeviceControlRequest
	GetControlRequest() *DeviceControlRequestControlRequest
	SetDeviceInfo(v *DeviceControlRequestDeviceInfo) *DeviceControlRequest
	GetDeviceInfo() *DeviceControlRequestDeviceInfo
}

type DeviceControlRequest struct {
	ControlRequest *DeviceControlRequestControlRequest `json:"ControlRequest,omitempty" xml:"ControlRequest,omitempty" type:"Struct"`
	// This parameter is required.
	DeviceInfo *DeviceControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s DeviceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequest) GetControlRequest() *DeviceControlRequestControlRequest {
	return s.ControlRequest
}

func (s *DeviceControlRequest) GetDeviceInfo() *DeviceControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *DeviceControlRequest) SetControlRequest(v *DeviceControlRequestControlRequest) *DeviceControlRequest {
	s.ControlRequest = v
	return s
}

func (s *DeviceControlRequest) SetDeviceInfo(v *DeviceControlRequestDeviceInfo) *DeviceControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeviceControlRequest) Validate() error {
	return dara.Validate(s)
}

type DeviceControlRequestControlRequest struct {
	// example:
	//
	// false
	Muted *bool `json:"Muted,omitempty" xml:"Muted,omitempty"`
	// example:
	//
	// 10
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DeviceControlRequestControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequestControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestControlRequest) GetMuted() *bool {
	return s.Muted
}

func (s *DeviceControlRequestControlRequest) GetVolume() *int32 {
	return s.Volume
}

func (s *DeviceControlRequestControlRequest) SetMuted(v bool) *DeviceControlRequestControlRequest {
	s.Muted = &v
	return s
}

func (s *DeviceControlRequestControlRequest) SetVolume(v int32) *DeviceControlRequestControlRequest {
	s.Volume = &v
	return s
}

func (s *DeviceControlRequestControlRequest) Validate() error {
	return dara.Validate(s)
}

type DeviceControlRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1*****2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeviceControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeviceControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeviceControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *DeviceControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeviceControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeviceControlRequestDeviceInfo) SetEncodeKey(v string) *DeviceControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetEncodeType(v string) *DeviceControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetId(v string) *DeviceControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetIdType(v string) *DeviceControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) SetOrganizationId(v string) *DeviceControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeviceControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
