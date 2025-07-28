// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvent(v []*DescribeUserEventTypeResponseBodyEvent) *DescribeUserEventTypeResponseBody
	GetEvent() []*DescribeUserEventTypeResponseBodyEvent
	SetRequestId(v string) *DescribeUserEventTypeResponseBody
	GetRequestId() *string
}

type DescribeUserEventTypeResponseBody struct {
	// The types and statistics of security events.
	Event []*DescribeUserEventTypeResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 177BA739-6512-5470-98C6-E***0BAA3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEventTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTypeResponseBody) GetEvent() []*DescribeUserEventTypeResponseBodyEvent {
	return s.Event
}

func (s *DescribeUserEventTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEventTypeResponseBody) SetEvent(v []*DescribeUserEventTypeResponseBodyEvent) *DescribeUserEventTypeResponseBody {
	s.Event = v
	return s
}

func (s *DescribeUserEventTypeResponseBody) SetRequestId(v string) *DescribeUserEventTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEventTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserEventTypeResponseBodyEvent struct {
	// The code of the security event.
	//
	// example:
	//
	// Event_LoginCollision
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The number of events.
	//
	// example:
	//
	// 0
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The parent type of the security event.
	//
	// example:
	//
	// EventType_Account
	EventParentType *string `json:"EventParentType,omitempty" xml:"EventParentType,omitempty"`
	// The type of the security event.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of security events.
	//
	// example:
	//
	// Event_AbnormalFrequency
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s DescribeUserEventTypeResponseBodyEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTypeResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTypeResponseBodyEvent) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeUserEventTypeResponseBodyEvent) GetEventCount() *int64 {
	return s.EventCount
}

func (s *DescribeUserEventTypeResponseBodyEvent) GetEventParentType() *string {
	return s.EventParentType
}

func (s *DescribeUserEventTypeResponseBodyEvent) GetEventType() *string {
	return s.EventType
}

func (s *DescribeUserEventTypeResponseBodyEvent) SetEventCode(v string) *DescribeUserEventTypeResponseBodyEvent {
	s.EventCode = &v
	return s
}

func (s *DescribeUserEventTypeResponseBodyEvent) SetEventCount(v int64) *DescribeUserEventTypeResponseBodyEvent {
	s.EventCount = &v
	return s
}

func (s *DescribeUserEventTypeResponseBodyEvent) SetEventParentType(v string) *DescribeUserEventTypeResponseBodyEvent {
	s.EventParentType = &v
	return s
}

func (s *DescribeUserEventTypeResponseBodyEvent) SetEventType(v string) *DescribeUserEventTypeResponseBodyEvent {
	s.EventType = &v
	return s
}

func (s *DescribeUserEventTypeResponseBodyEvent) Validate() error {
	return dara.Validate(s)
}
