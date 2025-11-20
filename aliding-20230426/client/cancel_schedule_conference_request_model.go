// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleConferenceId(v string) *CancelScheduleConferenceRequest
	GetScheduleConferenceId() *string
	SetTenantContext(v *CancelScheduleConferenceRequestTenantContext) *CancelScheduleConferenceRequest
	GetTenantContext() *CancelScheduleConferenceRequestTenantContext
}

type CancelScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string                                       `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContext        *CancelScheduleConferenceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CancelScheduleConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *CancelScheduleConferenceRequest) GetTenantContext() *CancelScheduleConferenceRequestTenantContext {
	return s.TenantContext
}

func (s *CancelScheduleConferenceRequest) SetScheduleConferenceId(v string) *CancelScheduleConferenceRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *CancelScheduleConferenceRequest) SetTenantContext(v *CancelScheduleConferenceRequestTenantContext) *CancelScheduleConferenceRequest {
	s.TenantContext = v
	return s
}

func (s *CancelScheduleConferenceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelScheduleConferenceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CancelScheduleConferenceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CancelScheduleConferenceRequestTenantContext) SetTenantId(v string) *CancelScheduleConferenceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CancelScheduleConferenceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
