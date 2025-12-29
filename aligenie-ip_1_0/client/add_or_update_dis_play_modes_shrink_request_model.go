// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDisPlayModesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelDeviceModeListShrink(v string) *AddOrUpdateDisPlayModesShrinkRequest
	GetHotelDeviceModeListShrink() *string
	SetHotelId(v string) *AddOrUpdateDisPlayModesShrinkRequest
	GetHotelId() *string
}

type AddOrUpdateDisPlayModesShrinkRequest struct {
	// This parameter is required.
	HotelDeviceModeListShrink *string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s AddOrUpdateDisPlayModesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDisPlayModesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) GetHotelDeviceModeListShrink() *string {
	return s.HotelDeviceModeListShrink
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) SetHotelDeviceModeListShrink(v string) *AddOrUpdateDisPlayModesShrinkRequest {
	s.HotelDeviceModeListShrink = &v
	return s
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) SetHotelId(v string) *AddOrUpdateDisPlayModesShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateDisPlayModesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
