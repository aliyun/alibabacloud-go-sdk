// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoomCheckOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *RoomCheckOutRequestDeviceInfo) *RoomCheckOutRequest
	GetDeviceInfo() *RoomCheckOutRequestDeviceInfo
	SetUserInfo(v *RoomCheckOutRequestUserInfo) *RoomCheckOutRequest
	GetUserInfo() *RoomCheckOutRequestUserInfo
}

type RoomCheckOutRequest struct {
	DeviceInfo *RoomCheckOutRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *RoomCheckOutRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s RoomCheckOutRequest) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutRequest) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequest) GetDeviceInfo() *RoomCheckOutRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *RoomCheckOutRequest) GetUserInfo() *RoomCheckOutRequestUserInfo {
	return s.UserInfo
}

func (s *RoomCheckOutRequest) SetDeviceInfo(v *RoomCheckOutRequestDeviceInfo) *RoomCheckOutRequest {
	s.DeviceInfo = v
	return s
}

func (s *RoomCheckOutRequest) SetUserInfo(v *RoomCheckOutRequestUserInfo) *RoomCheckOutRequest {
	s.UserInfo = v
	return s
}

func (s *RoomCheckOutRequest) Validate() error {
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

type RoomCheckOutRequestDeviceInfo struct {
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

func (s RoomCheckOutRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *RoomCheckOutRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *RoomCheckOutRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *RoomCheckOutRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *RoomCheckOutRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeKey(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetEncodeType(v string) *RoomCheckOutRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetId(v string) *RoomCheckOutRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetIdType(v string) *RoomCheckOutRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) SetOrganizationId(v string) *RoomCheckOutRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *RoomCheckOutRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type RoomCheckOutRequestUserInfo struct {
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

func (s RoomCheckOutRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutRequestUserInfo) GoString() string {
	return s.String()
}

func (s *RoomCheckOutRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *RoomCheckOutRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *RoomCheckOutRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *RoomCheckOutRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *RoomCheckOutRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeKey(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetEncodeType(v string) *RoomCheckOutRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetId(v string) *RoomCheckOutRequestUserInfo {
	s.Id = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetIdType(v string) *RoomCheckOutRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) SetOrganizationId(v string) *RoomCheckOutRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *RoomCheckOutRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
