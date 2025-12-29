// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *DeleteHotelSettingRequest
	GetHotelId() *string
	SetSettingType(v string) *DeleteHotelSettingRequest
	GetSettingType() *string
}

type DeleteHotelSettingRequest struct {
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// SCREENSAVER
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
}

func (s DeleteHotelSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteHotelSettingRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *DeleteHotelSettingRequest) SetHotelId(v string) *DeleteHotelSettingRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelSettingRequest) SetSettingType(v string) *DeleteHotelSettingRequest {
	s.SettingType = &v
	return s
}

func (s *DeleteHotelSettingRequest) Validate() error {
	return dara.Validate(s)
}
