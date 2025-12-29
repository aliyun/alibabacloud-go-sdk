// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelControlDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *ListHotelControlDeviceRequestUserInfo) *ListHotelControlDeviceRequest
	GetUserInfo() *ListHotelControlDeviceRequestUserInfo
}

type ListHotelControlDeviceRequest struct {
	UserInfo *ListHotelControlDeviceRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelControlDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequest) GetUserInfo() *ListHotelControlDeviceRequestUserInfo {
	return s.UserInfo
}

func (s *ListHotelControlDeviceRequest) SetUserInfo(v *ListHotelControlDeviceRequestUserInfo) *ListHotelControlDeviceRequest {
	s.UserInfo = v
	return s
}

func (s *ListHotelControlDeviceRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelControlDeviceRequestUserInfo struct {
	// This parameter is required.
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelControlDeviceRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListHotelControlDeviceRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListHotelControlDeviceRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListHotelControlDeviceRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListHotelControlDeviceRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeKey(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetEncodeType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetIdType(v string) *ListHotelControlDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) SetOrganizationId(v string) *ListHotelControlDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListHotelControlDeviceRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
