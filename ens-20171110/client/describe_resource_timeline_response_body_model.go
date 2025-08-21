// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableEvents(v []*DescribeResourceTimelineResponseBodyAvailableEvents) *DescribeResourceTimelineResponseBody
	GetAvailableEvents() []*DescribeResourceTimelineResponseBodyAvailableEvents
	SetBizEvents(v []*DescribeResourceTimelineResponseBodyBizEvents) *DescribeResourceTimelineResponseBody
	GetBizEvents() []*DescribeResourceTimelineResponseBodyBizEvents
	SetDesc(v string) *DescribeResourceTimelineResponseBody
	GetDesc() *string
	SetInventoryEvents(v []*DescribeResourceTimelineResponseBodyInventoryEvents) *DescribeResourceTimelineResponseBody
	GetInventoryEvents() []*DescribeResourceTimelineResponseBodyInventoryEvents
	SetMsg(v string) *DescribeResourceTimelineResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeResourceTimelineResponseBody
	GetRequestId() *string
	SetReserveEvents(v []*DescribeResourceTimelineResponseBodyReserveEvents) *DescribeResourceTimelineResponseBody
	GetReserveEvents() []*DescribeResourceTimelineResponseBodyReserveEvents
}

type DescribeResourceTimelineResponseBody struct {
	AvailableEvents []*DescribeResourceTimelineResponseBodyAvailableEvents `json:"AvailableEvents,omitempty" xml:"AvailableEvents,omitempty" type:"Repeated"`
	BizEvents       []*DescribeResourceTimelineResponseBodyBizEvents       `json:"BizEvents,omitempty" xml:"BizEvents,omitempty" type:"Repeated"`
	Desc            *string                                                `json:"Desc,omitempty" xml:"Desc,omitempty"`
	InventoryEvents []*DescribeResourceTimelineResponseBodyInventoryEvents `json:"InventoryEvents,omitempty" xml:"InventoryEvents,omitempty" type:"Repeated"`
	Msg             *string                                                `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserveEvents   []*DescribeResourceTimelineResponseBodyReserveEvents   `json:"ReserveEvents,omitempty" xml:"ReserveEvents,omitempty" type:"Repeated"`
}

func (s DescribeResourceTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponseBody) GetAvailableEvents() []*DescribeResourceTimelineResponseBodyAvailableEvents {
	return s.AvailableEvents
}

func (s *DescribeResourceTimelineResponseBody) GetBizEvents() []*DescribeResourceTimelineResponseBodyBizEvents {
	return s.BizEvents
}

func (s *DescribeResourceTimelineResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *DescribeResourceTimelineResponseBody) GetInventoryEvents() []*DescribeResourceTimelineResponseBodyInventoryEvents {
	return s.InventoryEvents
}

func (s *DescribeResourceTimelineResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeResourceTimelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceTimelineResponseBody) GetReserveEvents() []*DescribeResourceTimelineResponseBodyReserveEvents {
	return s.ReserveEvents
}

func (s *DescribeResourceTimelineResponseBody) SetAvailableEvents(v []*DescribeResourceTimelineResponseBodyAvailableEvents) *DescribeResourceTimelineResponseBody {
	s.AvailableEvents = v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetBizEvents(v []*DescribeResourceTimelineResponseBodyBizEvents) *DescribeResourceTimelineResponseBody {
	s.BizEvents = v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetDesc(v string) *DescribeResourceTimelineResponseBody {
	s.Desc = &v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetInventoryEvents(v []*DescribeResourceTimelineResponseBodyInventoryEvents) *DescribeResourceTimelineResponseBody {
	s.InventoryEvents = v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetMsg(v string) *DescribeResourceTimelineResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetRequestId(v string) *DescribeResourceTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceTimelineResponseBody) SetReserveEvents(v []*DescribeResourceTimelineResponseBodyReserveEvents) *DescribeResourceTimelineResponseBody {
	s.ReserveEvents = v
	return s
}

func (s *DescribeResourceTimelineResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceTimelineResponseBodyAvailableEvents struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceTimelineResponseBodyAvailableEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponseBodyAvailableEvents) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) GetName() *string {
	return s.Name
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) GetType() *string {
	return s.Type
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) SetName(v string) *DescribeResourceTimelineResponseBodyAvailableEvents {
	s.Name = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) SetOccurrenceTime(v string) *DescribeResourceTimelineResponseBodyAvailableEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) SetReason(v string) *DescribeResourceTimelineResponseBodyAvailableEvents {
	s.Reason = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) SetType(v string) *DescribeResourceTimelineResponseBodyAvailableEvents {
	s.Type = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyAvailableEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceTimelineResponseBodyBizEvents struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceTimelineResponseBodyBizEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponseBodyBizEvents) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) GetName() *string {
	return s.Name
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) GetType() *string {
	return s.Type
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) SetName(v string) *DescribeResourceTimelineResponseBodyBizEvents {
	s.Name = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) SetOccurrenceTime(v string) *DescribeResourceTimelineResponseBodyBizEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) SetReason(v string) *DescribeResourceTimelineResponseBodyBizEvents {
	s.Reason = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) SetType(v string) *DescribeResourceTimelineResponseBodyBizEvents {
	s.Type = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyBizEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceTimelineResponseBodyInventoryEvents struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceTimelineResponseBodyInventoryEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponseBodyInventoryEvents) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) GetName() *string {
	return s.Name
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) GetType() *string {
	return s.Type
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) SetName(v string) *DescribeResourceTimelineResponseBodyInventoryEvents {
	s.Name = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) SetOccurrenceTime(v string) *DescribeResourceTimelineResponseBodyInventoryEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) SetReason(v string) *DescribeResourceTimelineResponseBodyInventoryEvents {
	s.Reason = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) SetType(v string) *DescribeResourceTimelineResponseBodyInventoryEvents {
	s.Type = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyInventoryEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceTimelineResponseBodyReserveEvents struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceTimelineResponseBodyReserveEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTimelineResponseBodyReserveEvents) GoString() string {
	return s.String()
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) GetName() *string {
	return s.Name
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) GetType() *string {
	return s.Type
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) SetName(v string) *DescribeResourceTimelineResponseBodyReserveEvents {
	s.Name = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) SetOccurrenceTime(v string) *DescribeResourceTimelineResponseBodyReserveEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) SetReason(v string) *DescribeResourceTimelineResponseBodyReserveEvents {
	s.Reason = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) SetType(v string) *DescribeResourceTimelineResponseBodyReserveEvents {
	s.Type = &v
	return s
}

func (s *DescribeResourceTimelineResponseBodyReserveEvents) Validate() error {
	return dara.Validate(s)
}
