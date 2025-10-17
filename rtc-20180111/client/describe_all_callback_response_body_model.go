// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbacks(v []*DescribeAllCallbackResponseBodyCallbacks) *DescribeAllCallbackResponseBody
	GetCallbacks() []*DescribeAllCallbackResponseBodyCallbacks
	SetRequestId(v string) *DescribeAllCallbackResponseBody
	GetRequestId() *string
}

type DescribeAllCallbackResponseBody struct {
	Callbacks []*DescribeAllCallbackResponseBodyCallbacks `json:"Callbacks,omitempty" xml:"Callbacks,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllCallbackResponseBody) GetCallbacks() []*DescribeAllCallbackResponseBodyCallbacks {
	return s.Callbacks
}

func (s *DescribeAllCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllCallbackResponseBody) SetCallbacks(v []*DescribeAllCallbackResponseBodyCallbacks) *DescribeAllCallbackResponseBody {
	s.Callbacks = v
	return s
}

func (s *DescribeAllCallbackResponseBody) SetRequestId(v string) *DescribeAllCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllCallbackResponseBody) Validate() error {
	if s.Callbacks != nil {
		for _, item := range s.Callbacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllCallbackResponseBodyCallbacks struct {
	// example:
	//
	// RecordEvent
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 录制回调
	Name     *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	SubEvent []*DescribeAllCallbackResponseBodyCallbacksSubEvent `json:"SubEvent,omitempty" xml:"SubEvent,omitempty" type:"Repeated"`
}

func (s DescribeAllCallbackResponseBodyCallbacks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllCallbackResponseBodyCallbacks) GoString() string {
	return s.String()
}

func (s *DescribeAllCallbackResponseBodyCallbacks) GetCategory() *string {
	return s.Category
}

func (s *DescribeAllCallbackResponseBodyCallbacks) GetName() *string {
	return s.Name
}

func (s *DescribeAllCallbackResponseBodyCallbacks) GetSubEvent() []*DescribeAllCallbackResponseBodyCallbacksSubEvent {
	return s.SubEvent
}

func (s *DescribeAllCallbackResponseBodyCallbacks) SetCategory(v string) *DescribeAllCallbackResponseBodyCallbacks {
	s.Category = &v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacks) SetName(v string) *DescribeAllCallbackResponseBodyCallbacks {
	s.Name = &v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacks) SetSubEvent(v []*DescribeAllCallbackResponseBodyCallbacksSubEvent) *DescribeAllCallbackResponseBodyCallbacks {
	s.SubEvent = v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacks) Validate() error {
	if s.SubEvent != nil {
		for _, item := range s.SubEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllCallbackResponseBodyCallbacksSubEvent struct {
	// example:
	//
	// 2000
	Event *int32 `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 录制开始
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllCallbackResponseBodyCallbacksSubEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllCallbackResponseBodyCallbacksSubEvent) GoString() string {
	return s.String()
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) GetEvent() *int32 {
	return s.Event
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) GetType() *int32 {
	return s.Type
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) SetEvent(v int32) *DescribeAllCallbackResponseBodyCallbacksSubEvent {
	s.Event = &v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) SetEventName(v string) *DescribeAllCallbackResponseBodyCallbacksSubEvent {
	s.EventName = &v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) SetType(v int32) *DescribeAllCallbackResponseBodyCallbacksSubEvent {
	s.Type = &v
	return s
}

func (s *DescribeAllCallbackResponseBodyCallbacksSubEvent) Validate() error {
	return dara.Validate(s)
}
