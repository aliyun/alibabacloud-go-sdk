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
	// Input parameters for volume control
	ControlRequest *DeviceControlRequestControlRequest `json:"ControlRequest,omitempty" xml:"ControlRequest,omitempty" type:"Struct"`
	// List of device ID information.
	//
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
	if s.ControlRequest != nil {
		if err := s.ControlRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeviceControlRequestControlRequest struct {
	// Indicates whether mute is enabled. If this field is set to true, you must also specify the volume value as 0.
	//
	// example:
	//
	// false
	Muted *bool `json:"Muted,omitempty" xml:"Muted,omitempty"`
	// Target volume value
	//
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
	// Value corresponding to the encoding type. Enter the Project ID of the project where the product resides. You can View this in the Tmall Genie AI platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding Type. Enter PROJECT_ID here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID. Enter the value of deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID:
	//
	// OPEN_ID: The default Device ID identity.
	//
	// UNION_ID: The organization-dimension Device ID identity. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID of the device. Required if IdType is UNION_ID.
	//
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
