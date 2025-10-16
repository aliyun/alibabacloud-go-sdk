// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*CreateMemoryEventRequestEvents) *CreateMemoryEventRequest
	GetEvents() []*CreateMemoryEventRequestEvents
}

type CreateMemoryEventRequest struct {
	// This parameter is required.
	Events []*CreateMemoryEventRequestEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
}

func (s CreateMemoryEventRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryEventRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryEventRequest) GetEvents() []*CreateMemoryEventRequestEvents {
	return s.Events
}

func (s *CreateMemoryEventRequest) SetEvents(v []*CreateMemoryEventRequestEvents) *CreateMemoryEventRequest {
	s.Events = v
	return s
}

func (s *CreateMemoryEventRequest) Validate() error {
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

type CreateMemoryEventRequestEvents struct {
	// This parameter is required.
	//
	// example:
	//
	// 33.63.112.139_29914368_MTEE3_1754878127662_117007
	EventId  *string                `json:"eventId,omitempty" xml:"eventId,omitempty"`
	Message  []map[string]*string   `json:"message,omitempty" xml:"message,omitempty" type:"Repeated"`
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fd0bbf26-3c0c-4d91-a392-42d59501e12b
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uid1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateMemoryEventRequestEvents) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryEventRequestEvents) GoString() string {
	return s.String()
}

func (s *CreateMemoryEventRequestEvents) GetEventId() *string {
	return s.EventId
}

func (s *CreateMemoryEventRequestEvents) GetMessage() []map[string]*string {
	return s.Message
}

func (s *CreateMemoryEventRequestEvents) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateMemoryEventRequestEvents) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateMemoryEventRequestEvents) GetUserId() *string {
	return s.UserId
}

func (s *CreateMemoryEventRequestEvents) SetEventId(v string) *CreateMemoryEventRequestEvents {
	s.EventId = &v
	return s
}

func (s *CreateMemoryEventRequestEvents) SetMessage(v []map[string]*string) *CreateMemoryEventRequestEvents {
	s.Message = v
	return s
}

func (s *CreateMemoryEventRequestEvents) SetMetadata(v map[string]interface{}) *CreateMemoryEventRequestEvents {
	s.Metadata = v
	return s
}

func (s *CreateMemoryEventRequestEvents) SetSessionId(v string) *CreateMemoryEventRequestEvents {
	s.SessionId = &v
	return s
}

func (s *CreateMemoryEventRequestEvents) SetUserId(v string) *CreateMemoryEventRequestEvents {
	s.UserId = &v
	return s
}

func (s *CreateMemoryEventRequestEvents) Validate() error {
	return dara.Validate(s)
}
