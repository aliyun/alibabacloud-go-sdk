// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneBookItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateHotelSceneBookItemRequest
	GetHotelId() *string
	SetUpdateHotelSceneBookReq(v *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) *UpdateHotelSceneBookItemRequest
	GetUpdateHotelSceneBookReq() *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq
}

type UpdateHotelSceneBookItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneBookReq
	//
	// This parameter is required.
	UpdateHotelSceneBookReq *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq `json:"UpdateHotelSceneBookReq,omitempty" xml:"UpdateHotelSceneBookReq,omitempty" type:"Struct"`
}

func (s UpdateHotelSceneBookItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelSceneBookItemRequest) GetUpdateHotelSceneBookReq() *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	return s.UpdateHotelSceneBookReq
}

func (s *UpdateHotelSceneBookItemRequest) SetHotelId(v string) *UpdateHotelSceneBookItemRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequest) SetUpdateHotelSceneBookReq(v *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) *UpdateHotelSceneBookItemRequest {
	s.UpdateHotelSceneBookReq = v
	return s
}

func (s *UpdateHotelSceneBookItemRequest) Validate() error {
	if s.UpdateHotelSceneBookReq != nil {
		if err := s.UpdateHotelSceneBookReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq struct {
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
}

func (s UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GetIcon() *string {
	return s.Icon
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GetId() *int64 {
	return s.Id
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GetName() *string {
	return s.Name
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) GetPrice() *int64 {
	return s.Price
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetIcon(v string) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Icon = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetId(v int64) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Id = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetName(v string) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Name = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) SetPrice(v int64) *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq {
	s.Price = &v
	return s
}

func (s *UpdateHotelSceneBookItemRequestUpdateHotelSceneBookReq) Validate() error {
	return dara.Validate(s)
}
