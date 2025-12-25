// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeThreatEventResponseBody
	GetRequestId() *string
	SetThreatEvents(v []*DescribeThreatEventResponseBodyThreatEvents) *DescribeThreatEventResponseBody
	GetThreatEvents() []*DescribeThreatEventResponseBodyThreatEvents
	SetTotalCount(v int64) *DescribeThreatEventResponseBody
	GetTotalCount() *int64
}

type DescribeThreatEventResponseBody struct {
	// example:
	//
	// 98106632-6865-5600-A834-3D909***
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ThreatEvents []*DescribeThreatEventResponseBodyThreatEvents `json:"ThreatEvents,omitempty" xml:"ThreatEvents,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeThreatEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeThreatEventResponseBody) GetThreatEvents() []*DescribeThreatEventResponseBodyThreatEvents {
	return s.ThreatEvents
}

func (s *DescribeThreatEventResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeThreatEventResponseBody) SetRequestId(v string) *DescribeThreatEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeThreatEventResponseBody) SetThreatEvents(v []*DescribeThreatEventResponseBodyThreatEvents) *DescribeThreatEventResponseBody {
	s.ThreatEvents = v
	return s
}

func (s *DescribeThreatEventResponseBody) SetTotalCount(v int64) *DescribeThreatEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeThreatEventResponseBody) Validate() error {
	if s.ThreatEvents != nil {
		for _, item := range s.ThreatEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeThreatEventResponseBodyThreatEvents struct {
	// example:
	//
	// 100%
	BlockRate *string `json:"BlockRate,omitempty" xml:"BlockRate,omitempty"`
	// example:
	//
	// 1768406400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// f439994c8ab39f84eced33490f0c4388
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// example:
	//
	// 3.3.3.3
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// example:
	//
	// Event_InternalLoginWeakPasswd
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// example:
	//
	// test.aliyundemo.com-waf
	EventTarget *string `json:"EventTarget,omitempty" xml:"EventTarget,omitempty"`
}

func (s DescribeThreatEventResponseBodyThreatEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventResponseBodyThreatEvents) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetBlockRate() *string {
	return s.BlockRate
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventSrc() *string {
	return s.EventSrc
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeThreatEventResponseBodyThreatEvents) GetEventTarget() *string {
	return s.EventTarget
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetBlockRate(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.BlockRate = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEndTime(v int64) *DescribeThreatEventResponseBodyThreatEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventId(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventId = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventLevel(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventLevel = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventSrc(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventSrc = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventTag(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventTag = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) SetEventTarget(v string) *DescribeThreatEventResponseBodyThreatEvents {
	s.EventTarget = &v
	return s
}

func (s *DescribeThreatEventResponseBodyThreatEvents) Validate() error {
	return dara.Validate(s)
}
