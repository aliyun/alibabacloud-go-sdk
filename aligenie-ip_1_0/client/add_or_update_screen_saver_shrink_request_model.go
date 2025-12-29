// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateScreenSaverShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *AddOrUpdateScreenSaverShrinkRequest
	GetHotelId() *string
	SetHotelScreenSaverShrink(v string) *AddOrUpdateScreenSaverShrinkRequest
	GetHotelScreenSaverShrink() *string
}

type AddOrUpdateScreenSaverShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	HotelScreenSaverShrink *string `json:"HotelScreenSaver,omitempty" xml:"HotelScreenSaver,omitempty"`
}

func (s AddOrUpdateScreenSaverShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateScreenSaverShrinkRequest) GetHotelScreenSaverShrink() *string {
	return s.HotelScreenSaverShrink
}

func (s *AddOrUpdateScreenSaverShrinkRequest) SetHotelId(v string) *AddOrUpdateScreenSaverShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateScreenSaverShrinkRequest) SetHotelScreenSaverShrink(v string) *AddOrUpdateScreenSaverShrinkRequest {
	s.HotelScreenSaverShrink = &v
	return s
}

func (s *AddOrUpdateScreenSaverShrinkRequest) Validate() error {
	return dara.Validate(s)
}
