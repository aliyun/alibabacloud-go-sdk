// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyApplyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCacheKey(v string) *FlightModifyApplyV2Request
	GetCacheKey() *string
	SetContactPhone(v string) *FlightModifyApplyV2Request
	GetContactPhone() *string
	SetIsvName(v string) *FlightModifyApplyV2Request
	GetIsvName() *string
	SetItemId(v string) *FlightModifyApplyV2Request
	GetItemId() *string
	SetOrderId(v int64) *FlightModifyApplyV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyApplyV2Request
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightModifyApplyV2Request
	GetOutSubOrderId() *string
	SetPassengerSegmentRelations(v []*FlightModifyApplyV2RequestPassengerSegmentRelations) *FlightModifyApplyV2Request
	GetPassengerSegmentRelations() []*FlightModifyApplyV2RequestPassengerSegmentRelations
	SetReason(v string) *FlightModifyApplyV2Request
	GetReason() *string
	SetSessionId(v string) *FlightModifyApplyV2Request
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyApplyV2Request
	GetVoluntary() *bool
}

type FlightModifyApplyV2Request struct {
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
	OutSubOrderId             *string                                                `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	PassengerSegmentRelations []*FlightModifyApplyV2RequestPassengerSegmentRelations `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty" type:"Repeated"`
	Reason                    *string                                                `json:"reason,omitempty" xml:"reason,omitempty"`
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

func (s FlightModifyApplyV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2Request) GetCacheKey() *string {
	return s.CacheKey
}

func (s *FlightModifyApplyV2Request) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightModifyApplyV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyApplyV2Request) GetItemId() *string {
	return s.ItemId
}

func (s *FlightModifyApplyV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyApplyV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyApplyV2Request) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightModifyApplyV2Request) GetPassengerSegmentRelations() []*FlightModifyApplyV2RequestPassengerSegmentRelations {
	return s.PassengerSegmentRelations
}

func (s *FlightModifyApplyV2Request) GetReason() *string {
	return s.Reason
}

func (s *FlightModifyApplyV2Request) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyApplyV2Request) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyApplyV2Request) SetCacheKey(v string) *FlightModifyApplyV2Request {
	s.CacheKey = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetContactPhone(v string) *FlightModifyApplyV2Request {
	s.ContactPhone = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetIsvName(v string) *FlightModifyApplyV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetItemId(v string) *FlightModifyApplyV2Request {
	s.ItemId = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetOrderId(v int64) *FlightModifyApplyV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetOutOrderId(v string) *FlightModifyApplyV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetOutSubOrderId(v string) *FlightModifyApplyV2Request {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetPassengerSegmentRelations(v []*FlightModifyApplyV2RequestPassengerSegmentRelations) *FlightModifyApplyV2Request {
	s.PassengerSegmentRelations = v
	return s
}

func (s *FlightModifyApplyV2Request) SetReason(v string) *FlightModifyApplyV2Request {
	s.Reason = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetSessionId(v string) *FlightModifyApplyV2Request {
	s.SessionId = &v
	return s
}

func (s *FlightModifyApplyV2Request) SetVoluntary(v bool) *FlightModifyApplyV2Request {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyApplyV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightModifyApplyV2RequestPassengerSegmentRelations struct {
	// example:
	//
	// 3243028
	PassengerId   *string   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s FlightModifyApplyV2RequestPassengerSegmentRelations) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2RequestPassengerSegmentRelations) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2RequestPassengerSegmentRelations) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightModifyApplyV2RequestPassengerSegmentRelations) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *FlightModifyApplyV2RequestPassengerSegmentRelations) SetPassengerId(v string) *FlightModifyApplyV2RequestPassengerSegmentRelations {
	s.PassengerId = &v
	return s
}

func (s *FlightModifyApplyV2RequestPassengerSegmentRelations) SetSegmentIdList(v []*string) *FlightModifyApplyV2RequestPassengerSegmentRelations {
	s.SegmentIdList = v
	return s
}

func (s *FlightModifyApplyV2RequestPassengerSegmentRelations) Validate() error {
	return dara.Validate(s)
}
