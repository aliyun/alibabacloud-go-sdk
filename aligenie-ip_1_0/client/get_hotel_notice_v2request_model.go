// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelNoticeV2RequestUserInfo) *GetHotelNoticeV2Request
	GetUserInfo() *GetHotelNoticeV2RequestUserInfo
}

type GetHotelNoticeV2Request struct {
	// This parameter is required.
	UserInfo *GetHotelNoticeV2RequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelNoticeV2Request) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2Request) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2Request) GetUserInfo() *GetHotelNoticeV2RequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelNoticeV2Request) SetUserInfo(v *GetHotelNoticeV2RequestUserInfo) *GetHotelNoticeV2Request {
	s.UserInfo = v
	return s
}

func (s *GetHotelNoticeV2Request) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelNoticeV2RequestUserInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelNoticeV2RequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2RequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2RequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelNoticeV2RequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelNoticeV2RequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelNoticeV2RequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelNoticeV2RequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelNoticeV2RequestUserInfo) SetEncodeKey(v string) *GetHotelNoticeV2RequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetEncodeType(v string) *GetHotelNoticeV2RequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetId(v string) *GetHotelNoticeV2RequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetIdType(v string) *GetHotelNoticeV2RequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) SetOrganizationId(v string) *GetHotelNoticeV2RequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelNoticeV2RequestUserInfo) Validate() error {
	return dara.Validate(s)
}
