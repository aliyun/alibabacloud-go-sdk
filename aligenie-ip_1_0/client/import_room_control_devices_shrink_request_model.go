// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomControlDevicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesShrinkRequest
	GetEnableInfraredDeviceImport() *string
	SetEnableMeshDeviceImport(v string) *ImportRoomControlDevicesShrinkRequest
	GetEnableMeshDeviceImport() *string
	SetHotelId(v string) *ImportRoomControlDevicesShrinkRequest
	GetHotelId() *string
	SetLocationDevicesShrink(v string) *ImportRoomControlDevicesShrinkRequest
	GetLocationDevicesShrink() *string
	SetRoomNo(v string) *ImportRoomControlDevicesShrinkRequest
	GetRoomNo() *string
}

type ImportRoomControlDevicesShrinkRequest struct {
	EnableInfraredDeviceImport *string `json:"EnableInfraredDeviceImport,omitempty" xml:"EnableInfraredDeviceImport,omitempty"`
	EnableMeshDeviceImport     *string `json:"EnableMeshDeviceImport,omitempty" xml:"EnableMeshDeviceImport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdgrefds
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	LocationDevicesShrink *string `json:"LocationDevices,omitempty" xml:"LocationDevices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ImportRoomControlDevicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesShrinkRequest) GetEnableInfraredDeviceImport() *string {
	return s.EnableInfraredDeviceImport
}

func (s *ImportRoomControlDevicesShrinkRequest) GetEnableMeshDeviceImport() *string {
	return s.EnableMeshDeviceImport
}

func (s *ImportRoomControlDevicesShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportRoomControlDevicesShrinkRequest) GetLocationDevicesShrink() *string {
	return s.LocationDevicesShrink
}

func (s *ImportRoomControlDevicesShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ImportRoomControlDevicesShrinkRequest) SetEnableInfraredDeviceImport(v string) *ImportRoomControlDevicesShrinkRequest {
	s.EnableInfraredDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetEnableMeshDeviceImport(v string) *ImportRoomControlDevicesShrinkRequest {
	s.EnableMeshDeviceImport = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetHotelId(v string) *ImportRoomControlDevicesShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetLocationDevicesShrink(v string) *ImportRoomControlDevicesShrinkRequest {
	s.LocationDevicesShrink = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) SetRoomNo(v string) *ImportRoomControlDevicesShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomControlDevicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
