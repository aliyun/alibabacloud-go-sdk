// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomEventAttributeResponseBody
	GetCode() *string
	SetCustomEvents(v *DescribeCustomEventAttributeResponseBodyCustomEvents) *DescribeCustomEventAttributeResponseBody
	GetCustomEvents() *DescribeCustomEventAttributeResponseBodyCustomEvents
	SetMessage(v string) *DescribeCustomEventAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomEventAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCustomEventAttributeResponseBody
	GetSuccess() *string
}

type DescribeCustomEventAttributeResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The event details.
	CustomEvents *DescribeCustomEventAttributeResponseBodyCustomEvents `json:"CustomEvents,omitempty" xml:"CustomEvents,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// userId:173651113438***	- and name:"BABEL_CHECK"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60912C8D-B340-4253-ADE7-61ACDFD25CFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomEventAttributeResponseBody) GetCustomEvents() *DescribeCustomEventAttributeResponseBodyCustomEvents {
	return s.CustomEvents
}

func (s *DescribeCustomEventAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomEventAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomEventAttributeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCustomEventAttributeResponseBody) SetCode(v string) *DescribeCustomEventAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetCustomEvents(v *DescribeCustomEventAttributeResponseBodyCustomEvents) *DescribeCustomEventAttributeResponseBody {
	s.CustomEvents = v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetMessage(v string) *DescribeCustomEventAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetRequestId(v string) *DescribeCustomEventAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetSuccess(v string) *DescribeCustomEventAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomEventAttributeResponseBodyCustomEvents struct {
	CustomEvent []*DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent `json:"CustomEvent,omitempty" xml:"CustomEvent,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventAttributeResponseBodyCustomEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBodyCustomEvents) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEvents) GetCustomEvent() []*DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	return s.CustomEvent
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEvents) SetCustomEvent(v []*DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) *DescribeCustomEventAttributeResponseBodyCustomEvents {
	s.CustomEvent = v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent struct {
	// The content of the event.
	//
	// example:
	//
	// requestId:4975A6F3-19AC-4C01-BAD2-034DA07FEBB5, info:{"autoPay":false,"autoUseCoupon":false,"bid":"26842","buyerId":118935342242****,"commodities":[{"aliyunProduceCode":"cms","chargeType":"PREPAY","commodityCode":"cms_call_num","components":[{"componentCode":"phone_count","instanceProperty":[{"code":"phone_count","value":"500"}],"moduleAttrStatus":1}],"duration":6,"free":false,"orderParams":{"aliyunProduceCode":"cms"},"orderType":"BUY","prePayPostCharge":false,"pricingCycle":"Month","quantity":1,"refundSpecCode":"","renewChange":false,"specCode":"cms_call_num","specUpdate":false,"syncToSubscription":false,"upgradeInquireFinancialValue":true}],"fromApp":"commonbuy","orderParams":{"priceCheck":"true"},"payerId":118935342242****,"requestId":"ade3ad32-f58b-45d7-add4-ac542be3d8ec","skipChannel":false,"userId":118935342242****}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The event name.
	//
	// example:
	//
	// BABEL_CHECK
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the event occurred.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1552199984000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GetContent() *string {
	return s.Content
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GetId() *string {
	return s.Id
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GetTime() *string {
	return s.Time
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetContent(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Content = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetGroupId(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetId(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Id = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetName(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetTime(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Time = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) Validate() error {
	return dara.Validate(s)
}
