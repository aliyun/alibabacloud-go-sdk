// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTracedEventsResponseBody
	GetCode() *string
	SetData(v *QueryTracedEventsResponseBodyData) *QueryTracedEventsResponseBody
	GetData() *QueryTracedEventsResponseBodyData
	SetMessage(v string) *QueryTracedEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTracedEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTracedEventsResponseBody
	GetSuccess() *bool
}

type QueryTracedEventsResponseBody struct {
	// The response code. Valid values:
	//
	// 	- Success: The request was successful.
	//
	// 	- Other codes: The request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryTracedEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// d9e4628b-8b34-4f33-82be-5aac50aac0ba
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTracedEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTracedEventsResponseBody) GetData() *QueryTracedEventsResponseBodyData {
	return s.Data
}

func (s *QueryTracedEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTracedEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTracedEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTracedEventsResponseBody) SetCode(v string) *QueryTracedEventsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetData(v *QueryTracedEventsResponseBodyData) *QueryTracedEventsResponseBody {
	s.Data = v
	return s
}

func (s *QueryTracedEventsResponseBody) SetMessage(v string) *QueryTracedEventsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetRequestId(v string) *QueryTracedEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetSuccess(v bool) *QueryTracedEventsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTracedEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTracedEventsResponseBodyData struct {
	// The events.
	Events []*QueryTracedEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryTracedEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBodyData) GetEvents() []*QueryTracedEventsResponseBodyDataEvents {
	return s.Events
}

func (s *QueryTracedEventsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryTracedEventsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryTracedEventsResponseBodyData) SetEvents(v []*QueryTracedEventsResponseBodyDataEvents) *QueryTracedEventsResponseBodyData {
	s.Events = v
	return s
}

func (s *QueryTracedEventsResponseBodyData) SetNextToken(v string) *QueryTracedEventsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventsResponseBodyData) SetTotal(v int32) *QueryTracedEventsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryTracedEventsResponseBodyData) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTracedEventsResponseBodyDataEvents struct {
	// The name of the event bus.
	//
	// example:
	//
	// test-custom-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 07E-1OCckaVzNB92BIFFh4xgydOF1wd
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the event was delivered to the event bus.
	//
	// example:
	//
	// 1661773573100
	EventReceivedTime *int64 `json:"EventReceivedTime,omitempty" xml:"EventReceivedTime,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// acs.resourcemanager
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	//
	// example:
	//
	// eventbridge:Events:HTTPEvent
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s QueryTracedEventsResponseBodyDataEvents) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventsResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBodyDataEvents) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryTracedEventsResponseBodyDataEvents) GetEventId() *string {
	return s.EventId
}

func (s *QueryTracedEventsResponseBodyDataEvents) GetEventReceivedTime() *int64 {
	return s.EventReceivedTime
}

func (s *QueryTracedEventsResponseBodyDataEvents) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryTracedEventsResponseBodyDataEvents) GetEventType() *string {
	return s.EventType
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventBusName(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventId(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventReceivedTime(v int64) *QueryTracedEventsResponseBodyDataEvents {
	s.EventReceivedTime = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventSource(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventType(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventType = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) Validate() error {
	return dara.Validate(s)
}
