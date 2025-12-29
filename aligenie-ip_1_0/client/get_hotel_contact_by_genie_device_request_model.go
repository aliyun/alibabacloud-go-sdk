// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByGenieDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetHotelContactByGenieDeviceRequestDeviceInfo) *GetHotelContactByGenieDeviceRequest
	GetDeviceInfo() *GetHotelContactByGenieDeviceRequestDeviceInfo
	SetUserInfo(v *GetHotelContactByGenieDeviceRequestUserInfo) *GetHotelContactByGenieDeviceRequest
	GetUserInfo() *GetHotelContactByGenieDeviceRequestUserInfo
}

type GetHotelContactByGenieDeviceRequest struct {
	DeviceInfo *GetHotelContactByGenieDeviceRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *GetHotelContactByGenieDeviceRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactByGenieDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequest) GetDeviceInfo() *GetHotelContactByGenieDeviceRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetHotelContactByGenieDeviceRequest) GetUserInfo() *GetHotelContactByGenieDeviceRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelContactByGenieDeviceRequest) SetDeviceInfo(v *GetHotelContactByGenieDeviceRequestDeviceInfo) *GetHotelContactByGenieDeviceRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetHotelContactByGenieDeviceRequest) SetUserInfo(v *GetHotelContactByGenieDeviceRequestUserInfo) *GetHotelContactByGenieDeviceRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelContactByGenieDeviceRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
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

type GetHotelContactByGenieDeviceRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
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
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactByGenieDeviceRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetEncodeKey(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetEncodeType(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetId(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetIdType(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) SetOrganizationId(v string) *GetHotelContactByGenieDeviceRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetHotelContactByGenieDeviceRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
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
	// 1***2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelContactByGenieDeviceRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetEncodeKey(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetEncodeType(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetId(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetIdType(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) SetOrganizationId(v string) *GetHotelContactByGenieDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
