// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *DeleteEventBusRequest
	GetEventBusName() *string
}

type DeleteEventBusRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s DeleteEventBusRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBusRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventBusRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *DeleteEventBusRequest) SetEventBusName(v string) *DeleteEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteEventBusRequest) Validate() error {
	return dara.Validate(s)
}
