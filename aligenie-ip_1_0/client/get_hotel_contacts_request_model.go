// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelContactsRequestUserInfo) *GetHotelContactsRequest
	GetUserInfo() *GetHotelContactsRequestUserInfo
}

type GetHotelContactsRequest struct {
	UserInfo *GetHotelContactsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactsRequest) GetUserInfo() *GetHotelContactsRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelContactsRequest) SetUserInfo(v *GetHotelContactsRequestUserInfo) *GetHotelContactsRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelContactsRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelContactsRequestUserInfo struct {
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

func (s GetHotelContactsRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactsRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelContactsRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelContactsRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelContactsRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelContactsRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelContactsRequestUserInfo) SetEncodeKey(v string) *GetHotelContactsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetEncodeType(v string) *GetHotelContactsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetId(v string) *GetHotelContactsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetIdType(v string) *GetHotelContactsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) SetOrganizationId(v string) *GetHotelContactsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelContactsRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
