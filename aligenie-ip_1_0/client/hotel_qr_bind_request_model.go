// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelQrBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *HotelQrBindRequest
	GetClientId() *string
	SetCode(v string) *HotelQrBindRequest
	GetCode() *string
	SetExtInfo(v string) *HotelQrBindRequest
	GetExtInfo() *string
	SetHotelId(v string) *HotelQrBindRequest
	GetHotelId() *string
	SetRoomNo(v string) *HotelQrBindRequest
	GetRoomNo() *string
}

type HotelQrBindRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// freuisghrtiesnvfkdsvbfuidslnvfs
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
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
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s HotelQrBindRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindRequest) GoString() string {
	return s.String()
}

func (s *HotelQrBindRequest) GetClientId() *string {
	return s.ClientId
}

func (s *HotelQrBindRequest) GetCode() *string {
	return s.Code
}

func (s *HotelQrBindRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *HotelQrBindRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelQrBindRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *HotelQrBindRequest) SetClientId(v string) *HotelQrBindRequest {
	s.ClientId = &v
	return s
}

func (s *HotelQrBindRequest) SetCode(v string) *HotelQrBindRequest {
	s.Code = &v
	return s
}

func (s *HotelQrBindRequest) SetExtInfo(v string) *HotelQrBindRequest {
	s.ExtInfo = &v
	return s
}

func (s *HotelQrBindRequest) SetHotelId(v string) *HotelQrBindRequest {
	s.HotelId = &v
	return s
}

func (s *HotelQrBindRequest) SetRoomNo(v string) *HotelQrBindRequest {
	s.RoomNo = &v
	return s
}

func (s *HotelQrBindRequest) Validate() error {
	return dara.Validate(s)
}
