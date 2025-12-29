// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelScreenSaverRequestUserInfo) *GetHotelScreenSaverRequest
	GetUserInfo() *GetHotelScreenSaverRequestUserInfo
}

type GetHotelScreenSaverRequest struct {
	// This parameter is required.
	UserInfo *GetHotelScreenSaverRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelScreenSaverRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverRequest) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequest) GetUserInfo() *GetHotelScreenSaverRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelScreenSaverRequest) SetUserInfo(v *GetHotelScreenSaverRequestUserInfo) *GetHotelScreenSaverRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelScreenSaverRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelScreenSaverRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
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
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetHotelScreenSaverRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelScreenSaverRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelScreenSaverRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelScreenSaverRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelScreenSaverRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeKey(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetEncodeType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetIdType(v string) *GetHotelScreenSaverRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) SetOrganizationId(v string) *GetHotelScreenSaverRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelScreenSaverRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
