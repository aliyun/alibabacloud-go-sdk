// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelHomeBackImageAndModesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelHomeBackImageAndModesRequestUserInfo) *GetHotelHomeBackImageAndModesRequest
	GetUserInfo() *GetHotelHomeBackImageAndModesRequestUserInfo
}

type GetHotelHomeBackImageAndModesRequest struct {
	// This parameter is required.
	UserInfo *GetHotelHomeBackImageAndModesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelHomeBackImageAndModesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequest) GetUserInfo() *GetHotelHomeBackImageAndModesRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelHomeBackImageAndModesRequest) SetUserInfo(v *GetHotelHomeBackImageAndModesRequestUserInfo) *GetHotelHomeBackImageAndModesRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelHomeBackImageAndModesRequestUserInfo struct {
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

func (s GetHotelHomeBackImageAndModesRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeKey(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetEncodeType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetIdType(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) SetOrganizationId(v string) *GetHotelHomeBackImageAndModesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
