// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkMeetingInfoRequestTenantContext) *GetDingtalkMeetingInfoRequest
	GetTenantContext() *GetDingtalkMeetingInfoRequestTenantContext
	SetConferenceId(v string) *GetDingtalkMeetingInfoRequest
	GetConferenceId() *string
	SetOrgId(v string) *GetDingtalkMeetingInfoRequest
	GetOrgId() *string
}

type GetDingtalkMeetingInfoRequest struct {
	TenantContext *GetDingtalkMeetingInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
}

func (s GetDingtalkMeetingInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoRequest) GetTenantContext() *GetDingtalkMeetingInfoRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkMeetingInfoRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingInfoRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingInfoRequest) SetTenantContext(v *GetDingtalkMeetingInfoRequestTenantContext) *GetDingtalkMeetingInfoRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkMeetingInfoRequest) SetConferenceId(v string) *GetDingtalkMeetingInfoRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingInfoRequest) SetOrgId(v string) *GetDingtalkMeetingInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkMeetingInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkMeetingInfoRequestTenantContext) SetTenantId(v string) *GetDingtalkMeetingInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkMeetingInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
