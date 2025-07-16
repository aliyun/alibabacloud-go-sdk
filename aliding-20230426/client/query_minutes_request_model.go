// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryMinutesRequestTenantContext) *QueryMinutesRequest
	GetTenantContext() *QueryMinutesRequestTenantContext
	SetConferenceId(v string) *QueryMinutesRequest
	GetConferenceId() *string
}

type QueryMinutesRequest struct {
	TenantContext *QueryMinutesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesRequest) GetTenantContext() *QueryMinutesRequestTenantContext {
	return s.TenantContext
}

func (s *QueryMinutesRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesRequest) SetTenantContext(v *QueryMinutesRequestTenantContext) *QueryMinutesRequest {
	s.TenantContext = v
	return s
}

func (s *QueryMinutesRequest) SetConferenceId(v string) *QueryMinutesRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryMinutesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMinutesRequestTenantContext) SetTenantId(v string) *QueryMinutesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryMinutesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
