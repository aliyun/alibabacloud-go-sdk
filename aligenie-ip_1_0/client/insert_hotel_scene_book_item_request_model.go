// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertHotelSceneBookItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddHotelSceneItemReq(v *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) *InsertHotelSceneBookItemRequest
	GetAddHotelSceneItemReq() *InsertHotelSceneBookItemRequestAddHotelSceneItemReq
	SetHotelId(v string) *InsertHotelSceneBookItemRequest
	GetHotelId() *string
}

type InsertHotelSceneBookItemRequest struct {
	// addHotelSceneItemReq
	//
	// This parameter is required.
	AddHotelSceneItemReq *InsertHotelSceneBookItemRequestAddHotelSceneItemReq `json:"AddHotelSceneItemReq,omitempty" xml:"AddHotelSceneItemReq,omitempty" type:"Struct"`
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s InsertHotelSceneBookItemRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemRequest) GetAddHotelSceneItemReq() *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	return s.AddHotelSceneItemReq
}

func (s *InsertHotelSceneBookItemRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *InsertHotelSceneBookItemRequest) SetAddHotelSceneItemReq(v *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) *InsertHotelSceneBookItemRequest {
	s.AddHotelSceneItemReq = v
	return s
}

func (s *InsertHotelSceneBookItemRequest) SetHotelId(v string) *InsertHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

func (s *InsertHotelSceneBookItemRequest) Validate() error {
	if s.AddHotelSceneItemReq != nil {
		if err := s.AddHotelSceneItemReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertHotelSceneBookItemRequestAddHotelSceneItemReq struct {
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 青椒肉丝
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1250
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InsertHotelSceneBookItemRequestAddHotelSceneItemReq) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GetIcon() *string {
	return s.Icon
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GetName() *string {
	return s.Name
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GetPrice() *int64 {
	return s.Price
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) GetType() *string {
	return s.Type
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetIcon(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Icon = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetName(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Name = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetPrice(v int64) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Price = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) SetType(v string) *InsertHotelSceneBookItemRequestAddHotelSceneItemReq {
	s.Type = &v
	return s
}

func (s *InsertHotelSceneBookItemRequestAddHotelSceneItemReq) Validate() error {
	return dara.Validate(s)
}
