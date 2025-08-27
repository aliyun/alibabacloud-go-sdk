// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOrderDetailV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightModifyOrderDetailV2Request
	GetIsvName() *string
	SetModifyApplyId(v string) *FlightModifyOrderDetailV2Request
	GetModifyApplyId() *string
	SetNeedQueryServiceFee(v bool) *FlightModifyOrderDetailV2Request
	GetNeedQueryServiceFee() *bool
	SetOrderId(v string) *FlightModifyOrderDetailV2Request
	GetOrderId() *string
	SetOutModifyApplyId(v string) *FlightModifyOrderDetailV2Request
	GetOutModifyApplyId() *string
	SetOutOrderId(v string) *FlightModifyOrderDetailV2Request
	GetOutOrderId() *string
}

type FlightModifyOrderDetailV2Request struct {
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1019195786853020
	ModifyApplyId *string `json:"modify_apply_id,omitempty" xml:"modify_apply_id,omitempty"`
	// example:
	//
	// false
	NeedQueryServiceFee *bool `json:"need_query_service_fee,omitempty" xml:"need_query_service_fee,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId          *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutModifyApplyId *string `json:"out_modify_apply_id,omitempty" xml:"out_modify_apply_id,omitempty"`
	OutOrderId       *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s FlightModifyOrderDetailV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyOrderDetailV2Request) GetModifyApplyId() *string {
	return s.ModifyApplyId
}

func (s *FlightModifyOrderDetailV2Request) GetNeedQueryServiceFee() *bool {
	return s.NeedQueryServiceFee
}

func (s *FlightModifyOrderDetailV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightModifyOrderDetailV2Request) GetOutModifyApplyId() *string {
	return s.OutModifyApplyId
}

func (s *FlightModifyOrderDetailV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyOrderDetailV2Request) SetIsvName(v string) *FlightModifyOrderDetailV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) SetModifyApplyId(v string) *FlightModifyOrderDetailV2Request {
	s.ModifyApplyId = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) SetNeedQueryServiceFee(v bool) *FlightModifyOrderDetailV2Request {
	s.NeedQueryServiceFee = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) SetOrderId(v string) *FlightModifyOrderDetailV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) SetOutModifyApplyId(v string) *FlightModifyOrderDetailV2Request {
	s.OutModifyApplyId = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) SetOutOrderId(v string) *FlightModifyOrderDetailV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyOrderDetailV2Request) Validate() error {
	return dara.Validate(s)
}
