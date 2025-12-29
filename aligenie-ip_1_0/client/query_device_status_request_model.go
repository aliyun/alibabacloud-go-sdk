// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *QueryDeviceStatusRequestPayload) *QueryDeviceStatusRequest
	GetPayload() *QueryDeviceStatusRequestPayload
	SetUserInfo(v *QueryDeviceStatusRequestUserInfo) *QueryDeviceStatusRequest
	GetUserInfo() *QueryDeviceStatusRequestUserInfo
}

type QueryDeviceStatusRequest struct {
	Payload  *QueryDeviceStatusRequestPayload  `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo *QueryDeviceStatusRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryDeviceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequest) GetPayload() *QueryDeviceStatusRequestPayload {
	return s.Payload
}

func (s *QueryDeviceStatusRequest) GetUserInfo() *QueryDeviceStatusRequestUserInfo {
	return s.UserInfo
}

func (s *QueryDeviceStatusRequest) SetPayload(v *QueryDeviceStatusRequestPayload) *QueryDeviceStatusRequest {
	s.Payload = v
	return s
}

func (s *QueryDeviceStatusRequest) SetUserInfo(v *QueryDeviceStatusRequestUserInfo) *QueryDeviceStatusRequest {
	s.UserInfo = v
	return s
}

func (s *QueryDeviceStatusRequest) Validate() error {
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

type QueryDeviceStatusRequestPayload struct {
	LocationDevices []*QueryDeviceStatusRequestPayloadLocationDevices `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty" type:"Repeated"`
	Properties      map[string]*string                                `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s QueryDeviceStatusRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusRequestPayload) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayload) GetLocationDevices() []*QueryDeviceStatusRequestPayloadLocationDevices {
	return s.LocationDevices
}

func (s *QueryDeviceStatusRequestPayload) GetProperties() map[string]*string {
	return s.Properties
}

func (s *QueryDeviceStatusRequestPayload) SetLocationDevices(v []*QueryDeviceStatusRequestPayloadLocationDevices) *QueryDeviceStatusRequestPayload {
	s.LocationDevices = v
	return s
}

func (s *QueryDeviceStatusRequestPayload) SetProperties(v map[string]*string) *QueryDeviceStatusRequestPayload {
	s.Properties = v
	return s
}

func (s *QueryDeviceStatusRequestPayload) Validate() error {
	if s.LocationDevices != nil {
		for _, item := range s.LocationDevices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDeviceStatusRequestPayloadLocationDevices struct {
	// example:
	//
	// night_light
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// light
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// room
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusRequestPayloadLocationDevices) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) GetLocation() *string {
	return s.Location
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceNumber(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceNumber = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetDeviceType(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) SetLocation(v string) *QueryDeviceStatusRequestPayloadLocationDevices {
	s.Location = &v
	return s
}

func (s *QueryDeviceStatusRequestPayloadLocationDevices) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceStatusRequestUserInfo struct {
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
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s QueryDeviceStatusRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryDeviceStatusRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryDeviceStatusRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *QueryDeviceStatusRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *QueryDeviceStatusRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeKey(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetEncodeType(v string) *QueryDeviceStatusRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetId(v string) *QueryDeviceStatusRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetIdType(v string) *QueryDeviceStatusRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) SetOrganizationId(v string) *QueryDeviceStatusRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *QueryDeviceStatusRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
