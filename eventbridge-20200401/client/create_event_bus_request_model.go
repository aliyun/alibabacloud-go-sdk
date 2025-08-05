// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventBusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEventBusRequest
	GetDescription() *string
	SetEventBusName(v string) *CreateEventBusRequest
	GetEventBusName() *string
}

type CreateEventBusRequest struct {
	// The description of the event bus.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the request is successful. The value true indicates that the request is successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s CreateEventBusRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventBusRequest) GoString() string {
	return s.String()
}

func (s *CreateEventBusRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventBusRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateEventBusRequest) SetDescription(v string) *CreateEventBusRequest {
	s.Description = &v
	return s
}

func (s *CreateEventBusRequest) SetEventBusName(v string) *CreateEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventBusRequest) Validate() error {
	return dara.Validate(s)
}
