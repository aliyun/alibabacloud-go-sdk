// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSceneBookItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *DeleteHotelSceneBookItemRequest
	GetHotelId() *string
	SetId(v int64) *DeleteHotelSceneBookItemRequest
	GetId() *int64
	SetName(v string) *DeleteHotelSceneBookItemRequest
	GetName() *string
}

type DeleteHotelSceneBookItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 11823
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteHotelSceneBookItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteHotelSceneBookItemRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteHotelSceneBookItemRequest) GetName() *string {
	return s.Name
}

func (s *DeleteHotelSceneBookItemRequest) SetHotelId(v string) *DeleteHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteHotelSceneBookItemRequest) SetId(v int64) *DeleteHotelSceneBookItemRequest {
	s.Id = &v
	return s
}

func (s *DeleteHotelSceneBookItemRequest) SetName(v string) *DeleteHotelSceneBookItemRequest {
	s.Name = &v
	return s
}

func (s *DeleteHotelSceneBookItemRequest) Validate() error {
	return dara.Validate(s)
}
