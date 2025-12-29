// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *GetHotelOrderDetailRequestPayload) *GetHotelOrderDetailRequest
	GetPayload() *GetHotelOrderDetailRequestPayload
}

type GetHotelOrderDetailRequest struct {
	// This parameter is required.
	Payload *GetHotelOrderDetailRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s GetHotelOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequest) GetPayload() *GetHotelOrderDetailRequestPayload {
	return s.Payload
}

func (s *GetHotelOrderDetailRequest) SetPayload(v *GetHotelOrderDetailRequestPayload) *GetHotelOrderDetailRequest {
	s.Payload = v
	return s
}

func (s *GetHotelOrderDetailRequest) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelOrderDetailRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 20220714150702000168270112410630
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
}

func (s GetHotelOrderDetailRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailRequestPayload) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailRequestPayload) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GetHotelOrderDetailRequestPayload) SetOrderNo(v string) *GetHotelOrderDetailRequestPayload {
	s.OrderNo = &v
	return s
}

func (s *GetHotelOrderDetailRequestPayload) Validate() error {
	return dara.Validate(s)
}
