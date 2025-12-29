// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *GetHotelContactByNumberRequest
	GetNumber() *string
	SetUserInfo(v *GetHotelContactByNumberRequestUserInfo) *GetHotelContactByNumberRequest
	GetUserInfo() *GetHotelContactByNumberRequestUserInfo
}

type GetHotelContactByNumberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	Number   *string                                 `json:"Number,omitempty" xml:"Number,omitempty"`
	UserInfo *GetHotelContactByNumberRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelContactByNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *GetHotelContactByNumberRequest) GetUserInfo() *GetHotelContactByNumberRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelContactByNumberRequest) SetNumber(v string) *GetHotelContactByNumberRequest {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberRequest) SetUserInfo(v *GetHotelContactByNumberRequestUserInfo) *GetHotelContactByNumberRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelContactByNumberRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelContactByNumberRequestUserInfo struct {
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

func (s GetHotelContactByNumberRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelContactByNumberRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelContactByNumberRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelContactByNumberRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelContactByNumberRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelContactByNumberRequestUserInfo) SetEncodeKey(v string) *GetHotelContactByNumberRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetEncodeType(v string) *GetHotelContactByNumberRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetId(v string) *GetHotelContactByNumberRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetIdType(v string) *GetHotelContactByNumberRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) SetOrganizationId(v string) *GetHotelContactByNumberRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelContactByNumberRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
