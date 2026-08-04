// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThirdImmediateMsgPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ThirdImmediateMsgPushRequest
	GetBizType() *string
	SetChangeDetail(v string) *ThirdImmediateMsgPushRequest
	GetChangeDetail() *string
	SetOrderId(v string) *ThirdImmediateMsgPushRequest
	GetOrderId() *string
	SetPsgIds(v string) *ThirdImmediateMsgPushRequest
	GetPsgIds() *string
	SetSource(v string) *ThirdImmediateMsgPushRequest
	GetSource() *string
	SetTrafficChangeType(v string) *ThirdImmediateMsgPushRequest
	GetTrafficChangeType() *string
	SetTrafficChangeTypeDesc(v string) *ThirdImmediateMsgPushRequest
	GetTrafficChangeTypeDesc() *string
	SetTrafficJourneyIds(v string) *ThirdImmediateMsgPushRequest
	GetTrafficJourneyIds() *string
	SetTrafficSubOrderIds(v string) *ThirdImmediateMsgPushRequest
	GetTrafficSubOrderIds() *string
	SetUserId(v string) *ThirdImmediateMsgPushRequest
	GetUserId() *string
}

type ThirdImmediateMsgPushRequest struct {
	// Business type (FLIGHT: flight, TRAIN: train)
	//
	// example:
	//
	// FLIGHT
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Change details (JSON-formatted change data)
	//
	// example:
	//
	// {}
	ChangeDetail *string `json:"ChangeDetail,omitempty" xml:"ChangeDetail,omitempty"`
	// Order ID
	//
	// example:
	//
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Passenger IDs to change
	//
	// example:
	//
	// psgIds
	PsgIds *string `json:"PsgIds,omitempty" xml:"PsgIds,omitempty"`
	// Request source
	//
	// example:
	//
	// source
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// change type (TRAIN_ISSUED, FLIGHT_CHANGED)
	//
	// example:
	//
	// TRAIN_ISSUED
	TrafficChangeType *string `json:"TrafficChangeType,omitempty" xml:"TrafficChangeType,omitempty"`
	// change type description (ticket issued, rebooked)
	//
	// example:
	//
	// 已出票
	TrafficChangeTypeDesc *string `json:"TrafficChangeTypeDesc,omitempty" xml:"TrafficChangeTypeDesc,omitempty"`
	// Journey IDs to change
	//
	// example:
	//
	// trafficJourneyIds
	TrafficJourneyIds *string `json:"TrafficJourneyIds,omitempty" xml:"TrafficJourneyIds,omitempty"`
	// sub-order ID of the changed train request
	//
	// example:
	//
	// trafficSubOrderIds
	TrafficSubOrderIds *string `json:"TrafficSubOrderIds,omitempty" xml:"TrafficSubOrderIds,omitempty"`
	// user ID
	//
	// example:
	//
	// userId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ThirdImmediateMsgPushRequest) String() string {
	return dara.Prettify(s)
}

func (s ThirdImmediateMsgPushRequest) GoString() string {
	return s.String()
}

func (s *ThirdImmediateMsgPushRequest) GetBizType() *string {
	return s.BizType
}

func (s *ThirdImmediateMsgPushRequest) GetChangeDetail() *string {
	return s.ChangeDetail
}

func (s *ThirdImmediateMsgPushRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *ThirdImmediateMsgPushRequest) GetPsgIds() *string {
	return s.PsgIds
}

func (s *ThirdImmediateMsgPushRequest) GetSource() *string {
	return s.Source
}

func (s *ThirdImmediateMsgPushRequest) GetTrafficChangeType() *string {
	return s.TrafficChangeType
}

func (s *ThirdImmediateMsgPushRequest) GetTrafficChangeTypeDesc() *string {
	return s.TrafficChangeTypeDesc
}

func (s *ThirdImmediateMsgPushRequest) GetTrafficJourneyIds() *string {
	return s.TrafficJourneyIds
}

func (s *ThirdImmediateMsgPushRequest) GetTrafficSubOrderIds() *string {
	return s.TrafficSubOrderIds
}

func (s *ThirdImmediateMsgPushRequest) GetUserId() *string {
	return s.UserId
}

func (s *ThirdImmediateMsgPushRequest) SetBizType(v string) *ThirdImmediateMsgPushRequest {
	s.BizType = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetChangeDetail(v string) *ThirdImmediateMsgPushRequest {
	s.ChangeDetail = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetOrderId(v string) *ThirdImmediateMsgPushRequest {
	s.OrderId = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetPsgIds(v string) *ThirdImmediateMsgPushRequest {
	s.PsgIds = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetSource(v string) *ThirdImmediateMsgPushRequest {
	s.Source = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetTrafficChangeType(v string) *ThirdImmediateMsgPushRequest {
	s.TrafficChangeType = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetTrafficChangeTypeDesc(v string) *ThirdImmediateMsgPushRequest {
	s.TrafficChangeTypeDesc = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetTrafficJourneyIds(v string) *ThirdImmediateMsgPushRequest {
	s.TrafficJourneyIds = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetTrafficSubOrderIds(v string) *ThirdImmediateMsgPushRequest {
	s.TrafficSubOrderIds = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) SetUserId(v string) *ThirdImmediateMsgPushRequest {
	s.UserId = &v
	return s
}

func (s *ThirdImmediateMsgPushRequest) Validate() error {
	return dara.Validate(s)
}
