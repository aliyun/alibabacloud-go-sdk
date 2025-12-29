// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetHotelSettingRequest
	GetHotelId() *string
	SetSettingType(v string) *GetHotelSettingRequest
	GetSettingType() *string
}

type GetHotelSettingRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
}

func (s GetHotelSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSettingRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelSettingRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *GetHotelSettingRequest) SetHotelId(v string) *GetHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelSettingRequest) SetSettingType(v string) *GetHotelSettingRequest {
	s.SettingType = &v
	return s
}

func (s *GetHotelSettingRequest) Validate() error {
	return dara.Validate(s)
}
