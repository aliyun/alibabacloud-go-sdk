// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCDNTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *CDNTriggerConfig
	GetEventName() *string
	SetEventVersion(v string) *CDNTriggerConfig
	GetEventVersion() *string
	SetFilter(v map[string][]*string) *CDNTriggerConfig
	GetFilter() map[string][]*string
	SetNotes(v string) *CDNTriggerConfig
	GetNotes() *string
}

type CDNTriggerConfig struct {
	// example:
	//
	// CdnDomainStarted
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// example:
	//
	// 1.0.0
	EventVersion *string              `json:"eventVersion,omitempty" xml:"eventVersion,omitempty"`
	Filter       map[string][]*string `json:"filter" xml:"filter"`
	// example:
	//
	// 缓存事件触发器
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
}

func (s CDNTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s CDNTriggerConfig) GoString() string {
	return s.String()
}

func (s *CDNTriggerConfig) GetEventName() *string {
	return s.EventName
}

func (s *CDNTriggerConfig) GetEventVersion() *string {
	return s.EventVersion
}

func (s *CDNTriggerConfig) GetFilter() map[string][]*string {
	return s.Filter
}

func (s *CDNTriggerConfig) GetNotes() *string {
	return s.Notes
}

func (s *CDNTriggerConfig) SetEventName(v string) *CDNTriggerConfig {
	s.EventName = &v
	return s
}

func (s *CDNTriggerConfig) SetEventVersion(v string) *CDNTriggerConfig {
	s.EventVersion = &v
	return s
}

func (s *CDNTriggerConfig) SetFilter(v map[string][]*string) *CDNTriggerConfig {
	s.Filter = v
	return s
}

func (s *CDNTriggerConfig) SetNotes(v string) *CDNTriggerConfig {
	s.Notes = &v
	return s
}

func (s *CDNTriggerConfig) Validate() error {
	return dara.Validate(s)
}
