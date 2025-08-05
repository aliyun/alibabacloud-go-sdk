// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventByEventIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *QueryTracedEventByEventIdRequest
	GetEventBusName() *string
	SetEventId(v string) *QueryTracedEventByEventIdRequest
	GetEventId() *string
	SetEventSource(v string) *QueryTracedEventByEventIdRequest
	GetEventSource() *string
}

type QueryTracedEventByEventIdRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1935debf-ddac-49dc-a090-d4f2857a046d
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// mse
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
}

func (s QueryTracedEventByEventIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventByEventIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryTracedEventByEventIdRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryTracedEventByEventIdRequest) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryTracedEventByEventIdRequest) SetEventBusName(v string) *QueryTracedEventByEventIdRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventByEventIdRequest) SetEventId(v string) *QueryTracedEventByEventIdRequest {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventByEventIdRequest) SetEventSource(v string) *QueryTracedEventByEventIdRequest {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventByEventIdRequest) Validate() error {
	return dara.Validate(s)
}
