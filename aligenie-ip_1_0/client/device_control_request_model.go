// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *DeviceControlRequestPayload) *DeviceControlRequest
	GetPayload() *DeviceControlRequestPayload
	SetUserInfo(v *DeviceControlRequestUserInfo) *DeviceControlRequest
	GetUserInfo() *DeviceControlRequestUserInfo
}

type DeviceControlRequest struct {
	Payload  *DeviceControlRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *DeviceControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeviceControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlRequest) GetPayload() *DeviceControlRequestPayload {
	return s.Payload
}

func (s *DeviceControlRequest) GetUserInfo() *DeviceControlRequestUserInfo {
	return s.UserInfo
}

func (s *DeviceControlRequest) SetPayload(v *DeviceControlRequestPayload) *DeviceControlRequest {
	s.Payload = v
	return s
}

func (s *DeviceControlRequest) SetUserInfo(v *DeviceControlRequestUserInfo) *DeviceControlRequest {
	s.UserInfo = v
	return s
}

func (s *DeviceControlRequest) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeviceControlRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// aircondition
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open
	Cmd *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// air_condition
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// {}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// room
	Location   *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s DeviceControlRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequestPayload) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestPayload) GetCategory() *string {
	return s.Category
}

func (s *DeviceControlRequestPayload) GetCmd() *string {
	return s.Cmd
}

func (s *DeviceControlRequestPayload) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *DeviceControlRequestPayload) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *DeviceControlRequestPayload) GetLocation() *string {
	return s.Location
}

func (s *DeviceControlRequestPayload) GetProperties() map[string]*string {
	return s.Properties
}

func (s *DeviceControlRequestPayload) SetCategory(v string) *DeviceControlRequestPayload {
	s.Category = &v
	return s
}

func (s *DeviceControlRequestPayload) SetCmd(v string) *DeviceControlRequestPayload {
	s.Cmd = &v
	return s
}

func (s *DeviceControlRequestPayload) SetDeviceNumber(v string) *DeviceControlRequestPayload {
	s.DeviceNumber = &v
	return s
}

func (s *DeviceControlRequestPayload) SetExtendInfo(v string) *DeviceControlRequestPayload {
	s.ExtendInfo = &v
	return s
}

func (s *DeviceControlRequestPayload) SetLocation(v string) *DeviceControlRequestPayload {
	s.Location = &v
	return s
}

func (s *DeviceControlRequestPayload) SetProperties(v map[string]*string) *DeviceControlRequestPayload {
	s.Properties = v
	return s
}

func (s *DeviceControlRequestPayload) Validate() error {
	return dara.Validate(s)
}

type DeviceControlRequestUserInfo struct {
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
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeviceControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeviceControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeviceControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeviceControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *DeviceControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeviceControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeviceControlRequestUserInfo) SetEncodeKey(v string) *DeviceControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetEncodeType(v string) *DeviceControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetId(v string) *DeviceControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetIdType(v string) *DeviceControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeviceControlRequestUserInfo) SetOrganizationId(v string) *DeviceControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeviceControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
