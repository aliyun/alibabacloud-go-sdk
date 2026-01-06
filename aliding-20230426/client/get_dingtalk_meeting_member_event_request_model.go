// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkMeetingMemberEventRequestTenantContext) *GetDingtalkMeetingMemberEventRequest
	GetTenantContext() *GetDingtalkMeetingMemberEventRequestTenantContext
	SetBeginTime(v int64) *GetDingtalkMeetingMemberEventRequest
	GetBeginTime() *int64
	SetConferenceId(v string) *GetDingtalkMeetingMemberEventRequest
	GetConferenceId() *string
	SetEndTime(v int64) *GetDingtalkMeetingMemberEventRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingMemberEventRequest
	GetOrgId() *string
	SetWorkNo(v string) *GetDingtalkMeetingMemberEventRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingMemberEventRequest struct {
	TenantContext *GetDingtalkMeetingMemberEventRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 1638360000000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 1638363600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 34242
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetDingtalkMeetingMemberEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventRequest) GetTenantContext() *GetDingtalkMeetingMemberEventRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkMeetingMemberEventRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDingtalkMeetingMemberEventRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMemberEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingMemberEventRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMemberEventRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingMemberEventRequest) SetTenantContext(v *GetDingtalkMeetingMemberEventRequestTenantContext) *GetDingtalkMeetingMemberEventRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) SetBeginTime(v int64) *GetDingtalkMeetingMemberEventRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) SetConferenceId(v string) *GetDingtalkMeetingMemberEventRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) SetEndTime(v int64) *GetDingtalkMeetingMemberEventRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) SetOrgId(v string) *GetDingtalkMeetingMemberEventRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) SetWorkNo(v string) *GetDingtalkMeetingMemberEventRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingMemberEventRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkMeetingMemberEventRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkMeetingMemberEventRequestTenantContext) SetTenantId(v string) *GetDingtalkMeetingMemberEventRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
