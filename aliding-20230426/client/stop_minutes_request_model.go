// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *StopMinutesRequestTenantContext) *StopMinutesRequest
	GetTenantContext() *StopMinutesRequestTenantContext
	SetConferenceId(v string) *StopMinutesRequest
	GetConferenceId() *string
}

type StopMinutesRequest struct {
	TenantContext *StopMinutesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s StopMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesRequest) GoString() string {
	return s.String()
}

func (s *StopMinutesRequest) GetTenantContext() *StopMinutesRequestTenantContext {
	return s.TenantContext
}

func (s *StopMinutesRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StopMinutesRequest) SetTenantContext(v *StopMinutesRequestTenantContext) *StopMinutesRequest {
	s.TenantContext = v
	return s
}

func (s *StopMinutesRequest) SetConferenceId(v string) *StopMinutesRequest {
	s.ConferenceId = &v
	return s
}

func (s *StopMinutesRequest) Validate() error {
	return dara.Validate(s)
}

type StopMinutesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StopMinutesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *StopMinutesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *StopMinutesRequestTenantContext) SetTenantId(v string) *StopMinutesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *StopMinutesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
