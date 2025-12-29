// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelNoticeRequestUserInfo) *GetHotelNoticeRequest
	GetUserInfo() *GetHotelNoticeRequestUserInfo
}

type GetHotelNoticeRequest struct {
	// This parameter is required.
	UserInfo *GetHotelNoticeRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeRequest) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeRequest) GetUserInfo() *GetHotelNoticeRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelNoticeRequest) SetUserInfo(v *GetHotelNoticeRequestUserInfo) *GetHotelNoticeRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelNoticeRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelNoticeRequestUserInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelNoticeRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelNoticeRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelNoticeRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelNoticeRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelNoticeRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelNoticeRequestUserInfo) SetEncodeKey(v string) *GetHotelNoticeRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetEncodeType(v string) *GetHotelNoticeRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetId(v string) *GetHotelNoticeRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetIdType(v string) *GetHotelNoticeRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) SetOrganizationId(v string) *GetHotelNoticeRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelNoticeRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
