// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventBusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *GetEventBusRequest
	GetEventBusName() *string
}

type GetEventBusRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s GetEventBusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusRequest) GoString() string {
	return s.String()
}

func (s *GetEventBusRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetEventBusRequest) SetEventBusName(v string) *GetEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *GetEventBusRequest) Validate() error {
	return dara.Validate(s)
}
