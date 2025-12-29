// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageGetHotelRoomDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *PageGetHotelRoomDevicesRequest
	GetHotelId() *string
	SetPageNumber(v int32) *PageGetHotelRoomDevicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *PageGetHotelRoomDevicesRequest
	GetPageSize() *int32
}

type PageGetHotelRoomDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageGetHotelRoomDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesRequest) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PageGetHotelRoomDevicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *PageGetHotelRoomDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageGetHotelRoomDevicesRequest) SetHotelId(v string) *PageGetHotelRoomDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *PageGetHotelRoomDevicesRequest) SetPageNumber(v int32) *PageGetHotelRoomDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *PageGetHotelRoomDevicesRequest) SetPageSize(v int32) *PageGetHotelRoomDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *PageGetHotelRoomDevicesRequest) Validate() error {
	return dara.Validate(s)
}
