// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventTracesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *QueryEventTracesRequest
	GetEventBusName() *string
	SetEventId(v string) *QueryEventTracesRequest
	GetEventId() *string
}

type QueryEventTracesRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1935debf-ddac-49dc-a090-d4f2857a046d
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s QueryEventTracesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventTracesRequest) GoString() string {
	return s.String()
}

func (s *QueryEventTracesRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryEventTracesRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryEventTracesRequest) SetEventBusName(v string) *QueryEventTracesRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryEventTracesRequest) SetEventId(v string) *QueryEventTracesRequest {
	s.EventId = &v
	return s
}

func (s *QueryEventTracesRequest) Validate() error {
	return dara.Validate(s)
}
