// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOSSTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*string) *OSSTriggerConfig
	GetEvents() []*string
	SetFilter(v *Filter) *OSSTriggerConfig
	GetFilter() *Filter
}

type OSSTriggerConfig struct {
	// The list of events. Enter an event related to OSS. For more information about event types,see [OSS events](https://help.aliyun.com/document_detail/2513613.html).
	Events []*string `json:"events" xml:"events" type:"Repeated"`
	// The event filtering rule.
	Filter *Filter `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s OSSTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s OSSTriggerConfig) GoString() string {
	return s.String()
}

func (s *OSSTriggerConfig) GetEvents() []*string {
	return s.Events
}

func (s *OSSTriggerConfig) GetFilter() *Filter {
	return s.Filter
}

func (s *OSSTriggerConfig) SetEvents(v []*string) *OSSTriggerConfig {
	s.Events = v
	return s
}

func (s *OSSTriggerConfig) SetFilter(v *Filter) *OSSTriggerConfig {
	s.Filter = v
	return s
}

func (s *OSSTriggerConfig) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}
