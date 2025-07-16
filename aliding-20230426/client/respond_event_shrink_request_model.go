// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *RespondEventShrinkRequest
	GetCalendarId() *string
	SetEventId(v string) *RespondEventShrinkRequest
	GetEventId() *string
	SetResponseStatus(v string) *RespondEventShrinkRequest
	GetResponseStatus() *string
	SetTenantContextShrink(v string) *RespondEventShrinkRequest
	GetTenantContextShrink() *string
}

type RespondEventShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RHN2REJFc2w4VHNiUUlvcVB0ejFydz09
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// accepted
	ResponseStatus      *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s RespondEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RespondEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *RespondEventShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RespondEventShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *RespondEventShrinkRequest) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *RespondEventShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *RespondEventShrinkRequest) SetCalendarId(v string) *RespondEventShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *RespondEventShrinkRequest) SetEventId(v string) *RespondEventShrinkRequest {
	s.EventId = &v
	return s
}

func (s *RespondEventShrinkRequest) SetResponseStatus(v string) *RespondEventShrinkRequest {
	s.ResponseStatus = &v
	return s
}

func (s *RespondEventShrinkRequest) SetTenantContextShrink(v string) *RespondEventShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *RespondEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
