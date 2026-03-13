// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannel interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *Channel
	GetDescription() *string
	SetName(v string) *Channel
	GetName() *string
	SetProperties(v map[string]interface{}) *Channel
	GetProperties() map[string]interface{}
	SetRequired(v bool) *Channel
	GetRequired() *bool
	SetSupportedChannelTypes(v []*string) *Channel
	GetSupportedChannelTypes() []*string
}

type Channel struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name                  *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties            map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Required              *bool                  `json:"Required,omitempty" xml:"Required,omitempty"`
	SupportedChannelTypes []*string              `json:"SupportedChannelTypes,omitempty" xml:"SupportedChannelTypes,omitempty" type:"Repeated"`
}

func (s Channel) String() string {
	return dara.Prettify(s)
}

func (s Channel) GoString() string {
	return s.String()
}

func (s *Channel) GetDescription() *string {
	return s.Description
}

func (s *Channel) GetName() *string {
	return s.Name
}

func (s *Channel) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *Channel) GetRequired() *bool {
	return s.Required
}

func (s *Channel) GetSupportedChannelTypes() []*string {
	return s.SupportedChannelTypes
}

func (s *Channel) SetDescription(v string) *Channel {
	s.Description = &v
	return s
}

func (s *Channel) SetName(v string) *Channel {
	s.Name = &v
	return s
}

func (s *Channel) SetProperties(v map[string]interface{}) *Channel {
	s.Properties = v
	return s
}

func (s *Channel) SetRequired(v bool) *Channel {
	s.Required = &v
	return s
}

func (s *Channel) SetSupportedChannelTypes(v []*string) *Channel {
	s.SupportedChannelTypes = v
	return s
}

func (s *Channel) Validate() error {
	return dara.Validate(s)
}
