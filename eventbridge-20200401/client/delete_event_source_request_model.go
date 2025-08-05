// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DeleteEventSourceRequest
	GetEventBusName() *string
	SetEventSourceName(v string) *DeleteEventSourceRequest
	GetEventSourceName() *string
}

type DeleteEventSourceRequest struct {
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	//
	// This parameter is required.
	//
	// example:
	//
	// myrabbitmq.source
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
}

func (s DeleteEventSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DeleteEventSourceRequest) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *DeleteEventSourceRequest) SetEventBusName(v string) *DeleteEventSourceRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteEventSourceRequest) SetEventSourceName(v string) *DeleteEventSourceRequest {
	s.EventSourceName = &v
	return s
}

func (s *DeleteEventSourceRequest) Validate() error {
	return dara.Validate(s)
}
