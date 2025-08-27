// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *TicketChangingApplyShrinkRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingApplyShrinkRequest
	GetDisSubOrderId() *string
	SetIsVoluntary(v int32) *TicketChangingApplyShrinkRequest
	GetIsVoluntary() *int32
	SetModifyFlightInfoListShrink(v string) *TicketChangingApplyShrinkRequest
	GetModifyFlightInfoListShrink() *string
	SetOtaItemId(v string) *TicketChangingApplyShrinkRequest
	GetOtaItemId() *string
	SetReason(v string) *TicketChangingApplyShrinkRequest
	GetReason() *string
	SetSessionId(v string) *TicketChangingApplyShrinkRequest
	GetSessionId() *string
	SetWhetherRetry(v bool) *TicketChangingApplyShrinkRequest
	GetWhetherRetry() *bool
}

type TicketChangingApplyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis1234
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mid1243
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	IsVoluntary   *int32  `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	ModifyFlightInfoListShrink *string `json:"modify_flight_info_list,omitempty" xml:"modify_flight_info_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1s8837sh991hsj92h
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// us88s2bsbin22hjusd8i
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	WhetherRetry *bool `json:"whether_retry,omitempty" xml:"whether_retry,omitempty"`
}

func (s TicketChangingApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingApplyShrinkRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingApplyShrinkRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *TicketChangingApplyShrinkRequest) GetModifyFlightInfoListShrink() *string {
	return s.ModifyFlightInfoListShrink
}

func (s *TicketChangingApplyShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *TicketChangingApplyShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *TicketChangingApplyShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *TicketChangingApplyShrinkRequest) GetWhetherRetry() *bool {
	return s.WhetherRetry
}

func (s *TicketChangingApplyShrinkRequest) SetDisOrderId(v string) *TicketChangingApplyShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetDisSubOrderId(v string) *TicketChangingApplyShrinkRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetIsVoluntary(v int32) *TicketChangingApplyShrinkRequest {
	s.IsVoluntary = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetModifyFlightInfoListShrink(v string) *TicketChangingApplyShrinkRequest {
	s.ModifyFlightInfoListShrink = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetOtaItemId(v string) *TicketChangingApplyShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetReason(v string) *TicketChangingApplyShrinkRequest {
	s.Reason = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetSessionId(v string) *TicketChangingApplyShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) SetWhetherRetry(v bool) *TicketChangingApplyShrinkRequest {
	s.WhetherRetry = &v
	return s
}

func (s *TicketChangingApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
