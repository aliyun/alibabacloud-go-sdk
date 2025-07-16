// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *StopMinutesShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *StopMinutesShrinkRequest
	GetConferenceId() *string
}

type StopMinutesShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StopMinutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopMinutesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StopMinutesShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StopMinutesShrinkRequest) SetTenantContextShrink(v string) *StopMinutesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StopMinutesShrinkRequest) SetConferenceId(v string) *StopMinutesShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *StopMinutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
