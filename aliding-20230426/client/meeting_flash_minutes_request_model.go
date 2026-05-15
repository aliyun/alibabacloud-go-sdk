// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *MeetingFlashMinutesRequestTenantContext) *MeetingFlashMinutesRequest
	GetTenantContext() *MeetingFlashMinutesRequestTenantContext
	SetConferenceId(v string) *MeetingFlashMinutesRequest
	GetConferenceId() *string
}

type MeetingFlashMinutesRequest struct {
	TenantContext *MeetingFlashMinutesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s MeetingFlashMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesRequest) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesRequest) GetTenantContext() *MeetingFlashMinutesRequestTenantContext {
	return s.TenantContext
}

func (s *MeetingFlashMinutesRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MeetingFlashMinutesRequest) SetTenantContext(v *MeetingFlashMinutesRequestTenantContext) *MeetingFlashMinutesRequest {
	s.TenantContext = v
	return s
}

func (s *MeetingFlashMinutesRequest) SetConferenceId(v string) *MeetingFlashMinutesRequest {
	s.ConferenceId = &v
	return s
}

func (s *MeetingFlashMinutesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MeetingFlashMinutesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *MeetingFlashMinutesRequestTenantContext) SetTenantId(v string) *MeetingFlashMinutesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *MeetingFlashMinutesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
