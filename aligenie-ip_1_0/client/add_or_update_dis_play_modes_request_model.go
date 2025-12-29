// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDisPlayModesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelDeviceModeList(v []*string) *AddOrUpdateDisPlayModesRequest
	GetHotelDeviceModeList() []*string
	SetHotelId(v string) *AddOrUpdateDisPlayModesRequest
	GetHotelId() *string
}

type AddOrUpdateDisPlayModesRequest struct {
	// This parameter is required.
	HotelDeviceModeList []*string `json:"HotelDeviceModeList,omitempty" xml:"HotelDeviceModeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s AddOrUpdateDisPlayModesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDisPlayModesRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesRequest) GetHotelDeviceModeList() []*string {
	return s.HotelDeviceModeList
}

func (s *AddOrUpdateDisPlayModesRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddOrUpdateDisPlayModesRequest) SetHotelDeviceModeList(v []*string) *AddOrUpdateDisPlayModesRequest {
	s.HotelDeviceModeList = v
	return s
}

func (s *AddOrUpdateDisPlayModesRequest) SetHotelId(v string) *AddOrUpdateDisPlayModesRequest {
	s.HotelId = &v
	return s
}

func (s *AddOrUpdateDisPlayModesRequest) Validate() error {
	return dara.Validate(s)
}
