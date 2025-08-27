// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyCancelV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightModifyCancelV2Request
	GetIsvName() *string
	SetOrderId(v int64) *FlightModifyCancelV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyCancelV2Request
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightModifyCancelV2Request
	GetOutSubOrderId() *string
	SetSubOrderId(v int64) *FlightModifyCancelV2Request
	GetSubOrderId() *int64
}

type FlightModifyCancelV2Request struct {
	IsvName       *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	OrderId       *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId    *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	SubOrderId    *int64  `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s FlightModifyCancelV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyCancelV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyCancelV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyCancelV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyCancelV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyCancelV2Request) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyCancelV2Request) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *FlightModifyCancelV2Request) SetIsvName(v string) *FlightModifyCancelV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyCancelV2Request) SetOrderId(v int64) *FlightModifyCancelV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyCancelV2Request) SetOutOrderId(v string) *FlightModifyCancelV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyCancelV2Request) SetOutSubOrderId(v string) *FlightModifyCancelV2Request {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyCancelV2Request) SetSubOrderId(v int64) *FlightModifyCancelV2Request {
	s.SubOrderId = &v
	return s
}

func (s *FlightModifyCancelV2Request) Validate() error {
	return dara.Validate(s)
}
