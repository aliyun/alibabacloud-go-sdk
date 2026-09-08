// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *PutEventsShrinkRequest
	GetEventBusName() *string
	SetEventListShrink(v string) *PutEventsShrinkRequest
	GetEventListShrink() *string
}

type PutEventsShrinkRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The list of events.
	EventListShrink *string `json:"EventList,omitempty" xml:"EventList,omitempty"`
}

func (s PutEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutEventsShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *PutEventsShrinkRequest) GetEventListShrink() *string {
	return s.EventListShrink
}

func (s *PutEventsShrinkRequest) SetEventBusName(v string) *PutEventsShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *PutEventsShrinkRequest) SetEventListShrink(v string) *PutEventsShrinkRequest {
	s.EventListShrink = &v
	return s
}

func (s *PutEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
