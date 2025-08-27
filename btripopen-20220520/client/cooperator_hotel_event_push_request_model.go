// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelEventPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderStatus(v int32) *CooperatorHotelEventPushRequest
	GetChangeOrderStatus() *int32
	SetChangeOrderStatusDesc(v string) *CooperatorHotelEventPushRequest
	GetChangeOrderStatusDesc() *string
	SetCooperatorOrderId(v string) *CooperatorHotelEventPushRequest
	GetCooperatorOrderId() *string
	SetEvent(v string) *CooperatorHotelEventPushRequest
	GetEvent() *string
	SetEventDesc(v string) *CooperatorHotelEventPushRequest
	GetEventDesc() *string
	SetEventTime(v string) *CooperatorHotelEventPushRequest
	GetEventTime() *string
	SetOrderId(v string) *CooperatorHotelEventPushRequest
	GetOrderId() *string
}

type CooperatorHotelEventPushRequest struct {
	// example:
	//
	// 11
	ChangeOrderStatus     *int32  `json:"change_order_status,omitempty" xml:"change_order_status,omitempty"`
	ChangeOrderStatusDesc *string `json:"change_order_status_desc,omitempty" xml:"change_order_status_desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// coop_123456
	CooperatorOrderId *string `json:"cooperator_order_id,omitempty" xml:"cooperator_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USER_LEAVE
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventDesc *string `json:"event_desc,omitempty" xml:"event_desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 204-09-01 10:55:20
	EventTime *string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s CooperatorHotelEventPushRequest) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelEventPushRequest) GoString() string {
	return s.String()
}

func (s *CooperatorHotelEventPushRequest) GetChangeOrderStatus() *int32 {
	return s.ChangeOrderStatus
}

func (s *CooperatorHotelEventPushRequest) GetChangeOrderStatusDesc() *string {
	return s.ChangeOrderStatusDesc
}

func (s *CooperatorHotelEventPushRequest) GetCooperatorOrderId() *string {
	return s.CooperatorOrderId
}

func (s *CooperatorHotelEventPushRequest) GetEvent() *string {
	return s.Event
}

func (s *CooperatorHotelEventPushRequest) GetEventDesc() *string {
	return s.EventDesc
}

func (s *CooperatorHotelEventPushRequest) GetEventTime() *string {
	return s.EventTime
}

func (s *CooperatorHotelEventPushRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *CooperatorHotelEventPushRequest) SetChangeOrderStatus(v int32) *CooperatorHotelEventPushRequest {
	s.ChangeOrderStatus = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetChangeOrderStatusDesc(v string) *CooperatorHotelEventPushRequest {
	s.ChangeOrderStatusDesc = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetCooperatorOrderId(v string) *CooperatorHotelEventPushRequest {
	s.CooperatorOrderId = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetEvent(v string) *CooperatorHotelEventPushRequest {
	s.Event = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetEventDesc(v string) *CooperatorHotelEventPushRequest {
	s.EventDesc = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetEventTime(v string) *CooperatorHotelEventPushRequest {
	s.EventTime = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) SetOrderId(v string) *CooperatorHotelEventPushRequest {
	s.OrderId = &v
	return s
}

func (s *CooperatorHotelEventPushRequest) Validate() error {
	return dara.Validate(s)
}
