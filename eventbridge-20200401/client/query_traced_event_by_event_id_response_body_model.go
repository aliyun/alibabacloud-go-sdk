// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventByEventIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTracedEventByEventIdResponseBody
	GetCode() *string
	SetData(v []*QueryTracedEventByEventIdResponseBodyData) *QueryTracedEventByEventIdResponseBody
	GetData() []*QueryTracedEventByEventIdResponseBodyData
	SetMessage(v string) *QueryTracedEventByEventIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTracedEventByEventIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTracedEventByEventIdResponseBody
	GetSuccess() *bool
}

type QueryTracedEventByEventIdResponseBody struct {
	// The response code. Valid values:
	//
	// Success: The request was successful.
	//
	// Other codes: The request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	Data []*QueryTracedEventByEventIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// A5D7B9F4-BF96-51A9-90B1-928955FABB5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTracedEventByEventIdResponseBody) GetData() []*QueryTracedEventByEventIdResponseBodyData {
	return s.Data
}

func (s *QueryTracedEventByEventIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTracedEventByEventIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTracedEventByEventIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTracedEventByEventIdResponseBody) SetCode(v string) *QueryTracedEventByEventIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetData(v []*QueryTracedEventByEventIdResponseBodyData) *QueryTracedEventByEventIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetMessage(v string) *QueryTracedEventByEventIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetRequestId(v string) *QueryTracedEventByEventIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetSuccess(v bool) *QueryTracedEventByEventIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTracedEventByEventIdResponseBodyData struct {
	// The events.
	Events []*QueryTracedEventByEventIdResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBodyData) GetEvents() []*QueryTracedEventByEventIdResponseBodyDataEvents {
	return s.Events
}

func (s *QueryTracedEventByEventIdResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryTracedEventByEventIdResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetEvents(v []*QueryTracedEventByEventIdResponseBodyDataEvents) *QueryTracedEventByEventIdResponseBodyData {
	s.Events = v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetNextToken(v string) *QueryTracedEventByEventIdResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetTotal(v int32) *QueryTracedEventByEventIdResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyData) Validate() error {
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

type QueryTracedEventByEventIdResponseBodyDataEvents struct {
	// The name of the event bus.
	//
	// example:
	//
	// default
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 37C-1P6Yn6EM7TcH37Vod8w7rbSeimJ
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
	// cert-api
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	//
	// example:
	//
	// eventbridge:Events:HTTPEvent
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBodyDataEvents) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) GetEventId() *string {
	return s.EventId
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) GetEventReceivedTime() *int64 {
	return s.EventReceivedTime
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) GetEventType() *string {
	return s.EventType
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventBusName(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventId(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventReceivedTime(v int64) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventReceivedTime = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventSource(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventType(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventType = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) Validate() error {
	return dara.Validate(s)
}
