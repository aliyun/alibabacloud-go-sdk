// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaItemDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightOtaItemDetailRequest
	GetIsvName() *string
	SetOrderId(v string) *FlightOtaItemDetailRequest
	GetOrderId() *string
	SetOtaItemId(v string) *FlightOtaItemDetailRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *FlightOtaItemDetailRequest
	GetOutOrderId() *string
}

type FlightOtaItemDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cheshi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68cdc6b37c87484c98b479b49306ffbb_0
	OtaItemId  *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s FlightOtaItemDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailRequest) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightOtaItemDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightOtaItemDetailRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightOtaItemDetailRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightOtaItemDetailRequest) SetIsvName(v string) *FlightOtaItemDetailRequest {
	s.IsvName = &v
	return s
}

func (s *FlightOtaItemDetailRequest) SetOrderId(v string) *FlightOtaItemDetailRequest {
	s.OrderId = &v
	return s
}

func (s *FlightOtaItemDetailRequest) SetOtaItemId(v string) *FlightOtaItemDetailRequest {
	s.OtaItemId = &v
	return s
}

func (s *FlightOtaItemDetailRequest) SetOutOrderId(v string) *FlightOtaItemDetailRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightOtaItemDetailRequest) Validate() error {
	return dara.Validate(s)
}
