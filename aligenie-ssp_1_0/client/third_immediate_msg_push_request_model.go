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
	// example:
	//
	// FLIGHT
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// changeDetail
	ChangeDetail *string `json:"ChangeDetail,omitempty" xml:"ChangeDetail,omitempty"`
	// example:
	//
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// psgIds
	PsgIds *string `json:"PsgIds,omitempty" xml:"PsgIds,omitempty"`
	// example:
	//
	// trafficChangeType
	TrafficChangeType *string `json:"TrafficChangeType,omitempty" xml:"TrafficChangeType,omitempty"`
	// example:
	//
	// trafficChangeTypeDesc
	TrafficChangeTypeDesc *string `json:"TrafficChangeTypeDesc,omitempty" xml:"TrafficChangeTypeDesc,omitempty"`
	// example:
	//
	// trafficJourneyIds
	TrafficJourneyIds *string `json:"TrafficJourneyIds,omitempty" xml:"TrafficJourneyIds,omitempty"`
	// example:
	//
	// trafficSubOrderIds
	TrafficSubOrderIds *string `json:"TrafficSubOrderIds,omitempty" xml:"TrafficSubOrderIds,omitempty"`
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
