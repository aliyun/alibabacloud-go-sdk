// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSampleUtterancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetHotelSampleUtterancesRequestUserInfo) *GetHotelSampleUtterancesRequest
	GetUserInfo() *GetHotelSampleUtterancesRequestUserInfo
}

type GetHotelSampleUtterancesRequest struct {
	UserInfo *GetHotelSampleUtterancesRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetHotelSampleUtterancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequest) GetUserInfo() *GetHotelSampleUtterancesRequestUserInfo {
	return s.UserInfo
}

func (s *GetHotelSampleUtterancesRequest) SetUserInfo(v *GetHotelSampleUtterancesRequestUserInfo) *GetHotelSampleUtterancesRequest {
	s.UserInfo = v
	return s
}

func (s *GetHotelSampleUtterancesRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelSampleUtterancesRequestUserInfo struct {
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

func (s GetHotelSampleUtterancesRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetHotelSampleUtterancesRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetHotelSampleUtterancesRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetHotelSampleUtterancesRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetHotelSampleUtterancesRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeKey(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetEncodeType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetIdType(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) SetOrganizationId(v string) *GetHotelSampleUtterancesRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetHotelSampleUtterancesRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
