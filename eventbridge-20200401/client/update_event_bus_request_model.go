// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventBusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateEventBusRequest
	GetDescription() *string
	SetEventBusName(v string) *UpdateEventBusRequest
	GetEventBusName() *string
}

type UpdateEventBusRequest struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// eventTest
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s UpdateEventBusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventBusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventBusRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEventBusRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateEventBusRequest) SetDescription(v string) *UpdateEventBusRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventBusRequest) SetEventBusName(v string) *UpdateEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventBusRequest) Validate() error {
	return dara.Validate(s)
}
