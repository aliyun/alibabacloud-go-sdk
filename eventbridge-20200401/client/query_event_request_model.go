// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *QueryEventRequest
	GetEventBusName() *string
	SetEventId(v string) *QueryEventRequest
	GetEventId() *string
	SetEventSource(v string) *QueryEventRequest
	GetEventSource() *string
}

type QueryEventRequest struct {
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
	// 	- This parameter is required if you query the system event bus.
	//
	// example:
	//
	// testEventSourceName
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
}

func (s QueryEventRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryEventRequest) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryEventRequest) SetEventBusName(v string) *QueryEventRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryEventRequest) SetEventId(v string) *QueryEventRequest {
	s.EventId = &v
	return s
}

func (s *QueryEventRequest) SetEventSource(v string) *QueryEventRequest {
	s.EventSource = &v
	return s
}

func (s *QueryEventRequest) Validate() error {
	return dara.Validate(s)
}
