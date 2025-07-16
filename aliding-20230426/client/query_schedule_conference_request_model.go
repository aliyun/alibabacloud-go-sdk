// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryScheduleConferenceRequestTenantContext) *QueryScheduleConferenceRequest
	GetTenantContext() *QueryScheduleConferenceRequestTenantContext
	SetScheduleConferenceId(v string) *QueryScheduleConferenceRequest
	GetScheduleConferenceId() *string
}

type QueryScheduleConferenceRequest struct {
	TenantContext *QueryScheduleConferenceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489c68-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
}

func (s QueryScheduleConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceRequest) GetTenantContext() *QueryScheduleConferenceRequestTenantContext {
	return s.TenantContext
}

func (s *QueryScheduleConferenceRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryScheduleConferenceRequest) SetTenantContext(v *QueryScheduleConferenceRequestTenantContext) *QueryScheduleConferenceRequest {
	s.TenantContext = v
	return s
}

func (s *QueryScheduleConferenceRequest) SetScheduleConferenceId(v string) *QueryScheduleConferenceRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceRequest) Validate() error {
	return dara.Validate(s)
}

type QueryScheduleConferenceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryScheduleConferenceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryScheduleConferenceRequestTenantContext) SetTenantId(v string) *QueryScheduleConferenceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryScheduleConferenceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
