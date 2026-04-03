// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHook interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *Hook
	GetApiVersion() *string
	SetDescription(v string) *Hook
	GetDescription() *string
	SetEnabled(v bool) *Hook
	GetEnabled() *bool
	SetEvent(v string) *Hook
	GetEvent() *string
	SetHeaders(v map[string]*string) *Hook
	GetHeaders() map[string]*string
	SetTimeout(v int32) *Hook
	GetTimeout() *int32
	SetUrl(v string) *Hook
	GetUrl() *string
}

type Hook struct {
	ApiVersion  *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// This parameter is required.
	Event   *string            `json:"event,omitempty" xml:"event,omitempty"`
	Headers map[string]*string `json:"headers" xml:"headers"`
	Timeout *int32             `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s Hook) String() string {
	return dara.Prettify(s)
}

func (s Hook) GoString() string {
	return s.String()
}

func (s *Hook) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *Hook) GetDescription() *string {
	return s.Description
}

func (s *Hook) GetEnabled() *bool {
	return s.Enabled
}

func (s *Hook) GetEvent() *string {
	return s.Event
}

func (s *Hook) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Hook) GetTimeout() *int32 {
	return s.Timeout
}

func (s *Hook) GetUrl() *string {
	return s.Url
}

func (s *Hook) SetApiVersion(v string) *Hook {
	s.ApiVersion = &v
	return s
}

func (s *Hook) SetDescription(v string) *Hook {
	s.Description = &v
	return s
}

func (s *Hook) SetEnabled(v bool) *Hook {
	s.Enabled = &v
	return s
}

func (s *Hook) SetEvent(v string) *Hook {
	s.Event = &v
	return s
}

func (s *Hook) SetHeaders(v map[string]*string) *Hook {
	s.Headers = v
	return s
}

func (s *Hook) SetTimeout(v int32) *Hook {
	s.Timeout = &v
	return s
}

func (s *Hook) SetUrl(v string) *Hook {
	s.Url = &v
	return s
}

func (s *Hook) Validate() error {
	return dara.Validate(s)
}
