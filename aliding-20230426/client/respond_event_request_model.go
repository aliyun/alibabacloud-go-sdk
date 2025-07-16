// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *RespondEventRequest
	GetCalendarId() *string
	SetEventId(v string) *RespondEventRequest
	GetEventId() *string
	SetResponseStatus(v string) *RespondEventRequest
	GetResponseStatus() *string
	SetTenantContext(v *RespondEventRequestTenantContext) *RespondEventRequest
	GetTenantContext() *RespondEventRequestTenantContext
}

type RespondEventRequest struct {
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
	ResponseStatus *string                           `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	TenantContext  *RespondEventRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s RespondEventRequest) String() string {
	return dara.Prettify(s)
}

func (s RespondEventRequest) GoString() string {
	return s.String()
}

func (s *RespondEventRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *RespondEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *RespondEventRequest) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *RespondEventRequest) GetTenantContext() *RespondEventRequestTenantContext {
	return s.TenantContext
}

func (s *RespondEventRequest) SetCalendarId(v string) *RespondEventRequest {
	s.CalendarId = &v
	return s
}

func (s *RespondEventRequest) SetEventId(v string) *RespondEventRequest {
	s.EventId = &v
	return s
}

func (s *RespondEventRequest) SetResponseStatus(v string) *RespondEventRequest {
	s.ResponseStatus = &v
	return s
}

func (s *RespondEventRequest) SetTenantContext(v *RespondEventRequestTenantContext) *RespondEventRequest {
	s.TenantContext = v
	return s
}

func (s *RespondEventRequest) Validate() error {
	return dara.Validate(s)
}

type RespondEventRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RespondEventRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s RespondEventRequestTenantContext) GoString() string {
	return s.String()
}

func (s *RespondEventRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *RespondEventRequestTenantContext) SetTenantId(v string) *RespondEventRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *RespondEventRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
