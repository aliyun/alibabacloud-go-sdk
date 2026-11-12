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
	SetTags(v []*UpdateEventBusRequestTags) *UpdateEventBusRequest
	GetTags() []*UpdateEventBusRequestTags
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
	// The list of tags.
	Tags []*UpdateEventBusRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *UpdateEventBusRequest) GetTags() []*UpdateEventBusRequestTags {
	return s.Tags
}

func (s *UpdateEventBusRequest) SetDescription(v string) *UpdateEventBusRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventBusRequest) SetEventBusName(v string) *UpdateEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventBusRequest) SetTags(v []*UpdateEventBusRequestTags) *UpdateEventBusRequest {
	s.Tags = v
	return s
}

func (s *UpdateEventBusRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEventBusRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// R&D
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventBusRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventBusRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateEventBusRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateEventBusRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateEventBusRequestTags) SetKey(v string) *UpdateEventBusRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateEventBusRequestTags) SetValue(v string) *UpdateEventBusRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateEventBusRequestTags) Validate() error {
	return dara.Validate(s)
}
