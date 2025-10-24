// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v *DescribeFreeUserEventCountResponseBodyEvent) *DescribeFreeUserEventCountResponseBody
	GetEvent() *DescribeFreeUserEventCountResponseBodyEvent
	SetRequestId(v string) *DescribeFreeUserEventCountResponseBody
	GetRequestId() *string
}

type DescribeFreeUserEventCountResponseBody struct {
	// The information about the security events that are detected by using the basic detection feature.
	Event *DescribeFreeUserEventCountResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0D9FB3BC-0DE9-58A8-9663-ACE56F24F405
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFreeUserEventCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventCountResponseBody) GetEvent() *DescribeFreeUserEventCountResponseBodyEvent {
	return s.Event
}

func (s *DescribeFreeUserEventCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFreeUserEventCountResponseBody) SetEvent(v *DescribeFreeUserEventCountResponseBodyEvent) *DescribeFreeUserEventCountResponseBody {
	s.Event = v
	return s
}

func (s *DescribeFreeUserEventCountResponseBody) SetRequestId(v string) *DescribeFreeUserEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFreeUserEventCountResponseBody) Validate() error {
	if s.Event != nil {
		if err := s.Event.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFreeUserEventCountResponseBodyEvent struct {
	// The number of high-risk events.
	//
	// example:
	//
	// 1
	EventHigh *int64 `json:"EventHigh,omitempty" xml:"EventHigh,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 12
	EventLow *int64 `json:"EventLow,omitempty" xml:"EventLow,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 3
	EventMedium *int64 `json:"EventMedium,omitempty" xml:"EventMedium,omitempty"`
	// The total number of security events.
	//
	// example:
	//
	// 16
	EventTotal *int64 `json:"EventTotal,omitempty" xml:"EventTotal,omitempty"`
}

func (s DescribeFreeUserEventCountResponseBodyEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventCountResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) GetEventHigh() *int64 {
	return s.EventHigh
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) GetEventLow() *int64 {
	return s.EventLow
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) GetEventMedium() *int64 {
	return s.EventMedium
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) GetEventTotal() *int64 {
	return s.EventTotal
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) SetEventHigh(v int64) *DescribeFreeUserEventCountResponseBodyEvent {
	s.EventHigh = &v
	return s
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) SetEventLow(v int64) *DescribeFreeUserEventCountResponseBodyEvent {
	s.EventLow = &v
	return s
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) SetEventMedium(v int64) *DescribeFreeUserEventCountResponseBodyEvent {
	s.EventMedium = &v
	return s
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) SetEventTotal(v int64) *DescribeFreeUserEventCountResponseBodyEvent {
	s.EventTotal = &v
	return s
}

func (s *DescribeFreeUserEventCountResponseBodyEvent) Validate() error {
	return dara.Validate(s)
}
