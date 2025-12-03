// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderSummaryForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetOrderSummaryForPartnerResponseBodyData) *GetOrderSummaryForPartnerResponseBody
	GetData() *GetOrderSummaryForPartnerResponseBodyData
	SetRequestId(v string) *GetOrderSummaryForPartnerResponseBody
	GetRequestId() *string
}

type GetOrderSummaryForPartnerResponseBody struct {
	Data *GetOrderSummaryForPartnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 56B009EB-AAA5-52C9-B05F-AF425E3885E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOrderSummaryForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrderSummaryForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderSummaryForPartnerResponseBody) GetData() *GetOrderSummaryForPartnerResponseBodyData {
	return s.Data
}

func (s *GetOrderSummaryForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrderSummaryForPartnerResponseBody) SetData(v *GetOrderSummaryForPartnerResponseBodyData) *GetOrderSummaryForPartnerResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBody) SetRequestId(v string) *GetOrderSummaryForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderSummaryForPartnerResponseBodyData struct {
	// example:
	//
	// UShMru9tcC0RHD/04O0x1OrVkEJOq804ccSkgXSuZH0/42FlhtzyCKa6FofPs8n0
	EncryptedBuyerId *string `json:"EncryptedBuyerId,omitempty" xml:"EncryptedBuyerId,omitempty"`
	// example:
	//
	// UShMru9tcC0RHD/04O0x1OrVkEJOq804ccSkgXSuZH0/42FlhtzyCKa6FofPs8n0
	EncryptedPayerId *string `json:"EncryptedPayerId,omitempty" xml:"EncryptedPayerId,omitempty"`
	// example:
	//
	// UShMru9tcC0RHD/04O0x1OrVkEJOq804ccSkgXSuZH0/42FlhtzyCKa6FofPs8n0
	EncryptedUserId *string `json:"EncryptedUserId,omitempty" xml:"EncryptedUserId,omitempty"`
	// example:
	//
	// 11350044
	OrderId    *string                                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderLines []*GetOrderSummaryForPartnerResponseBodyDataOrderLines `json:"OrderLines,omitempty" xml:"OrderLines,omitempty" type:"Repeated"`
}

func (s GetOrderSummaryForPartnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrderSummaryForPartnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrderSummaryForPartnerResponseBodyData) GetEncryptedBuyerId() *string {
	return s.EncryptedBuyerId
}

func (s *GetOrderSummaryForPartnerResponseBodyData) GetEncryptedPayerId() *string {
	return s.EncryptedPayerId
}

func (s *GetOrderSummaryForPartnerResponseBodyData) GetEncryptedUserId() *string {
	return s.EncryptedUserId
}

func (s *GetOrderSummaryForPartnerResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrderSummaryForPartnerResponseBodyData) GetOrderLines() []*GetOrderSummaryForPartnerResponseBodyDataOrderLines {
	return s.OrderLines
}

func (s *GetOrderSummaryForPartnerResponseBodyData) SetEncryptedBuyerId(v string) *GetOrderSummaryForPartnerResponseBodyData {
	s.EncryptedBuyerId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyData) SetEncryptedPayerId(v string) *GetOrderSummaryForPartnerResponseBodyData {
	s.EncryptedPayerId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyData) SetEncryptedUserId(v string) *GetOrderSummaryForPartnerResponseBodyData {
	s.EncryptedUserId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyData) SetOrderId(v string) *GetOrderSummaryForPartnerResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyData) SetOrderLines(v []*GetOrderSummaryForPartnerResponseBodyDataOrderLines) *GetOrderSummaryForPartnerResponseBodyData {
	s.OrderLines = v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyData) Validate() error {
	if s.OrderLines != nil {
		for _, item := range s.OrderLines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderSummaryForPartnerResponseBodyDataOrderLines struct {
	// example:
	//
	// ob4twsebrj1734
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 11350044
	OrderLineId *string `json:"OrderLineId,omitempty" xml:"OrderLineId,omitempty"`
	// example:
	//
	// UPGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s GetOrderSummaryForPartnerResponseBodyDataOrderLines) String() string {
	return dara.Prettify(s)
}

func (s GetOrderSummaryForPartnerResponseBodyDataOrderLines) GoString() string {
	return s.String()
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) GetOrderType() *string {
	return s.OrderType
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) SetInstanceId(v string) *GetOrderSummaryForPartnerResponseBodyDataOrderLines {
	s.InstanceId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) SetOrderLineId(v string) *GetOrderSummaryForPartnerResponseBodyDataOrderLines {
	s.OrderLineId = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) SetOrderType(v string) *GetOrderSummaryForPartnerResponseBodyDataOrderLines {
	s.OrderType = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponseBodyDataOrderLines) Validate() error {
	return dara.Validate(s)
}
