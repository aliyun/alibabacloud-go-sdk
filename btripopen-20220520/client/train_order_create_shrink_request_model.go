// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCreateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptNoSeat(v string) *TrainOrderCreateShrinkRequest
	GetAcceptNoSeat() *string
	SetBookTrainInfosShrink(v string) *TrainOrderCreateShrinkRequest
	GetBookTrainInfosShrink() *string
	SetBtripUserId(v string) *TrainOrderCreateShrinkRequest
	GetBtripUserId() *string
	SetBtripUserName(v string) *TrainOrderCreateShrinkRequest
	GetBtripUserName() *string
	SetBusinessInfoShrink(v string) *TrainOrderCreateShrinkRequest
	GetBusinessInfoShrink() *string
	SetContactInfoShrink(v string) *TrainOrderCreateShrinkRequest
	GetContactInfoShrink() *string
	SetForceMatch(v string) *TrainOrderCreateShrinkRequest
	GetForceMatch() *string
	SetIsPayNow(v bool) *TrainOrderCreateShrinkRequest
	GetIsPayNow() *bool
	SetOutOrderId(v string) *TrainOrderCreateShrinkRequest
	GetOutOrderId() *string
	SetPassengerOpenInfoSShrink(v string) *TrainOrderCreateShrinkRequest
	GetPassengerOpenInfoSShrink() *string
}

type TrainOrderCreateShrinkRequest struct {
	// example:
	//
	// 0
	AcceptNoSeat *string `json:"accept_no_seat,omitempty" xml:"accept_no_seat,omitempty"`
	// This parameter is required.
	BookTrainInfosShrink *string `json:"book_train_infos,omitempty" xml:"book_train_infos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12344321
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	BtripUserName      *string `json:"btrip_user_name,omitempty" xml:"btrip_user_name,omitempty"`
	BusinessInfoShrink *string `json:"business_info,omitempty" xml:"business_info,omitempty"`
	// This parameter is required.
	ContactInfoShrink *string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// example:
	//
	// 0
	ForceMatch *string `json:"force_match,omitempty" xml:"force_match,omitempty"`
	// example:
	//
	// false
	IsPayNow *bool `json:"is_pay_now,omitempty" xml:"is_pay_now,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	PassengerOpenInfoSShrink *string `json:"passenger_open_info_s,omitempty" xml:"passenger_open_info_s,omitempty"`
}

func (s TrainOrderCreateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateShrinkRequest) GetAcceptNoSeat() *string {
	return s.AcceptNoSeat
}

func (s *TrainOrderCreateShrinkRequest) GetBookTrainInfosShrink() *string {
	return s.BookTrainInfosShrink
}

func (s *TrainOrderCreateShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *TrainOrderCreateShrinkRequest) GetBtripUserName() *string {
	return s.BtripUserName
}

func (s *TrainOrderCreateShrinkRequest) GetBusinessInfoShrink() *string {
	return s.BusinessInfoShrink
}

func (s *TrainOrderCreateShrinkRequest) GetContactInfoShrink() *string {
	return s.ContactInfoShrink
}

func (s *TrainOrderCreateShrinkRequest) GetForceMatch() *string {
	return s.ForceMatch
}

func (s *TrainOrderCreateShrinkRequest) GetIsPayNow() *bool {
	return s.IsPayNow
}

func (s *TrainOrderCreateShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderCreateShrinkRequest) GetPassengerOpenInfoSShrink() *string {
	return s.PassengerOpenInfoSShrink
}

func (s *TrainOrderCreateShrinkRequest) SetAcceptNoSeat(v string) *TrainOrderCreateShrinkRequest {
	s.AcceptNoSeat = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetBookTrainInfosShrink(v string) *TrainOrderCreateShrinkRequest {
	s.BookTrainInfosShrink = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetBtripUserId(v string) *TrainOrderCreateShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetBtripUserName(v string) *TrainOrderCreateShrinkRequest {
	s.BtripUserName = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetBusinessInfoShrink(v string) *TrainOrderCreateShrinkRequest {
	s.BusinessInfoShrink = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetContactInfoShrink(v string) *TrainOrderCreateShrinkRequest {
	s.ContactInfoShrink = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetForceMatch(v string) *TrainOrderCreateShrinkRequest {
	s.ForceMatch = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetIsPayNow(v bool) *TrainOrderCreateShrinkRequest {
	s.IsPayNow = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetOutOrderId(v string) *TrainOrderCreateShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) SetPassengerOpenInfoSShrink(v string) *TrainOrderCreateShrinkRequest {
	s.PassengerOpenInfoSShrink = &v
	return s
}

func (s *TrainOrderCreateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
