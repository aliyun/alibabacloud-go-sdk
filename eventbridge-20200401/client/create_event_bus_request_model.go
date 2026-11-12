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
	SetResourceGroupId(v string) *CreateEventBusRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateEventBusRequestTags) *CreateEventBusRequest
	GetTags() []*CreateEventBusRequestTags
}

type CreateEventBusRequest struct {
	// The description.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the created event bus belongs to the default resource group of the account.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tags []*CreateEventBusRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *CreateEventBusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEventBusRequest) GetTags() []*CreateEventBusRequestTags {
	return s.Tags
}

func (s *CreateEventBusRequest) SetDescription(v string) *CreateEventBusRequest {
	s.Description = &v
	return s
}

func (s *CreateEventBusRequest) SetEventBusName(v string) *CreateEventBusRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventBusRequest) SetResourceGroupId(v string) *CreateEventBusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEventBusRequest) SetTags(v []*CreateEventBusRequestTags) *CreateEventBusRequest {
	s.Tags = v
	return s
}

func (s *CreateEventBusRequest) Validate() error {
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

type CreateEventBusRequestTags struct {
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

func (s CreateEventBusRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEventBusRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEventBusRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateEventBusRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateEventBusRequestTags) SetKey(v string) *CreateEventBusRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEventBusRequestTags) SetValue(v string) *CreateEventBusRequestTags {
	s.Value = &v
	return s
}

func (s *CreateEventBusRequestTags) Validate() error {
	return dara.Validate(s)
}
