// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyApplyV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheKey(v string) *FlightModifyApplyV2ShrinkRequest
	GetCacheKey() *string
	SetContactPhone(v string) *FlightModifyApplyV2ShrinkRequest
	GetContactPhone() *string
	SetIsvName(v string) *FlightModifyApplyV2ShrinkRequest
	GetIsvName() *string
	SetItemId(v string) *FlightModifyApplyV2ShrinkRequest
	GetItemId() *string
	SetOrderId(v int64) *FlightModifyApplyV2ShrinkRequest
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyApplyV2ShrinkRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightModifyApplyV2ShrinkRequest
	GetOutSubOrderId() *string
	SetPassengerSegmentRelationsShrink(v string) *FlightModifyApplyV2ShrinkRequest
	GetPassengerSegmentRelationsShrink() *string
	SetReason(v string) *FlightModifyApplyV2ShrinkRequest
	GetReason() *string
	SetSessionId(v string) *FlightModifyApplyV2ShrinkRequest
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyApplyV2ShrinkRequest
	GetVoluntary() *bool
}

type FlightModifyApplyV2ShrinkRequest struct {
	// example:
	//
	// 72e961f8-930b-43c1-a4ca-18a6f28349c6distributionModifyCacheInfo
	CacheKey *string `json:"cache_key,omitempty" xml:"cache_key,omitempty"`
	// example:
	//
	// 17816963077
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// fa2fb23a859a4e78b5ddb87a6a23094b_0
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 1017002195370467138
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195786853020
	OutSubOrderId                   *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	PassengerSegmentRelationsShrink *string `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty"`
	Reason                          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// sessionId
	//
	// example:
	//
	// a2ffebfe733742aab5c491d960ba3d59
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightModifyApplyV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2ShrinkRequest) GetCacheKey() *string {
	return s.CacheKey
}

func (s *FlightModifyApplyV2ShrinkRequest) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightModifyApplyV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyApplyV2ShrinkRequest) GetItemId() *string {
	return s.ItemId
}

func (s *FlightModifyApplyV2ShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyApplyV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyApplyV2ShrinkRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyApplyV2ShrinkRequest) GetPassengerSegmentRelationsShrink() *string {
	return s.PassengerSegmentRelationsShrink
}

func (s *FlightModifyApplyV2ShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *FlightModifyApplyV2ShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyApplyV2ShrinkRequest) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyApplyV2ShrinkRequest) SetCacheKey(v string) *FlightModifyApplyV2ShrinkRequest {
	s.CacheKey = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetContactPhone(v string) *FlightModifyApplyV2ShrinkRequest {
	s.ContactPhone = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetIsvName(v string) *FlightModifyApplyV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetItemId(v string) *FlightModifyApplyV2ShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetOrderId(v int64) *FlightModifyApplyV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetOutOrderId(v string) *FlightModifyApplyV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetOutSubOrderId(v string) *FlightModifyApplyV2ShrinkRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetPassengerSegmentRelationsShrink(v string) *FlightModifyApplyV2ShrinkRequest {
	s.PassengerSegmentRelationsShrink = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetReason(v string) *FlightModifyApplyV2ShrinkRequest {
	s.Reason = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetSessionId(v string) *FlightModifyApplyV2ShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) SetVoluntary(v bool) *FlightModifyApplyV2ShrinkRequest {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyApplyV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
